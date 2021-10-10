import axios from "axios";
import jwt_decode from "jwt-decode";

const state = {
  showLogin: false,
  user_id: 0,

  token: localStorage.getItem("access_token") || "",
  refreshTask: null,
  status: "",
}

const mutations = {
  reg_success(state) {
    state.status = "success";
  },
  auth_success(state, token) {
    state.status = "success";
    state.token = token;
  },
  auth_error(state) {
    state.status = "error";
  },
  logout(state) {
    state.status = "";
    state.token = "";
  },
  refreshTask(state, task) {
    state.refreshTask = task;
  },
  cancelTask(state) {
    clearTimeout(state.refreshTask);
  },
}

const actions = {
  async refreshTokens(context) {
    try {
      const response = await axios.post("http://localhost:8081/refresh")

      const access_token = response.data.access_token;
      const refresh_token = response.data.refresh_token;

      localStorage.setItem("access_token", access_token);
      localStorage.setItem("refresh_token", refresh_token);

      axios.defaults.headers.common["Authorization"] = `Bearer ${access_token}`;

      context.commit("auth_success", access_token);
      await context.dispatch("dropRefresh")
    } catch {
      context.commit("auth_error");
      localStorage.removeItem("access_token");
    }

    await context.dispatch("autoRefresh");
  },

  autoRefresh(context) {
    if (state.token) {
      let token = jwt_decode(state.token);
      let now = Math.round(Date.now() / 1000);
      let timeUntilRefresh = token.exp - now;
      const refreshTask = setTimeout(() => context.dispatch("refreshTokens"), timeUntilRefresh * 1000);
      context.commit("refreshTask", refreshTask);
    }
  },

  dropRefresh(context) {
    let refresh = localStorage.getItem("refresh_token");
    let token = jwt_decode(refresh);
    setTimeout(() => context.dispatch("logout"), token.exp * 1000);
  },

  login(context, user) {
    return new Promise((resolve, reject) => {
      axios
        .post("http://localhost:8081/login", user)
        .then(response => {
          const access_token = response.data.access_token;
          const refresh_token = response.data.refresh_token;

          localStorage.setItem("access_token", access_token);
          localStorage.setItem("refresh_token", refresh_token);

          axios.defaults.headers.common["Authorization"] = `Bearer ${access_token}`;

          context.commit("auth_success", access_token);
          context.dispatch("autoRefresh");
          context.dispatch("dropRefresh");

          resolve(response);
        })
        .catch(error => {
          context.commit('auth_error');
          localStorage.removeItem("access_token");
          reject(error);
        })
    })
  },

  registration(context, user) {
    return new Promise((resolve) => {
      axios
        .post("http://localhost:8081/createUser", user)
        .then(response => {
          state.user_id = response.data.id
          context.commit("reg_success");
          resolve(response);
        })
    })
  },

  logout(context) {
    return new Promise((resolve) => {
      context.commit("logout");
      axios.post("http://localhost:8081/logout");

      localStorage.removeItem("access_token");
      localStorage.removeItem("refresh_token");

      delete axios.defaults.headers.common["Authorization"];
      context.commit("cancelTask");
      resolve();
    })
  },
}

const getters = {
  isLoggedIn: state => !!state.token,
  id: state => state.user_id,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}