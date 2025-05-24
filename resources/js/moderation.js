




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
    const reportedPostsTable = document.getElementById('reportedPostsTable').getElementsByTagName('tbody')[0];
    
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

          console.log(posts);
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
            actionsCell.innerHTML = `
              <div class="action-buttons">
                <form action="/delete_report" method="POST">
                  <input type="hidden" name="report_id" value="${post.report_id}">
                  <button type="submit">Delete Report</button>
                </form>
                <form action="/delete_post" method="POST">
                  <input type="hidden" name="post_id" value="${post.post_id}">
                  <button type="submit">Delete Post</button>
                </form>
              </div>
            `;
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














  document.addEventListener("DOMContentLoaded", function() {
    // Existing report form handling
    const reportForms = document.querySelectorAll(".report-form");
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

    // Moderator reported posts handling
    const reportedPostsPopover = document.getElementById('reportedPostsPopover');
    const reportedPostsTable = document.getElementById('reportedPostsTable');
    
    if (reportedPostsPopover && reportedPostsTable) {
        const reportedPostsTableBody = reportedPostsTable.getElementsByTagName('tbody')[0];
        
        reportedPostsPopover.addEventListener('beforetoggle', function(event) {
            if (event.newState === "open") {
                fetchReportedPosts();
            }
        });

        function fetchReportedPosts() {
            fetch('/get_reported_posts')
                .then(response => response.json())
                .then(posts => {
                    reportedPostsTableBody.innerHTML = '';
                    
                    if (!posts || posts.length === 0) {
                        const row = reportedPostsTableBody.insertRow();
                        const cell = row.insertCell(0);
                        cell.colSpan = 6;
                        cell.textContent = "No reported posts found.";
                        cell.style.textAlign = "center";
                        return;
                    }

                    posts.forEach(post => {
                        const row = reportedPostsTableBody.insertRow();
                        
                        row.insertCell(0).textContent = post.title;
                        row.insertCell(1).textContent = post.category;
                        row.insertCell(2).textContent = post.report_reason;
                        row.insertCell(3).textContent = new Date(post.report_date).toLocaleString();
                        row.insertCell(4).textContent = post.status;
                        
                        const actionsCell = row.insertCell(5);
                        
                        // Show different actions based on status
                        if (post.status === "pending") {
                            actionsCell.innerHTML = `
                                <div class="action-buttons">
                                    <span class="pending-status">Waiting for admin approval</span>
                                </div>
                            `;
                        } else if (post.status === "approved") {
                            actionsCell.innerHTML = `
                                <div class="action-buttons">
                                    <form action="/delete_post" method="POST">
                                        <input type="hidden" name="post_id" value="${post.post_id}">
                                        <button type="submit" class="btn-delete-post">Delete Post</button>
                                    </form>
                                </div>
                            `;
                        } else if (post.status === "rejected") {
                            actionsCell.innerHTML = `
                                <div class="action-buttons">
                                    <form action="/delete_report" method="POST">
                                        <input type="hidden" name="report_id" value="${post.report_id}">
                                        <button type="submit" class="btn-delete-report">Delete Report</button>
                                    </form>
                                </div>
                            `;
                        }
                    });
                })
                .catch(error => {
                    console.error('Error fetching reported posts:', error);
                    reportedPostsTableBody.innerHTML = '';
                    const row = reportedPostsTableBody.insertRow();
                    const cell = row.insertCell(0);
                    cell.colSpan = 6;
                    cell.textContent = "Error loading reported posts. Please try again.";
                    cell.style.textAlign = "center";
                });
        }
    }

    // Admin report management
    const adminReportPopover = document.getElementById('adminReportRequests');
    const adminReportsTable = document.getElementById('adminReportsTable');
    
    if (adminReportPopover && adminReportsTable) {
        const adminReportsTableBody = adminReportsTable.getElementsByTagName('tbody')[0];
        
        adminReportPopover.addEventListener('beforetoggle', function(event) {
            if (event.newState === "open") {
                fetchPendingReports();
            }
        });

        function fetchPendingReports() {
            fetch('/get_pending_reports')
                .then(response => response.json())
                .then(reports => {
                    adminReportsTableBody.innerHTML = '';
                    
                    if (!reports || reports.length === 0) {
                        const row = adminReportsTableBody.insertRow();
                        const cell = row.insertCell(0);
                        cell.colSpan = 6;
                        cell.textContent = "No pending reports found.";
                        cell.style.textAlign = "center";
                        return;
                    }

                    reports.forEach(report => {
                        const row = adminReportsTableBody.insertRow();
                        
                        row.insertCell(0).textContent = report.post_title;
                        row.insertCell(1).textContent = report.report_reason;
                        row.insertCell(2).textContent = report.reporter_username;
                        row.insertCell(3).textContent = new Date(report.report_date).toLocaleString();
                        row.insertCell(4).textContent = "Delete Post"; // Requested action
                        
                        const decisionCell = row.insertCell(5);
                        decisionCell.innerHTML = `
                            <div class="admin-decision">
                                <form class="report-decision-form" method="POST" action="/admin_report_decision">
                                    <input type="hidden" name="report_id" value="${report.report_id}">
                                    <select name="decision" required>
                                        <option value="">Choose Action</option>
                                        <option value="approved">Approve (Allow Delete)</option>
                                        <option value="rejected">Reject (Keep Post)</option>
                                    </select>
                                    <button type="submit">Submit Decision</button>
                                </form>
                            </div>
                        `;
                    });
                })
                .catch(error => {
                    console.error('Error fetching pending reports:', error);
                    adminReportsTableBody.innerHTML = '';
                    const row = adminReportsTableBody.insertRow();
                    const cell = row.insertCell(0);
                    cell.colSpan = 6;
                    cell.textContent = "Error loading pending reports. Please try again.";
                    cell.style.textAlign = "center";
                });
        }
    }

    // Handle form submissions with event delegation
    document.addEventListener('submit', function(event) {
        const target = event.target;
        
        // Handle moderator actions
        if (target.action.includes('/delete_report') || target.action.includes('/delete_post')) {
            event.preventDefault();
            const formData = new FormData(target);
            const submitButton = target.querySelector('button[type="submit"]');
            if (submitButton) submitButton.disabled = true;
            
            fetch(target.action, {
                method: 'POST',
                body: formData
            })
            .then(response => {
                if (response.ok) {
                    setTimeout(() => {
                        fetchReportedPosts();
                    }, 500);
                } else {
                    throw new Error('Network response was not ok');
                }
            })
            .catch(error => {
                console.error('Error submitting form:', error);
                alert('An error occurred while processing your request. Please try again.');
                if (submitButton) submitButton.disabled = false;
            });
        }
        
        // Handle admin report decisions
        if (target.classList.contains('report-decision-form')) {
            event.preventDefault();
            const formData = new FormData(target);
            const submitButton = target.querySelector('button[type="submit"]');
            if (submitButton) submitButton.disabled = true;
            
            fetch(target.action, {
                method: 'POST',
                body: formData
            })
            .then(response => {
                if (response.ok) {
                    alert('Decision submitted successfully!');
                    setTimeout(() => {
                        fetchPendingReports();
                    }, 500);
                } else {
                    throw new Error('Network response was not ok');
                }
            })
            .catch(error => {
                console.error('Error submitting decision:', error);
                alert('An error occurred while processing your decision. Please try again.');
                if (submitButton) submitButton.disabled = false;
            });
        }
        
        // Handle moderator requests
        if (target.classList.contains('mod-request-form')) {
            event.preventDefault();
            const formData = new FormData(target);
            
            fetch('/handleRequest', {
                method: 'POST',
                body: formData
            })
            .then(response => {
                if (response.ok) {
                    const selectedRole = formData.get("role");
                    if (selectedRole === "user") {
                        const row = target.closest('tr');
                        if (row) row.remove();
                    }
                } else {
                    const errorText = response.text();
                    console.error(errorText);
                }
            })
            .catch(err => {
                console.error('Fetch error:', err);
                alert('Request failed. Please try again.');
            });
        }
    });

    // Add CSS styles
    const style = document.createElement('style');
    style.textContent = `
        #reportedPostsPopover, #adminReportRequests {
            padding: 20px;
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
            max-width: 95%;
            max-height: 80vh;
            overflow-y: auto;
        }
        
        #reportedPostsTable, #adminReportsTable {
            width: 100%;
            border-collapse: collapse;
            margin-top: 10px;
        }
        
        #reportedPostsTable th, #reportedPostsTable td,
        #adminReportsTable th, #adminReportsTable td {
            padding: 8px;
            border: 1px solid #ddd;
            text-align: left;
        }
        
        #reportedPostsTable th, #adminReportsTable th {
            background-color: #f2f2f2;
            font-weight: bold;
        }
        
        .action-buttons {
            display: flex;
            gap: 5px;
            flex-wrap: wrap;
        }
        
        .action-buttons button {
            padding: 5px 10px;
            cursor: pointer;
            border: none;
            border-radius: 4px;
        }
        
        .btn-delete-post {
            background-color: #dc3545;
            color: white;
        }
        
        .btn-delete-report {
            background-color: #6c757d;
            color: white;
        }
        
        .pending-status {
            color: #ffc107;
            font-style: italic;
            padding: 5px;
        }
        
        .admin-decision select {
            margin-right: 5px;
            padding: 3px;
        }
        
        .admin-decision button {
            padding: 5px 10px;
            background-color: #007bff;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
        }
        
        .admin-reports-container h3 {
            margin-top: 0;
            color: #333;
        }
    `;
    document.head.appendChild(style);
});