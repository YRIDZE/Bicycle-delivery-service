import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import VueFinalModal from "vue-final-modal";
import PortalVue from 'portal-vue'
import PulseLoader from 'vue-spinner/src/BeatLoader'


import Cart from './components/Cart'
import FilterList from './components/FilterPanel'
import GoBack from './components/GoBack'
import GoTop from './components/GoTop'
import Login from "./components/Login";
import Footer from "./components/Footer";
import Header from "./components/Header";
import ProductPopUp from "./components/ProductPopUp";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";

import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {library} from "@fortawesome/fontawesome-svg-core";
import {faFacebookF, faInstagram, faLinkedinIn, faTwitter} from "@fortawesome/free-brands-svg-icons"
import {
  faBicycle,
  faChevronCircleLeft,
  faChevronUp,
  faMinusCircle,
  faPlusCircle,
  faShoppingBasket,
  faTrash,
  faUser,
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faTwitter,
  faInstagram,
  faFacebookF,
  faLinkedinIn,
  faShoppingBasket,
  faUser,
  faPlusCircle,
  faMinusCircle,
  faTrash,
  faBicycle,
  faChevronCircleLeft,
  faChevronUp
);

Vue.config.productionTip = false;

Vue.component("font-awesome-icon", FontAwesomeIcon);
Vue.component("login", Login);
Vue.component("go-top", GoTop);
Vue.component("go-back", GoBack);
Vue.component("filter-list", FilterList);
Vue.component("cart", Cart);
Vue.component("bottom-footer", Footer);
Vue.component("header-top", Header);
Vue.component("product-popup", ProductPopUp);
Vue.component('pulse-loader', PulseLoader);


Vue.use(VueFinalModal)
Vue.use(PortalVue)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");

