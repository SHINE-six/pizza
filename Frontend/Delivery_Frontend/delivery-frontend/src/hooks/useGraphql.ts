import { GET_ORDER, UPDATE_ORDER_STATUS, orderType } from "@/lib/getOrder";
import { useQuery, useMutation } from "@apollo/client";
import { apolloClient } from "@/lib/graphql_client";


const useGetOrder = ({ status, Id }: { status?: string; Id?: string }) => {
  const { data, loading, error, refetch } = useQuery<orderType>(GET_ORDER, {
    variables: { status: status, deliveryStaffId: Id },
    client: apolloClient,
  });

  // const formattedData = data?.allOrders;

  return { data, loading, error, refetch };
};

const useUpdateOrderStatus = () => {
  const [updateOrderStatus] = useMutation(UPDATE_ORDER_STATUS, {
    client: apolloClient,
  });

  const updateOrder = async (orderId: string, status: string, deliveryStaffId?: string) => {
    const { data } = await updateOrderStatus({
      variables: {
        orderId: orderId,
        status: status,
        deliveryStaffId: deliveryStaffId,
      },
    });

    return data;
  }

  return updateOrder;
}

export { useGetOrder, useUpdateOrderStatus };