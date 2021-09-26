<template>
  <vue-final-modal
      v-bind="$attrs"
      name="cartTg"
      classes="modal-container"
      content-class="modal-content"
      transition="slide"
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
      <button class="cart-btn" style="width:50%"
              @click="this.$root.openC">Back
      </button>
      <button class="cart-btn" style="width:50%"
              @click="this.$root.openC">Confirm
      </button>
    </div>
  </vue-final-modal>
</template>

<script>
import Inputmask from "inputmask";

export default {
  name: "CartInfo",
  mounted() {
    let im = new Inputmask("+38(999)-999-99-99");
    im.mask(document.getElementById('phone'));
  },
}
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all .7s;
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
</style>