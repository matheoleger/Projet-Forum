let posts = document.querySelectorAll(".posts")

// cette boucle permet de remettre les flèches "like/dislike" à la bonne valeur pour chaque post
for(let post of posts) {

    let queryString = `div#${CSS.escape(post.id)}.posts`

    VerifyStateOfLike(queryString)
}

// cette fonction est appelé lors du clique sur un like/dislike d'un post
function LikedPost(id, isLiked) {
    fetch(`/like?post=${id}&isLiked=${isLiked}`)
    .then(response => response.json())
    .then((resp) => {
        changeAttributeLike(resp, id, "posts")
    })

}

// cette fonction est appelé lors du clique sur un like/dislike d'un commentaire
function LikedComment(id, isLiked) {
    fetch(`/like?comment=${id}&isLiked=${isLiked}`)
    .then(response => response.json())
    .then((resp) => {
        changeAttributeLike(resp, id, "comments")
    })

}

// cette fonction permet de changer les attributs liés au like : 
// valeurs des "data-likestates" et "data-isliked", ainsi que le nombre de like
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

// cette fonction permet de modifier la flèche en fonction de l'état du like
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

