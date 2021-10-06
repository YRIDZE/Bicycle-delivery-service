<template>
  <div class="card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <button type="button" @click="showAllProd" style="background-color: #e97d56;
        border: none;
        text-align: left;
        margin: 0;
        padding: 2px 20px 0;
        border-radius: unset;
        font-size: 16px !important;
        width: 250px"> ALL PRODUCTS
      </button>


      <div class="accordion-item" v-if="this.$router.currentRoute.fullPath === '/'">
        <h2 class="accordion-header" id="panel1">
          <button class="accordion-button" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelSupplierTypes" aria-expanded="true"
                  aria-controls="panelSupplierTypes">
            RESTAURANT TYPE
          </button>
        </h2>
        <div id="panelSupplierTypes" class="accordion-collapse collapse show"
             aria-labelledby="panel1">
          <div class="accordion-body">
            <div class="filter-content">
              <div class="card-body" v-for="(value) in this.$store.getters['supp/getSuppliersTypes']" :key="value"
                   style="padding: 5px !important;">
                <check :title="'restaurantType'" :name="value"></check>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item" v-else>
        <h2 class="accordion-header" id="panel2">
          <button class="accordion-button" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelProductTypes" aria-expanded="true"
                  aria-controls="panelProductTypes">
            PRODUCT TYPE
          </button>
        </h2>
        <div id="panelProductTypes" class="accordion-collapse collapse show"
             aria-labelledby="panel2">
          <div class="accordion-body">
            <div class="filter-content">
              <div class="card-body" v-for="(value) in this.$store.getters['item/getItemsTypes']" :key="value"
                   style="padding: 5px !important;">
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
    showAllProd() {
      this.$router.push('all');
    },
  },
}

</script>

<style scoped>

</style>