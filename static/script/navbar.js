let flipflop = true;

function showDiv() {

    if (flipflop) {
        document.getElementById('navbar_wrapper').style.visibility = "visible";
        document.querySelector('li').style.visibility = "visible";
        flipflop = false

    } else {
        document.getElementById('navbar_wrapper').style.visibility = "hidden";
        document.querySelector('li').style.visibility = "hidden";
        flipflop = true


    }
}