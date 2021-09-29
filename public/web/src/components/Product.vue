<template>
  <div class="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
       style="border-radius: 10px; max-height: 700px; max-width: 760px">
    <div class="icon-box">
      <img :src="image" class="img-fluid" data-bs-toggle="modal" data-bs-target="#item-modal"
           @click="this.showItem = true" alt="menu-item"/>
      <h4>{{ name }}</h4>
      <Teleport to="#pop-portal">
        <ProductPopUp :id="id" v-if="this.showItem" @close="this.showItem = false"></ProductPopUp>
      </Teleport>
    </div>
  </div>
</template>

<script>
import ProductPopUp from './ProductPopUp'


export default {
  data() {
    return {
      showItem: false,
    }
  },
  props: ['id', 'name', 'image'],
  name: "Product",
  components: {
    ProductPopUp
  }

}
</script>

<style scoped>

</style>