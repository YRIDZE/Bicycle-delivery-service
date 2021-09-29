<template>
  <div class="col-lg-2 col-md-6 d-flex align-items-stretch mt-4"
       style="border-radius: 10px; max-height: 700px; max-width: 760px">
    <div class="icon-box">
      <router-link :to="{ path : String(id) }"><img :src="image" class="img-fluid" alt="Pizza picture"></router-link>
      <h4><a href="#">{{ name }}</a></h4>
    </div>
  </div>
</template>

<script>
export default {
  props: ['id', 'name', 'image'],
  name: "Supplier",
  data() {
    return {
      loading: false,
    };
  },
}

</script>

<style scoped>

</style>