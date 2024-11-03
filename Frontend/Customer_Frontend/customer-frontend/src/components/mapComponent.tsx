'use client';

import React, { useEffect, useState, Suspense } from 'react';
import { MapContainer, TileLayer, Marker } from 'react-leaflet';
import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import useGPSWebSocket from '@/hooks/useGPSWebsocket';
import { useSearchParams } from 'next/navigation';

type latestMessageType = {
    staffId: number;
    latitude: number;
    longitude: number;
}

const MapComponent = () => {
    const searchParams = useSearchParams();
    const deliveryStaffId = searchParams.get('deliveryStaffID');
    const url = process.env.WS_WEBSOCKET_URL || 'wss://xqtvgb1k-8081.asse.devtunnels.ms';
    const uri = `${url}/ws/customer?deliveryStaffID=${deliveryStaffId}`;
    const { latestMessage } = useGPSWebSocket(uri);
    const [position, setPosition] = useState<[number, number]>([3.0, 101.5]);

    useEffect(() => {
        if (latestMessage) {
            // console.log(latestMessage);
            const latestMessageJson: latestMessageType = JSON.parse(latestMessage);
            setPosition([latestMessageJson.latitude, latestMessageJson.longitude]);
        }
    }, [latestMessage, setPosition]);

    return (
        <Suspense fallback={<div>Loading...</div>}>
            <div className='h-full w-full'>
                <MapContainer center={position} zoom={17} style={{ height: '100%', width: '100%' }}>
                    <TileLayer
                        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                    />
                    <Marker position={position} icon={L.icon({ iconUrl: './images/delivery_guy_on_motor.png', iconSize: [20, 20], iconAnchor: [10, 10] })}>
                    </Marker>
                </MapContainer>
            </div>
        </Suspense>
    )
}
export default MapComponent;