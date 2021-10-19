<template>
  <div class="card filter-card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <router-link to="/">
        <button class="all-items text-lg" @click="clearSuppFilter" type="button"> ALL SUPPLIERS
        </button>
      </router-link>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panel2">
          <button class="accordion-button text-lg" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelProductTypes" aria-expanded="true"
                  aria-controls="panelProductTypes">
            PRODUCT TYPE
          </button>
        </h2>
        <div id="panelProductTypes" class="accordion-collapse collapse show"
             aria-labelledby="panel2" style="background-color: #fafafa">
          <div class="accordion-body">
            <div class="filter-content">
              <div class="card-body p-1.5" v-for="(value) in productTypes" :key="value">
                <check :title="'productType'" :name="value"></check>
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
  methods: {
    clearSuppFilter() {
      this.$store.dispatch("filter/setSuppTimeFilter", null);
      this.$store.dispatch("filter/setSuppTypeFilter", []);
    },
  },
  mounted() {
    this.$store.dispatch('filter/setProdTypeFilter', []);
  },
  computed: {
    productTypes() {
      let path = this.$router.currentRoute.fullPath
      if (path !== "/all") {
        return this.$store.getters['prod/getProductsTypesBySupp']
      } else {
        return this.$store.getters['prod/getProductsTypes']
      }
    }
  }
}
</script>

<style scoped>

</style>