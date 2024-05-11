import React from 'react';

export default function CardCentered(props) {
    return (
        <div className="w-full rounded-lg bg-gray-800 shadow border border-gray-700 md:mt-0 sm:max-w-md xl:p-0">
            <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
                {props.children}
            </div>
        </div>
    )
}
