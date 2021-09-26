<template>
  <Product
      v-for="item in this.$store.state.items.filter(x => x.restaurantId == id)"
      :key="item.id"
      :id="item.id"
      :title="item.title"
      :logo="item.logo"
  ></Product>
</template>

<script>

import Product from "@/components/Product";

export default {
  props: ["id"],
  name: "ProductList",
  components: {
    Product
  }
}
</script>

<style scoped>

</style>