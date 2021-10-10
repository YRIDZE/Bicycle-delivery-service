import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "./assets/css/tailwind.css"


import VueFinalModal from "vue-final-modal";
import PortalVue from 'portal-vue'
import PulseLoader from 'vue-spinner/src/BeatLoader'


import Cart from './views/Cart'
import SupplierFilter from './components/SupplierFilter'
import ProductFilter from './components/ProductFilter'
import GoBack from './components/GoBack'
import GoTop from './components/GoTop'
import Login from "./views/Login";
import Footer from "./components/Footer";
import Header from "./components/Header";
import ProductPopUp from "./components/ProductPopUp";
import CartTr from './components/CartTr'
import Check from './components/Check'
import Product from "@/components/Product";
import Supplier from '@/components/Supplier'

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
  faSignInAlt,
  faSignOutAlt,
  faTrash
} from "@fortawesome/free-solid-svg-icons";

library.add(
  faTwitter,
  faInstagram,
  faFacebookF,
  faLinkedinIn,
  faShoppingBasket,
  faPlusCircle,
  faMinusCircle,
  faTrash,
  faSignInAlt,
  faSignOutAlt,
  faBicycle,
  faChevronCircleLeft,
  faChevronUp
);

Vue.config.productionTip = false;

Vue.component("font-awesome-icon", FontAwesomeIcon);
Vue.component("login", Login);
Vue.component("go-top", GoTop);
Vue.component("go-back", GoBack);
Vue.component("filter-list", SupplierFilter);
Vue.component("cart", Cart);
Vue.component("bottom-footer", Footer);
Vue.component("header-top", Header);
Vue.component("product-popup", ProductPopUp);
Vue.component('pulse-loader', PulseLoader);
Vue.component('cart-tr', CartTr);
Vue.component('check', Check);
Vue.component('product', Product);
Vue.component('supplier', Supplier);
Vue.component('product-filer', ProductFilter);

Vue.use(VueFinalModal)
Vue.use(PortalVue)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");

