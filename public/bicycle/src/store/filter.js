const state = {
  supplierTypeFilter: [],
  productTypeFilter: [],
}

const mutations = {
  setSuppTypeFilter(state, value) {
    state.supplierTypeFilter = value;
  },
  setProdTypeFilter(state, value) {
    state.productTypeFilter = value;
  },
}

const actions = {
  setSuppTypeFilter(context, value) {
    context.commit("setSuppTypeFilter", value)
  },
  setProdTypeFilter(context, value) {
    context.commit("setProdTypeFilter", value)
  },
}

const getters = {
  getSuppTypeFilter: (state) => {
    return state.supplierTypeFilter
  },
  getProdTypeFilter: (state) => {
    return state.productTypeFilter
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}