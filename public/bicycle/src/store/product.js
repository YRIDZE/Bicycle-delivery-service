import axios from "axios";

const state = {
  products: [],
  productTypes: [],
  showProduct: false,
}

const mutations = {
  setProduct(state, items) {
    state.products = items;
  },
  getProductTypes(state) {
    return new Promise((resolve, reject) => {
        axios
          .get("http://localhost:8081/getProductTypes")
          .then(response => {
            state.productTypes = response.data;
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
  setProduct(context, items) {
    context.commit("setProduct", items);
  },
  getProductTypes(context) {
    context.commit("getProductTypes");
  },
}

const getters = {
  getProducts: state => state.products,
  getProductsTypes: state => state.productTypes,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}