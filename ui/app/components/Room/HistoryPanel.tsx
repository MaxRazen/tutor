import React, { useState, useEffect } from 'react';
import { GradientBorderCard } from '../GradientBorderCard';
import { Message, MessageContext } from './History/Message';
import AudioTrack from './History/AudioTrack';
import WSConnection from '../../ws';

type HistoryPanelProps = {
    wsConnection: WSConnection
}

export default function HistoryPanel(props: HistoryPanelProps) {
    const [messages, setMessages] = useState<MessageContext[]>([
        {
            avatar: 'https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg',
            name: 'Scarlett',
            align: 'right',
            message: <AudioTrack source='/assets/FirstSnow-Emancipator.mp3'/>,
        },
        {
            avatar: 'https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg',
            name: 'John',
            align: 'left',
            message: <AudioTrack source='/assets/Anthem-Emancipator.mp3'/>,
        },
        {
            avatar: 'https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg',
            name: 'Scarlett',
            align: 'right',
            message: 'In this example, the ChatRoom component uses an Effect to stay connected to an external system defined in chat.js. '
                +'Press “Open chat” to make the ChatRoom component appear. '
                +'This sandbox runs in development mode, so there is an extra connect-and-disconnect cycle, as explained here.'
                +'Try changing the roomId and serverUrl using the dropdown and the input, and see how the Effect re-connects to the chat.'
                +'Press “Close chat” to see the Effect disconnect one last time.',
        },
        {
            avatar: 'https://res.cloudinary.com/dzgusx2vf/image/upload/v1716310225/tutor/avatar-scarlett.jpg',
            name: 'Scarlett',
            align: 'right',
            message: 'In this example, the ChatRoom component uses an Effect to stay connected to an external system defined in chat.js. '
                +'Press “Open chat” to make the ChatRoom component appear. '
                +'This sandbox runs in development mode, so there is an extra connect-and-disconnect cycle, as explained here.'
                +'Try changing the roomId and serverUrl using the dropdown and the input, and see how the Effect re-connects to the chat.'
                +'Press “Close chat” to see the Effect disconnect one last time.',
        },
    ]);

    props.wsConnection.onMessage(() => {
        // TODO: load messages
    })

    return (
        <GradientBorderCard
            containerClassName="w-full h-full"
            className="w-full h-full flex flex-col rounded-[22px] px-4 md:px-8 bg-gray-900 overflow-hidden"
            animate={false}
        >
            <div className="flex-grow overflow-y-auto no-scrollbar py-4 md:py-8 max-h-[65vh] md:max-h-max">
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
