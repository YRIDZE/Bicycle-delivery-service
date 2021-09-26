<template>
  <div class="col-lg-2 col-md-6 d-flex align-items-stretch mt-4"
       style="border-radius: 10px; max-height: 700px; max-width: 760px">
    <div class="icon-box">
      <router-link :to="{ path : String(id) }"><img :src="logo" class="img-fluid" alt="Pizza picture"></router-link>
      <h4><a href="#">{{ title }}</a></h4>
    </div>
  </div>
</template>

<script>
export default {
  props: ['id', 'title', 'logo'],
  name: "Supplier"
}
</script>

<style scoped>

</style>