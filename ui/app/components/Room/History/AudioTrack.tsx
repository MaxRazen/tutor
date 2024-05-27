import React, { useState, useEffect, useRef, Ref } from 'react';
import { PlayIcon, PauseIcon } from '@heroicons/react/24/solid';
import { formatTimeDuration } from '../../../utils';

export default function AudioTrack({ source }: { source: string }) {
    const audioEl: Ref<HTMLAudioElement> = React.createRef();
    const [playing, setPlaying] = useState(false);
    const [duration, setDuration] = useState('0:00');

    const togglePlaying = () => {
        const newState = !playing;
        setPlaying(newState);

        if (!audioEl?.current) {
            return
        }

        if (newState) {
            audioEl.current.play()
        } else {
            audioEl.current.pause();
        }
    }

    useEffect(() => {
        if (!audioEl.current) {
            return
        }
        const el: HTMLAudioElement = audioEl.current;

        el.addEventListener('durationchange', () => {
            const duration = el.duration;

            if (Number.isFinite(duration)) {
                setDuration(formatTimeDuration(duration * 1000));
            }
        })
        el.addEventListener('ended', () => {
            setPlaying(false);
        })
    })

    return (
        <div className="flex items-center space-x-2 rtl:space-x-reverse">
            <audio className="hidden" ref={audioEl} controls>
                <source src={source} type="audio/ogg"></source>
            </audio>

            <button
                type="button"
                className="inline-flex self-center items-center p-2 text-sm font-medium text-center text-white rounded-lg hover:bg-gray-600 focus:ring-4 focus:outline-none focus:ring-gray-600"
                onClick={togglePlaying}
            >
                {
                    playing
                    ? <PauseIcon className="size-4"/>
                    : <PlayIcon className="size-4"/>
                }
            </button>
            <svg
                className="w-[145px] md:w-[185px] md:h-[40px]"
                aria-hidden="true"
                viewBox="0 0 185 40"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <rect
                    y={17}
                    width={3}
                    height={6}
                    rx="1.5"
                    fill="#6B7280"
                    className="dark:fill-white"
                />
                <rect
                    x={7}
                    y="15.5"
                    width={3}
                    height={9}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={21}
                    y="6.5"
                    width={3}
                    height={27}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={14}
                    y="6.5"
                    width={3}
                    height={27}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={28}
                    y={3}
                    width={3}
                    height={34}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={35}
                    y={3}
                    width={3}
                    height={34}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={42}
                    y="5.5"
                    width={3}
                    height={29}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={49}
                    y={10}
                    width={3}
                    height={20}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={56}
                    y="13.5"
                    width={3}
                    height={13}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={63}
                    y={16}
                    width={3}
                    height={8}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={70}
                    y="12.5"
                    width={3}
                    height={15}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={77}
                    y={3}
                    width={3}
                    height={34}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={84}
                    y={3}
                    width={3}
                    height={34}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={91}
                    y="0.5"
                    width={3}
                    height={39}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={98}
                    y="0.5"
                    width={3}
                    height={39}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={105}
                    y={2}
                    width={3}
                    height={36}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={112}
                    y="6.5"
                    width={3}
                    height={27}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={119}
                    y={9}
                    width={3}
                    height={22}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={126}
                    y="11.5"
                    width={3}
                    height={17}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={133}
                    y={2}
                    width={3}
                    height={36}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={140}
                    y={2}
                    width={3}
                    height={36}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={147}
                    y={7}
                    width={3}
                    height={26}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={154}
                    y={9}
                    width={3}
                    height={22}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={161}
                    y={9}
                    width={3}
                    height={22}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={168}
                    y="13.5"
                    width={3}
                    height={13}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={175}
                    y={16}
                    width={3}
                    height={8}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect
                    x={182}
                    y="17.5"
                    width={3}
                    height={5}
                    rx="1.5"
                    fill="#E5E7EB"
                    className="dark:fill-gray-500"
                />
                <rect x={0} y={16} width={8} height={8} rx={4} fill="#1C64F2" />
            </svg>
            <span className="inline-flex self-center items-center p-2 text-sm font-medium text-white">
                {duration}
            </span>
        </div>

    )
}