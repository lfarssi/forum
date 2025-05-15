




document.addEventListener("DOMContentLoaded", () => {
    const reportForms = document.querySelectorAll(".report-form")

    reportForms.forEach(form => {
        form.addEventListener("submit", async (e) => {
            e.preventDefault();

            const postId = form.querySelector("input[name='post_id']").value;
            const categorySelect = form.querySelector("select[name='category_report_id']");
            const categoryId = categorySelect.value;

            // Validation
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
                    body: formData,  // sending the FormData object directly
                });

                if (response.ok) {
                    alert("Report submitted successfully!");
                    categorySelect.selectedIndex = 0; // Reset
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










document.addEventListener("DOMContentLoaded", function () {
  const viewReportsBtn = document.getElementById('viewReportsBtn');
  const adminModRequestsPopup = document.getElementById('adminModRequests');
  const closePopupBtn = document.getElementById('closePopupBtn');
   const table = document.getElementById('reportedPostsTable');
if (!table || !closePopupBtn || !adminModRequestsPopup || !viewReportsBtn) {
    console.warn("Some elements are missing from the DOM.");
    return;
  }

  const reportedPostsTable = table.getElementsByTagName('tbody')[0];
  if (!reportedPostsTable) {
    console.warn("Tbody not found inside reportedPostsTable.");
    return;
  }

  // Show the popup when the button is clicked
  viewReportsBtn.addEventListener('click', function() {
    // Fetch reported posts
    fetchReportedPosts();
    // Show the popup
    adminModRequestsPopup.style.display = 'block';
  });

  // Close the popup when the close button is clicked
  closePopupBtn.addEventListener('click', function() {
    adminModRequestsPopup.style.display = 'none';
  });

  // Function to fetch reported posts and populate the table
  function fetchReportedPosts() {
    fetch('/get_reported_posts') // Replace with your actual endpoint
      .then(response => response.json())
      .then(posts => {
        // Clear the table body
        reportedPostsTable.innerHTML = '';

        // Loop through the posts and add them to the table
        posts.forEach(post => {
          const row = reportedPostsTable.insertRow();

          const titleCell = row.insertCell(0);
          titleCell.textContent = post.title;

          const categoryCell = row.insertCell(1);
          categoryCell.textContent = post.category;

          const reportDateCell = row.insertCell(2);
          reportDateCell.textContent = post.report_date;

          const statusCell = row.insertCell(3);
          statusCell.textContent = post.status;

          const actionsCell = row.insertCell(4);
          actionsCell.innerHTML = `
            <form action="/delete_report" method="POST">
              <input type="hidden" name="report_id" value="${post.id}">
              <button type="submit">Delete Report</button>
            </form>
            <form action="/delete_post" method="POST">
              <input type="hidden" name="post_id" value="${post.id}">
              <button type="submit">Delete Post</button>
            </form>
          `;
        });
      })
      .catch(error => {
        console.error('Error fetching reported posts:', error);
      });
  }
});
