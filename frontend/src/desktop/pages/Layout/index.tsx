import { Fragment, ReactNode } from "react"
import styled from "styled-components"

const NavBar = styled('div')`
    position: fixed;
    top: 15px;
    left: 50%;
    transform: translateX(-50%);
    display: flex;
    justify-content: space-between;
    align-items: center;

    background-color: var(--nav-background-color);

    width: 100%;
    max-width: 800px; 
    padding: 5px 20px;
    border-radius: 12px;

    .company {
        font-size: 2em;
        letter-spacing: 3px;
        font-family: "Gloria Hallelujah", cursive;
        font-weight: 400;
    }

    .item-tab{
        font-weight: 400;
    }
`

export interface IParams {
    children?: ReactNode
}

const App = (params: IParams): JSX.Element => {

    const items: string[] = [
        "Resources",
        "Pricing",
        "Docs",
    ]

    return (
        <Fragment>
            <NavBar>
                <p className="company">Godis</p>

                <div className="row" style={{ gap: 25 }}>
                    <div className="row">
                        {items.map((v, i) => <p key={i} className="item-tab">{v}</p>)}
                    </div>

                    <button>
                        Download
                    </button>
                </div>

            </NavBar>

            {params.children}
        </Fragment>
    )
}

export default App