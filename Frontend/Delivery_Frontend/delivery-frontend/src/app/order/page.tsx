'use client';

import { useState, useEffect, useRef } from 'react';
import { useGetOrder } from '@/hooks/useGraphql';
import useUserStore from '../../../store/store';
import { orderType } from '@/lib/getOrder';
import PendingOrder from '@/components/pendingOrder';
import YourDeliveryOrder from '@/components/yourDeliveryOrder';
import useGPSWebSocket from '@/hooks/useGPSWebsocket';
import Image from 'next/image';

const Order = () => {
    const [pullDownDistance, setPullDownDistance] = useState(0);
    const startY = useRef(0);
    const [localData, setLocalData] = useState<orderType>();
    const { Id, CurrentPage } = useUserStore() as { Id: string, CurrentPage: string };
    const url = process.env.WS_WEBSOCKET_URL || 'wss://xqtvgb1k-8081.asse.devtunnels.ms';
    const uri = `/ws/delivery?staffId=${Id}`;
    const endpoint = url + uri;
    const { startTrackingLocation } = useGPSWebSocket(endpoint);

    const orderParams = CurrentPage == 'Pending' ? { status: "pending" } : { status: "delivering", Id: Id };
    const { data, loading, error, refetch } = useGetOrder(orderParams);
    
    useEffect(() => {
        startTrackingLocation();
    }, [startTrackingLocation]);

    useEffect(() => {
        if (data) {
            setLocalData(data);
        }
    }, [data]);

    // Pull down to refresh
    useEffect(() => {
        const handleTouchStart = (event: TouchEvent) => {
            startY.current = event.touches[0].clientY;
        }

        const handleTouchMove = (event: TouchEvent) => {
            const currentY = event.touches[0].clientY; // Get current Y position of the touch
            const diff = currentY - startY.current; // Calculate the difference
            if (diff > 0) {
                setPullDownDistance(diff);
            }
        };

        const handleTouchEnd = (event: TouchEvent) => {
            const endY = event.changedTouches[0].clientY; // Get end Y position of the touch
            // Check if the touch was a pull down at the top of the page
            if (endY - startY.current > 50) {
                console.log('Refresh!');
                refetch();
            }
            setPullDownDistance(0); // Reset the pull down distance
        };

        document.addEventListener('touchstart', handleTouchStart, false);
        document.addEventListener('touchmove', handleTouchMove, false);
        document.addEventListener('touchend', handleTouchEnd, false);

        return () => {
            document.removeEventListener('touchstart', handleTouchStart);
            document.removeEventListener('touchmove', handleTouchMove);
            document.removeEventListener('touchend', handleTouchEnd);
        }
    }, [refetch]);

    return (
        <div>
            {pullDownDistance > 0 && (
                <div style={{
                    transform: `translateY(${Math.min(pullDownDistance, 40)}px)`, // Limit to 100px, adjust as needed
                }} className='w-full h-full flex justify-center'>
                    {/* Refresh wheel or indicator component */}
                    <Image 
                        src="/loading.gif"
                        alt="refresh"
                        width={50}
                        height={50}
                    />
                </div>
            )}
            <div style={{ marginTop: `${Math.min(pullDownDistance, 50)}px` }}>
                {loading && <div>Loading...</div>}
                {error && <div>Error...</div>}
                {localData && (CurrentPage == "Pending") && localData.allOrders.map((order) => {
                    return <PendingOrder key={order.id} order={order}/>
                })}
                {localData && (CurrentPage == "Your Delivery") && localData.allOrders.map((order) => {
                    return <YourDeliveryOrder key={order.id} order={order}/>
                })}
            </div>
            
        </div>
    )
}
export default Order;