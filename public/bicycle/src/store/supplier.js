const state = {
  suppliers: [],
  loading: false,
}

const mutations = {
  setSuppliers(state, suppliers) {
    state.suppliers = suppliers;
  },
}

const actions = {
  setSuppliers(context, suppliers) {
    context.commit("setSuppliers", suppliers)
  },
}

const getters = {
  getSuppliers: state => state.suppliers,

  getSuppliersTypes: (state) => {
    let result = new Set();
    state.suppliers.forEach(element => result.add(element.type));
    return result;
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}