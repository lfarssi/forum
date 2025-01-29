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
  
    const cooldownTime = 1000;
    let likepost = document.querySelectorAll(".likepost")
    likepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const numDislike = post.querySelector(".numdislikepost");
            const numLike = post.querySelector(".numlikepost");
            btn.disabled = true;

            setTimeout(() => {
                btn.disabled = false;
            }, cooldownTime);

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
                    const result = await response.json()
                    console.log(result);
                    result.forEach((res)=>{
                        if(res.id==postId){
                            numLike.textContent=res.likes
                            numDislike.textContent=res.dislikes
                        
                        }
                    })
                    
                    console.log("OK");
                } else {
                    console.log("eerr");
         
                }
            } catch(e)  {
               console.log(e);
               
                
            }
            
        })
    })

    //dislike post  
    let dislikepost = document.querySelectorAll(".dislikepost")
    dislikepost.forEach(btn=>{
        btn.addEventListener("click", async (e) =>{
            btn.disabled = true;

            setTimeout(() => {
                btn.disabled = false;
            }, cooldownTime);
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const data = new FormData()
            data.append("post_id", postId)
            data.append("status", "dislike")
            data.append("sender", "post")
            const numDislike = post.querySelector(".numdislikepost");

            const numLike = post.querySelector(".numlikepost");
          
         
           

            try {
                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {

                    const result = await response.json()
                    console.log(result);
                    result.forEach((res)=>{
                        if(res.id==postId){
                            numLike.textContent=res.likes
                            numDislike.textContent=res.dislikes
                        
                        }
                    })
                    
                    console.log("OK");
                    
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
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            btn.disabled = true;
            const post = e.target.closest('.post')
            const postId = post.dataset.id

            setTimeout(() => {
                btn.disabled = false;
            }, cooldownTime);
            const data = new FormData()
            data.append("comment_id", commentId)
            data.append("status", "like")
            data.append("sender", "comment")
            
            const numLike = comment.querySelector(".numlikecomment");
            const numDislike = comment.querySelector(".numdislikecomment");

            try {


                const response = await fetch("/react",{
                    method: "POST", 
                    body: data,
                })
                if (response.ok) {
                    const result = await response.json()
                    result.forEach((res)=>{
                        if(res.id==postId){
                            if (Array.isArray(res.comments)){
                                res.comments.forEach((comm)=>{
                                    
                                    if(comm.id==commentId){
                                        console.log(comm);
                                        console.log(comm.dislikes);
                                        numLike.textContent=comm.likes
                                        numDislike.textContent=comm.dislikes
                                    }
                                })
                            }
                        }
                    })
                    
                    console.log("OK");
                    
                } else {
                    console.log("eerr");
                   // updateReactionCount(numLike, updatedLikes - 1, "islikedpost");
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
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const post = e.target.closest('.post')
            const postId = post.dataset.id
            const numDislike = comment.querySelector(".numdislikecomment");
            const numLike = comment.querySelector(".numlikecomment");

             btn.disabled = true;

             setTimeout(() => {
                 btn.disabled = false;
             }, cooldownTime);
            
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
                    const result = await response.json();
                    console.log("Full API Response:", result); // Log the full response
    
                    result.forEach((res) => {
                        console.log("Checking post:", res.id, "Post ID:", postId);
                        if (res.id == postId) {
                            console.log("Inside post:", res);
    
                            if (Array.isArray(res.comments)) {
                                console.log("Looping through comments...");
    
                                res.comments.forEach((comm) => {
                                    console.log("Checking comment ID:", comm.id, "Looking for:", commentId);
    
                                    if (comm.id == commentId) {
                                        console.log("Found matching comment:", comm);
    
                                        numLike.textContent = comm.likes;
                                        numDislike.textContent = comm.dislikes;
                                    }
                                });
                            } else {
                                console.log("res.comments is not an array:", res.comments);
                            }
                        }
                    });
    
                    console.log("OK");
                } else {
                    console.log("eerr");
                  //  updateReactionCount(numDislike, updatedDislikes - 1, "isdislikedpost");
                    
                }
            } catch {
                console.log("err ldislike comment");
                
            }
           
        })
    })




});