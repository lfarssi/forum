// Function to toggle the visibility of comments when the comments button is clicked
function toggleComments() {
    const commentBtns = document.querySelectorAll(".comments-btn"); // Select all comment buttons
    commentBtns.forEach(btn => {
        const comments = btn.closest('.post').querySelector(".comments"); // Find the closest comments section in the post
        comments.style.display = "none"; // Initially hide the comments section

        btn.addEventListener('click', () => { // Add event listener for click on the button
            if (comments.style.display === "none" || comments.style.display === "") {
                comments.style.display = "block"; // Show comments if they are hidden
            } else {
                comments.style.display = "none"; // Hide comments if they are visible
            }
        });
    });
}

toggleComments(); // Call the toggleComments function to initialize the behavior

// Function to show flash alerts with custom messages
const showFlashAlert = (message) => {
    const flashAlert = document.getElementById('flashAlert');
    const flashMessage = document.getElementById('flashMessage');

    flashMessage.textContent = message; // Set the message content
    flashAlert.classList.add('show'); // Show the flash alert

    setTimeout(() => {
        flashAlert.classList.remove('show'); // Hide the flash alert after 2 seconds
    }, 2000); 
};

// Handle comment form submission
const commentForm = document.querySelectorAll(".comment-form");
commentForm.forEach((form) => {
    form.addEventListener("submit", async (e) => {
        e.preventDefault() // Prevent form submission from reloading the page
        const post = e.target.closest(".post");
        const postId = post.dataset.id; // Get the post ID
        const content = post.querySelector(".content-comment"); // Get the comment content

        let data = new FormData();
        data.append("post_id", postId);
        data.append("content", content.value); // Append the comment content to the FormData object

        console.log(content);

        if (content.value == "") {
            showFlashAlert("Comment Is Empty"); // Show alert if the comment is empty
            console.log("error");
            return;
        } else if (content.value.length > 10000) {
            showFlashAlert("Comment Too Large"); // Show alert if the comment is too long
            return;
        }
        try {
            const response = await fetch("/create_comment", {
                method: "POST",
                body: data,
            });

            if (response.ok) {
                window.location.href = "/"; // Redirect to the homepage if successful
            } else {
                const data = await response.json();
                showFlashAlert(data); // Show an alert with the error message if the response is not OK
            }
        } catch {
            console.log("error from front"); // Log error if fetch fails
        }
    });
});

// Handle post form submission
const postForm = document.querySelector('#formPost');
postForm.addEventListener("submit", async (e) => {
    e.preventDefault(); // Prevent page reload on form submission
    const title = postForm.querySelector(".title");
    const categories = postForm.querySelectorAll(".categories:checked"); // Get all checked categories
    const content = postForm.querySelector(".content");

    let data = new FormData();
    data.append("title", title.value);
    data.append("content", content.value);

    categories.forEach(category => {
        console.log(category.value); 
        data.append("categories", category.value); // Append selected categories to FormData
    });

    if (content.value == "" || title.value == "") {
        showFlashAlert("Post Info Is Empty"); // Show alert if any required field is empty
        return;
    } else if (content.value.length > 10000 || title.value.length > 255) {
        showFlashAlert("Post Info Too Large"); // Show alert if any field exceeds the length limit
        return;
    }

    try {
        const response = await fetch("/create_post", {
            method: "POST",
            body: data,
        });

        if (response.ok) {
            window.location.href = "/"; // Redirect to the homepage if successful
        } else {
            const data = await response.json();
            showFlashAlert(data); // Show alert with error message if response is not OK
        }
    } catch {
        window.location.href = "/login"; // Redirect to login if fetch fails
    }
});

// Event listener for post reactions (like/dislike)
document.addEventListener("DOMContentLoaded", () => {
    // Like post
    let likepost = document.querySelectorAll(".likepost");
    likepost.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const post = e.target.closest('.post');
            const postId = post.dataset.id; // Get the post ID
            const numDislike = post.querySelector(".numdislikepost");
            const numLike = post.querySelector(".numlikepost");
            const dislikeBtn = post.querySelector(".dislikepost");

            const data = new FormData();
            data.append("post_id", postId);
            data.append("status", "like");
            data.append("sender", "post");

            try {
                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                });

                if (response.ok) {
                    const result = await response.json();
                    result.forEach((res) => {
                        if (res.id == postId) {
                            numLike.textContent = res.likes;
                            numDislike.textContent = res.dislikes;
                            if (res.IsLiked) {
                                btn.classList.add("isReacted");
                                dislikeBtn.classList.remove("isReacted");
                            } else {
                                btn.classList.remove("isReacted");
                            }
                        }
                    });

                } else {
                    console.log("error in like post");
                }
            } catch (e) {
                window.location.href = "/login"; // Redirect to login on error
            }
        });
    });

    // Dislike post
    let dislikepost = document.querySelectorAll(".dislikepost");
    dislikepost.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const post = e.target.closest('.post');
            const postId = post.dataset.id; // Get the post ID
            const data = new FormData();
            data.append("post_id", postId);
            data.append("status", "dislike");
            data.append("sender", "post");

            const numDislike = post.querySelector(".numdislikepost");
            const likeBtn = post.querySelector(".likepost");
            const numLike = post.querySelector(".numlikepost");

            try {
                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                });

                if (response.ok) {
                    const result = await response.json();
                    result.forEach((res) => {
                        if (res.id == postId) {
                            numLike.textContent = res.likes;
                            numDislike.textContent = res.dislikes;
                            if (res.IsDisliked) {
                                btn.classList.add("isReacted");
                                likeBtn.classList.remove("isReacted");
                            } else {
                                btn.classList.remove("isReacted");
                            }
                        }
                    });
                } else {
                    console.log("error in dislike post");
                }
            } catch {
                console.log("error in dislike post");
                window.location.href = "/login"; // Redirect to login on error
            }
        });
    });

    // Like comment
    let likecomment = document.querySelectorAll(".likecomment");
    likecomment.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const post = e.target.closest('.post');
            const postId = post.dataset.id;

            const data = new FormData();
            data.append("comment_id", commentId);
            data.append("status", "like");
            data.append("sender", "comment");

            const numLike = comment.querySelector(".numlikecomment");
            const numDislike = comment.querySelector(".numdislikecomment");
            const dislikebtn = comment.querySelector(".dislikecomment");

            try {
                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                });

                if (response.ok) {
                    const result = await response.json();
                    result.forEach((res) => {
                        if (res.id == postId) {
                            if (Array.isArray(res.comments)) {
                                res.comments.forEach((comm) => {
                                    if (comm.id == commentId) {
                                        numLike.textContent = comm.likes;
                                        numDislike.textContent = comm.dislikes;
                                        if (comm.IsLiked) {
                                            btn.classList.add("isReacted");
                                            dislikebtn.classList.remove("isReacted");
                                        } else {
                                            btn.classList.remove("isReacted");
                                        }
                                    }
                                });
                            }
                        }
                    });
                } else {
                    console.log("error in like comment");
                }
            } catch {
                console.log("error in like comment");
                window.location.href = "/login"; // Redirect to login on error
            }
        });
    });

    // Dislike comment
    let dislikecomment = document.querySelectorAll(".dislikecomment");
    dislikecomment.forEach(btn => {
        btn.addEventListener("click", async (e) => {
            const comment = e.target.closest('.comment');
            const commentId = comment.dataset.id;
            const post = e.target.closest('.post');
            const postId = post.dataset.id;

            const data = new FormData();
            data.append("comment_id", commentId);
            data.append("status", "dislike");
            data.append("sender", "comment");

            const numDislike = comment.querySelector(".numdislikecomment");
            const numLike = comment.querySelector(".numlikecomment");
            const likebtn = comment.querySelector(".likecomment");

            try {
                const response = await fetch("/react", {
                    method: "POST",
                    body: data,
                });

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
                                            btn.classList.add("isReacted");
                                            likebtn.classList.remove("isReacted");
                                        } else {
                                            btn.classList.remove("isReacted");
                                        }
                                    }
                                });
                            }
                        }
                    });
                } else {
                    console.log("error in dislike comment");
                }
            } catch {
                console.log("error in dislike comment");
                window.location.href = "/login"; // Redirect to login on error
            }
        });
    });
});
