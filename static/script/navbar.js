let flipflop = true;

function showDiv() {

    if (flipflop) {
        document.getElementById('navbar_menu').style.visibility = "visible";
        document.querySelector('.navbar_elements').style.visibility = "visible";
        flipflop = false

    } else {
        document.getElementById('navbar_menu').style.visibility = "hidden";
        document.querySelector('.navbar_elements').style.visibility = "hidden";
        flipflop = true


    }
 }