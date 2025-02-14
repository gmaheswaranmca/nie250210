import DepartmentList from "./departments/DepartmentList"
import DepartmentCreate from "./departments/DepartmentCreate"
import DepartmentEdit from "./departments/DepartmentEdit"

import { BrowserRouter, Route, Routes } from 'react-router-dom'

function App() {
  return (
    <>     
      <div>
        <BrowserRouter>
          <Routes>
            <Route path="" element={<DepartmentList/>}/>
            <Route path="/departments/list" element={<DepartmentList/>}/>
            <Route path="/departments/create" element={<DepartmentCreate/>}/>
            <Route path="/departments/edit/:id" element={<DepartmentEdit/>}/>
          </Routes>
        </BrowserRouter>
      </div>
      
    </>
  )
}

export default App
