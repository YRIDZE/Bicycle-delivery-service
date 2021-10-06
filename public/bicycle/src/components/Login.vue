<template>
  <vue-final-modal
      v-model="$store.state.user.showLogin"
      classes="modal-container"
      content-class="modal-content"
      :esc-to-close="true"
      @before-open="hide"
  >
    <div class="modal-body">
      <div class="login-container"
           :class="{ 'right-panel-active': signUpMode }" id="login-container">
        <div class="form-container sign-up-container">
          <form class="l-r-form" v-on:submit.prevent="registUser">
            <h1><b>Create Account</b></h1>
            <input class="login-input" v-model="registrationForm.first_name" type="text" placeholder="Firstname"/>
            <input class="login-input" v-model="registrationForm.last_name" type="text" placeholder="Lastname"/>
            <input class="login-input" v-model="registrationForm.email" type="email" placeholder="Email"/>
            <input class="login-input" v-model="registrationForm.password" type="password" placeholder="Password"/>
            <button class="sign-up bt" style="margin-top: 7px">SIGN UP</button>
          </form>
        </div>
        <div class="form-container  sign-in-container">
          <form class="l-r-form" v-on:submit.prevent="loginUser">
            <h1><b>Sign in</b></h1>
            <input class="login-input" v-model="loginForm.email" type="email" placeholder="Email"/>
            <input class="login-input" v-model="loginForm.password" type="password" placeholder="Password"/>
            <a href="#" style="text-decoration: none; font-size: 14px; margin: 15px 0; color:#000000">Forgot your
              password?</a>
            <button class="sign-in bt">SIGN IN</button>
          </form>
        </div>
        <div class="overlay-container">
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
  </vue-final-modal>
</template>

<script>
import HideModals from '../mixins/hideModals'
import {mapActions} from "vuex";


export default {
  mixins: [HideModals],

  data() {
    return {
      signUpMode: false,
      registrationForm: {
        first_name: '',
        last_name: '',
        email: '',
        password: '',
      },
      loginForm: {
        email: '',
        password: '',
      },
    };
  },
  methods: {
    ...mapActions('user', ['login', 'registration']),

    registUser() {
      this.registration(this.registrationForm)
          .catch(err => console.log(err));
      this.signUpMode = false;
      this.registrationForm = '';
    },
    loginUser: function () {
      this.login(this.loginForm)
          .catch(err => console.log(err));
      this.loginForm = '';
      this.hide()
    },
  },
};
</script>

<style>
.slide-enter-active,
.slide-leave-active {
  transition: all 0.7s;
}

.slide-enter-from,
.slide-leave-to {
  transform: translateY(-100vh);
}

.modal-container {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>
