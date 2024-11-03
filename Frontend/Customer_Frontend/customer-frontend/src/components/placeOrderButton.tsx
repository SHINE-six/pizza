import { useAuth } from "@/hooks/useAuth";
import Link from "next/link";

const PlaceOrderButtom = () => {
    const isAuthenticated = useAuth();
    
    return (
        <div className='tertiary-background px-[2rem] py-[1rem] rounded-md border-2 cursor-pointer'>{isAuthenticated ? <Link href='/checkout'>Place Order</Link>
        : <Link href='/login'>Login to Place Order</Link>}
        </div>
    )
}
export default PlaceOrderButtom;