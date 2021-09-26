<template>
  <Supplier
      v-for="restaurant in this.$store.state.restaurants"
      :key="restaurant.id"
      :id="restaurant.id"
      :title="restaurant.title"
      :logo="restaurant.logo"
  ></Supplier>
</template>

<script>
import Supplier from './Supplier'

export default {
  name: "SupplierList",
  components: {
    Supplier
  }
}
</script>

<style scoped>

</style>