// let cart = localStorage.getItem('cart');

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
  saveCart(state) {
    localStorage.setItem('cart', JSON.stringify(state.cartList));
  }
}

const actions = {
  addItem(context, quantity) {
    context.commit("addItem", quantity)
    context.commit("saveCart");
  },
  removeItem(context, id) {
    context.commit("removeItem", id)

    if (state.cartList.length > 0) {
      context.commit("saveCart");
    } else localStorage.removeItem('cart')
  },
}

const getters = {
  getCartList: state => state.cartList,
  getCurrentItem: state => state.currentItem,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}