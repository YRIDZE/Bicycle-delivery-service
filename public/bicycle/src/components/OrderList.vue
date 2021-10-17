<template>
  <div class="list-group-item list-group-item p-3">
    <div class="d-flex w-100 justify-content-between">
      <b class="mb-1 w-3/4">{{ item.address }}</b>
      <small>{{ item.created_at }}</small>
    </div>
    <ol class="list-group list-group-numbered">
      <order-list-products v-for="(i) in item.products" :key="i" :item="i"></order-list-products>
    </ol>
    <p class="text-right font-black">{{ total(item.products) }}$</p>

  </div>
</template>

<script>
export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
    index: {
      type: Number,
      required: true,
    }
  },
  methods: {
    total(products) {
      let total = 0;
      products.forEach(product => total += product.price * product.quantity);
      return total.toFixed(2).toString().replace(/\B(?=(\d{3})+$)/g, ',');
    },
  }
}
</script>

<style scoped>

</style>