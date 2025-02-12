import CarList from './cars/CarList'
import CarCreate from './cars/CarCreate'
import CarView from './cars/CarView'

import { BrowserRouter, Route, Routes } from 'react-router-dom'

function App() {
  return (
    <>
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<CarList/>}/>
            <Route path="/cars/list" element={<CarList/>}/>
            <Route path="/cars/create" element={<CarCreate/>}/>
            <Route path="/cars/view/:id" element={<CarView/>}/>
          </Routes>
        </BrowserRouter>
      </div>
    </>
  );
}

export default App;
