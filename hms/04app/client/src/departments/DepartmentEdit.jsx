import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import { useNavigate, useParams } from "react-router-dom";
import axios from "axios";

function DepartmentEdit() {
    const [department, setDepartment] = useState({id:'', name:'', description:''})
    const OnBoxChange = (event) => {
        const newDepartment = {...department};
        newDepartment[event.target.id] = event.target.value;
        setDepartment(newDepartment);
    }
    const params = useParams();
    const readDepartmentById = async () => {
        //alert(params.id);
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.get(`${baseUrl}/departments/${params.id}`);
            setDepartment(response.data);
            
        } catch(error) {
            alert('Server Error');
        }
    };
    useEffect(()=>{ readDepartmentById(); },[]);
    const navigate = useNavigate();
    const OnUpdate = async () => {
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.put(`${baseUrl}/departments/${params.id}`, {...department});
            alert(response.data.message)
            navigate('/departments/list');
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader  PageNumber={1}/>
            <h3><a href="/departments/list" className="btn btn-light">Go Back</a>Edit Department Ticket Price</h3>
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
                <button className="btn btn-warning"
                    onClick={OnUpdate}>Update Price</button>
            </div>
        </>
    );
}

export default DepartmentEdit;