/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        API_GATEWAY: process.env.API_GATEWAY,
        WS_WEBSOCKET_URL: process.env.WS_WEBSOCKET_URL,
    },
    images: {
        unoptimized: true,
    },
};

export default nextConfig;
