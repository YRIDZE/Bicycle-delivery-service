<template>
  <div id="main-content">
    <header-top></header-top>

    <div class="container-fluid" id="content"
         style="padding-left: 70px !important; padding-right: 70px !important; font-family: 'Montserrat', sans-serif;">
      <div class="col-md-12">
        <h2>Food
          <go-back></go-back>
        </h2>
        <p style="margin-bottom: 0 !important;">...is any substance consumed to provide nutritional support for an
          organism. Food is usually of plant, animal
          or fungal origin, and contains essential nutrients, such as carbohydrates, fats, proteins, vitamins, or
          minerals. The substance is ingested by an organism and assimilated by the organism's cells to provide energy,
          maintain life, or stimulate growth. </p>
      </div>

      <div class="d-flex flex-row mb-5 mt-1">
        <router-view name="filter"></router-view>

        <div class="d-flex align-items-center justify-content-center flex-grow-1" v-if="this.$store.state.supp.loading">
          <pulse-loader class="flex-grow-1" :color="'#e97d56'" style="position: absolute"></pulse-loader>
        </div>

        <section id="services" class="services flex-shrink-1" style="padding-bottom: 10px"
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
    <bottom-footer></bottom-footer>
  </div>
</template>

<script>
import '../../bicycle/public/css/main-page.css'
import '../../bicycle/public/css/cart.css'
import '../../bicycle/public/css/login-registration.css'
import '../../bicycle/public/css/menu-item-page.css'

export default {
  name: "App",

  methods: {
    async fetchedSupplierProducts() {
      this.$store.state.supp.loading = true
      await fetch("http://localhost:8081/getSuppliers",)
          .then(response => response.json())
          .then(async data => {
            await this.$store.dispatch('supp/setSuppliers', data)
            await fetch("http://localhost:8081/getProducts",)
                .then(response => response.json())
                .then(data => this.$store.dispatch('item/setItem', data));
          })
      this.$store.state.supp.loading = false
    },
  },

  computed: {
    isLoggedIn: function () {
      return this.$store.getters["user/isLoggedIn"];
    },
  },
  created() {
    this.fetchedSupplierProducts()
  },
  mounted() {
    this.$store.dispatch('dropRefresh');
    if (this.isLoggedIn) {
      this.$store.dispatch('cart/getCart').catch(err => console.log(err));
    }
  }

};

</script>
