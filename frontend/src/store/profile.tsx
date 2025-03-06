import { create } from 'zustand';

export interface IProfile {

}

export interface State {
    profile: IProfile
    setProfile: (profile: IProfile) => void;
}

const useStore = create<State>((set) => ({
    profile: {

    },
    setProfile: (profile: IProfile) => set({ profile }),
}));

export default useStore;