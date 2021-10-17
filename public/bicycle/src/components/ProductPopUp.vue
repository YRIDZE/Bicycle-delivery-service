<template>
  <vue-final-modal
      v-model="$store.state.prod.showProduct"
      classes="modal-container"
      content-class="modal-content"
      name="item"
      :esc-to-close="true"
      @before-open="hide; quantity = 1"
  >
    <div class="modal-body">
      <div class="d-flex flex-column item-container product-from">
        <div class="img-part gallery"><img :src="item.image" class="img-f" alt="item"></div>
        <div class="product-data">
          <div class="data">
            <h1>{{ item.name }}</h1>
            <b>Ingredients: </b>
            <a>{{ item.ingredients.join(', ') }}</a>
          </div>
        </div>
        <div class="d-flex bd-highlight justify-content-between item-overlay-panel p-2">
          <div class="p-2 px-5">
            <p>Price: <strong>{{ item.price }}$</strong></p>
          </div>
          <div>
            <div class="px-2.5">
              <div class="number ml-36" data-step="1" data-min="1" data-max="100">
                <input v-model="quantity" class="number-text" type="text" name="quantity" readonly>
                <a @click="reduce()" class="number-minus">−</a>
                <a @click="add()" class="number-plus">+</a>
              </div>
            </div>
          </div>
          <div class="px-5 d-flex justify-content-end">
            <p class="px-2 pt-1">Total: <strong>{{ total }}$</strong></p>
              <button class="add-to-cart-btn"
                      @click="addToCart(item.id, quantity, item.price)">ADD TO CART
              </button>
          </div>
        </div>
      </div>
    </div>
  </vue-final-modal>
</template>

<script>
import HideModals from "@/mixins/hideModals";
import {mapActions} from "vuex";

export default {
  props: {
    item: {
      type: Object,
      required: true,
    },
  },
  mixins: [HideModals],
  data() {
    return {
      quantity: 1,
    };
  },

  computed: {
    total: function () {
      return (this.item.price * parseInt(this.quantity, 10)).toFixed(2);
    },
  },

  methods: {
    ...mapActions('cart', ['addProduct']),

    addToCart: function (id, quantity, price) {
      if (!this.$store.getters["user/isLoggedIn"]) {
        this.$store.state.user.showLogin = true;
        return
      }
      this.addProduct({
            cart_id: this.$store.getters["cart/getCartId"],
            product_id: id,
            quantity: quantity,
            price: price,
          }
      );
      this.$vfm.hide("item");
    },
    reduce: function () {
      if (parseInt(this.quantity, 10) === 1) return;
      this.quantity--;
    },
    add: function () {
      this.quantity++;
    },
  },
};
</script>

<style scoped>
.gallery img {
  width: 800px !important;
  height: 404px;
  object-fit: cover;
}
</style>
