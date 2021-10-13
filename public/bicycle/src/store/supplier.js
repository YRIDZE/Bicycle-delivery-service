import axios from "axios";

const state = {
  suppliers: [],
  supplierTypes: [],
  loading: false,
}

const mutations = {
  setSuppliers(state, suppliers) {
    state.suppliers = suppliers;
  },
  getSupplierTypes(state) {
    return new Promise((resolve, reject) => {
        axios
          .get("http://localhost:8081/getSupplierTypes")
          .then(response => {
            state.supplierTypes = response.data;
            resolve(response);
          })
          .catch(error => {
            reject(error);
          })
      }
    )
  }
}

const actions = {
  setSuppliers(context, suppliers) {
    context.commit("setSuppliers", suppliers)
  },
  getSupplierTypes(context) {
    context.commit("getSupplierTypes")
  }
}

const getters = {
  getSuppliers: state => state.suppliers,
  getSuppliersTypes: state => state.supplierTypes,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}