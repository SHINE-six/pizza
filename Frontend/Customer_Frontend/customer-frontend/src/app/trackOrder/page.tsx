'use client';

import dynamic from 'next/dynamic';
import React, { Suspense } from 'react';
// import MapComponent from '@/components/mapComponent';

const MapComponent = dynamic(() => import('@/components/mapComponent'), { ssr: false });

const TrackOrder = () => {

    return (
        <Suspense fallback={<div>Loading...</div>}>
            <div className='h-[40rem] fillPage flex justify-center items-center'>
                <MapComponent/>
            </div>
        </Suspense>
    )
}
export default TrackOrder;