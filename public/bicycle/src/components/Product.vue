<template>
  <div class="mt-4 mx-1 product">
    <div class="icon-box gallery">
      <img :src="item.image" class="img-fluid" @click="showProduct" alt="Product picture"/>
      <div>
        <p class="short w-4/5 float-left">{{ item.name }}</p>
        <p class="float-right text-right"> {{ item.price }}$</p>
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

.product {
  border-radius: 10px;
  min-height: 215px;
  min-width: 330px;
  max-width: 330px
}
</style>