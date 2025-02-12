import PageHeader from "../header/PageHeader";

function CarList() {
    return (
        <>
            <PageHeader />
            <h3>List of Cars</h3>
            <div className="container">
                <table className="table table-success table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">ID</th>
                            <th scope="col">Car Number</th>
                            <th scope="col">Model</th>
                            <th scope="col">Type</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <th scope="row">67a5e5733c7771e781094636</th>
                            <td>KA03 A1010</td>
                            <td>Maruti Suzuki</td>
                            <td>CUV</td>
                            <td><a href="/cars/view/67a5e5733c7771e781094636" 
                                className="btn btn-success">View</a></td>
                        </tr>
                        <tr>
                            <th scope="row">67a5e5943c7771e781094637</th>
                            <td>KA25 Z5599</td>
                            <td>Toyoto Innova</td>
                            <td>SUV</td>
                            <td><a href="/cars/view/67a5e5943c7771e781094637" 
                                className="btn btn-success">View</a></td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default CarList;