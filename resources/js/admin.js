

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




document.addEventListener('submit', async function (e) {
  
  const form = e.target.closest('.report-post-admin');
  if (!form) return; 

  e.preventDefault(); 

  const formData = new FormData(form);
  console.log([...formData.entries()]); 

  const submitButton = form.querySelector('button[type="submit"]');
  if (submitButton) submitButton.disabled = true; 

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
      alert("Request sent successfully");
    } else {
      if (submitButton) submitButton.disabled = false; 
      const errorText = await res.text();
      console.error(errorText);
    }
  } catch (err) {
    console.error('Fetch error:', err);
    alert('Request failed. Please try again.');
    if (submitButton) submitButton.disabled = false; 
  }
});







document.addEventListener('submit', async function (e) {
  
  const form = e.target.closest('.addCategoryForm');
  if (!form) return; 

  e.preventDefault(); 

  const formData = new FormData(form);
  console.log([...formData.entries()]); 

  const submitButton = form.querySelector('button[type="submit"]');
  if (submitButton) submitButton.disabled = true; 

  try {
    const res = await fetch('/add-categorie-report', {
      method: 'POST',
      body: formData
    });

    if (res.ok) {
      const selectedRole = formData.get("role");
      if (selectedRole === "user") {
        const row = form.closest('tr');
        if (row) row.remove();
      }
      alert("Request sent successfully");
    } else {
      if (submitButton) submitButton.disabled = false; 
      const errorText = await res.text();
      console.error(errorText);
    }
  } catch (err) {
    console.error('Fetch error:', err);
    alert('Request failed. Please try again.');
    if (submitButton) submitButton.disabled = false; 
  }
});

document.addEventListener('submit', async function (e) {
  
  const form = e.target.closest('.delete-category-form');
  if (!form) return; 

  e.preventDefault(); 

  const formData = new FormData(form);
  console.log([...formData.entries()]); 

  const submitButton = form.querySelector('button[type="submit"]');
  if (submitButton) submitButton.disabled = true; 

  try {
    const res = await fetch('/delete-categorie-report', {
      method: 'POST',
      body: formData
    });

    if (res.ok) {
      alert("Request sent successfully");
    } else {
      if (submitButton) submitButton.disabled = false; 
      const errorText = await res.text();
      console.error(errorText);
    }
  } catch (err) {
    console.error('Fetch error:', err);
    alert('Request failed. Please try again.');
    if (submitButton) submitButton.disabled = false; 
  }
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

   
    reportedPostsPopover.addEventListener('beforetoggle', function(event) {
      if (event.newState === "open") {
        fetchReportedPosts();
      }
    });

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







