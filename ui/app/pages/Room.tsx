import React, {useState} from 'react';
import {useParams, useNavigate} from 'react-router-dom';
import Navigation from '../components/Navigation';
import AudioRecorder from '../recorder';
import Toolbar from '../components/Room/Toolbar';
import HistoryPanel from '../components/Room/HistoryPanel';
import CallerPanel from '../components/Room/CallerPanel';
import {WSConnection} from '../ws';

type ReceivePayload = {
    type: 'audio' | 'translation' | 'feedback'
    content: string
}

export default function Room () {
    const {roomId} = useParams();
    const navigate = useNavigate();

    // Toolbar
    const [historyShown, setHistoryShown] = useState(false);
    const toolbarHandlers= {
        leaveRoom() {            
            wsConn.disconnect();
            navigate('/');
        },
        openContextSettings() {
            alert('TODO');
        },
        toggleHistoryPanel() {
            setHistoryShown(!historyShown);
        },
    };
    const toolbarStates = {
        historyShown,
    }

    const socketEndpoint = `ws://localhost:3000/ws/room/${roomId}`;
    const wsConn: WSConnection = new WSConnection(socketEndpoint);
    wsConn.onConnect((e) => {
        console.log('Connected to WS | Room:', roomId);
    });
    //wsConn.connect();

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


    const recorder: AudioRecorder = new AudioRecorder();

    return (
        <main>
            <Navigation/>

            <section className="container py-4 px-4 md:px-0">
                <Toolbar
                    handlers={toolbarHandlers}
                    states={toolbarStates}
                ></Toolbar>
            </section>

            <section className="container py-8 flex flex-col md:flex-row gap-8 md:gap-16 px-8 md:px-0">
                <CallerPanel
                    historyShown={historyShown}
                ></CallerPanel>

                {
                    historyShown && (
                        <HistoryPanel></HistoryPanel>
                    )
                }
            </section>
        </main>
    )
}
