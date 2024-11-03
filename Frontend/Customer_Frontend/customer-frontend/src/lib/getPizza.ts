import { gql } from "@apollo/client";

type pizzaMenu = {
  AvailablePizzaSelection: {
    sizes: {
      id: string;
      name: string;
      multiplier: number;
    }[];
    bases: {
      id: string;
      name: string;
      price: number;
    }[];
    toppings: {
      id: string;
      name: string;
      price: number;
    }[];
  };
};

const GET_PIZZA_MENU = gql`
  query {
    AvailablePizzaSelection {
      sizes {
        id
        name
        multiplier
      }
      bases {
        id
        name
        price
      }
      toppings {
        id
        name
        price
      }
    }
  }
`;

type postPizzaOrderType = {
  createPizza: {
    orderId: string;
  };
};

const POST_PIZZA_ORDER = gql`
  mutation PostPizzaOrder($baseId: ID!, $sizeId: ID!, $toppingIds: [ID!]!, $customerId: ID!) {
    createPizza(baseId: $baseId, sizeId: $sizeId, toppingIds: $toppingIds, customerId: $customerId) {
      orderId
    }
  }
`;


export { GET_PIZZA_MENU, POST_PIZZA_ORDER };
export type { pizzaMenu, postPizzaOrderType };