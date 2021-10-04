<template>
  <vue-final-modal
      v-model="$store.state.showLogin"
      classes="modal-container"
      content-class="modal-content"
      :esc-to-close="true"
      @before-open="hide"
  >
    <div class="modal-body">
      <div class="login-container"
           :class="{ 'right-panel-active': signUpMode }" id="login-container">
        <div class="form-container sign-up-container">
          <form class="l-r-form" v-on:submit.prevent="login">
            <h1><b>Create Account</b></h1>
            <input class="login-input" v-model="form.first_name" type="text" placeholder="Firstname"/>
            <input class="login-input" v-model="form.last_name" type="text" placeholder="Lastname"/>
            <input class="login-input" v-model="form.email" type="email" placeholder="Email"/>
            <input class="login-input" v-model="form.password" type="password" placeholder="Password"/>
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

export default {
  mixins: [HideModals],

  data() {
    return {
      signUpMode: false,
      form: {
        first_name: '',
        last_name: '',
        email: '',
        password: '',
      }
    };
  },
  methods: {
    login() {
      // let entry = {
      //   first_name: document.getElementsByName("name"),
      //   last_name: document.getElementsByName("surname"),
      //   email: document.getElementsByName("email"),
      //   password: document.getElementsByName("password"),
      // };

      const entry = {
        first_name: "jfjfjf",
        last_name: "A2",
        email: "29@gmail.com",
        password: "password"
      }
      // const headers = {
      //   "Access-Control-Allow-Methods": "POST, GET, OPTIONS, PUT, DELETE",
      //   "Access-Control-Allow-Headers": "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
      //   "Access-Control-Allow-Origin": "*",
      // };
      // axios.post("http://localhost:8081/createUser", entry, )
      //     .then(response => this.id = response.data.id)
      //     .catch(error => {
      //       this.errorMessage = error.message;
      //       console.error("There was an error!", error);
      //     });


      const requestOptions = {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",

        },
        body: JSON.stringify(entry)
      };
      fetch("http://localhost:8081/createUser", requestOptions)
          .then(response => response.json())
          .then(data => console.log(data.id));
    }
  }
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
