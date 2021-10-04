import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state() {
    return {
      restaurants: [],
      items: [],
      cartList: [],
      currentProduct: null,
      loading: false,
      showLogin: false,
      showCart: false,
      showProduct: false,
    };
  },
  mutations: {
    addProduct(state, payload) {
      let entry = state.cartList.find((x) => x.id == payload.id);
      if (entry == null) {
        entry = {
          id: payload.id,
          quantity: payload.quantity,
        };
        state.cartList.push(entry);
      } else {
        entry.quantity += payload.quantity;
      }
    },

    remove(state, payload) {
      state.cartList.splice(payload.index, 1);
    },

    setCart(state, cart) {
      state.cartList.push(cart);
    }
  },
  actions: {
    addProduct(context, quantity) {
      context.commit("addProduct", quantity)
    },

    remove(context, id) {
      context.commit("remove", id)
    }
  },

  getters: {
    getCartList: (state) => {
      return state.cartList
    }
  },
});
