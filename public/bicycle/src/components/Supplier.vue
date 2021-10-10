<template>
  <div class="mt-4 mx-1 supplier">
    <div class="icon-box gallery">
      <router-link :to="{ path : String(supp.id) }"><img :src="supp.image" class="img-fluid" alt="Pizza picture">
      </router-link>
      <p class="short">{{ supp.name }}</p>
      <p class="time"> {{ supp.workingHours.opening.substr(0, 5) }}-{{ supp.workingHours.closing.substr(0, 5) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ['supp'],
}

</script>

<style scoped>

.gallery img {
  width: 260px;
  height: 260px;
  object-fit: cover;
}

.supplier {
  border-radius: 10px;
  min-height: 260px;
  min-width: 260px;
}

.time {
  margin: 0 5px 0 5px !important;
}

</style>