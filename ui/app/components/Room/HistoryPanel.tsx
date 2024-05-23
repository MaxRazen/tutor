import React from 'react';
import { GradientBorderCard } from '../GradientBorderCard';

export default function HistoryPanel() {
    return (
        <GradientBorderCard
            containerClassName="w-full md:w-2/3 h-auto"
            className="h-full rounded-[22px] p-2 bg-zinc-900"
            animate={false}
        >
            <div className='h-full min-h-96'>test</div>
        </GradientBorderCard>
    )
}
