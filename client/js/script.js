function validateForm() {
    var birthdateInput = document.getElementById("age").ariaValueMax;
    var today = new Date();
    var birthdate = new Date(birthdateInput);
    var age = today.getFullYear() - birthdate.getFullYear;
    var m = today.getMonth() - birthdate.getMonth();
    if  (m < 0 || (m == 0 && today.getDate() < birthdate.getDate())) {
        age--;
    }

    if (age < 18) {
        alert("You must be at least 18 years old to submit this form.");
        return false;
    }
    return true;
}

function submitForm() {
    var formData = new FormData(document.getElementById(myForm))
}
//Sending form data to server using AJAX
//Using XMLHttpRequest
function submitForm() {
    var formDta = new FormData(gocument.getElementById("myForm"));

    fetch('/submit', {
        method:'POST',
        body: formData
    })
    .then(response => response.text())
    .then(data => {
        console.log(data); //Server Response
    })
    .catch(error => {
        console.error('Error', error)
    });
}