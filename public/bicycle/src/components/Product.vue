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
  props: {
    item: {
      type: Object,
      required: true,
    },
  },

  methods: {
    showProduct() {
      this.$store.state.cart.currentItem = this.item
      this.$store.state.prod.showProduct = true;
    },
  },
}
</script>

<style scoped>

.gallery img {
  width: 365px !important;
  height: 235px;
  object-fit: cover;
}

.product {
  border-radius: 10px;
  min-height: 235px;
  min-width: 365px;
  max-width: 365px
}
</style>