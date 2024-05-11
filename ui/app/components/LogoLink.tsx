import React from 'react';
import Logo from './Logo';

export default function LogoLink() {
    return (
        <a
            href="/"
            className="flex items-center text-2xl font-semibold text-white"
            >
            <Logo className={'w-8 h-8 mr-2'}/>
            AI Tutor
            <sup className="text-gray-300 pl-2 font-normal text-xs">by Hexnet</sup>
        </a>
    )
    
}
