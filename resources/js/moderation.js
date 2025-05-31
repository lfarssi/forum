




document.addEventListener("DOMContentLoaded", () => {
    const reportForms = document.querySelectorAll(".report-form")

    reportForms.forEach(form => {
        form.addEventListener("submit", async (e) => {
            e.preventDefault();

            const postId = form.querySelector("input[name='post_id']").value;
            const categorySelect = form.querySelector("select[name='category_report_id']");
            const categoryId = categorySelect.value;

            if (!categoryId) {
                alert("Please select a report reason before submitting.");
                categorySelect.focus();
                return;
            }

            const formData = new FormData();
            formData.append("post_id", postId);
            formData.append("category_report_id", categoryId);
            console.log(formData);
            
            try {
                const response = await fetch("/report_post", {
                    method: "POST",
                    body: formData,  
                });

                if (response.ok) {
                    alert("Report submitted successfully!");
                    categorySelect.selectedIndex = 0; 
                    window.location.href = "/"; 
                } else {
                    const errorText = await response.text();
                   console.log(errorText);
                   
                }

            } catch (err) {
                console.error("Error reporting post:", err);
                alert("An error occurred. Please try again.");
            }
        });
    });
});



document.addEventListener('submit', function (event) {
  const target = event.target;

  if (target.classList.contains('delete-post-form')) {
    event.preventDefault();

    
    const formData = new FormData(target);

 
    const submitButton = target.querySelector('button[type="submit"]');
    if (submitButton) submitButton.disabled = true;

    fetch('/delete_post', {
      method: 'POST',
      body: formData
    })
    .then(response => {
      if (!response.ok) {
        throw new Error('Failed to delete post');
      }
      
      alert("post deleted succ")
      const postElement = target.closest('.post');
      if (postElement) {
        postElement.remove();
      }
    })
    .catch(error => {
      console.error('Error deleting post:', error);
      alert('An error occurred while deleting the post.');
      if (submitButton) submitButton.disabled = false;
    });
  }
});


document.addEventListener("DOMContentLoaded", function() {
    const reportedPostsPopover = document.getElementById('reportedPostsPopover');



 const table = document.getElementById('reportedPostsTable');
if (!table ) {
    console.warn("Some elements are missing from the DOM.");
    return;
  }

  const reportedPostsTable = table.getElementsByTagName('tbody')[0];
  if (!reportedPostsTable) {
    console.warn("Tbody not found inside reportedPostsTable.");
    return;
  }

   

    reportedPostsPopover.addEventListener('beforetoggle', function(event) {
      if (event.newState === "open") {
        fetchReportedPosts();
      }
    });

    // Function to fetch reported posts 
    function fetchReportedPosts() {
      fetch('/get_reported_posts')
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(posts => {
        
          reportedPostsTable.innerHTML = '';

          if (posts==null) {
            const row = reportedPostsTable.insertRow();
            const cell = row.insertCell(0);
            cell.colSpan = 5;
            cell.textContent = "No reported posts found.";
            cell.style.textAlign = "center";
            return;
          }
          

     
          posts.forEach(post => {
          console.log(post.status);

            const row = reportedPostsTable.insertRow();

            const titleCell = row.insertCell(0);
            titleCell.textContent = post.title;

            const categoryCell = row.insertCell(1);
            categoryCell.textContent = post.category;
            const reason = row.insertCell(2);
            reason.textContent = post.report_reason;

            const reportDateCell = row.insertCell(3);
            reportDateCell.textContent = new Date(post.report_date).toLocaleString();

            const statusCell = row.insertCell(4);
            statusCell.textContent = post.status;
            const actionsCell = row.insertCell(5);
            if(post.status=="pending"){
              actionsCell.innerHTML="no responce from admin"
            }else if(post.status=="rejected"){
              actionsCell.innerHTML = `
              <div class="action-buttons">
                <form action="/delete_report" method="POST">
                  <input type="hidden" name="report_id" value="${post.report_id}">
                  <button type="submit">Delete Report</button>
                </form>
              </div>
            `;
            }else if(post.status=="accepted"){
              actionsCell.innerHTML = `
              <div class="action-buttons">
                <form action="/delete_post" method="POST">
                  <input type="hidden" name="post_id" value="${post.post_id}">
                   <button type="submit">Delete Post</button>
                 </form>
              </div>
            `;
            }
          });
        })
        .catch(error => {
          console.error('Error fetching reported posts:', error);
          reportedPostsTable.innerHTML = '';
          const row = reportedPostsTable.insertRow();
          const cell = row.insertCell(0);
          cell.colSpan = 5;
          cell.textContent = "Error loading reported posts. Please try again.";
          cell.style.textAlign = "center";
        });
    }

   document.addEventListener('submit', function(event) {
  const target = event.target;

  if (target.action.includes('/delete_report') || target.action.includes('/delete_post')) {
    event.preventDefault();

    console.log('Submitting form:', target.action);

    const formData = new FormData(target);
console.log('Form data contents:');
    for (let [key, value] of formData.entries()) {
      console.log(key, value);
    }
    // Disable the button to prevent double submission
    const submitButton = target.querySelector('button[type="submit"]');
    if (submitButton) submitButton.disabled = true;
    console.log(formData);
    
    fetch(target.action, {
      method: 'POST',
      body: formData
    })  
    .then(response => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
     setTimeout(() => {
        fetchReportedPosts();
      }, 1500);
    })
    .catch(error => {
      console.error('Error submitting form:', error);
      alert('An error occurred while processing your request. Please try again.');
    });
  }
});
  });









