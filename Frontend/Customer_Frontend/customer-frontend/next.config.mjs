/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        API_GATEWAY: process.env.API_GATEWAY,
        WS_WEBSOCKET_URL: process.env.WS_WEBSOCKET_URL,
    },
    experimental: {
        missingSuspenseWithCSRBailout: false,
    },
    reactStrictMode: true,
    images: {
        unoptimized: true,
        remotePatterns: [
            {
                protocol: 'https',
                hostname: 'api.dicebear.com',
                port: '',
                pathname: '**',
            }
        ]
    }
};

export default nextConfig;
