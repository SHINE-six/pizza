import {GET_PIZZA_MENU, POST_PIZZA_ORDER, pizzaMenu, postPizzaOrderType } from "@/lib/getPizza";
import { useQuery, useMutation } from "@apollo/client";
import { apolloClient } from "@/lib/graphql_client";


const useGetPizzaMenu = () => {
  const { data, loading, error } = useQuery<pizzaMenu>(GET_PIZZA_MENU, {
    client: apolloClient,
  });

  const formattedData = data?.AvailablePizzaSelection;

  return { data: formattedData, loading, error };
};

const usePostPizzaOrder = () => {
  const [postPizzaOrder] = useMutation(POST_PIZZA_ORDER, {
    client: apolloClient,
  });

  const submitOrder = async (baseId: string, sizeId: string, toppingIds: string[], customerId: string) => {
    const { data } = await postPizzaOrder({
      variables: {
        baseId,
        sizeId,
        toppingIds,
        customerId,
      },
    });

    // convert data into type postPizzaOrderType
    const formattedData = data as postPizzaOrderType['createPizza'];


    return formattedData;
  }

  return submitOrder;
}

export { useGetPizzaMenu, usePostPizzaOrder };