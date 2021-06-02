const API_URL = "http://localhost:1204";

const resetForm = (idForm) =>{
    document.getElementById(`${idForm}`).reset();
}

export {API_URL, resetForm}