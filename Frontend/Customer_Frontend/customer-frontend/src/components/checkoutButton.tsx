import { pizzaMenu } from "@/lib/getPizza";
import { usePostPizzaOrder } from "@/hooks/useGraphql";
import { usePizzaMenuStore, useUserStore } from "../../store/store";
import Link from "next/link";

const CheckoutButton = () => {
    const { st_size, st_base, st_toppings } = usePizzaMenuStore() as { st_size: any, st_base: any, st_toppings: pizzaMenu['AvailablePizzaSelection']['toppings']};
    const { customerId } = useUserStore() as { customerId: string };
    const postPizzaOrder = usePostPizzaOrder();


    const handleOrderSubmit = async () => {
        try {
            const data = await postPizzaOrder(st_base.id, st_size.id, st_toppings.map(topping => topping.id), customerId);
            console.log(data);
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <Link href='/orderSuccess'><div className='tertiary-background px-[2rem] py-[1rem] rounded-md border-2 cursor-pointer' onClick={handleOrderSubmit}>GO</div></Link>
    )
}
export default CheckoutButton;