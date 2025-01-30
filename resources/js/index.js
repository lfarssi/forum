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

toggleComments();




const commentForm = document.querySelectorAll(".comment-form");
commentForm.forEach((btn)=>{
    btn.addEventListener("click", async (e)=>{
        const post = e.target.closest(".post");
        const postId =post.dataset.id;
        const content = post.querySelector(".content")
        let data = new FormData()
        data.append("post_id", postId)
        data.append("content", content.value)
        if (content.value == "") {
            alert("Comment is empty")
            return;
        } else if (content.value.length > 10000) {
            alert("Comment is too long")
            return
        }
        try{
            const response = await fetch("/create_comment",{
                method: "POST",
                body: data,
            })
            if (response.ok) {
                window.location.href = "/" 
            } else {
                if (response.error) {
                    console.log(response.error);
                }
            }
        }catch{

        }
    });
})


const postForm = document.querySelector('post-form');
postForm.addEventListener("submit", async (e)=>{

    e.preventDefault();
    const title = postForm.querySelector(".title")
    const categories = postForm.querySelectorAll(".categories")
    const content = postForm.querySelector(".content")
    let data = new FormData()
    data.append("title", title.value)
    data.append("categories", categories.value)
    data.append("content", content.value)
    if (content.value == "") {
        alert("Post is empty")
        return;
    } else if (content.value.length > 10000) {
        alert("Post is too long")
        return
    }
    try{
        const response = await fetch("/create_post",{
            method: "POST",
            body: data,
        })
        if (response.ok) {
            window.location.href = "/" 
        } else {
            if (response.error) {
                console.log(response.error);
            }
        }
    }catch{

    }
})



document.addEventListener("DOMContentLoaded", () => {

    let likepost = document.querySelectorAll(".likepost")
    likepost.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const numDislike = post.querySelector(".numdislikepost");
            const numLike = post.querySelector(".numlikepost");
            const dislikeBtn = post.querySelector(".dislikepost");

            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "like")
            data.append("sender", "post")
            try {

                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                })
                if (response.ok) {
                    const result = await response.json()
                    result.forEach((res) => {
                        if (res.id == postId) {
                            numLike.textContent = res.likes
                            numDislike.textContent = res.dislikes
                            if (res.IsLiked) {
                                btn.classList.add("isReacted")
                                dislikeBtn.classList.remove("isReacted");
                            } else {
                                btn.classList.remove("isReacted")
                            }


                        }
                    })

                } else {
                    console.log("eerr");

                }
            } catch(e) {
                window.location.href= "/login"

            }

        })
    })

    //dislike post  
    let dislikepost = document.querySelectorAll(".dislikepost")
    dislikepost.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "dislike")
            data.append("sender", "post")
            const numDislike = post.querySelector(".numdislikepost");
            const likeBtn = post.querySelector(".likepost");
            const numLike = post.querySelector(".numlikepost");




            try {
                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                })
                if (response.ok) {

                    const result = await response.json()
                    result.forEach((res) => {
                        if (res.id == postId) {
                            numLike.textContent = res.likes
                            numDislike.textContent = res.dislikes
                            if (res.IsDisliked) {
                                btn.classList.add("isReacted")
                                likeBtn.classList.remove("isReacted")
                            } else {
                                btn.classList.remove("isReacted")
                            }

                        }
                    })


                } else {
                    console.log("eerr");
                }
            } catch {
                console.log("err dislike post");
                window.location.href= "/login"


            }

        })
    })


    // like comment
    let likecomment = document.querySelectorAll(".likecomment")
    likecomment.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const post = e.target.closest('.post')
            const postId = post.dataset.id

            const data = new FormData()
            data.append("comment_id", commentId)
            data.append("status", "like")
            data.append("sender", "comment")

            const numLike = comment.querySelector(".numlikecomment");
            const numDislike = comment.querySelector(".numdislikecomment");
            const dislikebtn = comment.querySelector(".dislikecomment")

            try {


                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                })
                if (response.ok) {
                    const result = await response.json()
                    result.forEach((res) => {
                        if (res.id == postId) {
                            if (Array.isArray(res.comments)) {
                                res.comments.forEach((comm) => {
                                    if (comm.id == commentId) {
                                        numLike.textContent = comm.likes
                                        numDislike.textContent = comm.dislikes
                                        if (comm.IsLiked) {
                                            btn.classList.add("isReacted")
                                            dislikebtn.classList.remove("isReacted");
                                        } else {
                                            btn.classList.remove("isReacted")
                                        }
                                    }
                                })
                            }
                        }
                    })


                } else {
                    console.log("eerr");
                    // updateReactionCount(numLike, updatedLikes - 1, "islikedpost");
                }
            } catch {
                console.log("err like comment");
                window.location.href= "/login"

            }

        })
    })


    // dislike comment
    let dislikecomment = document.querySelectorAll(".dislikecomment")
    dislikecomment.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const numDislike = comment.querySelector(".numdislikecomment");
            const numLike = comment.querySelector(".numlikecomment");
            const likebtn = comment.querySelector(".likecomment")

            const data = new FormData()
            data.append("comment_id", commentId)
            data.append("status", "dislike")
            data.append("sender", "comment")
            try {

                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                })
                if (response.ok) {
                    const result = await response.json();

                    result.forEach((res) => {
                        if (res.id == postId) {

                            if (Array.isArray(res.comments)) {

                                res.comments.forEach((comm) => {

                                    if (comm.id == commentId) {

                                        numLike.textContent = comm.likes;
                                        numDislike.textContent = comm.dislikes;
                                        if (comm.IsDisliked) {
                                            btn.classList.add("isReacted")
                                            likebtn.classList.remove("isReacted");
                                        } else {
                                            btn.classList.remove("isReacted")
                                        }
                                    }
                                });
                            } else {
                                console.log("res.comments is not an array:", res.comments);
                            }
                        }
                    });

                } else {
                    console.log("eerr");
                    //  updateReactionCount(numDislike, updatedDislikes - 1, "isdislikedpost");

                }
            } catch {
                console.log("err ldislike comment");
                window.location.href= "/login"


            }
        })
    })

});