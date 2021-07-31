const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');
const closeButton = document.getElementById('close')

signUpButton.addEventListener('click', () => {
    container.classList.add("right-panel-active");
    closeButton.classList.add("left")
});

signInButton.addEventListener('click', () => {
    container.classList.remove("right-panel-active");
    closeButton.classList.remove("left")

});
