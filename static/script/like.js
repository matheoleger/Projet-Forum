let isAlreadyLiked = false;
let isAlreadyDisLiked = false;

function LikedPost(id, isLiked) {
    fetch(`/like?post=${id}&isLiked=${isLiked}`)

    if (!isAlreadyLiked) {
        isAlreadyLiked = true;
    } else {
        isAlreadyLiked = false;
    }
}

function LikedComment(id, isLiked) {
    fetch(`/like?comment=${id}&isLiked=${isLiked}`)

    if (!isAlreadyDisLiked) {
        isAlreadyDisLiked = true;
    } else {
        isAlreadyDisLiked = false;
    }
}

