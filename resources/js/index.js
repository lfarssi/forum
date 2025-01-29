function toggleComments() {
    const commentBtns = document.querySelectorAll(".comments-btn");
    commentBtns.forEach(btn => {
        const comments = btn.closest('.post').querySelector(".comments");
        comments.style.display = "none"; 

        btn.addEventListener('click', () => {
            if (comments.style.display === "none" || comments.style.display === "") {
                comments.style.display = "block";
            } else {
                comments.style.display = "none";
            }
        });
    });
}

// Call the function when needed
toggleComments();


document.addEventListener("DOMContentLoaded", () => {
    function updateReactionCount(element, newCount, toggleClass) {
        element.textContent = newCount;
        element.closest("button").classList.toggle(toggleClass);
    }
    let likepost = document.querySelectorAll(".likepost")
    likepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const numDislike = post.querySelector(".numdislikepost");
            let updatedDislikes = parseInt(numDislike.textContent);
            const numLike = post.querySelector(".numlikepost");
            let updatedLikes = parseInt(numLike.textContent);
            const oppositeDislikeButton = post.querySelector(".dislikepost");

            if (btn.classList.contains("islikedpost")) {
                updatedLikes -= 1;
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                oppositeDislikeButton.classList.remove("isdislikedpost"); // Remove dislike class
            } else {
                updatedLikes += 1;
                updatedDislikes -= 1;
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                oppositeDislikeButton.classList.remove("isdislikedpost"); // Remove dislike class
            }
            btn.disabled = true;
            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "like")
            data.append("sender", "post")
            try {

                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {
                    console.log("OK");
                    console.log(response.likes);
                } else {
                    console.log("eerr");
                    updateReactionCount(numLike, updatedLikes - 1, "islikedpost");
                }
            } catch  {
                console.log("err like post");
                
            }
            btn.disabled = false;
        })
    })

    //dislike post  
    let dislikepost = document.querySelectorAll(".dislikepost")
    dislikepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "dislike")
            data.append("sender", "post")
            const numDislike = post.querySelector(".numdislikepost");
            let updatedDislikes = parseInt(numDislike.textContent);
            const oppositeLikeButton = post.querySelector(".likepost");
            const numLike = post.querySelector(".numlikepost");
            let updatedLikes = parseInt(numLike.textContent);
            if (btn.classList.contains("isdislikedpost")) {
                updatedDislikes -= 1;
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                oppositeLikeButton.classList.remove("islikedpost"); // Remove like class
            } else {
                updatedDislikes += 1;
                updatedLikes -= 1;
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                oppositeLikeButton.classList.remove("islikedpost"); // Remove like class
            }
            btn.disabled = true;

            try {
                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {

                    console.log("sucess");
                    
                } else {
                    console.log("eerr");
                    updateReactionCount(numDislike, updatedDislikes - 1, "isdislikedpost");
                }
            } catch {
                console.log("err dislike post");
                
            }
            btn.disabled = false;
        })
    })


    // like comment
    let likecomment = document.querySelectorAll(".likecomment")
    likecomment.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            
            const data = new FormData()
            data.append("comment_id", commentId)
            data.append("status", "like")
            data.append("sender", "comment")
            const numDislike = comment.querySelector(".numdislikecomment");
            
            let updatedDislikes = parseInt(numDislike.textContent);
            const numLike = comment.querySelector(".numlikecomment");
            let updatedLikes = parseInt(numLike.textContent);
            const oppositeDislikeButton = comment.querySelector(".dislikecomment");

            if (btn.classList.contains("islikedpost")) {
                updatedLikes -= 1;
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                oppositeDislikeButton.classList.remove("isdislikedpost"); // Remove dislike class
            } else {
                updatedLikes += 1;
                updatedDislikes -= 1;
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                oppositeDislikeButton.classList.remove("isdislikedpost"); // Remove dislike class
            }

            btn.disabled = true;


            try {


                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {
                    console.log("sucess");
                    
                } else {
                    console.log("eerr");
                    updateReactionCount(numLike, updatedLikes - 1, "islikedpost");
                }
            } catch {
                console.log("err like comment");
            }
            btn.disabled = false;
        })
    })


    // dislike comment
    let dislikecomment = document.querySelectorAll(".dislikecomment")
    dislikecomment.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const numDislike = comment.querySelector(".numdislikecomment");
            const numLike = comment.querySelector(".numlikecomment");
            let updatedLikes = parseInt(numLike.textContent);
            let updatedDislikes = parseInt(numDislike.textContent);
             const oppositeLikeButton = comment.querySelector(".likecomment");

            if (btn.classList.contains("isdislikedpost")) {
                updatedDislikes -= 1;
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                oppositeLikeButton.classList.remove("islikedpost"); // Remove like class
            } else {
                updatedDislikes += 1;
                updatedLikes -= 1;
                updateReactionCount(numLike, updatedLikes, "islikedpost");
                updateReactionCount(numDislike, updatedDislikes, "isdislikedpost");
                oppositeLikeButton.classList.remove("islikedpost"); // Remove like class
            }
            btn.disabled = true;  
             const data = new FormData()
            data.append("comment_id", commentId)
            data.append("status", "dislike")
            data.append("sender", "comment")
            try {

                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {
                    console.log("sucess");
                    
                } else {
                    console.log("eerr");
                    updateReactionCount(numDislike, updatedDislikes - 1, "isdislikedpost");
                    
                }
            } catch {
                console.log("err like post");
                
            }
            btn.disabled = false;
        })
    })




});