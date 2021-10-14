<template>
  <vue-final-modal
      v-model="$store.state.orders.showOrders"
      classes="modal-container"
      content-class="modal-content"
      :esc-to-close="true"
      @before-open="hide"
  >
    <ol class="list-group list-block">
      <order-list
          v-for="(item, index) in $store.getters['orders/getReverseOrders']"
          :item="item" :key="index" :index="index">
      </order-list>
      <div class="text-center italic" v-if="this.$store.getters['orders/getOrders'].length === 0">not a single
        order has been made
      </div>
    </ol>
  </vue-final-modal>
</template>

<script>
import HideModals from "@/mixins/hideModals";

export default {
  mixins: [HideModals],
}
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all 0.7s;
}

.slide-enter-from,
.slide-leave-to {
  transform: translateY(-100vh);
}

.modal-container {
  display: flex;
  justify-content: center;
  align-items: center;
}

.list-block {
  overflow: auto;
  width: 37.5rem;
  height: auto;
  max-height: 37.5rem;
}
</style>