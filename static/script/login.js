// function showLoginOrRegister(isLoginBtn) {

//     // let loginBtn = document.getElementById("choice_login_btn");
//     // let registerBtn = document.getElementById("choice_register_btn");

//     let loginForm = document.querySelector(".loginForm");
//     let registerForm = document.querySelector(".registerForm");

//     if(isLoginBtn) {
//         loginForm.style.display = "";
//         registerForm.style.display = "none";
//     } else {
//         loginForm.style.display = "none";
//         registerForm.style.display = "";
//     }


// }

function connexionForm() {
    window.location.href = "?loginForm=connexion"
    // showLoginForm()
}

function inscriptionForm() {
    window.location.href = "?loginForm=inscription"
    // showLoginForm()
}

// function showLoginForm() {

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
// }