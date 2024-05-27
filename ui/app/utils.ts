import { ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
 
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

export function formatTime(dt: Date): string {
    const h = dt.getHours();
    const m = dt.getMinutes();
    return `${h}:${m}`;
}

export function formatTimeDuration(durationMs: number): string {
    const mins = Math.floor(durationMs / 1000 / 60);
    const minToken = mins > 0 ? mins.toString() : '0';
    const seconds = Math.ceil(60 * (durationMs / 1000 / 60 - mins));
    const secToken = (seconds < 10 ? '0' : '') + seconds.toFixed();

    return `${minToken}:${secToken}`;
}
