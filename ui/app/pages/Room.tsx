import React, {useState} from 'react';
import {useParams} from 'react-router-dom';
import Navigation from '../components/Navigation';
import AudioRecorder from '../recorder';
import { GradientBorderCard } from '../components/GradientBorderCard';
import { BackgroundGradientAnimation } from '../components/BackgroundGradientAnimation';
import { RecordButton } from '../components/RecordButton';
import ChatIcon from '../components/Icons/ChatIcon';

type ReceivePayload = {
    type: 'audio' | 'translation' | 'feedback'
    content: string
}

export default function Room () {
    const {roomId} = useParams();
    const [historyShown, setHistoryShown] = useState(false);

    /*
    const socket = new WebSocket("ws://localhost:3000/ws/room/1001");
    console.log('ws connection initializing...');

    socket.onopen = (e) => {
        console.log('ws connection is established');
    }
    
    socket.onmessage = (event) => {
        try {
            const data: ReceivePayload = JSON.parse(event.data);
        } catch (e) {
            const msg = 'Received unexpected data type in socket connection';
            console.warn(msg, typeof event.data);
            alert(msg);
            return;
        }

        const payload = JSON.parse(event.data);

        console.log('received message', payload);
    }

    socket.onerror = (e) => {
        console.warn(e);
    }

    const onClickHandler = () => {
        console.log('onClickHandler');

        socket.send('some-test-str');
    };

    const startRecordingHandler = async () => {
        recorder.start()
        console.log('recording started');
    }
    */

    const stopRecordingHandler = async () => {
        recorder.stop().then((blob: Blob): void => {
            console.log('recording finished', typeof blob);
            const el: HTMLAudioElement|null = document.querySelector('#audio');
            if (el) {
                el.src = URL.createObjectURL(blob);
            }
            // socket.send(blob);
        })
    }

    const recorder = new AudioRecorder();

    return (
        <main>
            <Navigation/>

            <section className="container py-8 flex flex-col md:flex-row gap-8 md:gap-16 px-8 md:px-0">
                <GradientBorderCard
                    containerClassName={`w-full md:w-1/3 md:mx-auto h-full ${!historyShown && 'md:mx-auto'}`}
                    className="h-full rounded-[22px] p-2 bg-zinc-900"
                    animate={false}
                >
                    <div>
                        <BackgroundGradientAnimation
                            containerClassName="w-full h-full rounded-[16px] p-4"
                            interactive={false}
                        >
                            <div className="flex flex-col" style={{minHeight: '75vh'}}>
                                <div className="flex-grow h-full flex flex-col items-center justify-center text-white font-bold px-4 pointer-events-none text-lg text-center md:text-xl lg:text-2xl">
                                    <div className="border-2 border-white rounded-full p-0.5 mb-4">
                                        <img
                                            src="https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-jane.jpg"
                                            width={128}
                                            height={128}
                                            alt="Avatar"
                                            className="rounded-full"
                                        />
                                    </div>

                                    <p className="bg-clip-text text-transparent drop-shadow-2xl bg-gradient-to-b from-white/80 to-white/20">
                                        Calling
                                    </p>
                                </div>
                                <div className="relative w-full flex flex-row z-10 gap-4 pb-4">
                                    <div className="w-1/3"></div>
                                    <div className="w-1/3">
                                        <RecordButton></RecordButton>
                                    </div>
                                    {/* <button
                                        onClick={() => setHistoryShown(!historyShown)}
                                    >Toggle Context Window</button> */}
                                    <div className="w-1/3 flex items-center justify-center">
                                        <button
                                            type="button"
                                            className="relative inline-flex items-center justify-center p-1 me-2 overflow-hidden text-sm rounded-full group bg-gradient-to-br from-purple-600 to-blue-500 group-hover:from-purple-600 group-hover:to-blue-500 hover:text-white text-white focus:ring-4 focus:outline-none focus:ring-blue-800"
                                            onClick={() => setHistoryShown(!historyShown)}
                                        >
                                            <span className="relative p-3 transition-all ease-in duration-75 bg-gray-900 rounded-full group-hover:bg-opacity-0">
                                                <ChatIcon className="w-6 h-6"/>
                                            </span>
                                        </button>
                                    </div>
                                    
                                </div>
                            </div>
                        </BackgroundGradientAnimation>
                    </div>
                </GradientBorderCard>

                {
                    historyShown && (
                        <GradientBorderCard
                            containerClassName="w-full md:w-2/3 h-auto"
                            className="h-full rounded-[22px] p-2 bg-zinc-900"
                            animate={false}
                        >
                            <div className='h-full'>test</div>
                        </GradientBorderCard>
                    )
                    
                }
            </section>
        </main>
    )
}
