// let isAlreadyLiked = false;
// let isAlreadyDisLiked = false;

function LikedPost(id, isLiked) {
    // fetch(`/like?post=${id}&isLiked=${isLiked}`)

    let queryString = `div#${CSS.escape(id)}.posts`
    let postEl = document.querySelector(queryString)

    let likeState = postEl.dataset.likestate
    let isLikedData = postEl.dataset.isliked

    console.log(queryString)
    console.log(postEl.dataset.likestate)
    console.log(postEl.dataset.isliked)

    if (likeState == "false") {

        postEl.dataset.likestate = "true"
        postEl.dataset.isliked = isLiked 

    } else if (likeState == "true" && isLiked == isLikedData) {

        postEl.dataset.likestate = "false"
    
    } else if (likeState == "true" && isLikedData != isLiked) {

        postEl.dataset.isliked = isLiked
    }

    if (postEl.dataset.likestate == "true" && postEl.dataset.isliked == "true") {
        
        let imgLike = document.querySelector(`${queryString} .like-btn-img`)
        imgLike.setAttribute("src", "../static/img/arrow-like-coloring.png")

    } else if (postEl.dataset.likestate == "true" && postEl.dataset.isliked == "false") {
        
        let imgLike = document.querySelector(`${queryString} .dislike-btn-img`)
        imgLike.setAttribute("src", "../static/img/arrow-dislike-coloring.png")
    }

    // if (!isAlreadyLiked) {
    //     isAlreadyLiked = true;
    // } else {
    //     isAlreadyLiked = false;
    // }
}

function LikedComment(id, isLiked) {
    fetch(`/like?comment=${id}&isLiked=${isLiked}`)
}

