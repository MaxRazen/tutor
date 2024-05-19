import React from 'react';
import {useNavigate} from 'react-router-dom';
import Application from '../Application';
import Navigation from '../components/Navigation';

export default function Home () {
    Application.setTitle('Dashboard');
    const navigator = useNavigate();

    const onCallModeClick = async () => {
        const roomId = await Application.createRoom({
            mode: 'call',
        });
        if (roomId) {
            navigator(`/room/${roomId}`);
        }
    }
    const onChatModeClick = async () => {
        const roomId: Number|null = await Application.createRoom({
            mode: 'chat',
        });
        if (roomId) {
            navigator(`/room/${roomId}`);
        }
    }

    return (
        <main>
            <Navigation/>

            <section className="container py-8">
                <div className="grid gap-4 grid-cols-3">
                    <button
                        type="button"
                        className="block max-w-sm p-6 bg-gray-800 border border-gray-700 rounded-lg shadow hover:bg-gray-700"
                        onClick={onCallModeClick}
                    >
                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-white">Call Mode</h5>
                        <p className="font-normal text-gray-400">Try a call mode with AI assistant playing in a tutor role</p>
                    </button>
                    <button
                        type="button"
                        className="block max-w-sm p-6 bg-gray-800 border border-gray-700 rounded-lg shadow hover:bg-gray-700"
                        onClick={onChatModeClick}
                    >
                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-white">Assistant</h5>
                        <p className="font-normal text-gray-400">Ask latest GPT model any question in chat mode</p>
                    </button>
                </div>
            </section>
        </main>
    )
}
