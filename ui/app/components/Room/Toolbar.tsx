import React from 'react';

interface Handlers {
    leaveRoom(): void
    openContextSettings(): void
    toggleHistoryPanel(): void
}

export default function Toolbar(props) {
    const handlers: Handlers = props.handlers;
    const {states} = props;
    
    return (
        <div className="w-full h-16 bg-gray-700 border-gray-600 border rounded-full">
            <div className="flex flex-row h-full mx-auto">
                <button
                    type="button"
                    className="inline-flex flex-col items-center justify-center px-5 rounded-s-full hover:bg-gray-800 group md:min-w-32"
                    title="Leave room"
                    onClick={handlers.leaveRoom}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={1.5}
                        stroke="currentColor"
                        className="w-5 h-5 mb-1 text-gray-400 group-hover:text-blue-500"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            d="M9 15L3 9m0 0l6-6M3 9h12a6 6 0 010 12h-3"
                        />
                    </svg>
                    <span className="sr-only">Leave Room</span>
                </button>

                <div className="w-full"></div>

                <button
                    type="button"
                    className="inline-flex flex-col items-center justify-center px-5 hover:bg-gray-800 group md:min-w-32"
                    title="Context Settings"
                    onClick={handlers.openContextSettings}
                >
                    <svg
                        className="w-5 h-5 mb-1 text-gray-400 group-hover:text-blue-500"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 20 20"
                    >
                        <path
                            stroke="currentColor"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M4 12.25V1m0 11.25a2.25 2.25 0 0 0 0 4.5m0-4.5a2.25 2.25 0 0 1 0 4.5M4 19v-2.25m6-13.5V1m0 2.25a2.25 2.25 0 0 0 0 4.5m0-4.5a2.25 2.25 0 0 1 0 4.5M10 19V7.75m6 4.5V1m0 11.25a2.25 2.25 0 1 0 0 4.5 2.25 2.25 0 0 0 0-4.5ZM16 19v-2"
                        />
                    </svg>
                    <span className="sr-only">Context Settings</span>
                </button>

                <button
                    type="button"
                    className={`inline-flex flex-col items-center justify-center px-5 rounded-e-full group md:min-w-32 ${
                        states.historyShown ? 'bg-gray-800' : 'hover:bg-gray-800'
                    }`}
                    title="History Panel"
                    onClick={handlers.toggleHistoryPanel}
                >
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth={2}
                        stroke="currentColor"
                        className="w-5 h-5 mb-1 text-gray-400 group-hover:text-blue-500"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            d="M20.25 8.511c.884.284 1.5 1.128 1.5 2.097v4.286c0 1.136-.847 2.1-1.98 2.193-.34.027-.68.052-1.02.072v3.091l-3-3a49.5 49.5 0 01-4.02-.163 2.115 2.115 0 01-.825-.242m9.345-8.334a2.126 2.126 0 00-.476-.095 48.64 48.64 0 00-8.048 0c-1.131.094-1.976 1.057-1.976 2.192v4.286c0 .837.46 1.58 1.155 1.951m9.345-8.334V6.637c0-1.621-1.152-3.026-2.76-3.235A48.455 48.455 0 0011.25 3c-2.115 0-4.198.137-6.24.402-1.608.209-2.76 1.614-2.76 3.235v6.226c0 1.621 1.152 3.026 2.76 3.235.577.075 1.157.14 1.74.194V21l4.155-4.155"
                        />
                    </svg>
                    <span className="sr-only">History Panel</span>
                </button>
            </div>
        </div>
    )
}
