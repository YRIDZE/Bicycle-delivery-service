<template>
  <div class="card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <router-link to="/">
        <button class="all-items text-base" type="button"> ALL SUPPLIERS
        </button>
      </router-link>
      <div class="accordion-item">
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
              <div class="card-body p-1.5" v-for="(value) in this.$store.getters['prod/getProductsTypes']" :key="value">
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
  name: "ProductFilter"
}
</script>

<style scoped>

</style>