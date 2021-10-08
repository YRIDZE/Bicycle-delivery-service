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
        id: state.cart.id,
        product_id: payload.product_id,
        quantity: payload.quantity,
      };
      state.cart.products.push(entry);
    } else {
      if (entry.quantity === 1 && payload.quantity < 0) return;
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
        .post("http://localhost:8081/createCartProduct", cart)
        .then(response => {
          resolve(response);
        })
        .catch(error => {
          reject(error);
        })
    })
  },
  updateCartProduct(state, payload) {
    let cart = {
      id: 1,
      products: state.cart.products.filter(prod => prod.product_id == payload.product_id),
    };
    return new Promise((resolve, reject) => {
      axios
        .put("http://localhost:8081/updateCart", cart)
        .then(response => {
          resolve(response);
        })
        .catch(error => {
          reject(error);
        })
    })
  },
  createCart() {
    return new Promise((resolve, reject) => {
        axios
          .post("http://localhost:8081/createCart", {user_id: user.state.user_id})
          .then(response => {
            state.cart.id = response.data.id
            resolve(response);
          })
          .catch(error => {
            reject(error);
          })
      }
    )
  },
  getCart() {
    return new Promise((resolve, reject) => {
        axios
          .post("http://localhost:8081/getCartProducts")
          .then(response => {
            resolve(response);
          })
          .catch(error => {
            reject(error);
          })
      }
    )
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
    context.commit("updateCartProduct", payload);
  },
  createCart(context) {
    context.commit("createCart");
  },
}

const getters = {
  getCartList: state => state.cart.products,
  getCurrentItem: state => state.currentItem,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}