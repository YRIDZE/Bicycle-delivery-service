<template>
  <div>
    <div class="modal fade" id="cart-modal" tabindex="-1" aria-hidden="true"
         style="    font-family: 'Montserrat', sans-serif;">
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content"
             style="margin: 0 auto !important;
         border-radius: 20px !important;
         border-color:rgba(117, 190, 218, 0.0);
         width: auto;">
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
            <button class="cart-btn" data-bs-target="#order-data-model" data-bs-toggle="modal"
                    data-bs-dismiss="modal">Order
            </button>
            <p style="margin-top: 15px !important;">Total <strong>{{ total }}</strong></p>
          </div>
        </div>
      </div>
    </div>
    <div class="modal fade" id="order-data-model" aria-hidden="true" tabindex="-1">
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content"
             style="border-radius:20px !important;
         border-color:rgba(117, 190, 218, 0.0);">
          <div class="modal-body">
            <div class="row g-3">

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
            <button class="cart-btn" style="width:50%"
                    data-bs-target="#cart-modal"
                    data-bs-toggle="modal"
                    data-bs-dismiss="modal">Back
            </button>
            <button class="cart-btn" style="width:50%"
                    data-bs-target="#cart-modal"
                    data-bs-toggle="modal"
                    data-bs-dismiss="modal">Confirm
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CartTr from './CartTr'
import Inputmask from 'inputmask';

export default {
  data() {
    return {
      productItem: []
    }
  },
  mounted () {
    let im = new Inputmask("+38(999)-999-99-99");
    im.mask(document.getElementById('phone'));
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