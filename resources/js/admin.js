


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
