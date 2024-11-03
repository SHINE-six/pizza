import { create } from 'zustand';
import { pizzaMenu } from '@/lib/getPizza';

const useUserStore = create(set => ({
    username: '',
    email: '',
    customerId: '',
    setUsername: (value: string) => set(() => ({ username: value })),
    setEmail: (value: string) => set(() => ({ email: value })),
    setCustomerId: (value: string) => set(() => ({ customerId: value })),
}));

const useNavBarStore = create(set => ({
    isNavBarOpen: true,
    setIsNavBarOpen: (value: boolean) => set(() => ({ isNavBarOpen: value })),
}));


type PizzaMenuStoreState = {
    st_size: pizzaMenu['AvailablePizzaSelection']['sizes'][number],
    st_base: pizzaMenu['AvailablePizzaSelection']['bases'][number],
    st_toppings: pizzaMenu['AvailablePizzaSelection']['toppings'],
    st_price: number,
};

const usePizzaMenuStore = create(set => ({
    st_size: {} as pizzaMenu['AvailablePizzaSelection']['sizes'][number],
    st_base: {} as pizzaMenu['AvailablePizzaSelection']['bases'][number],
    st_toppings: [] as pizzaMenu['AvailablePizzaSelection']['toppings'],
    st_price: 0,
    setSize: (value: pizzaMenu['AvailablePizzaSelection']['sizes'][number]) => set(() => ({ st_size: value })),
    setBase: (value: pizzaMenu['AvailablePizzaSelection']['bases'][number]) => set(() => ({ st_base: value })),
    setTopping: (value: pizzaMenu['AvailablePizzaSelection']['toppings'][number]) => set((state: PizzaMenuStoreState) => ({ st_toppings: [...state.st_toppings, value] })),
    setPrice: (value: number) => set(() => ({ st_price: value })),

    removeTopping: (value: pizzaMenu['AvailablePizzaSelection']['toppings'][number]) => set((state: PizzaMenuStoreState) => ({
       st_toppings: state.st_toppings.filter(topping => topping !== value)
    })),
}));

export { useUserStore, useNavBarStore, usePizzaMenuStore };