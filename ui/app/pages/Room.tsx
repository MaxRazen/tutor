import React from 'react';
import Navigation from '../components/Navigation';
import AudioRecorder from '../recorder';

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

    const stopRecordingHandler = async () => {
        recorder.stop().then((blob: Blob): void => {
            console.log('recording finished', typeof blob);
            const el: HTMLAudioElement|null = document.querySelector('#audio');
            if (el) {
                el.src = URL.createObjectURL(blob);
            }
            socket.send(blob);
        })
    }

    const recorder = new AudioRecorder();

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

                    <hr />

                    <button
                        className="btn-link"
                        onClick={startRecordingHandler}
                    >Start recording</button>
                    <button
                        className="btn-link"
                        onClick={stopRecordingHandler}
                    >Stop recording</button>

                    <hr />

                    <audio
                        id="audio"
                        src=""
                        controls
                    ></audio>
                </div>


            </section>
        </main>
    )
}
