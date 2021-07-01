function connexionForm() {
    window.location.href = "?loginForm=connexion"
}

function inscriptionForm() {
    window.location.href = "?loginForm=inscription"
}

const loginFormElement = document.querySelector(".loginForm");
const registerForm = document.querySelector(".registerForm");

const queryString = window.location.search

const urlParams = new URLSearchParams(queryString)

const loginForm = urlParams.get("loginForm")
console.log(loginForm)
const errType = urlParams.get("err")

if (loginForm == "connexion") {
    
    loginFormElement.style.display = "";
    registerForm.style.display = "none";

} else if (loginForm == "inscription") {
    
    loginFormElement.style.display = "none";
    registerForm.style.display = "";

}

if (errType == "wrong_PW") {
    alert("Erreur : Vous avez renseigné un mauvais mot de passe...")
} else if (errType == "wrong_name") {
    alert("Erreur : Vous avez renseigné un mauvais nom d'utilisateur...")
} else if (errType == "wrong_EMAIL") {
    alert("Erreur : Vous avez renseigné un email qui existe déjà ... ")
}
