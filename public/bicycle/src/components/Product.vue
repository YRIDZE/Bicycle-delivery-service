<template>
  <div class="card mt-4 mx-1 product icon-box">
    <div class="gallery">
      <img :src="item.image" class="img-fluid" style="background-color: #FFFFFF" @click="showProduct"
           alt="Product picture"/>
      <div class="card-body  d-flex flex-row align-items-center">
        <div class="short text-xl">{{ item.name }}</div>
        <div class="ml-auto text-2xl">{{ item.price }}$</div>
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
};
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
  max-width: 365px;
}
</style>