<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <supplier
        v-for="supplier in filteredSuppList"
        :key="supplier.id"
        :supp="supplier"
    ></supplier>
  </div>
</template>

<script>


export default {
  computed: {
    filteredSuppList: function () {
      let suppliers = this.$store.getters["supp/getSuppliers"]

      if (this.$store.getters["filter/getSuppTypeFilter"].length !== 0) {
        suppliers = suppliers
            .filter(value => this.$store.getters["filter/getSuppTypeFilter"]
                .includes(value.type))
      }

      if (this.$store.getters["filter/getSuppTimeFilter"] != null) {
        let timestamp = this.$store.getters["filter/getSuppTimeFilter"]
        suppliers = suppliers.filter(value => (value.workingHours.opening >= timestamp.opening && value.workingHours.closing <= timestamp.closing))
      }

      return suppliers

    },
  },

}
</script>


<style scoped>


</style>