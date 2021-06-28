// let isAlreadyLiked = false;
// let isAlreadyDisLiked = false;

function LikedPost(id, isLiked) {
    fetch(`/like?post=${id}&isLiked=${isLiked}`)

    let queryString = `div#${CSS.escape(id)}.posts`

    ChangeAttributeLike(queryString, isLiked)

}

function LikedComment(id, isLiked) {
    fetch(`/like?comment=${id}&isLiked=${isLiked}`)
    
    let queryString = `div#${CSS.escape(id)}.posts`

    ChangeAttributeLike(queryString, isLiked)
}

let posts = document.querySelectorAll(".posts")

for(let post of posts) {

    let queryString = `div#${CSS.escape(post.id)}.posts`

    VerifyStateOfLike(queryString)
}

function ChangeAttributeLike(queryString, isLiked) {

    let postEl = document.querySelector(queryString)

    if (postEl.dataset.likestate == "false") {

        postEl.dataset.likestate = "true"
        postEl.dataset.isliked = isLiked 

    } else if (postEl.dataset.likestate == "true" && isLiked == postEl.dataset.isliked) {

        console.log("je suis bien allé ici")

        postEl.dataset.likestate = "false"
    
    } else if (postEl.dataset.likestate == "true" && postEl.dataset.isliked != isLiked) {

        postEl.dataset.isliked = isLiked
    }

    VerifyStateOfLike(queryString)
}

function VerifyStateOfLike(queryString) {

    let imgDisLike = document.querySelector(`${queryString} .dislike-btn-img`)
    let imgLike = document.querySelector(`${queryString} .like-btn-img`)
    let postEl = document.querySelector(queryString)

    if (postEl.dataset.likestate == "true" && postEl.dataset.isliked == "true") {
              
        imgLike.setAttribute("src", "../static/img/arrow-like-coloring.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike.png")

    } else if (postEl.dataset.likestate == "true" && postEl.dataset.isliked == "false") {
        
        imgLike.setAttribute("src", "../static/img/arrow-like.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike-coloring.png")

    } else if (postEl.dataset.likestate == "false") {
        imgLike.setAttribute("src", "../static/img/arrow-like.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike.png")
    }
}

