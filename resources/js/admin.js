


  document.querySelectorAll('.mod-request-form').forEach(form => {
    form.addEventListener('submit', async function (e) {
      e.preventDefault();
      const formData = new FormData(form);

      try {
        const res = await fetch('/handleRequest', {
          method: 'POST',
          body: formData
        });

        if (res.ok) {
          const selectedRole = formData.get("role");
          
          // Remove row only if role is 'user' (refused)
          if (selectedRole === "user") {
            const row = form.closest('tr');
            if (row) row.remove();
          }
          alert("request sent succ")
        } else {
          const errorText = await res.text();
          console.error(errorText);
        }
      } catch (err) {
        console.error('Fetch error:', err);
        alert('Request failed. Please try again.');
      }
    });
  });





    document.querySelectorAll('.report-post-admin').forEach(form => {
      const select = form.querySelector('select[name="desicion"]');
  const submitButton = form.querySelector('button[type="submit"]');
  submitButton.disabled = true;

  select.addEventListener('change', () => {
    if (select.value !== "") {
      submitButton.disabled = false;
    }
  });
    form.addEventListener('submit', async function (e) {

      
      e.preventDefault();
      const formData = new FormData(form);
submitButton.disabled = true;
      try {
        const res = await fetch('/repot-post-responce', {
          method: 'POST',
          body: formData
        });
       
        
        if (res.ok) {
          const selectedRole = formData.get("role");
          
          
          if (selectedRole === "user") {
            const row = form.closest('tr');
            if (row) row.remove();
          }
          alert("request sent succ")
        } else {
          submitButton.disabled = false;
          const errorText = await res.text();
          console.error(errorText);
        }
      } catch (err) {
        console.error('Fetch error:', err);
        alert('Request failed. Please try again.');
      }
    });
  });






document.addEventListener("DOMContentLoaded", function() {
    const reportedPostsPopover = document.getElementById('reportedPostsPopover-admin');



 const table = document.getElementById('reportedPostsTable-admin');
if (!table ) {
    console.warn("Some elements are missing from the DOM.");
    return;
  }

  const reportedPostsTable = table.getElementsByTagName('tbody')[0];
  if (!reportedPostsTable) {
    console.warn("Tbody not found inside reportedPostsTable-admin.");
    return;
  }

    // Add CSS for the popover if it's shown programmatically
    const style = document.createElement('style');
    style.textContent = `
      #reportedPostsPopover-admin {
        padding: 20px;
        background: white;
        border-radius: 8px;
        box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
        max-width: 90%;
        max-height: 80vh;
        overflow-y: auto;
      }
      
      #reportedPostsTable-admin {
        width: 100%;
        border-collapse: collapse;
      }
      
      #reportedPostsTable-admin th, #reportedPostsTable-admin td {
        padding: 8px;
        border: 1px solid #ddd;
        text-align: left;
      }
      
      #reportedPostsTable-admin th {
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

            const desicion = row.insertCell(4);
            desicion.innerHTML =`
            <div>
            <form class="report-post-admin">
                <input type="hidden" name="report-id" value="${post.report_id}">
                <select name="desicion">
                  <option value="" selected disabled>option</option>
                  <option value="rejected">Refuse report</option>
                  <option value="accepted">Accept report</option>
                </select>
                <button type="submit">Submit</button>
              </form>
            </div>
            `;
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







