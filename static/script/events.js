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
