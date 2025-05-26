


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
    form.addEventListener('submit', async function (e) {
      console.log(e);
      
      e.preventDefault();
      const formData = new FormData(form);

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
          const errorText = await res.text();
          console.error(errorText);
        }
      } catch (err) {
        console.error('Fetch error:', err);
        alert('Request failed. Please try again.');
      }
    });
  });
