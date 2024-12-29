function HandleSubmit(event) {
    // Prevent the default form submission
    event.preventDefault();

    const title = document.getElementById('title');
    const content = document.getElementById('content');
    const categoryCheckboxes = document.querySelectorAll('input[name="category_ids"]:checked');
    const errordiv = document.getElementById('error');

    // Clear previous error messages
    errordiv.textContent = '';

    // Check if title is empty
    if (!title.value) {
        errordiv.textContent = 'Title is required.';
        return;
    }

    // Check if content is empty
    if (!content.value) {
        errordiv.textContent = 'Content is required.';
        return;
    }

    // Check title length
    if (title.value.length > 100) {
        errordiv.textContent = 'Title must not exceed 100 characters.';
        return;
    }

    // Check content length
    if (content.value.length > 500) {
        errordiv.textContent = 'Content must not exceed 500 characters.';
        return;
    }

    // Check if at least one category is selected
    if (categoryCheckboxes.length === 0) {
        errordiv.textContent = 'At least one category must be selected.';
        return;
    }

    // If all validations pass, submit the form
    document.querySelector('form').submit();
}
