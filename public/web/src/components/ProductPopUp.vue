<template>
  <div class="modal fade" id="item-modal" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered modal-lg">
      <div class="modal-content"
           style="
         background-color: rgba(117, 190, 218, 0.0);
         border-color:rgba(117, 190, 218, 0.0);">
        <div class="modal-body">
          <div class="d-flex flex-column item-container product-from">
            <div class="img-part">
              <img :src="item.image" class="img-f" alt="item">
            </div>
            <div class="product-data">
              <div class="data">
                <h1>{{ item.name }}</h1>
                <b>Ingredients: </b>
                <text v-for="ingredient in item.ingredients" :key="ingredient">
                  <text>{{ ingredient }}{{ ", " }}</text>
                </text>
              </div>
            </div>
            <div class="d-flex bd-highlight item-overlay-panel">
              <div class="me-auto p-2 bd-highlight">
                <div style="padding: 0 10px ">
                  <div class="number" data-step="1" data-min="1" data-max="100">
                    <input v-model="quantity" class="number-text" type="text" name="quantity" readonly>
                    <a @click="reduce()" class="number-minus">−</a>
                    <a @click="add()" class="number-plus">+</a>
                  </div>
                </div>
              </div>
              <div class="p-2 bd-highlight">
                <p style="margin-top: 15px !important;">Total: <strong>{{ total }}$</strong></p>
              </div>
              <div class="p-2 bd-highlight">
                <div style="padding: 0 10px ">
                  <button @click="addToCart(id, this.quantity)" data-bs-dismiss="modal"
                          class="add-to-cart-btn">ADD TO CART
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      quantity: 1,
    };
  },
  props: ["id"],
  computed: {
    item: function () {
      return this.$store.state.items.find(x => x.id == this.id)
    },
    total: function () {
      return this.item.price * parseInt(this.quantity, 10)
    }
  },
  methods: {
    addToCart: function (id, quantity) {
      this.$store.commit("add", {
            id: id,
            quantity: quantity
          }
      )
    },

    reduce: function () {
      if (parseInt(this.quantity, 10) === 1) return;
      this.quantity--;
    },
    add: function () {
      this.quantity++;
    },
    init: function () {
      this.quantity = 1
    }
  },
  name: "ProductPopUp"
}
</script>

<style scoped>

</style>