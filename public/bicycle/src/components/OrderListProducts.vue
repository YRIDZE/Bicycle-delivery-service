<template>
  <li class="d-flex w-100">
    <p class="mb-1">{{ "&nbsp;" }} {{ product.name }} х {{ item.quantity }}</p>
  </li>
</template>

<script>
export default {
  props: ['item'],
  computed: {
    product: function () {
      return this.$store.getters["prod/getProducts"].find(x => x.id == this.item.product_id)
    },
  },
}
</script>

<style scoped>

</style>