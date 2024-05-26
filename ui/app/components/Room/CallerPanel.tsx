import React, { useState } from 'react';
import { GradientBorderCard } from '../GradientBorderCard';
import { BackgroundGradientAnimation } from '../BackgroundGradientAnimation';
import RecordButton from '../RecordButton';
import WSConnection from '../../ws';
import { AudioRecorder, RecordingResult } from '../../recorder';
import { formatTimeDuration } from '../../utils';


type CallerPanelProps = {
    historyShown: boolean
    wsConnection: WSConnection
}

export default function CallerPanel(props: CallerPanelProps) {
    const [status, setStatus] = useState('Calling');
    const [isRecordingDisabled, setIsRecordingDisabled] = useState(false);
    const [duration, setDuration] = useState('');
    const [durationTimer, setDurationTimer] = useState(0);
    const [recorder] = useState(new AudioRecorder());

    props.wsConnection.onMessage(() => {
        if (status === 'Calling') {
            setStatus('Connected');
        }
    });

    const onRecordingToggle = (isRecording: boolean) => {
        if (isRecording) {
            startVoiceRecording();

            setDuration(formatTimeDuration(0));
            const timerId = setInterval(() => {
                setDuration(formatTimeDuration(recorder.getDuration()));
            }, 1000);

            setDurationTimer(timerId);
        } else {
            stopVoiceRecoring();
            clearInterval(durationTimer);
            setDuration('');
        }
    }

    const startVoiceRecording = () => {
        recorder.start();
    }

    const stopVoiceRecoring = () => {
        recorder.stop().then((result: RecordingResult): void => {
            setIsRecordingDisabled(true);

            console.log('stopVoiceRecoring', typeof result.audio, result.audio.type, result.audio.size);
            props.wsConnection.send(result.audio);
            setTimeout(() => setIsRecordingDisabled(false), 1000);
        })
    }

    return (
        <GradientBorderCard
            containerClassName="w-full h-auto md:w-1/3 md:mx-auto"
            className="w-full h-full rounded-[22px] p-2 bg-zinc-900 min-h-[65vh] md:min-h-auto"
            animate={false}
        >
            <BackgroundGradientAnimation
                containerClassName="w-full h-full rounded-[16px]"
                className="w-full h-full"
                interactive={false}
            >
                <div className="relative min-h-[65vh] md:min-h-auto">
                    <div className="absolute w-full h-full flex flex-col p-4 z-10">
                        <div className="flex-grow h-full flex flex-col items-center justify-center text-white font-bold px-4 pointer-events-none text-lg text-center md:text-xl lg:text-2xl">
                            <div className="border-2 border-white rounded-full p-0.5 mb-4">
                                <img
                                    src="https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg"
                                    width={128}
                                    height={128}
                                    alt="Avatar"
                                    className="rounded-full"
                                />
                            </div>

                            <p className="bg-clip-text text-transparent drop-shadow-2xl bg-gradient-to-b from-white/80 to-white/20">
                                { status }
                            </p>
                        </div>
                        <div className="w-full flex flex-row justify-center z-10 gap-4 pb-4">
                            <div className="relative">
                                {
                                    duration && (
                                        <span className="absolute w-full text-center -top-8">{duration}</span>
                                    )
                                }
                                <RecordButton
                                    disabled={isRecordingDisabled}
                                    onRecordingToggle={onRecordingToggle}
                                ></RecordButton>
                            </div>
                        </div>
                    </div>
                </div>
            </BackgroundGradientAnimation>
        </GradientBorderCard>
    )
}
