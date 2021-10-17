<template>
  <font-awesome-icon class="go-back" style="color: #e96b56; " @click="goBack" :icon="['fa', 'chevron-circle-left']"/>

</template>

<script>
export default {
  methods: {
    goBack() {
        return this.$router.go(-1)
    }
  }
}
</script>

<style scoped>
.go-back {
  position: fixed;
  width: 40px;
  height: 40px;
  bottom: 100px;
  left: 15px;
  border-radius: 100%;
  align-items: center;
  justify-content: center;
  display: flex;
  background-color: #545454;
  cursor: pointer;
  z-index: 1000;
  opacity: .5;
  border: 0;
}

.go-back:hover,
.go-back:focus {
  filter: brightness(1.15);
  opacity: 1;
}
</style>