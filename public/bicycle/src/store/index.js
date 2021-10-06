import Vue from "vue";
import Vuex from "vuex";
import cartState from "./cart"
import supplier from "./supplier"
import filter from "./filter"
import item from "./item"

Vue.use(Vuex);

const modules = {
  cart: cartState,
  supp: supplier,
  filter: filter,
  item: item,
}

export default new Vuex.Store({
  modules,
  state() {
    return {
      loading: false,
      showLogin: false,
      accessToken: "",
    };
  },

});
