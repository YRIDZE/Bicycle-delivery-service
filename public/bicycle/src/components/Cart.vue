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
            <cart-tr
                v-for="(item, index) in this.$store.getters['cart/getCartList']"
                :item="item" :key="index" :index="index">
            </cart-tr>
            </tbody>
          </table>
        </div>
      </div>
      <div class="d-flex flex-row-reverse">
        <button class="cart-btn" @click="showCartInfo = true">Order</button>
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
        <div class="row g-3" style="margin: 1px; width: 800px !important;">
          <form>
            <div class="col-sm-12 fields">
              <input type="text" class="login-input" style="width: 49% !important;" v-model="orderForm.customer_name"
                     placeholder="Name" value=""
                     required="">
              <input type="text" class="login-input" style="width: 49% !important;"
                     v-model="orderForm.customer_lastname" autocomplete='off'
                     placeholder="Surname" value="" required="">
            </div>

            <div class="cart-col-12">
              <input type="text" class="login-input" autocomplete='off' v-model="orderForm.address"
                     placeholder="Kharkiv, st. Academica Pavlova 154, apt. 12" required="">
              <div class="invalid-feedback"> Please enter your address.</div>
            </div>

            <div class="cart-col-12 fields">
              <input type="text" id="phone" v-model="orderForm.phone_number" class="login-input" autocomplete='off'
                     style="width: 49%" placeholder="Phone number" required/>
              <select class="login-input" v-model="orderForm.payment_method" autocomplete='off' style="width: 49%"
                      required="">
                <option value="" disabled selected hidden>Payment method</option>
                <option>Credit Card</option>
                <option>Cash</option>
              </select>
            </div>

          </form>
        </div>
      </div>
      <div class="btn-group">
        <button class="cart-btn" style="width:50%">Back</button>
        <button class="cart-btn" @click="orderSet" style="width:50%">Confirm</button>
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
      orderForm: {
        address: '',
        phone_number: '',
        customer_name: '',
        customer_lastname: '',
        payment_method: '',
        order_cost: 1.0,
        products: {},
      },
    };
  },
  mounted() {
    let im = new Inputmask("+38(999)-999-99-99");
    im.mask(document.getElementById('phone'));
  },

  methods: {
    orderSet() {
      this.orderForm.products = this.$store.getters["cart/getCartList"];
      this.orderForm.phone_number = this.orderForm.phone_number.replace(/[^0-9]/g, '');

      axios.post("http://localhost:8081/createOrder", JSON.stringify(this.orderForm), {
        headers:
            {Authorization: `Bearer ${this.$store.state.accessToken}`}
      })
          .then(response => console.log(response.data.id));
      this.hide();
    },
  },

  computed: {
    total: function () {
      let total = 0;
      this.$store.getters["cart/getCartList"].forEach(cartItem => total += this.$store.getters["item/getItems"].find(x => x.id === cartItem.product_id).price * cartItem.quantity);
      return total.toFixed(2).toString().replace(/\B(?=(\d{3})+$)/g, ',');
    },
  },
}
</script>

<style scoped>

.fields {
  display: flex;
  justify-content: space-between;
}

</style>