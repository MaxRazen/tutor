import React from 'react';
import { GradientBorderCard } from '../GradientBorderCard';
import { BackgroundGradientAnimation } from '../BackgroundGradientAnimation';
import { RecordButton } from '../RecordButton';

export default function CallerPanel({historyShown}) {
    return (
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
                    <div className="flex flex-col" style={{minHeight: '70vh'}}>
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
                            <div className="w-1/3 relative">
                                <span className="absolute w-full text-center -top-8">0:23</span>
                                <RecordButton></RecordButton>
                            </div>                                    
                        </div>
                    </div>
                </BackgroundGradientAnimation>
            </div>
        </GradientBorderCard>
    )
}
