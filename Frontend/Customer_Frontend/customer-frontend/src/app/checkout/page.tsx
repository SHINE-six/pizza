'use client';

import { usePizzaMenuStore, useUserStore } from "../../../store/store";
import { pizzaMenu } from "@/lib/getPizza";
import CheckoutButton from "@/components/checkoutButton";

const Checkout = () => {
    const { st_size, st_base, st_toppings, st_price } = usePizzaMenuStore() as { st_size: pizzaMenu['AvailablePizzaSelection']['sizes'][number], st_base: pizzaMenu['AvailablePizzaSelection']['bases'][number], st_toppings: pizzaMenu['AvailablePizzaSelection']['toppings'], st_price: any };
    const { username } = useUserStore() as { username: any };


    return (
        <div className='fillPage grid sm:grid-cols-2 grid-cols-1 font-bold'>
            <div className="m-[1rem] p-[1rem] space-y-[2rem] tertiary-background h-fit rounded-sm border-2">
                <div className="flex flex-row">
                    <div>Name</div>
                    <div className="flex-1 flex justify-center">{username}</div>
                </div>
                <div className="flex flex-row">
                    <div>Pizza</div>
                    <div className="flex-1 flex flex-col justify-center">
                        <div className="grid grid-cols-3 gap-x-[4rem] gap-y-[0.5rem]">
                            <>
                            <div className="text-right">Size</div>
                            <div>{st_size.name}</div>
                            <div>x{st_size.multiplier}</div>
                            </>

                            <>
                            <div className="text-right">Base</div>
                            <div>{st_base.name}</div>
                            <div>RM {st_base.price}</div>
                            </>
                            
                            <>
                            <div className="text-right">Topping</div>
                            <div>{st_toppings.map((topping: any) => topping.name).join(', ')}</div>
                            <div>RM {st_toppings.reduce((acc: number, topping: any) => acc + topping.price, 0)}</div>
                            </>

                            <div className="col-start-3">
                                <hr />
                                RM {st_price}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div>payment thingy</div>
            <div className="col-span-2 flex justify-center items-center">
                <CheckoutButton />
            </div>
        </div>
    )
}
export default Checkout;