import React, {useState} from 'react';
import {useParams, useNavigate} from 'react-router-dom';
import Navigation from '../components/Navigation';
import Toolbar from '../components/Room/Toolbar';
import HistoryPanel from '../components/Room/HistoryPanel';
import CallerPanel from '../components/Room/CallerPanel';
import WSConnection from '../ws';

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
                    wsConnection={wsConn}
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
