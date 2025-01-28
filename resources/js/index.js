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

    const numlikecomment = document.querySelectorAll(".numlikecomment")
    const numdislikecomment = document.querySelectorAll(".numdislikecomment")
    
    // like post
    let likepost = document.querySelectorAll(".likepost")
    likepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const post = e.target.closest('.post')
            const postId = post.dataset.id
          const numlikepost = post.querySelector(".numlikepost")
          const dislikebtn = post.querySelector(".dislikepost")
            
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
                    const reponse = await response.json()
                    console.log(response.likes);
                    
                    // btn.classList.toggle("likedpost")
                    // const isLiked = btn.classList.contains("likedpost")
                    // numlikepost.textContent = parseInt(numlikepost.textContent) + (isLiked ? 1 : -1);
                  
                    // if (dislikebtn.classList.contains("dislikepost")) {
                    //    const numdislikepost = post.querySelector("numdislikepost")
                    //    dislikebtn.classList.remove("dislikepost")
                    //    numdislikepost.textContent = parseInt(numdislikepost.textContent) - 1;

                    // }
                
                    
                } else {
                    console.log("eerr");
                    
                }
            } catch {
                console.log("err like post");
                
            }
        })
    })

    //dislike post  
    let dislikepost = document.querySelectorAll(".dislikepost")
    dislikepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const postId = e.target.closest('.post').dataset.id
            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "dislike")
            data.append("sender", "post")
            try {
                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {

                    console.log("sucess");
                    
                } else {
                    console.log("eerr");
                    
                }
            } catch {
                console.log("err dislike post");
                
            }
        })
    })


    // like comment
    let likecomment = document.querySelectorAll(".likecomment")
    likecomment.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const commentid = e.target.closest('.comment').dataset.id
            console.log(commentid);
            
            const data = new FormData()
            data.append("comment_id", commentid)
            data.append("status", "like")
            data.append("sender", "comment")
            try {
console.log("fdsfs");

                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {
                    console.log("sucess");
                    
                } else {
                    console.log("eerr");
                    
                }
            } catch {
                console.log("err like comment");
            }
        })
    })


    // dislike comment
    let dislikecomment = document.querySelectorAll(".dislikecomment")
    dislikecomment.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const commentid = e.target.closest('.comment').dataset.id
            const data = new FormData()
            data.append("comment_id", commentid)
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
                    
                }
            } catch {
                console.log("err like post");
                
            }
        })
    })




});