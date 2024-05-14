import React from 'react';
import { Link } from 'react-router-dom';
import Navigation from '../components/Navigation';

export default function Home () {
    return (
        <main>
            <Navigation/>

            <section className="container py-8">
                <div className="grid gap-4 grid-cols-3">
                    <Link
                        to={'/room/1001'}
                        className="block max-w-sm p-6 bg-gray-800 border border-gray-700 rounded-lg shadow hover:bg-gray-700"
                    >
                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-white">Call Mode</h5>
                        <p className="font-normal text-gray-400">Try a new call mode with AI assistant as tutor mode</p>
                    </Link>
                    <a
                        href="#"
                        className="block max-w-sm p-6 bg-gray-800 border border-gray-700 rounded-lg shadow hover:bg-gray-700"
                    >
                        <h5 className="mb-2 text-2xl font-bold tracking-tight text-white">Assistant</h5>
                        <p className="font-normal text-gray-400">Ask latest GPT any question in chat mode</p>
                    </a>
                </div>
            </section>
        </main>
    )
}
