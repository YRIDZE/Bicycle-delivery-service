<template>
  <label class="form-check">
    <input class="form-check-input" type="checkbox" :value="name" v-model="checked">
    <span class="form-check-label">{{ name.charAt(0).toUpperCase() + (name.replace('_', ' ').slice(1)) }}</span>
  </label>
</template>

<script>
export default {
  props: ['name', 'title'],
  computed: {
    checked: {
      get() {
        return this.title === 'restaurantType' ? this.$store.getters["filter/getSuppTypeFilter"] : this.$store.getters["filter/getProdTypeFilter"];
      },
      set(value) {
        if (this.title === 'restaurantType') {
          return this.$store.dispatch('filter/setSuppTypeFilter', value)
        }
        this.$store.state.prod.pageNumber = 0;
        return this.$store.dispatch('filter/setProdTypeFilter', value)
      }
    }
  }
}
</script>

<style scoped>

</style>