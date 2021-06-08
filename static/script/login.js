function showLoginOrRegister(isLoginBtn) {

    // let loginBtn = document.getElementById("choice_login_btn");
    // let registerBtn = document.getElementById("choice_register_btn");

    let loginForm = document.querySelector(".loginForm");
    let registerForm = document.querySelector(".registerForm");

    if(isLoginBtn) {
        loginForm.style.display = "";
        registerForm.style.display = "none";
    } else {
        loginForm.style.display = "none";
        registerForm.style.display = "";
    }
}