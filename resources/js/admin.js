


document.querySelectorAll('.mod-request-form').forEach(form => {
  form.addEventListener('submit', async function (e) {
    e.preventDefault(); // Stop default form submit

    const formData = new FormData(form);
    // const userId = formData.get('user_id');
    // const action = formData.get('action');
for (const [key, value] of formData.entries()) {
  console.log(`${key}: ${value}`);
}
    try {
      const res = await fetch('/handleRequest', {
        method: 'POST',
        body: formData
      });

      if (res.ok) {
        // Remove the table row containing this form
        const row = form.closest('tr');
        if (row) row.remove();
      } else {
        const errorText = await res.text();
        console.log(errorText);
        
      }
    } catch (err) {
      console.error('Fetch error:', err);
      alert('Request failed. Please try again.');
    }
  });
});