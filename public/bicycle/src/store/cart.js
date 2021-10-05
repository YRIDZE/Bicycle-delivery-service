const state = {
  cartList: [],
  currentItem: null,
  showCart: false,
}

const mutations = {
  addItem(state, payload) {
    let entry = state.cartList.find(x => x.product_id == payload.product_id);
    if (entry == null) {
      entry = {
        product_id: payload.product_id,
        quantity: payload.quantity,
      };
      state.cartList.push(entry);
    } else {
      if (entry.quantity === 1 && payload.quantity < 0) return;
      entry.quantity += payload.quantity;
    }
  },
  removeItem(state, payload) {
    state.cartList.splice(payload.index, 1);
  },
}

const actions = {
  addItem(context, quantity) {
    context.commit("addItem", quantity)
  },
  removeItem(context, id) {
    context.commit("removeItem", id)
  },
}

const getters = {
  getCartList: (state) => {
    return state.cartList
  },
  getCurrentItem: (state) => {
    return state.currentItem
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}