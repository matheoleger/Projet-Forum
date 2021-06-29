let flipflop = false;
let nav = document.querySelector('nav');

function setupFlipflop() {
    
    document.querySelector('.headerToNavbar_button').addEventListener('click', function(event) {

        event.preventDefault()
        event.stopPropagation()
        flipflop = !flipflop
        
        if (flipflop) {
            enlargeDiv()
        } else { 
            reduceDiv()
        }
    });
    
    document.body.addEventListener('click', function(event){

        if (event.target != nav) {
            reduceDiv()
            flipflop = false
        }
        
    });
    
}

setupFlipflop();


function vider()
{
	document.getElementById("filtre").value = "";
	return false;
};



// function dispatchEvent (eventClick) {

//     if (flipflop = 0)

//         (document.getElementById('headerToNavbar_button')) {

//         enlargeDiv();
//         flipflop = flipflop + 1;

//         console.log(flipflop);
//     }

//     if (flipflop = 1)

//     (MouseEvent.click((document.getElementById('headerToNavbar_button')))) {

//         reduceDiv();
//         flipflop = flipflop - 1;

//         console.log(flipflop);
//     }

// if (MouseEvent.click((document.getElementById('headerToNavbar_button')))) {
//     reduceDiv()
// }