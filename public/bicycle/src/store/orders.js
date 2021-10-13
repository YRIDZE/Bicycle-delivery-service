import axios from "axios";

const state = {
  showOrders: false,
  orders: []
}

const mutations = {
  getOrders() {
    return new Promise((resolve, reject) => {
        axios
          .post("http://localhost:8081/getOrders")
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
  getOrders: state => state.orders.reverse(),

}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}