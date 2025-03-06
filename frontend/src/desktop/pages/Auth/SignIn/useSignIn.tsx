import { useForm } from "react-hook-form";

import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useState } from "react";

const schema = z.object({
    email: z.string()
        .min(1, { message: "Email es requerido" })
        .email({ message: "Email no es válido" }),
    password: z.string()
        .min(6, { message: "La contraseña debe tener al menos 6 caracteres" })
});

export type LoginForm = z.infer<typeof schema>;

const useSignIn = () => {

    const [showPassword, setShowPassword] = useState(false);

    const { register, handleSubmit, formState: { errors } } = useForm<LoginForm>({
        resolver: zodResolver(schema),
    });

    const onSubmit = (data: LoginForm) => {
        console.log({ data });
    };

    return {
        showPassword,
        setShowPassword,
        register, 
        errors, 
        handleSubmit: handleSubmit(onSubmit)
    }
}

export default useSignIn