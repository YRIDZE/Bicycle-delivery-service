<template>
  <div class="list-group-item list-group-item">
    <div class="d-flex w-100 justify-content-between">
      <b class="mb-1 w-3/4">{{ item.address }}</b>
      <small>{{ item.created_at }}</small>
    </div>
    <ol class="list-group list-group-numbered">
      <order-list-products v-for="(i) in item.products" :key="i" :item="i"></order-list-products>
    </ol>
  </div>
</template>

<script>
export default {
  props: ['item', 'index'],
}
</script>

<style scoped>

</style>