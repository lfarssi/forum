const form =document.getElementById("Form");

form.addEventListener("submit", function (event) {

    event.preventDefault();

    const title = document.getElementById('title');
    const content = document.getElementById('content');
    const categoryCheckboxes = document.querySelectorAll('input[name="categories"]:checked');
    const errordiv = document.getElementById('error');

    errordiv.textContent = '';

    if (!title.value) {
        errordiv.textContent = 'Title is required.';
        return;
    }

    if (!content.value) {
        errordiv.textContent = 'Content is required.';
        return;
    }

    if (categoryCheckboxes.length === 0) {
        console.log(categoryCheckboxes);
        console.log(categoryCheckboxes.length);
        
        
        console.log("Please select at least one category");
        
        errordiv.textContent = 'At least one category must be selected.';
        return;
    }
// console.log("add post ");
alert("form submited")
    document.querySelector('form').submit();
//    window.location.href='/home'
}
)