<template>
  <label class="form-check">
    <input class="form-check-input" type="checkbox" :value="name" v-model="checked">
    <span class="form-check-label">{{ name.charAt(0).toUpperCase() + (name.replace('_', ' ').slice(1)) }}</span>
  </label>
</template>

<script>
export default {
  props: ['name'],
  computed: {
    checked: {
      get() {
        return this.$store.getters["filter/getFilter"].type
      },
      set(value) {
        this.$store.dispatch('filter/setFilter', value)
      }
    }
  }
}
</script>

<style scoped>

</style>