'use client';

import React, { useEffect, useState } from 'react';
import ThemeToggleButton from './themeButton';
import { useNavBarStore, useUserStore } from '../../store/store';
import { useAuth } from '@/hooks/useAuth';
import Link from 'next/link';
import Image from 'next/image';

const Navbar: React.FC = () => {
    const [authenticated, setAuthenticated] = useState(false);
    const { isNavBarOpen } = useNavBarStore() as { isNavBarOpen: boolean };
    const { username } = useUserStore() as { username: string };
    const { isAuthenticated } = useAuth();

    useEffect(() => {
        setAuthenticated(isAuthenticated);
    }, [isAuthenticated]);
    // useEffect(() => {
    //     const verifyUser = async () => {
    //         try {
    //             const data = await verifyCookie();
    //             console.log('Cookie verified:', data);
    //         }
    //         catch (error) {
    //             console.log('Failed to send cookie to server: ', error);
    //         }
    //     }
    //     verifyUser();
    // }, [verifyCookie]);

    if (!isNavBarOpen) {
        return null;
    }

    return (
        <nav>
            <div className='flex flex-row justify-between h-[5rem] w-full px-[2rem] items-center secondary-background shadow-2xl fixed top-0'>
                <div>
                    <Link href={'./'}><h1 className='text-3xl font-bold active:font-extrabold'>Pizza</h1></Link>
                </div>
                <div className='flex flex-row space-x-[3rem] items-center w-fit text-lg font-semibold'>
                    <div>
                        <Link href={'./menu'}><div className=' active:font-extrabold'>Menu</div></Link>
                    </div>
                    <div>
                        {authenticated? (
                            // <Link href={'./profile'}>
                                <div className='tooltip'>
                                    <Image
                                        src={`https://api.dicebear.com/9.x/adventurer/png?seed=${username}`}
                                        alt="Profile"
                                        width={50}
                                        height={50}
                                    />
                                    <span className='tooltiptext'>{username}</span>
                                </div>
                            // </Link>
                        ) : (
                            <Link href={'./login'}><div className='active:font-extrabold'>Login</div></Link>
                        )}
                    </div>
                    <div className='text-3xl'>
                        <ThemeToggleButton />
                    </div>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;