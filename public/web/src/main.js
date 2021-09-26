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
            restaurants: [
                {
                    id: 1,
                    title: 'My journey with Vue',
                    type: "type 1",
                    logo: 'http://cdn.shopify.com/s/files/1/1578/1589/files/colorf-01_198x200.png'
                },
                {
                    id: 2,
                    title: 'Blogging with Vue',
                    type: "type 2",
                    logo: 'https://1000logos.net/wp-content/uploads/2021/04/Target-logo.png'
                },
                {
                    id: 3,
                    title: 'Why Vue is so fun',
                    type: "type 3",
                    logo: 'https://eda.ua/images/506509-195-195-burger_club_harkov.jpg'
                },
                {
                    id: 4,
                    title: 'Blogging with Vue',
                    type: "type 1",
                    logo: 'https://image.freepik.com/free-vector/sushi-restaurant-logo_8169-12.jpg'
                },
                {
                    id: 5,
                    title: 'Blogging with Vue',
                    type: "type 4",
                    logo: 'https://lineacaffe.com/wp-content/themes/lineacaffe/images/linea-logo.svg'
                },
                {
                    id: 6,
                    title: 'Blogging with Vue',
                    type: "type 5",
                    logo: 'https://play-lh.googleusercontent.com/qMewibe3u5Wvq3fBf3Ca3_QItjHCOKeGrOAzVXWxqzgRpMwxYlD5CA6M2M5L78SwNA'
                }
            ],
            items: [
                {
                    id: 1,
                    restaurantId: 1,
                    title: 'My journey with Vue',
                    ingredients: ["Salmon", "Nori", "Rice", "Cucumber", "Cream cheese", "Unagi sauce", "Japanese tamago"],
                    price: 122,
                    logo: 'https://roll-club.kh.ua/wp-content/uploads/2021/04/okean-1.jpg.webp'
                },
                {
                    id: 2,
                    restaurantId: 1,
                    title: 'My journey with Vue 2',
                    ingredients: ["Cheese", "Caramelized onions", "Tomatoes", "Original sauce"],
                    price: 122,
                    logo: 'https://roll-club.kh.ua/wp-content/uploads/2015/09/4-syra.jpg.webp'
                },
                {
                    id: 3,
                    restaurantId: 1,
                    title: 'My journey with Vue 3',
                    ingredients: ["Mozarella", "Peperoni", "Tomatoes", "BBQ sauce"],
                    price: 122,
                    logo: 'https://roll-club.kh.ua/wp-content/uploads/2019/03/kapricheza.jpg.webp'
                }
            ],
            cartList: []
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
