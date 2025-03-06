import styled from "styled-components"
import Layout from "../Layout"
import FloatProfile from "../FloatProfile"
import useGitHubStore from "@store/useGitHubStore"
import useDownload from "@/hooks/useDownload"

const Container = styled('div')`
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  gap: 25px;
  height: 100vh;

  h1 {
    font-size: 3em;
    font-family: "Gloria Hallelujah", cursive;
    font-weight: 400;
  }

  h4 {
    font-weight: 500;
  }

  h1, h4 {
    text-align: center;
  }

  .container{
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    gap: 25px;
    width: 40%;
  }
`

const App = (): JSX.Element => {

  const github = useGitHubStore(state => state.data)

  const download = useDownload()

  const handleDownload = () => {

    if (github.assets.length === 0) {
      return
    }

    const data = github.assets[0]

    download({
      url: data.browser_download_url, 
      name: data.name
    })
  }

  return (
    <Layout>
      <Container>
        <div className="container">
          <h1>Storage in-memory</h1>
          <h4>Godis es  un sistema de base de datos en memoria de alta velocidad y baja latencia, escrita desde cero utilizando las características y ventajas del lenguaje Go</h4>

          <button onClick={handleDownload}>
            Download {github.tag_name}
          </button>
        </div>
      </Container>

      <FloatProfile />
    </Layout>
  )
}

export default App