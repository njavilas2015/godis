import Desktop from '@/desktop'
import useGitHub from '@hooks/useGitHub'

const App = (): JSX.Element => {
 
  useGitHub()
  
  return (
    <Desktop>

    </Desktop>
  )
}

export default App