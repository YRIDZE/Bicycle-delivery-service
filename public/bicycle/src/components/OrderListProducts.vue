<template>
  <li class="d-flex bd-highlight text-lg py-1">
    <p>{{ "&nbsp;" }} {{ product.name }}</p>
    <p class="ms-auto">{{ item.quantity }}х{{ item.price }}$</p>
  </li>
</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
  },
  computed: {
    product: function () {
      return this.$store.getters["prod/getProducts"].find(x => x.id == this.item.product_id)
    },
  },
}
</script>

<style scoped>

</style>