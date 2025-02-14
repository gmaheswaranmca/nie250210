import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import axios from 'axios';

function DepartmentList() {
    const [departments, setDepartments] = useState([]);
    const readAllDepartments = async () => {
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.get(`${baseUrl}/departments`);
            setDepartments(response.data);
            
        } catch(error) {
            alert('Server Error');
        }
    };
    useEffect(()=>{ readAllDepartments(); },[]);
    const OnDelete = async (id) => {
        if(!confirm("Are you sure to delete?")) {
            return;
        }
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.delete(`${baseUrl}/departments/${id}`);
            alert(response.data.message)
            readAllDepartments();
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
            <PageHeader PageNumber={1}/>
            <h3>List of Departments</h3>
            <div className="container">
                <table className="table table-primary table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Name</th>
                            <th scope="col">Description</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        { departments.map( (department, index) => {
                            return (
                            <tr>
                                <th scope="row">{index + 1}</th>
                                <td>{department.name}</td>
                                <td>{department.description}</td>
                                <td>
                                    <a href={`/departments/edit/${department.id}`} className="btn btn-warning">Edit</a>
                                    <button className="btn btn-danger"
                                        onClick={()=>{OnDelete(department.id);}}>Delete</button>
                                </td>
                            </tr>
                            );
                        } ) 
                        }
                        
                        
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default DepartmentList;
