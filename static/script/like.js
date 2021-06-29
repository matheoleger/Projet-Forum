let posts = document.querySelectorAll(".posts")

for(let post of posts) {

    let queryString = `div#${CSS.escape(post.id)}.posts`

    VerifyStateOfLike(queryString)
}


function LikedPost(id, isLiked) {
    fetch(`/like?post=${id}&isLiked=${isLiked}`)
    .then(response => response.json())
    .then((resp) => {
        changeAttributeLike(resp, id, "posts")
    })

}

function LikedComment(id, isLiked) {
    fetch(`/like?comment=${id}&isLiked=${isLiked}`)
    .then(response => response.json())
    .then((resp) => {
        changeAttributeLike(resp, id, "comments")
    })

}

function changeAttributeLike(resp, id, elementType) {
    let queryString = `div#${CSS.escape(id)}.${elementType}`
    console.log(queryString)
    let element = document.querySelector(queryString)

    // let btnForLike = element.querySelector('.class-for-like')
    let nbrLikeElement = element.querySelector('.btn-for-like > h2')

    console.log(resp)
    
    element.dataset.likestate = resp.LikeState
    element.dataset.isliked = resp.IsLiked
    nbrLikeElement.textContent = resp.Number_like

    VerifyStateOfLike(queryString)
}

function VerifyStateOfLike(queryString) {

    let imgDisLike = document.querySelector(`${queryString} .dislike-btn-img`)
    let imgLike = document.querySelector(`${queryString} .like-btn-img`)
    let element = document.querySelector(queryString)

    if (element.dataset.likestate == "true" && element.dataset.isliked == "true") {
              
        imgLike.setAttribute("src", "../static/img/arrow-like-coloring.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike.png")

    } else if (element.dataset.likestate == "true" && element.dataset.isliked == "false") {
        
        imgLike.setAttribute("src", "../static/img/arrow-like.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike-coloring.png")

    } else if (element.dataset.likestate == "false") {
        imgLike.setAttribute("src", "../static/img/arrow-like.png")
        imgDisLike.setAttribute("src", "../static/img/arrow-dislike.png")
    }
}

