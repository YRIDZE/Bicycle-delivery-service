const state = {
  items: [],
  showProduct: false,
}

const mutations = {
  setItem(state, items) {
    state.items = items;
  },
}

const actions = {
  setItem(context, items) {
    context.commit("setItem", items)
  },
}

const getters = {
  getItems: (state) => {
    return state.items
  },
  getItemsTypes: (state) => {
    let result = new Set();
    state.items.forEach(element => result.add(element.type));
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