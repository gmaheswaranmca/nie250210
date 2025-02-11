import PageHeader from "../header/PageHeader";

function FlightCreate() {
    return (
        <>
            <PageHeader PageNumber={2}/>
            <h3><a href="/flights/list" className="btn btn-light">Go Back</a>New Flight</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label htmlFor="number" className="form-label">Flight Number:</label>
                    <input type="text" className="form-control" id="number" placeholder="Please enter flight number"/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="airline_name" className="form-label">Airline Name:</label>
                    <input type="text" className="form-control" id="airline_name" placeholder="Please enter airline name"/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="source" className="form-label">Source:</label>
                    <input type="text" className="form-control" id="source" placeholder="Please enter source"/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="destination" className="form-label">Destination:</label>
                    <input type="text" className="form-control" id="destination" placeholder="Please enter destination"/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="capacity" className="form-label">Capacity(Number of Seats):</label>
                    <input type="text" className="form-control" id="capacity" placeholder="Please enter capacity"/>
                </div>
                <div className="form-group mb-3">
                    <label htmlFor="price" className="form-label">Ticket Price:</label>
                    <input type="text" className="form-control" id="price" placeholder="Please enter ticket price"/>
                </div>
                <button className="btn btn-success">Create Flight</button>
            </div>
        </>
    );
}

export default FlightCreate;