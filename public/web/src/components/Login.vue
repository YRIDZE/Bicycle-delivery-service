<template>
  <div class="modal fade" id="login-modal" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered modal-lg">
      <div class="modal-content"
           style="
             background-color: rgba(117, 190, 218, 0.0);
             border-color:rgba(117, 190, 218, 0.0);">
        <div class="modal-body">
          <div class="login-container"
               :class="{ 'right-panel-active': signUpMode }" id="login-container">
            <div class="form-container sign-up-container">
              <form class="l-r-form" action="#">
                <h1><b>Create Account</b></h1>
                <input class="login-input" type="text" placeholder="Firstname"/>
                <input class="login-input" type="text" placeholder="Lastname"/>
                <input class="login-input" type="email" placeholder="Email"/>
                <input class="login-input" type="password" placeholder="Password"/>
                <button class="sign-up bt" style="margin-top: 7px">SIGN UP</button>
              </form>
            </div>
            <div class="form-container  sign-in-container">
              <form class="l-r-form" action="#">
                <h1><b>Sign in</b></h1>
                <input class="login-input" type="email" placeholder="Email"/>
                <input class="login-input" type="password" placeholder="Password"/>
                <a href="#" style="text-decoration: none; font-size: 14px; margin: 15px 0; color:#000000">Forgot your
                  password?</a>
                <button class="sign-in bt">SIGN IN</button>
              </form>
            </div>
            <div class="overlay-container">
              <div id="close" class="close" :class="{ 'left': signUpMode }">
                <a href="#" data-bs-dismiss="modal"><i style="color: #ffffff;" class="fas fa-times"></i></a>
              </div>
              <div class="overlay">
                <div class="overlay-panel overlay-left">
                  <h1><b>Welcome Back!</b></h1>
                  <p class="container-details">To keep connected with us please login with your personal info</p>
                  <button class="ghost bt" id="signIn" @click="signUpMode = false">SIGN IN</button>
                </div>
                <div style="right: 0" class="overlay-panel overlay-right">
                  <h1><b>Hello, Friend!</b></h1>
                  <p class="container-details">Enter your personal details and start journey with us</p>
                  <button class="ghost bt" id="signUp" @click="signUpMode = true">SIGN UP</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      signUpMode: false
    }
  },
  name: 'Login',
}
</script>

<style scoped>

</style>