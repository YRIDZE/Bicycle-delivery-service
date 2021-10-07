import Vue from "vue";
import Vuex from "vuex";
import cartState from "./cart"
import supplier from "./supplier"
import filter from "./filter"
import item from "./item"
import user from "./user"

Vue.use(Vuex);

const modules = {
  cart: cartState,
  supp: supplier,
  filter: filter,
  item: item,
  user: user,
}

export default new Vuex.Store({
  modules,
});
