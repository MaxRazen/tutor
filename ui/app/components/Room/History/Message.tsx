import React, {useState} from 'react';
import { EyeIcon, EyeSlashIcon, LanguageIcon } from '@heroicons/react/24/solid'


export type MessageContext = {
    avatar: string
    name: string
    time: string
    message: string | React.JSX.Element
    transcription: string
}

export type MessageProps = {
    context: MessageContext
}

export function Message(props: MessageProps) {
    const [isTranscriptionView, setIsTranscriptionView] = useState(false);
    
    const {context: msg} = props;

    return (
        <div className="flex items-start gap-2">
            <img
                className="w-8 h-8 rounded-full"
                src={msg.avatar}
                alt={msg.name}
            />
            <div className="flex flex-col w-full max-w-full md:max-w-[320px] leading-1.5">
                <div className="flex items-center space-x-2 rtl:space-x-reverse">
                    <span className="text-sm font-semibold :text-white">
                        {msg.name}
                    </span>
                    <span className="text-sm font-normal text-gray-400">
                        {msg.time}
                    </span>
                </div>
                <div className="text-sm font-normal py-2 text-white">
                    {
                        isTranscriptionView
                        ? msg.transcription
                        : msg.message
                    }
                </div>
                <div className="flex flex-row items-center gap-2">
                    <button
                        type="button"
                        className="inline-flex self-center items-center p-1 text-center text-gray-400 bg-gray-900 rounded-lg hover:bg-gray-800 focus:ring-4 focus:outline-none focus:ring-gray-600"
                        onClick={() => setIsTranscriptionView(!isTranscriptionView)}
                    >
                        {
                            isTranscriptionView
                            ? <EyeSlashIcon className="size-4"/>
                            : <EyeIcon className="size-4"/>
                        }
                    </button>
                    <button
                        type="button"
                        className="inline-flex self-center items-center p-1 text-center text-gray-400 bg-gray-900 rounded-lg hover:bg-gray-800 focus:ring-4 focus:outline-none focus:ring-gray-600"
                    >
                        <LanguageIcon className="size-4"/>
                    </button>
                </div>
            </div>
        </div>
    )
}
