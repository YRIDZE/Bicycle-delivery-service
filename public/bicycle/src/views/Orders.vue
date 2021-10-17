<template>
  <vue-final-modal
      v-model="$store.state.orders.showOrders"
      classes="modal-container"
      content-class="modal-content"
      :esc-to-close="true"
      @before-open="hide"
  >
    <ol class="list-group list-block text-lg justify-content-center">
      <order-list
          v-for="(item, index) in $store.getters['orders/getReverseOrders']"
          :item="item" :key="index" :index="index">
      </order-list>
      <div class="text-center italic justify-center" v-if="$store.getters['orders/getOrders'] === null">not a single
        order has been made
      </div>
    </ol>
  </vue-final-modal>
</template>

<script>
import HideModals from "@/mixins/hideModals";

export default {
  mixins: [HideModals],
};
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
  width: 40rem;
  min-height: 6.25rem;
  max-height: 40rem;
}
</style>