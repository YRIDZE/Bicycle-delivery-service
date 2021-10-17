import axios from "axios";
import jwt_decode from "jwt-decode";

const state = {
  showLogin: false,
  user_id: 0,

  refresh_token: localStorage.getItem("refresh_token") || "",
  access_token: localStorage.getItem("access_token") || "",

  refreshTask: null,
  logoutTask: null,

  status: "",
}

const mutations = {
  reg_success(state) {
    state.status = "success";
  },
  auth_success(state, token) {
    axios.defaults.headers.common["Authorization"] = `Bearer ${token.access_token}`;
    localStorage.setItem("access_token", token.access_token);
    localStorage.setItem("refresh_token", token.refresh_token);
    state.access_token = token.access_token;
    state.refresh_token = token.refresh_token;
    state.status = "success";
  },
  auth_error(state) {
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    state.status = "error";
  },
  logout(state) {
    delete axios.defaults.headers.common["Authorization"];
    localStorage.removeItem("access_token");
    localStorage.removeItem("refresh_token");
    state.access_token = "";
    state.refresh_token = "";
    state.status = "";
  },
  refreshTask(state, task) {
    state.refreshTask = task;
  },
  logoutTask(state, task) {
    state.logoutTask = task;
  },
  cancelTask(state) {
    clearTimeout(state.logoutTask)
    clearTimeout(state.refreshTask);
  },
}

export function getTokenTimeUntilRefresh(t) {
  let token = jwt_decode(t);
  let now = Math.round(Date.now() / 1000);
  return token.exp - now;
}

const actions = {
  async refreshTokens(context) {
    try {
      const response = await axios.get("refresh");
      const access_token = response.data.access_token;
      const refresh_token = response.data.refresh_token;

      context.commit("auth_success", {
        access_token: access_token,
        refresh_token: refresh_token,
      });
      context.commit("cancelTask");
    } catch {
      context.commit("auth_error");
    }

    await context.dispatch("dropRefresh");
    await context.dispatch("autoRefresh");
  },

  async autoRefresh(context) {
    if (state.access_token) {
      let timeUntilRefresh = getTokenTimeUntilRefresh(state.access_token)

      if (timeUntilRefresh < 1)
        await context.dispatch("refreshTokens")

      const refreshTask = setTimeout(() => context.dispatch("refreshTokens"), timeUntilRefresh * 1000);
      context.commit("refreshTask", refreshTask);
    }
  },

  async dropRefresh(context) {
    if (state.refresh_token) {
      let timeUntilRefresh = getTokenTimeUntilRefresh(state.refresh_token);
      const logoutTask = setTimeout(() => context.dispatch("logout"), timeUntilRefresh * 1000);
      context.commit("logoutTask", logoutTask);
    }
  },

  login(context, user) {
    return new Promise((resolve, reject) => {
      axios
        .post("login", user)
        .then(response => {
          const access_token = response.data.access_token;
          const refresh_token = response.data.refresh_token;

          context.commit("auth_success", {
            access_token: access_token,
            refresh_token: refresh_token
          });

          context.dispatch("autoRefresh");
          context.dispatch("dropRefresh");

          resolve(response);
        })
        .catch(error => {
          context.commit('auth_error');
          localStorage.removeItem("access_token");
          reject(error);
        });
    });
  },

  registration(context, user) {
    return new Promise((resolve, reject) => {
      axios
        .post("createUser", user)
        .then(response => {
          state.user_id = response.data.id;
          context.commit("reg_success");
          resolve(response);
        })
        .catch(error => reject(error));
    });
  },

  logout(context) {
    if (state.refresh_token) {
      return new Promise((resolve) => {
        axios.post("logout");
        context.commit("logout");
        context.commit("cancelTask");
        resolve();
      });
    }
  },
}

const getters = {
  isLoggedIn: (state) => !!state.refresh_token,
  id: (state) => state.user_id,
  accessToken: (state) => state.access_token,
  refreshToken: (state) => state.refresh_token,
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
}