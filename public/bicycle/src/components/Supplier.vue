<template>
  <div class="mt-4 mx-1 supplier">
    <div class="icon-box gallery">
      <router-link :to="{ path : String(supp.id) }">
        <img :src="supp.image" style="background-color: #FFFFFF !important;" class="img-fluid" alt="Supplier picture">
      </router-link>
      <div class="supp-cont">
        <p class="short">{{ supp.name }}</p>
        <p> {{ supp.workingHours.opening.substr(0, 5) }}-{{ supp.workingHours.closing.substr(0, 5) }}</p>
      </div>
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
  min-height: 290px;
  min-width: 290px;
}

.supp-cont {
  margin: 10px 15px 0 15px !important;
}

</style>