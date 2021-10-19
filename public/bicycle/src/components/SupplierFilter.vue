<template>
  <div class="card filter-card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <router-link to="all">
        <button class="all-items text-lg" type="button"> ALL PRODUCTS
        </button>
      </router-link>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panel2">
          <button class="accordion-button text-lg" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelSupplierTime" aria-expanded="true"
                  aria-controls="panelSupplierTime">
            WORKING HOURS
          </button>
        </h2>
        <div id="panelSupplierTime" class="accordion-collapse collapse show"
             aria-labelledby="panel2" style="background-color: #fafafa">
          <div class="accordion-body">
            <div class="filter-content">
              <p class="pb-2">Open <input class="w-40 text-right outline-none" id="from" type="time" value="00:00"
                                          @change="time" style="background-color: #fafafa"></p>
              <p>Close <input class="w-40 text-right outline-none" id="to" type="time" value="23:59" @change="time"
                              style="background-color: #fafafa">
              </p>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panel1">
          <button class="accordion-button text-lg" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelSupplierTypes" aria-expanded="true"
                  aria-controls="panelSupplierTypes">
            RESTAURANT TYPE
          </button>
        </h2>
        <div id="panelSupplierTypes" class="accordion-collapse collapse show"
             aria-labelledby="panel1" style="background-color: #fafafa">
          <div class="accordion-body">
            <div class="filter-content">
              <div class="card-body p-1.5" v-for="(value) in this.$store.getters['supp/getSuppliersTypes']"
                   :key="value">
                <check :title="'restaurantType'" :name="value"></check>
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
    time: function () {
      this.$store.dispatch("filter/setSuppTimeFilter", {
        opening: document.getElementById("from").value,
        closing: document.getElementById("to").value,
      })
    },
  },
}

</script>

<style scoped>

</style>