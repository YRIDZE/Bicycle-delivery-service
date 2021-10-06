const state = {
  filter: {
    type: [],
  },
}

const mutations = {
  setFilter(state, value) {
    state.filter.type = value;
  },
}

const actions = {
  setFilter(context, value) {
    context.commit("setFilter", value)
  },
}

const getters = {
  getFilter: (state) => {
    return state.filter
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}