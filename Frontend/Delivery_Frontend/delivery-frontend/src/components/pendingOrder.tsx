import { useEffect, useState, useCallback} from 'react';
import { orderType } from '@/lib/getOrder';
import { useUpdateOrderStatus } from '@/hooks/useGraphql';
import useUserStore from '../../store/store';


const PendingOrder = ({ order } : { order: orderType['allOrders'][number] }) => {
    const [checked, setChecked] = useState(false);
    const { Id } = useUserStore() as { Id: string };
    const updateOrderStatus = useUpdateOrderStatus();

    const updateOrderStatusToDelivering = useCallback(async () => {
        try {
            const data = await updateOrderStatus(order.id, 'delivering', Id);
            console.log(data);
        } catch (error) {
            console.error(error);
        }
    }, [order.id, Id, updateOrderStatus]);

    const toggleCheckBox = () => {
        setChecked(!checked);
    }

    useEffect(() => {
        if (checked) {
            updateOrderStatusToDelivering();
        }
    }, [checked, updateOrderStatusToDelivering]);

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
                return (
                    <div key={pizza.id} className='border-2 border-black p-2 m-2'>
                        <div>Size: {pizza.size.name}</div>
                        <div>Base: {pizza.base.name}</div>
                        <div>Toppings: {pizza.toppings.map((topping) => {
                            return (
                                <div key={topping.id} className='border-2 border-black p-2 m-2'>
                                    <div>Name: {topping.name}</div>
                                    <div>Price: {topping.price}</div>
                                </div>
                            )
                        })}</div>
                    </div>
                )
            })}</div>
        </div>
    )
}
export default PendingOrder;