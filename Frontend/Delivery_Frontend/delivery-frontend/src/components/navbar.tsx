'use client';

import { useState, useEffect } from 'react';
import { useSearchParams } from 'next/navigation';
import useUserStore from '../../store/store';

const Navbar = () => {
    const [currentlySelected, setCurrentlySelected] = useState('Pending');
    const { setName, setId, setCurrentPage } = useUserStore() as { setName: (value: string) => void, setId: (value: string) => void, setCurrentPage: (value: string) => void };
    const searchParams = useSearchParams();
    const id = searchParams.get('id');
    const name = searchParams.get('name');

    useEffect(() => {
        if (id) {
            setId(id);
        }
        if (name) {
            setName(name);
        }
        setCurrentPage(currentlySelected);
    }, [currentlySelected, id, name, setId, setName, setCurrentPage]);

    return (
        <div className='h-[6rem] bg-yellow-300 w-screen flex flex-col justify-between px-[1rem] py-[0.7rem]'>
            <div className='flex flex-row justify-between'>
                <div className='text-2xl font-semibold'>Order</div>
                <div>{name}</div>
            </div>
            <div className='flex flex-row justify-around'>
                <div className={`cursor-pointer ${currentlySelected == 'Pending' ? 'underline underline-offset-4 font-semibold' : ''}`} onClick={() => setCurrentlySelected('Pending')}>Pending</div>
                <div className={`cursor-pointer ${currentlySelected == 'Your Delivery' ? 'underline underline-offset-4 font-semibold' : ''}`} onClick={() => setCurrentlySelected('Your Delivery')}>Your Delivery</div>
            </div>
        </div>
    )
}
export default Navbar;