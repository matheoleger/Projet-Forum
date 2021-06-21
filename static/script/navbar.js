function enlargeDiv() {

    document.querySelector('.navbar_wrapper').classList.add('navbar_wrapper_appear')
    console.log("coucou")
}

function reduceDiv () {

    document.querySelector('.navbar_wrapper').classList.remove('navbar_wrapper_appear')
    console.log("cya")

}

let isShowCreatePost = true;

function showCreatePost() {
    
    if(!isShowCreatePost) {
        document.querySelector('.creationpost-bg').style.display = "none";
        isShowCreatePost = true;
    } else {
        document.querySelector('.creationpost-bg').style.display = "flex";
        isShowCreatePost = false;
    }
}

