import { FieldValues, Path, UseFormRegister } from "react-hook-form"
import styled from "styled-components"

const Container = styled('div')`
    .input-container {
        position: relative;
        width: 300px;
    }

    .input-icon {
        position: absolute;
        top: 50%;
        left: 10px; 
        transform: translateY(-50%);
        width: 20px; 
        height: 20px;
    }

    .input-field {
        width: 100%;
        padding: 10px 10px 10px 40px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 16px;
    }
`

export interface IInputIcon<T extends FieldValues> {
    htmlFor: string
    label: string
    icon: string
    id: string
    type?: React.HTMLInputTypeAttribute
    error_message?: string
    register: UseFormRegister<T>;
    register_name: Path<T>
}

const InputIcon = <T extends FieldValues,>(params: IInputIcon<T>): JSX.Element => {

    return (
        <Container>
            <label htmlFor={params.htmlFor}>{params.label}</label>

            <div className="input-container">
                <img src={params.icon} alt="Icono" className="input-icon" />
                <input
                    className="input-field"
                    type={params.type || 'text'}
                    id={params.id}
                    {...params.register(params.register_name)}
                />
            </div>
            {params.error_message && <p style={{ color: 'red' }}>{params.error_message}</p>}
        </Container>
    )
}

export default InputIcon