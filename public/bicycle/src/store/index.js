import Vue from "vue";
import Vuex from "vuex";
import cartState from "./cart"

Vue.use(Vuex);

const modules = {
  cart: cartState
}

export default new Vuex.Store({
  state() {
    return {
      restaurants: [],
      items: [],
      loading: false,
      showLogin: false,
      showProduct: false,
    };
  },
  modules,
  getters() {
    return {
      getItems: (state) => {
        return state.items
      },
      getRestaurants: (state) => {
        return state.restaurants
      },
    };
  },
});
