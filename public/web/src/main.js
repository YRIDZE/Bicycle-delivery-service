import {createApp} from 'vue'
import {createStore} from 'vuex'

import App from './App.vue'
import router from './router'
import VueFinalModal from "vue-final-modal";

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

library.add(faTwitter, faInstagram, faFacebookF, faLinkedinIn, faShoppingBasket, faUser, faPlusCircle, faMinusCircle, faTrash, faBicycle, faChevronCircleLeft, faChevronUp);

const store = createStore({
    state() {
        return {
            restaurants: [],
            items: [],
            cartList: [],
            loading: false,
        }
    },
    mutations: {
        add(state, payload) {
            let entry = state.cartList.find(x => x.id == payload.id)
            if (entry == null) {
                entry = {
                    id: payload.id,
                    quantity: payload.quantity
                }
                state.cartList.push(entry)
            } else {
                entry.quantity += payload.quantity
            }

        },

        removeFromCart(state, payload) {
            state.cartList.splice(payload.index, 1)
        },
    },
})


createApp(App)
    .component("font-awesome-icon", FontAwesomeIcon)
    .use(router).use(store).use(VueFinalModal)
    .mount('#app');


