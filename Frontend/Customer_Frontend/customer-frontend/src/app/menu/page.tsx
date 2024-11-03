'use client';

import { useGetPizzaMenu } from "@/hooks/useGraphql";
import { useEffect, useState, useCallback } from "react";
import Image from "next/image";
import { usePizzaMenuStore } from "../../../store/store";
import { pizzaMenu } from "@/lib/getPizza";
import PlaceOrderButtom from "@/components/placeOrderButton";

const MenuPage = () => {
    const { data, loading, error } = useGetPizzaMenu();
    const [lowestPrice, setLowestPrice] = useState<number>(0);
    const { st_size, st_base, st_toppings, st_price, setSize, setBase, setTopping, setPrice, removeTopping } = usePizzaMenuStore() as { st_size: any, st_base: any, st_toppings: pizzaMenu['AvailablePizzaSelection']['toppings'], st_price: any, setSize: (value: any) => void, setBase: (value: any) => void, setTopping: (value: any) => void, setPrice: (value: number) => void, removeTopping: (value: any) => void };

    const calculateLowestPrice = useCallback(() => {
        let lowestPrice: number = 0, basePrice: number = 0, toppingPrice: number = 0, sizeMultiplier: number = 0;
        // get the lowest price from the base
        data?.bases?.forEach((base) => {
            if (base.price < basePrice || basePrice === 0) {
                basePrice = base.price;
            }
        });
        // get the lowest price from the toppings
        data?.toppings?.forEach((topping) => {
            if (topping.price < toppingPrice || toppingPrice === 0) {
                toppingPrice = topping.price;
            }
        });
        // get the lowest multiplier from the sizes
        data?.sizes?.forEach((size) => {
            if (size.multiplier < sizeMultiplier || sizeMultiplier === 0) {
                sizeMultiplier = size.multiplier;
            }
        });

        lowestPrice = (basePrice + toppingPrice) * sizeMultiplier;

        return lowestPrice;
    }, [data]);

    const calculatePrice = useCallback(() => {
        const price = st_size.multiplier * (st_base.price + st_toppings.reduce((acc, topping) => acc + topping.price, 0));
        setPrice(price);
    }, [st_size, st_base, st_toppings, setPrice]);
    
    useEffect(() => {
        if (data) {
            const price = calculateLowestPrice();
            setLowestPrice(price);
        }
    }, [data, calculateLowestPrice, setLowestPrice]);

    useEffect(() => {
        calculatePrice();
    }, [calculatePrice]);

    function handleSizeClick(size: pizzaMenu['AvailablePizzaSelection']['sizes'][number]) {
        // set the size
        if (size.id === st_size.id) {
            setSize({});
            return;
        }
        setSize(size);
    }

    function handleBaseClick(base: pizzaMenu['AvailablePizzaSelection']['bases'][number]) {
        // set the base
        if (base.id === st_base.id) {
            setBase({});
            return;
        }
        setBase(base);
    }

    function handleToppingClick(topping: pizzaMenu['AvailablePizzaSelection']['toppings'][number]) {
        // set the topping
        if (st_toppings.some(st_topping => st_topping.id === topping.id)) {
            console.log('removing topping');
            removeTopping(topping);
            return;
        }
        if (st_toppings.length >= 3) {
            console.log('topping limit reached');
            return;
        }
        setTopping(topping);
    }

    
    if (loading) return <p>Loading...</p>;
    if (error) return <p>Error</p>;
    
    return (
        <div className="fillPage p-[2rem] tertiary-background">
            <div className="border rounded-lg p-[1rem] primary-background">
                <div className="text-xl font-semibold">Pizza</div>
                <div className="text-sm">Starting from RM {lowestPrice}</div>
                <div className="fixed top-[8rem] right-[4rem] border-2 rounded-lg w-fit px-[2rem] py-[0.5rem] secondary-background">RM {st_price}</div>
                <div className="mt-[2rem]">
                    <div className="tertiary-background p-[1rem] text-lg font-semibold rounded-md border-2">1. Choose your size</div>
                    <div className="grid grid-cols-1 sm:grid-cols-4 gap-[2rem] mt-[1rem]">
                        {data?.sizes?.map((size) => {
                            let dimension = SizeDimensions[size.name as keyof typeof SizeDimensions];
                            const isSelected = size.id === st_size.id ;
                            return (
                                <div key={size.id} className={`rounded-lg h-[6.5rem] flex flex-row items-center cursor-pointer ${isSelected ? 'tertiary-background border-2' : 'secondary-background'}`} onClick={() => handleSizeClick(size)}>
                                    <Image 
                                        src="/images/pizza.png" 
                                        alt="Pizza" 
                                        width={dimension} 
                                        height={dimension} />
                                    <div className="flex-1">
                                        {size.name}
                                    </div>
                                    <div className="self-end mb-[1rem] mr-[1rem] text-xs">
                                        {size.multiplier}x
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
                <div className="mt-[4rem]">
                    <div className="tertiary-background p-[1rem] text-lg font-semibold rounded-md border-2">2. Choose your crust</div>
                    <div className="grid grid-cols-1 sm:grid-cols-4 gap-[2rem] mt-[1rem]">
                        {data?.bases?.map((base) => {
                            const isSelected = base.id === st_base.id ;
                            return (
                                <div key={base.id} className={`rounded-lg h-[6.5rem] flex flex-row items-center cursor-pointer ${isSelected ? 'tertiary-background border-2' : 'secondary-background'}`} onClick={() => handleBaseClick(base)}>
                                    <div className="flex-1 flex justify-center items-center text-2xl font-semibold">
                                        {base.name}
                                    </div>
                                    <div className="self-end mb-[1rem] mr-[1rem] text-xs">
                                        RM {base.price}
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
                <div className="mt-[4rem]">
                    <div className="tertiary-background p-[1rem] text-lg font-semibold rounded-md border-2">3. Choose your topping (up to 3)</div>
                    <div className="grid grid-cols-2 sm:grid-cols-4 gap-[2rem] mt-[1rem]">
                        {data?.toppings?.map((topping) => {
                            const isSelected = st_toppings.some(st_topping => st_topping.id === topping.id);
                            return (
                                <div key={topping.id} className={`rounded-lg h-[6.5rem] flex flex-row items-center cursor-pointer ${isSelected ? 'tertiary-background border-2' : 'secondary-background'}`} onClick={() => handleToppingClick(topping)}>
                                    <div className="flex-1 flex justify-center items-center text-2xl font-semibold">
                                        {topping.name}
                                    </div>
                                    <div className="self-end mb-[1rem] mr-[1rem] text-xs">
                                        RM {topping.price}
                                    </div>
                                </div>
                            )
                        })}
                    </div>
                </div>
                <div className="flex justify-center items-center mt-[4rem]">
                    <PlaceOrderButtom />
                </div>
            </div>
        </div>
    );
};

export default MenuPage;

enum SizeDimensions {
    Small = 90,
    Medium = 115,
    Large = 140
}
