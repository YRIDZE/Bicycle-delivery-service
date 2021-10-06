import axios from "axios";

const state = {
  status: '',
  token: localStorage.getItem('access_token') || '',
}

const mutations = {
  auth_request(state) {
    state.status = 'loading';
  },
  auth_success(state, token) {
    state.status = 'success';
    state.token = token;
  },
  auth_error(state) {
    state.status = 'error';
  },
  logout(state) {
    state.status = '';
    state.token = '';
  },
}

const actions = {
  login({commit}, user) {
    return new Promise((resolve, reject) => {
      axios
        .post("http://localhost:8081/login", user)
        .then(response => {
          const token = response.data.access_token;
          localStorage.setItem('access-token', response.data.access_token);
          axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
          commit('auth_success', token);
          resolve(response);
        })
        .catch(error => {
          commit('auth_error');
          localStorage.removeItem('access-token');
          reject(error);
        })
    })
  },

  registration({commit}, user) {
    return new Promise((resolve, reject) => {
      commit('auth_request');
      axios
        .post("http://localhost:8081/createUser", user)
        .then(response => {
          const token = response.data.access_token;
          localStorage.setItem('access-token', token);
          axios.defaults.headers.common['Authorization'] = `Bearer ${token}`;
          commit('auth_success', token);
          resolve(response);
        })
        .catch(err => {
          commit('auth_error', err)
          localStorage.removeItem('access-token')
          reject(err)
        })
    })
  },
  logout({commit}) {
    return new Promise((resolve) => {
      commit('logout');
      axios.post("http://localhost:8081/logout").then()
      localStorage.removeItem('access-token');
      delete axios.defaults.headers.common['Authorization'];
      resolve();
    })
  },
}

const getters = {
  isLoggedIn: state => !!state.token,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}