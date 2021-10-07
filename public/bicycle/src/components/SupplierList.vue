<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <supplier
        v-for="supplier in filteredSuppList"
        :key="supplier.id"
        :supp = "supplier"
    ></supplier>
  </div>
</template>

<script>


export default {
  computed: {
    filteredSuppList: function () {
      if (this.$store.getters["filter/getSuppTypeFilter"].length === 0) return this.$store.getters["supp/getSuppliers"]
      return this.$store.getters["supp/getSuppliers"].filter(value => this.$store.getters["filter/getSuppTypeFilter"].includes(value.type))
    },
  },

}
</script>


<style scoped>


</style>