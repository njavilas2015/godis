import { Routes, Route } from 'react-router-dom';
import useRouter, { IRoute } from '@desktop/routes/useRouter';

function App() {

  const routes: IRoute[] = useRouter()

  return (
    <Routes>
      {routes.map((v, i) => <Route key={i} path={v.path} element={v.element} />)}
    </Routes>
  );
}

export default App;
