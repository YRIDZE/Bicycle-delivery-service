import axios from "axios";

const state = {
  showOrders: false,
  orders: []
}

const mutations = {
  getOrders() {
    return new Promise((resolve, reject) => {
        axios
          .get("/getOrders")
          .then(response => {
            state.orders = response.data;
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
  getOrders(context) {
    context.commit("getOrders");
  },
}

const getters = {
  getOrders: state => state.orders,

}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}