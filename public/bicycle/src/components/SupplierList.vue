<template>
<div class="d-flex flex-wrap bd-highlight align-items-stretch">
  <Supplier
      v-for="restaurant in this.$store.state.restaurants"
      :key="restaurant.id"
      :id="restaurant.id"
      :name="restaurant.name"
      :image="restaurant.image"
  ></Supplier>
</div>
</template>

<script>
import Supplier from './Supplier'

export default {
  components: {
    Supplier
  },

  methods: {
    async fetchedSupplierProducts() {
      this.$store.state.loading = true
      await fetch("http://localhost:8081/getSuppliers",)
          .then(response => response.json())
          .then(async data => {
            this.$store.state.restaurants = data

            await fetch("http://localhost:8081/getProducts",)
                .then(response => response.json())
                .then(data => this.$store.state.items = data);
          })
      this.$store.state.loading = false
    },
  },

  created() {
    this.fetchedSupplierProducts()
  }

}
</script>


<style scoped>


</style>