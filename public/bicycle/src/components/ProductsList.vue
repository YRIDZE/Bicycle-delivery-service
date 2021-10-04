<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <Product
        v-for="item in this.$store.state.items.filter(x => x.supplier_id === parseInt(id, 10))"
        :key="item.id"
        :item="item"
    ></Product>
  </div>
</template>

<script>

import Product from "@/components/Product";

export default {
  props: ["id"],

  components: {
    Product,
  },
};
</script>

<style scoped>

</style>