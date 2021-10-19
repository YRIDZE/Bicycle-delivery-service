<template id="backToTop">
  <div class="vue-back-to-top" v-show="visible"
       @click="scrollTop">
    <font-awesome-icon class="w-7 h-7" style="color: #545454" :icon="['fa', 'chevron-up']"/>

  </div>
</template>


<script>
export default {
  data() {
    return {
      visible: false
    }
  },
  methods: {
    scrollTop: function () {
      window.scroll(0, 0)
    },
    scrollListener: function () {
      this.visible = window.scrollY > 250
    }
  },
  mounted: function () {
    window.addEventListener('scroll', this.scrollListener)
  },
  unmounted: function () {
    window.removeEventListener('scroll', this.scrollListener)
  }
}
</script>

<style scoped>
.vue-back-to-top {
  position: fixed;
  width: 40px;
  height: 40px;
  bottom: 100px;
  right: 15px;
  border-radius: 100%;
  align-items: center;
  justify-content: center;
  display: flex;
  background-color: #e96b56;
  cursor: pointer;
  z-index: 1000;
  opacity: .5;
}

.vue-back-to-top:hover,
.vue-back-to-top:focus {
  filter: brightness(1.15);
  opacity: 1;
}
</style>