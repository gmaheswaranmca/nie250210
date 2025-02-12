import PageHeader from "../header/PageHeader";

function CarView() {
    return(
        <>
            <PageHeader/>
            <h3><a href="/cars/list" className="btn btn-light">Go Back</a>View Car</h3>
            <div className="container">
                <div className="form-group mb-3">
                    <label for="number" className="form-label">Car Number:</label>
                    <div className="form-control" id="number">???</div>
                </div>
                <div className="form-group mb-3">
                    <label for="model" className="form-label">Car Model:</label>
                    <div className="form-control" id="model">???</div>
                </div>
                <div className="form-group mb-3">
                    <label for="type" className="form-label">Car Type(SUV/ CUV/ Sedan):</label>
                    <div className="form-control" id="type">???</div>
                </div>
            </div>
        </>
    );
}

export default CarView;