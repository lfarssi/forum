


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

function fetchReportedPosts() {
  fetch("/get_reported_posts", {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  })
  .then(response => {
    if (!response.ok) throw new Error("Failed to fetch reports.");
    return response.json();
  })
  .then(data => {
    const tableBody = document.querySelector("#reportedPostsTable tbody");
    tableBody.innerHTML = ""; // Clear previous rows
console.log(data);

    if (data.length === 0) {
      const row = document.createElement("tr");
      row.innerHTML = `<td colspan="6" style="text-align: center; padding: 1em; color: gray;">No reports to display.</td>`;
      tableBody.appendChild(row);
      return;
    }

    data.forEach(report => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td>${report.PostTitle}</td>
        <td>${report.Category}</td>
        <td>${report.Reason}</td>
        <td>${new Date(report.ReportDate).toLocaleString()}</td>
        <td>${report.Status}</td>
        <td>
          ${report.Status === "pending" ? `
            <form method="POST" class="report-action-form">
              <input type="hidden" name="report_id" value="${report.ID}">
              <select name="decision" class="decision-select">
                <option value="accepted">Accept</option>
                <option value="refused">Refuse</option>
              </select>
              <button type="submit" class="submit-btn">Submit</button>
            </form>
          ` : report.Status}
        </td>
      `;
      tableBody.appendChild(row);
    });

    // Attach submit handlers
    document.querySelectorAll(".report-action-form").forEach(form => {
      form.addEventListener("submit", function (e) {
        e.preventDefault();

        const formData = new FormData(form);
        fetch("/moderator/handle_report", {
          method: "POST",
          body: formData
        })
        .then(res => {
          if (!res.ok) throw new Error("Action failed.");
          return res.text();
        })
        .then(() => fetchReportedPosts()) // Refresh table
        .catch(err => alert(err.message));
      });
    });
  })
  .catch(err => {
    console.error(err);
    alert("Could not load reports.");
  });
}
