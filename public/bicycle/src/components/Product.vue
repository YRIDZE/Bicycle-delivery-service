<template>
  <div class="mt-4 mx-1" style="border-radius: 10px; min-height: 220px; min-width: 330px">
    <div class="icon-box gallery">
      <img :src="this.item.image" class="img-fluid" @click="showProduct" alt="menu-item"/>
      <h4>{{ this.item.name }}</h4>
    </div>
  </div>
</template>

<script>

export default {
  props: ['item'],

  methods: {
    showProduct() {
      this.$store.state.cart.currentItem = this.item
      this.$store.state.showProduct = true;
    },
  },
}
</script>

<style scoped>

.gallery img {
  width: 330px !important;
  height: 215px;
  object-fit: cover;
}
</style>