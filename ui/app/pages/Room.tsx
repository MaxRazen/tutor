import React from 'react';
import Navigation from '../components/Navigation';

type ReceivePayload = {
    type: 'audio' | 'translation' | 'feedback'
    content: string
}

export default function Room () {
    const socket = new WebSocket("ws://localhost:3000/ws/room/1001");
    console.log('ws connection initializing...');

    socket.onopen = (e) => {
        console.log('ws connection is established');
    }
    
    socket.onmessage = (event) => {
        if (typeof event.data !== 'string') {
            console.warn('Received unexpected data type in socket connection', typeof event.data);
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

    return (
        <main>
            <Navigation/>

            <section className="container py-8">
                <h1 className="text-lg text-white">Room #1001</h1>

                <div className="w-1/4">
                    <button
                        className="btn-link"
                        onClick={onClickHandler}
                    >Init Room</button>

                    <audio id="audio">
                    </audio>
                </div>
            </section>
        </main>
    )
}
