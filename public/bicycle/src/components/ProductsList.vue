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
  props: ["id"],
  computed: {
    filteredProdList: function () {
      let products = this.$store.getters["prod/getProducts"]

      if (this.$store.getters["filter/getProdTypeFilter"].length !== 0) {
        products = products
            .filter(value => this.$store.getters["filter/getProdTypeFilter"]
                .includes(value.type))
      }

      if (this.id == "all")
        return products

      return products.filter(x => x.supplier_id === parseInt(this.id, 10))
    },
  },
};
</script>

<style scoped>

</style>