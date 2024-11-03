import { useState} from 'react';
import { orderType } from '@/lib/getOrder';
import { useUpdateOrderStatus } from '@/hooks/useGraphql';


const YourDeliveryOrder = ({ order } : { order: orderType['allOrders'][number] }) => {
    const [checked, setChecked] = useState(false);
    const [minimizedPizzas, setMinimizedPizzas] = useState<{ [key: string]: boolean }>({});
    const updateOrderStatus = useUpdateOrderStatus();

    const togglePizzaMinimize = (pizzaId: string) => {
        setMinimizedPizzas(prevState => ({
            ...prevState,
            [pizzaId]: !prevState[pizzaId]
        }))
    }

    async function updateOrderStatusToDelivering() {
        try {
            const data = await updateOrderStatus(order.id, 'delivered');
            console.log(data);
        } catch (error) {
            console.error(error);
        }
    }

    const toggleCheckBox = () => {
        setChecked(!checked);
        if (!checked) {
            updateOrderStatusToDelivering();
        }
    }

    return (
        <div className='border-2 border-black p-2 m-2'>
            <div className='p-1'>
                <input 
                    type="checkbox" 
                    className='scale-150' 
                    checked={checked}
                    onChange={() => toggleCheckBox()}/>
            </div>
            <div>Order Id: {order.id}</div>
            <div>Customer: {order.customer.username}</div>
            <div>Delivery Staff: {order.deliveryStaff.name}</div>
            <div>Status: {order.Status}</div>
            <div>Created At: {order.createdAt}</div>
            <div>Total Price: {order.totalPrice}</div>
            <div>Pizzas: {order.pizzas.map((pizza) => {
                const isMinimized = minimizedPizzas[pizza.id];
                if (isMinimized) {
                    return (
                        <div key={pizza.id} className='border-2 border-black p-2 m-2' onClick={() => togglePizzaMinimize(pizza.id)}>
                            <div>Pizza ID: {pizza.id}</div>
                        </div>
                    )
                }
                return (
                    <div key={pizza.id} className='border-2 border-black p-2 m-2' onClick={() => togglePizzaMinimize(pizza.id)}>
                        <div>Pizza ID: {pizza.id}</div>
                        <div>Size: {pizza.size.name}</div>
                        <div>Base: {pizza.base.name}</div>
                        <div>Toppings: {pizza.toppings.map((topping) => {
                            return (
                                <div key={topping.id} className='border-2 border-black p-2 m-2'>
                                    <div>Name: {topping.name}</div>
                                </div>
                            )
                        })}</div>
                    </div>
                )
            })}</div>
        </div>
    )
}
export default YourDeliveryOrder;