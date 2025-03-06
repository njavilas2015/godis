import { useEffect } from "react"
import { NavigateFunction, useNavigate } from 'react-router-dom';

import About from "@desktop/pages/About"
import Contact from "@desktop/pages/Contact"
import Home from "@desktop/pages/Home"
import SignIn from "@desktop/pages/Auth/SignIn"

import storeAuth from '@store/auth'

export interface IRoute {
    path: string
    element: JSX.Element
}

const useRouter = (): IRoute[] => {

    const navigate: NavigateFunction = useNavigate();

    const isAuthenticated: boolean = storeAuth(state => state.isAuthenticated)

    const routes: Array<IRoute> = [
        { path: "/", element: <Home /> },
        { path: "/sign-in", element: <SignIn /> },
        { path: "/about", element: <About /> },
        { path: "/contact", element: <Contact /> },
    ]

    useEffect(() => {

        if (!isAuthenticated) {
            //navigate('/sign-in');
        }

    }, [navigate, isAuthenticated])

    return routes
}

export default useRouter