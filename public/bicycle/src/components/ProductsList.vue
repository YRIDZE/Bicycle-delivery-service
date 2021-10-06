<template>
  <div class="d-flex flex-wrap bd-highlight align-items-stretch">
    <product
        v-for="item in this.$store.getters['item/getItems'].filter(x => x.supplier_id === parseInt(id, 10))"
        :key="item.id"
        :item="item"
    ></product>
  </div>
</template>

<script>


export default {
  props: ["id"],
};
</script>

<style scoped>

</style>