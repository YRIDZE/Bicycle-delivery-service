<template>
  <div id="main-content">
    <header-top></header-top>

    <div class="container-fluid px-16" id="content"
         style="font-family: 'Montserrat', sans-serif;">
      <div class="col-md-12">
        <h2>Food
          <go-back
              v-if="this.$router.currentRoute.fullPath !== '/' && this.$router.currentRoute.fullPath !== '/all'"></go-back>
        </h2>
        <p class="-mb-0">...is any substance consumed to provide nutritional support for an
          organism. Food is usually of plant, animal
          or fungal origin, and contains essential nutrients, such as carbohydrates, fats, proteins, vitamins, or
          minerals. The substance is ingested by an organism and assimilated by the organism's cells to provide energy,
          maintain life, or stimulate growth. </p>
      </div>

      <div class="d-flex flex-row mb-5 mt-1">
        <router-view name="filter"></router-view>

        <div class="d-flex align-items-center justify-content-center flex-grow-1 "
             v-if="this.$store.state.supp.loading">
          <pulse-loader class="flex-grow-1 absolute" :color="'#e97d56'"></pulse-loader>
        </div>

        <section id="services" class="services flex-shrink-1 pb-2.5"
                 v-if="!this.$store.state.supp.loading">
          <div class="row">
            <router-view name="content"></router-view>
          </div>
        </section>
      </div>
    </div>
    <product-popup v-if="this.$store.getters['cart/getCurrentItem']"
                   :item="this.$store.getters['cart/getCurrentItem']"></product-popup>

    <go-top></go-top>
    <notifications group="log-reg" position="bottom center" :max="3"/>
    <bottom-footer></bottom-footer>
  </div>
</template>

<script>
import './assets/css/main-page.css'
import './assets/css/cart.css'
import './assets/css/login-registration.css'
import './assets/css/menu-item-page.css'
import axios from "axios";
import {getTokenTimeUntilRefresh} from "@/store/user";

export default {
  name: "App",

  methods: {
    async fetchedSupplierProducts() {
      this.$store.state.supp.loading = true
      await fetch("getSuppliers",)
          .then(response => response.json())
          .then(async data => {
            await this.$store.dispatch('supp/setSuppliers', data)
            await this.$store.dispatch('supp/getSupplierTypes');

            await fetch("getProducts",)
                .then(response => response.json())
                .then(data => {
                  this.$store.dispatch('prod/setProduct', data)
                  this.$store.dispatch('prod/getProductTypes');
                });
          })
      this.$store.state.supp.loading = false
    },
  },
  async mounted() {
    if (this.$store.getters["user/isLoggedIn"]) {
      axios.defaults.headers.common["Authorization"] = `Bearer ${this.$store.getters["user/accessToken"]}`;
      if (getTokenTimeUntilRefresh(this.$store.getters["user/refreshToken"]) > 0) {
        await this.$store.dispatch('user/dropRefresh');
        await this.$store.dispatch('user/autoRefresh');
        await this.$store.dispatch('cart/getCart');
      } else {
        this.$store.commit("user/logout");
      }
    }
  },
  created() {
    this.fetchedSupplierProducts()
  },
};

</script>