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
          <table class="table table-striped custom-table m-0">
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
            <cart-tr
                v-for="(item, index) in this.$store.getters['cart/getCartList']"
                :item="item" :key="index" :index="index">
            </cart-tr>
            </tbody>
          </table>
          <div class="text-center mb-3 italic" v-if="this.$store.getters['cart/getCartList'].length === 0">empty cart
          </div>
        </div>
      </div>
      <div class="d-flex flex-row-reverse my-2" v-if="this.$store.getters['cart/getCartList'].length !== 0">
        <button class="cart-btn font-black mr-3" @click="showCartInfo = true">Order</button>
        <p class="mx-2">Total <strong>{{ total }}</strong>$</p>
      </div>
    </vue-final-modal>

    <vue-final-modal
        v-model="showCartInfo"
        name="cartInfo"
        classes="modal-container"
        content-class="modal-content"
        :esc-to-close="true"
    >
      <div class="modal-body">
        <form v-on:submit.prevent="showCartOrder = true">
          <div class="row g-3 m-px" style="width: 800px !important;">
            <div class="col-sm-12 fields">
              <input type="text" class="login-input m-0" style="width: 49%" v-model="orderForm.customer_name"
                     placeholder="Name" autocomplete="off" required>
              <input type="text" class="login-input m-0" style="width: 49%" autocomplete="off"
                     v-model="orderForm.customer_lastname" placeholder="Surname" required>
            </div>

            <div class="cart-col-12">
              <input type="text" class="login-input m-0" v-model="orderForm.address" autocomplete="off"
                     placeholder="Kharkiv, st. Academica Pavlova 154, apt. 12" required>
              <div class="invalid-feedback"> Please enter your address.</div>
            </div>

            <div class="cart-col-12 fields">
              <input type="text" id="phone" v-model="orderForm.phone_number" class="login-input m-0"
                     style="width: 49%" placeholder="Phone number" autocomplete="off" required/>
              <select class="login-input m-0" v-model="orderForm.payment_method" style="width: 49%" required>
                <option value="" disabled selected hidden>Payment method</option>
                <option>Credit Card</option>
                <option>Cash</option>
              </select>
            </div>
            <button class="cart-btn font-black my-2" type="submit" style="font-size: 16px">Confirm</button>
          </div>
        </form>
      </div>
    </vue-final-modal>
    <vue-final-modal
        v-model="showCartOrder"
        name="cartInfo"
        classes="modal-container"
        content-class="modal-content"
        :esc-to-close="true"
    >
      <div class="modal-body">

        <div class="row g-3 m-px" style="width: 800px !important;">
          <p>{{ orderForm.customer_lastname }} {{ orderForm.customer_name }}: {{ orderForm.address }}</p>
          <small class="m-0">{{ orderForm.phone_number }}</small>
          <ol class="list-group list-group-numbered px-2">
            <order-list-products v-for="(i) in this.$store.getters['cart/getCartList']" :key="i"
                                 :item="i"></order-list-products>
          </ol>
          <p class="text-right font-black mt-2">by {{ orderForm.payment_method }} <strong>{{ total }}</strong>$</p>
          <button class="cart-btn font-black my-2" @click="createOrder" type="submit" style="font-size: 16px">
            ORDER
          </button>
        </div>

      </div>
    </vue-final-modal>

  </div>
</template>

<script>
import HideModals from '../mixins/hideModals'
import Inputmask from "inputmask";
import axios from "axios";

export default {
  mixins: [HideModals],

  data() {
    return {
      showCartInfo: false,
      showCartOrder: false,
      orderForm: {
        address: '',
        phone_number: '',
        customer_name: '',
        customer_lastname: '',
        payment_method: '',
        products: {},
      },
    };
  },
  mounted() {
    let im = new Inputmask("+38(999)-999-99-99");
    im.mask(document.getElementById('phone'));
  },

  methods: {
    createOrder() {
      this.orderForm.products = this.$store.getters["cart/getCartList"];
      this.orderForm.phone_number = this.orderForm.phone_number.replace(/[^0-9]/g, '');

      axios
          .post("http://localhost:8081/createOrder", JSON.stringify(this.orderForm))
          .then(() => {
            this.$store.dispatch("cart/deleteAllFromCart");
            this.$store.state.cart.cart.products = [];
          })
      this.hide();

      this.orderForm.phone_number = '';
      this.orderForm.address = '';
      this.orderForm.customer_lastname = '';
      this.orderForm.customer_name = '';
      this.orderForm.payment_method = '';
    },
  },

  computed: {
    total: function () {
      let total = 0;
      this.$store.getters["cart/getCartList"]
          .forEach(cartItem => total += this.$store.getters["prod/getProducts"]
              .find(x => x.id === cartItem.product_id).price * cartItem.quantity);

      return total.toFixed(2).toString().replace(/\B(?=(\d{3})+$)/g, ',');
    },
  },
};
</script>

<style scoped>

.fields {
  display: flex;
  justify-content: space-between;
}

</style>