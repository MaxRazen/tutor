import React, {useState} from 'react';
import MicIcon from './Icons/MicIcon';

export const RecordButton = () => {
    const [isRecording, setIsRecording] = useState(false);

    const toggleRecording = () => {
      setIsRecording(!isRecording);
    };

    return (
        <div className="flex items-center justify-center">
            <button
                onClick={toggleRecording}
                className={`relative w-16 h-16 rounded-full flex items-center justify-center focus:outline-none transition-transform duration-300 ${
                    isRecording ? 'bg-gradient-to-l from-red-500 to-orange-500' : 'bg-gradient-to-r from-cyan-500 to-blue-500'
                }`}
            >
                <MicIcon className={`w-8 h-8 transition-transform duration-300 text-white ${
                    isRecording ? 'animate-pulse' : ''
                }`}></MicIcon>

                <div
                    className={`absolute w-16 h-16 rounded-full bg-red-500 opacity-50 animate-ping ${
                        isRecording ? 'block' : 'hidden'
                    }`}
                ></div>
            </button>
        </div>
    )
}
 