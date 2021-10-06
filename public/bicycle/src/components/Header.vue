<template>
  <header id="header">
    <div class="header-bar fixed-top">
      <nav class="navbar navbar-expand-sm navbar-dark bg-dark" style="
            background-color:#545454ff !important;
            padding: 1em !important;">

        <a class="navbar-brand" href="/" style="padding:0 30px">
          <font-awesome-icon :icon="['fas', 'bicycle']"/>
          <b> Bicycle</b></a>

        <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#toggleMenu"
            aria-controls="toggleMenu"
            aria-expanded="false"
            aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="toggleMenu">
          <ul class="navbar-nav ms-auto text-center" style="padding:0 30px">
            <li class="nav-item"><a class="nav-link" href="#">About</a></li>
            <li class="nav-item" v-if="isLoggedIn"><a class="nav-link" id="logout" @click="logout" type="button">
              <font-awesome-icon :icon="['fas', 'sign-out-alt']"/>
              Logout</a></li>
            <li class="nav-item" v-else><a class="nav-link" @click="$store.state.user.showLogin = true"
                                           type="button">
              <font-awesome-icon :icon="['fas', 'sign-in-alt']"/>
              Login</a></li>
            <li class="nav-item"><a class="nav-link" @click="showCart" type="button">
              <font-awesome-icon :icon="['fas', 'shopping-basket']"/>
              Cart</a></li>
          </ul>
        </div>
      </nav>
    </div>
    <login></login>
    <cart></cart>

    <div class="cont-custom text-center bg-light"
         :style="{ backgroundImage: 'url(' + require(`@/assets/header-pizza.jpg`) + ')',
           marginBottom : '1rem',
           backgroundSize: 'cover',
           padding: '6rem 3rem 0 3rem !important',
           height: '26rem'}">

      <h1 class="mb-3"><b>Bicycle</b></h1>
      <h4 class="mb-3">Subheading</h4>

      <div class="search-container">
        <input type="text" class="input" placeholder="what are you looking for...">
        <input type="button" class="close-btn" value="Search">
      </div>
    </div>
  </header>

</template>

<script>

export default {
  name: "Header",
  computed: {
    isLoggedIn: function () {
      return this.$store.getters["user/isLoggedIn"];
    },
  },
  methods: {
    logout() {
      this.$store.dispatch('user/logout');
    },
    showCart() {
      if (this.$store.getters["user/isLoggedIn"]) {
        this.$store.state.cart.showCart = true
      } else {
        this.$store.state.user.showLogin = true
      }
    }
  },
}
</script>

<style scoped>

</style>