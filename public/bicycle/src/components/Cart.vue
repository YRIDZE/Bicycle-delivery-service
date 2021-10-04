<template>
  <div>
    <vue-final-modal
        v-model="$store.state.cart.showCart"
        classes="modal-container"
        content-class="modal-content"
        name="cart"
        :esc-to-close="true"
        @before-open="hide"

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
                v-for="(item, index) in this.$store.getters['cart/getCartList']"
                :item="item" :key="index" :index="index">
            </CartTr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="d-flex flex-row-reverse">
        <button class="cart-btn" @click="showCartInfo = true">Order
        </button>
        <p style="margin-top: 15px !important;">Total <strong>{{ total }}</strong>$</p>
      </div>
    </vue-final-modal>

    <vue-final-modal
        v-model="showCartInfo"
        classes="modal-container"
        content-class="modal-content"
        :esc-to-close="true"
    >
      <div class="modal-body">
        <div class="row g-3" style="margin: 1px">

          <div class="col-sm-6">
            <input type="text" class="login-input" placeholder="Name" value="" required="">
          </div>

          <div class="col-sm-6">
            <input type="text" class="login-input" autocomplete='off' placeholder="Surname" value="" required="">
          </div>

          <div class="cart-col-12">
            <input type="text" id="phone" class="login-input" autocomplete='off' placeholder="Phone number"
                   required/>
          </div>

          <div class="cart-col-12">
            <input type="text" class="login-input" autocomplete='off'
                   placeholder="Kharkiv, st. Academician Pavlova 154, apt. 12"
                   required="">
            <div class="invalid-feedback">
              Please enter your address.
            </div>
          </div>

          <div class="col-md-7">
            <form>
              <select class="login-input" autocomplete='off' required="">
                <option value="" disabled selected hidden>Payment method</option>
                <option>Credit Card</option>
                <option>Cash</option>
              </select>
            </form>
          </div>
        </div>
      </div>
      <div class="btn-group">
        <button class="cart-btn" style="width:50%">Back
        </button>
        <button class="cart-btn" style="width:50%">Confirm
        </button>
      </div>
    </vue-final-modal>

  </div>
</template>

<script>
import CartTr from './CartTr'
import HideModals from '../mixins/hideModals'
import Inputmask from "inputmask";

export default {
  mixins: [HideModals],
  components: {
    CartTr,
  },

  data() {
    return {
      showCartInfo: false,
    };
  },
  mounted() {
    let im = new Inputmask("+38(999)-999-99-99");
    im.mask(document.getElementById('phone'));
  },

  computed: {
    total: function () {
      let total = 0;
      this.$store.state.cart.cartList.forEach(cartItem => total += this.$store.state.items.find(x => x.id === cartItem.id).price * cartItem.quantity)
      return total.toFixed(2).toString().replace(/\B(?=(\d{3})+$)/g, ',');
    },
  },
}
</script>

<style scoped>

</style>