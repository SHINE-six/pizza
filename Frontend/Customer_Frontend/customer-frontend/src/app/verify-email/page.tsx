'use client';

import { useCallback, useEffect, useState, Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { useNavBarStore } from "../../../store/store";
import Image from "next/image";
import { useRouter } from "next/navigation";
import Link from "next/link";

const VerifyEmail = () => {
    const [returnMessage, setReturnMessage] = useState('');
    const router = useRouter();
    const { setIsNavBarOpen } = useNavBarStore() as { setIsNavBarOpen: (value: boolean) => void };
    // Get email and verification token from URL
    const searchParams = useSearchParams();
    const email = searchParams.get('email');
    const token = searchParams.get('token');

    const sentToBackend = useCallback(async () => {
        const url = process.env.API_GATEWAY || 'https://xqtvgb1k-8080.asse.devtunnels.ms';
        const uri = '/rest/v1/user/verify_email';
        const endpoint = url + uri;
        try {
            const url = new URL(endpoint);

            // Add email and token to query
            url.searchParams.append('email', email ?? '');
            url.searchParams.append('token', token ?? '');

            // API call here - GET request to /rest/v1/user/verify_email, with email and password in query
            const response = await fetch(url.toString(), {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                },
            });
            const data = await response.json();
            
            if (response.ok) {
                console.log('Login successful', data);
            } else {
                // Check if the data object has an error property
                if (data.error) {
                    // Extract the description from the error message
                    const descMatch = data.error.match(/desc = (.+)$/);
                    if (descMatch) {
                        console.log('Error description:', descMatch[1]);
                        // Optionally, you can use the description for user feedback
                        // For example, setting an error message state to display in the UI
                        setReturnMessage(descMatch[1]);
                    }
                }
                console.log('Sign up failed', data);
                return;
            }
        } catch (error) {
            setReturnMessage('An error occurred' + error);
            console.error('Error:', error);
        }
        
        // Redirect to menu
        router.push('/menu');
    }, [email, token, router]);

    useEffect(() => {
        setIsNavBarOpen(false);
        sentToBackend();

        return () => {
            setIsNavBarOpen(true);
        }
    }, [sentToBackend, setIsNavBarOpen]);

    return (
        <Suspense fallback={<div>Loading...</div>}>
            <div className="flex flex-col justify-start items-center font-bold text-center h-screen pt-[10rem]">
                {
                    returnMessage === '' ? (
                        <>
                            <Image
                                src="/images/tick.png"
                                alt="Email Sent"
                                width={200}
                                height={200}
                            />
                            <div className="pt-[2rem]">Email Verified Successful. Click Log In and start 🍕</div>
                            <Link href="/login"><div className="bg-secondary w-fit rounded-lg mt-[2rem] justify-self-center p-[0.35rem]">Log In</div></Link>
                        </>
                    ) : (
                        <>
                            <Image
                                src="/images/cross.png"
                                alt="Email Sent"
                                width={200}
                                height={200}
                            />
                            <div className="pt-[2rem]"> {returnMessage} Please try Sign Up again</div>
                            <Link href="/login"><div className="bg-secondary w-fit rounded-lg mt-[2rem] justify-self-center p-[0.35rem]">Sign Up</div></Link>
                        </>
                    )
                }
            </div>
        </Suspense>
    )
}
export default VerifyEmail;