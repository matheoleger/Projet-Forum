/*
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
*/

function enlargeDiv (
    menu_slider = document.getElementById('navbar_wrapper');
)

$(menu_slider.addClass(".enlargeDiv"));

function reduceDiv (
    menu_slider = document.getElementById('navbar_wrapper');
)

$(menu_slider.addClass(".reduceDiv"));


// $(menu_slider.removeClass(".enlargeDiv"));

/*
function enlargeDiv() {
    menu_slider.addClass(".enlargeDiv")
    menu_slider.removeClass(".enlargeDiv")
    menu_slider = true
}

function reduceDiv() {
    menu_slider.removeClass(".enlargeDiv")
    menu_slider.addClass(".reduceDiv")
    menu_slider = false
}
*/