import { useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate } from "react-router-dom";
import axios from "axios";

function DepartmentCreate() {
    const [department, setDepartment] = useState({id:'', name:'', description:''})
    const OnBoxChange = (event) => {
        const newDepartment = {...department};
        newDepartment[event.target.id] = event.target.value;
        setDepartment(newDepartment);
    }
    const navigate = useNavigate();
    const OnCreate = async () => {
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.post(`${baseUrl}/departments`, {...department});
            alert(response.data.message)
            navigate('/departments/list');
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader PageNumber={2}/>
            <h3><a href="/departments/list" className="btn btn-light">Go Back</a>New Department</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="name" className="form-label">Department Name:</label>
                    <input type="text" className="form-control" id="name" 
                        placeholder="Please enter department name"
                        value={department.name} onChange={OnBoxChange}/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="description" className="form-label">Description:</label>
                    <input type="text" className="form-control" id="description" 
                        placeholder="Please enter description"
                        value={department.description} onChange={OnBoxChange}/>
                </div>
                <button className="btn btn-success"
                    onClick={OnCreate}>Create Department</button>
            </div>
        </>
    );
}

export default DepartmentCreate;