<template>
  <div class="mt-4 mx-1" style="border-radius: 10px; min-height: 225px; min-width: 225px">
    <div class="icon-box gallery">
      <router-link :to="{ path : String(id) }"><img :src="image" class="img-fluid" alt="Pizza picture"></router-link>
      <h4>{{ name }}</h4>
    </div>
  </div>
</template>

<script>
export default {
  props: ['id', 'name', 'image'],
  data() {
    return {
      loading: false,
    };
  },
}

</script>

<style scoped>

.gallery img {
  width: 225px;
  height: 225px;
  object-fit: cover;
}

</style>