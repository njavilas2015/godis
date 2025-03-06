import { useEffect } from "react";
import styled from "styled-components";

import InputIcon from "@components/InputIcon";

import useSignIn, { LoginForm } from "./useSignIn";


const Container = styled('div')`
   
`

const App = (): JSX.Element => {

    const { errors, handleSubmit, register } = useSignIn()


    useEffect(() => {

        const controller: AbortController = new AbortController()

        return () => {

            controller.abort()
        }

    }, [])

    return (
        <Container>
            <form onSubmit={handleSubmit} noValidate={true}>
                <InputIcon<LoginForm>
                    htmlFor="email"
                    label="Email"
                    icon="account.svg"
                    id="email"
                    type="email"
                    register={register}
                    register_name="email"
                    error_message={errors.email?.message}
                />

                <InputIcon<LoginForm>
                    htmlFor="password"
                    label="Contraseña"
                    icon="vite.svg"
                    id="password"
                    type="password"
                    register={register}
                    register_name="password"
                    error_message={errors.password?.message}
                />

                <button type="submit">Iniciar Sesión</button>
            </form>
        </Container>
    )
}

export default App