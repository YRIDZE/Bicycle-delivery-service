import axios from "axios";

const state = {
  products: [],
  productTypes: [],
  productTypesBySupp: [],
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
  getProductTypesBySupp(state, payload) {
    return new Promise((resolve, reject) => {
      axios
        .get(`getProductTypesBySupp?id=${payload}`)
        .then((response) => {
          state.productTypesBySupp = response.data;
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
  getProductTypesBySupp(context, payload) {
    context.commit("getProductTypesBySupp", payload);
  },
  changePageNumber(context, value) {
    context.commit("changePageNumber", value);
  },
};

const getters = {
  getProducts: (state) => state.products,
  getProductsTypes: (state) => state.productTypes,
  getProductsTypesBySupp: (state) => state.productTypesBySupp,
  getPage: (state) => state.pageNumber,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
