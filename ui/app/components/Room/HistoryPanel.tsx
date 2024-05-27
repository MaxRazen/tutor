import React, { useState, useEffect } from 'react';
import { GradientBorderCard } from '../GradientBorderCard';
import { Message, MessageContext } from './History/Message';
import AudioTrack from './History/AudioTrack';
import WSConnection from '../../ws';
import {TUTOR_NAME, TUTOR_AVATAR} from '../../Application';
import store from '../../store';
import {formatTime} from '../../utils';

type HistoryPanelProps = {
    wsConnection: WSConnection
}

type ResponseMessage = {
    type: "recording" | "error"
	ts: number
	authorship: "user" | "tutor" | "system"
 	content: string
    transcription: string
}

const messageFactory = (msg: ResponseMessage): MessageContext => {
    const user = store.getUser();
    let name: string = TUTOR_NAME;
    let avatar: string = TUTOR_AVATAR;
    const time = formatTime(new Date(msg.ts * 1000));

    if (msg.authorship === 'user' && user) {
        name = user.name;
        avatar = user.avatar;
    }

    let message: string | React.JSX.Element = '';
    switch (msg.type) {
        case 'recording':
            message = <AudioTrack source={msg.content}/>
            break;
        default:
            message = msg.content;
    }

    return {
        avatar,
        name,
        time,
        message,
        transcription: msg.transcription,
    }
}

export default function HistoryPanel(props: HistoryPanelProps) {
    const [messages, setMessages] = useState<MessageContext[]>([]);

    props.wsConnection.onMessage((e: MessageEvent) => {
        console.log('HistoryPanel :: onMessage');
        console.log(e.data, typeof e.data);
        const msg: ResponseMessage = JSON.parse(e.data);

        setMessages([
            ...messages,
            messageFactory(msg),
        ]);
    })

    return (
        <GradientBorderCard
            containerClassName="w-full h-full"
            className="w-full h-full flex flex-col rounded-[22px] px-2 md:px-6 bg-gray-900 overflow-hidden"
            animate={false}
        >
            <div className="flex-grow overflow-y-auto no-scrollbar py-2 md:py-6 max-h-[65vh] md:max-h-max">
                <div className="flex flex-col gap-4 md:gap-8">
                    {
                        messages.map((message, i) => (
                            <Message
                                key={i}
                                context={message}
                            ></Message>
                        ))
                    }
                </div>
            </div>
        </GradientBorderCard>
    )
}
