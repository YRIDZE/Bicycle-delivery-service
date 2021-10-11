import Vue from "vue";
import Vuex from "vuex";
import cartState from "./cart"
import supplier from "./supplier"
import filter from "./filter"
import item from "./item"
import user from "./user"
import orders from "./orders"

Vue.use(Vuex);

const modules = {
  cart: cartState,
  supp: supplier,
  filter: filter,
  item: item,
  user: user,
  orders: orders
}

export default new Vuex.Store({
  modules,
});
