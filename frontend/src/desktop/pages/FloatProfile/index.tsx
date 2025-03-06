import useGitHubStore from "@store/useGitHubStore";
import styled from "styled-components";

const Container = styled('div')`
    position: fixed;
    bottom: 0px;
    right: 0px;
    background-color: var(--nav-background-color);
    padding: 10px 20px;
    border-radius: 20px 0px 0px;

    display: flex;
    align-items: center;
    justify-content: start;
    gap: 15px;

    img {
        width: 45px;
        height: 45px;
        border-radius: 50%;
    }

    .username {
        font-size: .8em;
    }
`

const App = (): JSX.Element => {

    const github = useGitHubStore(state => state.data)

    const handleClick = () => {
        window.open(github.author.html_url);
    };

    return (
        <Container onClick={handleClick}>
            <img src={github.author.avatar_url} />
            <div>
                <p >Javier Avila</p>
                <p className="username">@{github.author.login}</p>
            </div>
        </Container>
    )
}

export default App