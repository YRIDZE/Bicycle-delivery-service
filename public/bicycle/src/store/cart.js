import axios from "axios";
import user from "./user"

const state = {
  cart: {
    id: 0,
    products: [],
  },
  currentItem: null,
  showCart: false,
}

const mutations = {

  addItem(state, payload) {
    let entry = state.cart.products.find(x => x.product_id == payload.product_id);
    if (entry == null) {
      entry = {
        cart_id: state.cart.id,
        product_id: payload.product_id,
        quantity: payload.quantity,
        price: payload.price,
      };
      state.cart.products.push(entry);
    } else {
      if (entry.quantity === 1 && payload.quantity < 0)
        return;

      entry.quantity += payload.quantity;
    }
  },

  removeItem(state, payload) {
    state.cart.products.splice(payload.index, 1);
  },

  addCartProduct(state, payload) {
    let cart = {
      products: state.cart.products.filter(prod => prod.product_id == payload.product_id)
    }

    return new Promise((resolve, reject) => {
      axios
        .post("createCartProduct", cart)
        .then((response) => resolve(response))
        .catch((error) => reject(error))
    })
  },

  updateCartProduct(state, payload) {
    let cart = {
      id: state.cart.id,
      products: state.cart.products.filter(prod => prod.product_id == payload.product_id),
    };
    return new Promise((resolve, reject) => {
      axios
        .put("updateCart", cart)
        .then((response) => resolve(response))
        .catch((error) => reject(error))
    })
  },

  deleteCartProduct(state, payload) {
    return new Promise((resolve, reject) => {
      axios
        .delete(`deleteCartProduct?productId=${payload.product_id}`)
        .then((response) => resolve(response))
        .catch((error) => reject(error))
    })
  },
  deleteAllCartProducts() {
    return new Promise((resolve, reject) => {
      axios
        .delete("deleteAllCartProducts")
        .then((response) => resolve(response))
        .catch((error) => reject(error))
    })
  },

  createCart() {
    return new Promise((resolve, reject) => {
        axios
          .post("createCart", {user_id: user.state.user_id})
          .then((response) => resolve(response))
          .catch((error) => reject(error))
      }
    )
  },

  getCart() {
    return new Promise((resolve, reject) => {
      axios
        .get("getCartProducts")
        .then((response) => {
          state.cart.id = response.data[0].id;

          if (response.data[0].products != null)
            state.cart.products = response.data[0].products;

          resolve(response);
        })
        .catch((error) => reject(error))
    })
  },
}

const actions = {
  addProduct(context, payload) {
    context.commit("addItem", payload);
    context.commit("addCartProduct", payload);
  },

  addItem(context, payload) {
    context.commit("addItem", payload);
    context.commit("updateCartProduct", payload);
  },
  removeItem(context, payload) {
    context.commit("removeItem", payload.product_id);
    context.commit("deleteCartProduct", payload);
  },
  createCart(context) {
    context.commit("createCart");
  },
  deleteAllFromCart(context) {
    context.commit("deleteAllCartProducts");
  },
  getCart(context) {
    context.commit("getCart");
  },
}

const getters = {
  getCartList: (state) => state.cart.products,
  getCurrentItem: (state) => state.currentItem,
  getCartId: (state) => state.cart.id,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}