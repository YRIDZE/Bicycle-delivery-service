<template>
  <div>
    <vue-final-modal
        v-bind="$attrs"
        name="cart"
        classes="modal-container"
        content-class="modal-content"
        transition="slide"
    >
      <div class="modal-body">
        <div class="table-responsive">
          <table class="table table-striped custom-table">
            <thead>
            <tr>
              <th>#</th>
              <th>Product Name</th>
              <th>From</th>
              <th>Price</th>
              <th>Quantity</th>
            </tr>
            </thead>
            <tbody>
            <CartTr
                v-for="(item, index) in this.$store.state.cartList"
                :item="item" :key="index" :index="index">
            </CartTr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="d-flex flex-row-reverse">
        <button class="cart-btn" @click="openI">Order
        </button>
        <p style="margin-top: 15px !important;">Total <strong>{{ total }}</strong>$</p>
      </div>
    </vue-final-modal>
  </div>
</template>

<script>
import CartTr from './CartTr'
import CartInfo from './CartInfo'

export default {
  data() {
    return {
      showCartInfo: false,
      productItem: []
    }
  },

  methods: {
    openI() {
      this.$vfm.hide("cart")
      this.$root.showCart = false

      this.$vfm.show({
        component: CartInfo,
      });
      this.showCartInfo = true
    },
  },
  computed: {
    total: function () {
      let total = 0;
      this.$store.state.cartList.forEach(cartItem => total += this.$store.state.items.find(x => x.id == cartItem.id).price * cartItem.quantity)
      return total.toString().replace(/\B(?=(\d{3})+$)/g, ',');
    },
  },
  name: "Cart",
  components: {
    CartTr
  }
}
</script>

<style scoped>

</style>