<template>
  <header id="header">
    <div class="header-bar fixed-top">
      <nav class="navbar navbar-expand-sm navbar-dark bg-dark p-4" style="
            background-color:#545454ff !important;">
        <a class="navbar-brand px-7 py-0" href="/">
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
          <ul class="navbar-nav ms-auto text-center px-7 py-0">
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
            <li class="nav-item"><a class="nav-link" @click="showOrders" v-if="isLoggedIn" type="button">
              <font-awesome-icon :icon="['fas', 'stream']"/>
              Orders</a></li>
          </ul>
        </div>
      </nav>
    </div>

    <login></login>
    <cart></cart>
    <orders></orders>

    <div class="text-center bg-light h-96 pt-28 px-12 pb-0 mb-4 bg-cover"
         :style="{ backgroundImage: 'url(' + require(`@/assets/header-pizza.jpg`) + ')'}">
      <p class="font-bold italic text-8xl	mb-3 mt-8">Bicycle</p>
      <p class="italic text-3xl	">This is the way</p>
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
        this.$store.state.cart.showCart = true;
      } else {
        this.$store.state.user.showLogin = true;
      }
    },
    showOrders() {
      this.$store.dispatch('orders/getOrders')
      this.$store.state.orders.showOrders = true;
    },
  },
}
</script>

<style scoped>
span {
  position: absolute;
  right: 0.1rem;
  top: 0.7rem;
  color: white;
  border-radius: 50%;
  font-size: 90%;
  padding: 0.4rem;
}
</style>