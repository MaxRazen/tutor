import { ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
 
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

export function formatTimeDuration(durationMs: number): string {
    const mins = +(durationMs / 1000 / 60).toFixed();
    const minToken = mins > 0 ? mins.toString() : '0';
    const seconds = Math.round(60 * (durationMs / 1000 / 60 - mins));
    const secToken = (seconds < 10 ? '0' : '') + seconds.toFixed();

    return `${minToken}:${secToken}`;
}
