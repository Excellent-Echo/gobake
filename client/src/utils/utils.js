const API_URL = "http://localhost:1204";

const resetForm = (idForm) =>{
    document.getElementById(`${idForm}`).reset();
}

function numberWithCommas(x) {
    x = x.toString();
    var pattern = /(-?\d+)(\d{3})/;
    while (pattern.test(x))
        x = x.replace(pattern, "$1,$2");
    return x;
}

export {API_URL, resetForm, numberWithCommas}