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
          <form class="l-r-form" v-on:submit.prevent="registrationUser">
            <h1><b>Create Account</b></h1>
            <input class="login-input" v-model="registrationForm.first_name" type="text" placeholder="Firstname"/>
            <input class="login-input" v-model="registrationForm.last_name" type="text" placeholder="Lastname"/>
            <input class="login-input" v-model="registrationForm.email" type="email" placeholder="Email"/>
            <input class="login-input" v-model="registrationForm.password" type="password" placeholder="Password"/>
            <button class="sign-up bt my-3">SIGN UP</button>
          </form>
        </div>
        <div class="form-container  sign-in-container">
          <form class="l-r-form" v-on:submit.prevent="loginUser">
            <h1><b>Sign in</b></h1>
            <input class="login-input" v-model="loginForm.email" type="email" placeholder="Email"/>
            <input class="login-input" v-model="loginForm.password" type="password" placeholder="Password"/>
            <a class="my-4 mt-0 no-underline text-sm" href="#" style="color:#000000">Forgot your
              password?</a>
            <button class="sign-in bt mb-3">SIGN IN</button>
          </form>
        </div>
        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel overlay-left">
              <h1><b class="mb-3">Welcome Back!</b></h1>
              <p class="container-details mb-3">To keep connected with us please login with your personal info</p>
              <button class="ghost bt" id="signIn" @click="signUpMode = false">SIGN IN</button>
            </div>
            <div class="overlay-panel overlay-right right-0">
              <h1><b class="mb-3">Hello, Friend!</b></h1>
              <p class="container-details mb-3">Enter your personal details and start journey with us</p>
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
    ...mapActions("user", ["login", "registration"]),

    registrationUser() {
      this.registration(this.registrationForm)
          .then(() => {
            this.$store.dispatch('cart/createCart').catch(err => console.log(err));
          })
          .catch(err => console.log(err));
      this.signUpMode = false;

      this.registrationForm.email = '';
      this.registrationForm.password = '';
      this.registrationForm.last_name = '';
      this.registrationForm.first_name = '';
    },

    loginUser() {
      this.login(this.loginForm)
          .then(() => {
            this.$store.dispatch('cart/getCart').catch(err => console.log(err));
          })
          .catch(err => console.log(err));
      this.hide()

      this.loginForm.email = '';
      this.loginForm.password = '';
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
