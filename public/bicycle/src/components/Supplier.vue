<template>
  <div class="mt-4 mx-1 supplier">
    <div class="icon-box gallery">
      <router-link :to="{ path : String(supp.id) }">
        <img :src="supp.image" class="img-fluid" alt="Supplier picture">
      </router-link>
      <p class="short">{{ supp.name }}</p>
      <p class="time"> {{ supp.workingHours.opening.substr(0, 5) }}-{{ supp.workingHours.closing.substr(0, 5) }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    supp: {
      type: Object,
      required: true,
    },
  },

};

</script>

<style scoped>

.gallery img {
  width: 290px;
  height: 290px;
  object-fit: cover;
}

.supplier {
  border-radius: 10px;
  min-height: 275px;
  min-width: 275px;
}

.time {
  margin: 0 5px 0 5px !important;
}

</style>