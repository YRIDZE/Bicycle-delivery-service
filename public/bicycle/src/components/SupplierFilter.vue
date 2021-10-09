<template>
  <div class="card">
    <div class="accordion accordion-flush" id="accordionFlushExample">
      <router-link to="all">
        <button type="button" style="background-color: #e97d56;
        border: none;
        text-align: left;
        margin: 0;
        padding: 2px 20px 0;
        border-radius: unset;
        font-size: 16px !important;
        width: 250px"> ALL PRODUCTS
        </button>
      </router-link>
      <div class="accordion-item">
        <h2 class="accordion-header" id="panel2">
          <button class="accordion-button" type="button" data-bs-toggle="collapse"
                  data-bs-target="#panelSupplierTime" aria-expanded="true"
                  aria-controls="panelSupplierTime">
            WORKING HOURS
          </button>
        </h2>
        <div id="panelSupplierTime" class="accordion-collapse collapse show"
             aria-labelledby="panel2">
          <div class="accordion-body">
            <div class="filter-content">
              <p>from <input id="from" type="time" value="00:00" @change="time" style="width: 150px"></p>
              <p>to <input id="to" type="time" value="23:59" @change="time" style="width: 150px"></p>
            </div>
          </div>
        </div>
      </div>
      <div class="accordion-item">
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
    }
  }
}

</script>

<style scoped>

</style>