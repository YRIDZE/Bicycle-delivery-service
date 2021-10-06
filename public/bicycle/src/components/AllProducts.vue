<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <product
        v-for="item in filteredProdList"
        :key="item.id"
        :item="item"
    ></product>
  </div>
</template>

<script>
export default {
  name: "AllProducts",
  computed: {
    filteredProdList: function () {
      if (this.$store.getters["filter/getProdTypeFilter"].length === 0) return this.$store.getters["item/getItems"]
      return this.$store.getters["item/getItems"].filter(value => this.$store.getters["filter/getProdTypeFilter"].includes(value.type))
    },
  },

}
</script>

<style scoped>

</style>