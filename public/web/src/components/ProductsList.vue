<template>
  <Product
      v-for="item in this.$store.state.items.filter(x => x.supplier_id == id)"
      :key="item.id"
      :id="item.id"
      :name="item.name"
      :image="item.image"
  ></Product>
</template>

<script>

import Product from "@/components/Product";

export default {
  props: ["id"],
  name: "ProductList",
  components: {
    Product
  },
}
</script>

<style scoped>

</style>