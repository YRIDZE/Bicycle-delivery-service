<template>
  <div class="card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <router-link to="/">
        <button type="button" style="background-color: #e97d56;
        border: none;
        text-align: left;
        margin: 0;
        padding: 2px 20px 0;
        border-radius: unset;
        font-size: 16px !important;
        width: 250px"> ALL SUPPLIERS
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
  name: "ProductFilter"
}
</script>

<style scoped>

</style>