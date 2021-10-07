<template>
  <div class="mt-4 mx-1" style="border-radius: 10px; min-height: 215px; min-width: 330px; max-width: 330px">
    <div class="icon-box gallery">
      <img :src="item.image" class="img-fluid" @click="showProduct" alt="menu-item"/>
      <div>
        <p class="short" style="float: left; width: 80%">{{ item.name }}</p>
        <p style="float: right; text-align: right !important;"> {{ item.price }}$</p>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  props: ['item'],

  methods: {
    showProduct() {
      this.$store.state.cart.currentItem = this.item
      this.$store.state.item.showProduct = true;
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