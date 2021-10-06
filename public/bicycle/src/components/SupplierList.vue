<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <supplier
        v-for="supplier in filteredSuppList"
        :key="supplier.id"
        :id="supplier.id"
        :name="supplier.name"
        :image="supplier.image"
    ></supplier>
  </div>
</template>

<script>


export default {
  computed: {
    filteredSuppList: function () {
      if (this.$store.getters["filter/getFilter"].type.length === 0) return this.$store.getters["supp/getSuppliers"]
      return this.$store.getters["supp/getSuppliers"].filter(value => this.$store.getters["filter/getFilter"].type.includes(value.type))
    },
  },

}
</script>


<style scoped>


</style>