let flipflop = true;
let postBool = true;

function showDiv() {

    if (flipflop) {
        // document.getElementById('navbar_wrapper').style.visibility = "visible";
        // document.querySelector('li').style.visibility = "visible";
        document.querySelector('.navbar_wrapper').classList.add('navbar_wrapper_appear')
        flipflop = false
        console.log("coucou")

    } else {
        // document.getElementById('navbar_wrapper').style.visibility = "hidden";
        // document.querySelector('li').style.visibility = "hidden";

        document.querySelector('.navbar_wrapper').classList.remove('navbar_wrapper_appear')
        flipflop = true
    }
}

function displayPost() {
    if (postBool == true) {
        postBool = false
        console.log("true")
        document.querySelector('.Ypost').style.visibility = "visible"
        
        //console.log("true")
    } else {
        postBool = true
        console.log("false")
        document.querySelector('.Ypost').style.visibility = "hidden"
        
        
    }
}

