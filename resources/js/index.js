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
    const numlikepost = document.querySelectorAll(".numlikepost")
    const numdislikepost = document.querySelectorAll(".numlikepost")
    const numlikepost = document.querySelectorAll(".numlikepost")
    const numlikepost = document.querySelectorAll(".numlikepost")
    
    // like post
    let likepost = document.querySelectorAll(".likepost")
    likepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const postId = e.target.closest('.post').dataset.id
            
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
                    console.log("sucess");
                    
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