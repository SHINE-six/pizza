import { create } from 'zustand';

const useUserStore = create(set => ({
    name: '',
    Id: '',
    CurrentPage: '',
    setName: (value: string) => set(() => ({ name: value })),
    setId: (value: string) => set(() => ({ Id: value })),
    setCurrentPage: (value: string) => set(() => ({ CurrentPage: value })),
}));

export default useUserStore;