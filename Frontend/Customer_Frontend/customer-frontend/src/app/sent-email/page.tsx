'use client';

import { useSearchParams } from "next/navigation";
import { useEffect, Suspense } from "react";
import { useNavBarStore } from "../../../store/store";
import Image from 'next/image';

const SentEmail = () => {
    const { setIsNavBarOpen } = useNavBarStore() as { setIsNavBarOpen: (value: boolean) => void };
    const searchParams = useSearchParams();

    const email = searchParams.get('email');

    useEffect(() => {
        setIsNavBarOpen(false);

        // Cleanup
        return () => {
            setIsNavBarOpen(true);
        }
    }, [setIsNavBarOpen]);
    
    return (
        <Suspense fallback={<div>Loading...</div>}>
            <div className="flex flex-col justify-center items-center h-screen space-y-[2rem]">
                <Image
                    src="/images/email-sent.png"
                    alt="Email Sent"
                    width={200}
                    height={200}
                />
                <div className="text-3xl font-bold">Verify your email</div>
                <div className="text-center">We&#39;ve sent an email to : <p className="underline inline-block font-bold text-[#3975cb]">{email}</p> to vefify
                your email address and activate your account. <br/>The link in the email will expire in 3 minutes.</div>
            </div>
        </Suspense>
    )
}
export default SentEmail;