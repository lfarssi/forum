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
