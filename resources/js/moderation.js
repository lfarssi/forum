




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

          // const categoryCell = row.insertCell(1);
          // categoryCell.textContent = post.category;

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



document.addEventListener("DOMContentLoaded", function() {
    const reportedPostsPopover = document.getElementById('reportedPostsPopover');
  //   const reportedPostsTable = document.getElementById('reportedPostsTable').getElementsByTagName('tbody')[0];
  //    if (!reportedPostsTable) {
  //   console.warn("Tbody not found inside reportedPostsTable.");
  //   return;
  // }


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

    // Add CSS for the popover if it's shown programmatically
    const style = document.createElement('style');
    style.textContent = `
      #reportedPostsPopover {
        padding: 20px;
        background: white;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
        max-width: 90%;
        max-height: 80vh;
        overflow-y: auto;
      }
      
      #reportedPostsTable {
        width: 100%;
        border-collapse: collapse;
      }
      
      #reportedPostsTable th, #reportedPostsTable td {
        padding: 8px;
        border: 1px solid #ddd;
        text-align: left;
      }
      
      #reportedPostsTable th {
        background-color: #f2f2f2;
      }
      
      .action-buttons {
        display: flex;
        gap: 5px;
      }
      
      .action-buttons button {
        padding: 5px 10px;
        cursor: pointer;
      }
    `;
    document.head.appendChild(style);

    // Event listener for the popover showing
    reportedPostsPopover.addEventListener('beforetoggle', function(event) {
      if (event.newState === "open") {
        fetchReportedPosts();
      }
    });

    // Function to fetch reported posts and populate the table
    function fetchReportedPosts() {
      fetch('/get_reported_posts')
        .then(response => {
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          return response.json();
        })
        .then(posts => {
          // Clear the table body
          reportedPostsTable.innerHTML = '';

          if (posts==null) {
            const row = reportedPostsTable.insertRow();
            const cell = row.insertCell(0);
            cell.colSpan = 5;
            cell.textContent = "No reported posts found.";
            cell.style.textAlign = "center";
            return;
          }
          

          // Loop through the posts and add them to the table
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
            //       <input type="hidden" name="post_id" value="${post.post_id}">
            //       <button type="submit">Delete Post</button>
            //     </form>
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

    // Add event delegation for form submissions
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









