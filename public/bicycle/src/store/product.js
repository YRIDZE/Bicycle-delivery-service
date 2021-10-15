import axios from "axios";

const state = {
  products: [],
  productTypes: [],
  pageNumber: 0,
  showProduct: false,
};

const mutations = {
  setProduct(state, items) {
    state.products = items;
  },
  changePageNumber(state, value) {
    state.pageNumber += value;
  },
  getProductTypes(state) {
    return new Promise((resolve, reject) => {
      axios
        .get("getProductTypes")
        .then((response) => {
          state.productTypes = response.data;
          resolve(response);
        })
        .catch((error) => reject(error));
    });
  },
};

const actions = {
  setProduct(context, items) {
    context.commit("setProduct", items);
  },
  getProductTypes(context) {
    context.commit("getProductTypes");
  },
  changePageNumber(context, value) {
    context.commit("changePageNumber", value);
  },
};

const getters = {
  getProducts: (state) => state.products,
  getProductsTypes: (state) => state.productTypes,
  getPage: (state) => state.pageNumber,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
