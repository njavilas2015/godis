import { create } from 'zustand';

interface State {
    token?: string
    loading: boolean
    isAuthenticated: boolean

    error?: string
    isError: boolean

    setAuth: (token: string) => void;
    setLogout: () => void;
}

const useStore = create<State>((set) => ({
    loading: false,
    isAuthenticated: false,

    isError: false,
    setAuth: (token: string) => set({ token }),
    setLogout: () => set({
        loading: false,
        isAuthenticated: false,

        isError: false,
    }),
}));

export default useStore;