import { gql } from "@apollo/client";

type pizzaMenu = {
  id: string;
  createdAt: string;
  updatedAt: string;
  Price: number;
  size: {
    id: string;
    name: string;
    multiplier: number;
  };
  base: {
    id: string;
    name: string;
    price: number;
  };
  toppings: {
    id: string;
    name: string;
    price: number;
  }[];
};

type orderType = {
  allOrders: {
    id: string,
    customer: {
      id: string,
      email: string,
      username: string
    },
    deliveryStaff: {
      id: string,
      name: string,
      email: string
    },
    Status: string,
    createdAt: string,
    totalPrice: number,
    pizzas: pizzaMenu[]
  }[]
};

const GET_ORDER = gql`
query AllOrders($status: String, $deliveryStaffId: ID) {
  allOrders(status: $status, deliveryStaffId: $deliveryStaffId) {
    id
    customer {
      id
      email
      username
    }
    deliveryStaff {
      id
      name
      email
    }
    Status
    createdAt
    totalPrice
    pizzas {
      id
      createdAt
      updatedAt
      Price
      base {
        id
        name
        price
      }
      size {
        name
        multiplier
        id
      }
      toppings {
        id
        name
        price
      }
    }
  }
}
`;

const UPDATE_ORDER_STATUS = gql`
mutation UpdateOrderStatus($orderId: ID!, $status: String!, $deliveryStaffId: ID) {
  updateOrderStatus(orderId: $orderId, status: $status, deliveryStaffId: $deliveryStaffId) {
    id
    Status
  }
}
`;

export { GET_ORDER, UPDATE_ORDER_STATUS };
export type { orderType };