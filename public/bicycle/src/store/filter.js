const state = {
  supplierTypeFilter: [],
  productTypeFilter: [],
  supplierTimeFilter: null,
};

const mutations = {
  setSuppTypeFilter(state, value) {
    state.supplierTypeFilter = value;
  },
  setProdTypeFilter(state, value) {
    state.productTypeFilter = value;
  },
  setSuppTimeFilter(state, value) {
    state.supplierTimeFilter = value;
  },
};

const actions = {
  setSuppTypeFilter(context, value) {
    context.commit("setSuppTypeFilter", value);
  },
  setProdTypeFilter(context, value) {
    context.commit("setProdTypeFilter", value);
  },
  setSuppTimeFilter(context, value) {
    context.commit("setSuppTimeFilter", value);
  },
};

const getters = {
  getSuppTypeFilter: (state) => state.supplierTypeFilter,
  getProdTypeFilter: (state) => state.productTypeFilter,
  getSuppTimeFilter: (state) => state.supplierTimeFilter,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
