import { useEffect, useRef, useState } from 'react';



const useGPSWebSocket = (serverUrl: string) => {
    const webSocket = useRef<WebSocket | null>(null);

    const [latestMessage, setLatestMessage] = useState<string>();

    useEffect(() => {
        // Initialize WebSocket connection
        webSocket.current = new WebSocket(serverUrl);

        webSocket.current.onopen = () => console.log('WebSocket connected');
        webSocket.current.onerror = (error) => console.error('WebSocket error:', error);
        webSocket.current.onclose = () => console.log('WebSocket disconnected');

        // Listen for messages
        webSocket.current.onmessage = (event) => {
            // const msg = JSON.parse(event.data);
            setLatestMessage(event.data);
        };

        return () => {
            if (webSocket.current) {
                webSocket.current.close();
            }
        };
    }, [serverUrl]);

    
    function sentErrorMessage() {
        if (webSocket.current && webSocket.current.readyState === WebSocket.OPEN) {
            if ('geolocation' in navigator) {
                navigator.geolocation.getCurrentPosition((position) => {
                    webSocket.current?.send(position.coords.latitude + ' ' + position.coords.longitude);
                }, (error) => {
                    webSocket.current?.send(error.message);
                    console.error('Error getting current position:', error);
                });
                webSocket.current?.send('Geolocation is supported by this browser.');
            } else {
                webSocket.current?.send('Geolocation is not supported by this browser.');
            }

        }
    }
    return { latestMessage };
};

export default useGPSWebSocket;