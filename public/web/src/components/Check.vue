<template>
  <label class="form-check">
    <input class="form-check-input" type="checkbox" value="">
    <span class="form-check-label">{{ name }}</span>
  </label>
</template>

<script>
export default {
  props: ['name'],
  name: "Check"
}
</script>

<style scoped>

</style>