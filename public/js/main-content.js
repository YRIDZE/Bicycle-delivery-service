const App = {
    data() {
        return {
            showLogin: false,
            showCart: false,

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
}

const app = Vue.createApp(App)

const Login = {
    data() {
        return {
            signUpMode: false
        }
    },
    template: `
      <div class="modal fade" id="login-modal" tabindex="-1" aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content"
             style="
             background-color: rgba(117, 190, 218, 0.0);
             border-color:rgba(117, 190, 218, 0.0);">
          <div class="modal-body">
            <div class="login-container"
                 :class="{ 'right-panel-active': signUpMode }" id="login-container">
              <div class="form-container sign-up-container">
                <form class="l-r-form" action="#">
                  <a href="main-page.html"> <i style="color: #545454" class="fas fa-bicycle fa-3x"></i></a>
                  <h1><b>Create Account</b></h1>
                  <input class="login-input" type="text" placeholder="Firstname"/>
                  <input class="login-input" type="text" placeholder="Lastname"/>
                  <input class="login-input" type="email" placeholder="Email"/>
                  <input class="login-input" type="password" placeholder="Password"/>
                  <button class="sign-up bt" style="margin-top: 7px">SIGN UP</button>
                </form>
              </div>
              <div class="form-container  sign-in-container">
                <form class="l-r-form" action="#">
                  <a href="main-page.html"> <i style="color: #e97d56" class="fas fa-bicycle fa-3x"></i></a>
                  <h1><b>Sign in</b></h1>
                  <input class="login-input" type="email" placeholder="Email"/>
                  <input class="login-input" type="password" placeholder="Password"/>
                  <a href="#" style="text-decoration: none; font-size: 14px; margin: 15px 0; color:#000000">Forgot your
                    password?</a>
                  <button class="sign-in bt">SIGN IN</button>
                </form>
              </div>
              <div class="overlay-container">
                <div id="close" class="close" :class="{ 'left': signUpMode }">
                  <a href="#" data-bs-dismiss="modal"><i style="color: #ffffff;" class="fas fa-times"></i></a>
                </div>
                <div class="overlay">
                  <div class="overlay-panel overlay-left">
                    <h1><b>Welcome Back!</b></h1>
                    <p class="container-details">To keep connected with us please login with your personal info</p>
                    <button class="ghost bt" id="signIn" @click="signUpMode = false">SIGN IN</button>
                  </div>
                  <div style="right: 0" class="overlay-panel overlay-right">
                    <h1><b>Hello, Friend!</b></h1>
                    <p class="container-details">Enter your personal details and start journey with us</p>
                    <button class="ghost bt" id="signUp" @click="signUpMode = true">SIGN UP</button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      </div>
    `,
}

const CartTr = {
    props: ['item', 'index'],
    template: `
      <tr>
      <td>{{ index + 1 }}</td>
      <td>{{ i.title }}</td>
      <td>{{ f.title }}</td>
      <td>{{ i.price }}</td>
      <td>
        <a style="color: #3b3b3b" @click="reduce()" :disable="item.count === 1"><i
            class="fas fa-minus-circle"></i></a>
        {{ item.count }}
        <a style="color: #3b3b3b" @click="add()"><i class="fas fa-plus-circle"></i></a>
      </td>
      <td></td>
      <td>
        <a @click="del(index)"><i class="far fa-trash-alt"></i></a>
      </td>
      </tr>
    `,
    methods: {
        reduce: function () {
            if (this.item.count === 1) return;
            this.item.count--;
        },
        add: function () {
            this.item.count++;
        },
        del: function (index) {
            this.$root.cartList.splice(index, 1);
        }
    },
    computed: {
        i: function () {
            return this.$root.items.find(x => x.id == this.item.id)
        },
        f: function () {
            return this.$root.restaurants.find(x => x.id == this.i.restaurantId)
        }
    }
}

const Cart = {
    data() {
        return {
            productItem: []
        }
    },
    template: `
      <div class="modal fade" id="cart-modal" tabindex="-1" aria-hidden="true"
           style="    font-family: 'Montserrat', sans-serif;">
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content"
             style="margin: 0 auto !important;
         border-radius: 20px !important;
         border-color:rgba(117, 190, 218, 0.0);
         width: auto;">
          <div class="modal-body">
            <div class="table-responsive">
              <table class="table table-striped custom-table">
                <thead>
                <tr>
                  <th>#</th>
                  <th>Product Name</th>
                  <th>From</th>
                  <th>Price</th>
                  <th>Quantity</th>
                </tr>
                </thead>
                <tbody>
                <cart-tr
                    v-for="(item, index) in this.$root.cartList"
                    :item="item" :index="index">
                </cart-tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="d-flex flex-row-reverse">
            <button class="cart-btn" data-bs-target="#order-data-model" data-bs-toggle="modal"
                    data-bs-dismiss="modal">Order
            </button>
            <p style="margin-top: 15px !important;">Total <strong>{{ total }}</strong></p>
          </div>
        </div>
      </div>
      </div>
      <div class="modal fade" id="order-data-model" aria-hidden="true" tabindex="-1">
      <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content"
             style="border-radius:20px !important;
         border-color:rgba(117, 190, 218, 0.0);">
          <div class="modal-body">
            <div class="row g-3">

              <div class="col-sm-6">
                <input type="text" class="login-input" placeholder="Name" value="" required="">
              </div>

              <div class="col-sm-6">
                <input type="text" class="login-input" autocomplete='off' placeholder="Surname" value="" required="">
              </div>

              <div class="cart-col-12">
                <input type="text" id="phone" class="login-input" autocomplete='off' placeholder="Phone number"
                       required/>

                <script>jQuery(function ($) {
                  $("#phone").mask("+38(999) 999-99-99");
                });</script>

              </div>

              <div class="cart-col-12">
                <input type="text" class="login-input" autocomplete='off'
                       placeholder="Kharkiv, st. Academician Pavlova 154, apt. 12"
                       required="">
                <div class="invalid-feedback">
                  Please enter your address.
                </div>
              </div>

              <div class="col-md-7">
                <form>
                  <select class="login-input" autocomplete='off' required="">
                    <option value="" disabled selected hidden>Payment method</option>
                    <option>Credit Card</option>
                    <option>Cash</option>
                  </select>
                </form>
              </div>
            </div>
          </div>
          <div class="btn-group">
            <button class="cart-btn" style="width:50%"
                    data-bs-target="#cart-modal"
                    data-bs-toggle="modal"
                    data-bs-dismiss="modal">Back
            </button>
            <button class="cart-btn" style="width:50%"
                    data-bs-target="#cart-modal"
                    data-bs-toggle="modal"
                    data-bs-dismiss="modal">Confirm
            </button>
          </div>
        </div>
      </div>
      </div>
    `,
    computed: {
        total: function () {
            let total = 0;
            this.$root.cartList.forEach(cartItem => total += this.$root.items.find(x => x.id == cartItem.id).price * cartItem.count)
            return total.toString().replace(/\B(?=(\d{3})+$)/g, ',');
        },
    }
}

const FiltersPanel = {
    template: `
        <div class="card">
            <div class="accordion accordion-flush" id="accordionFlushExample">
                <div class="accordion-item">
                    <h2 class="accordion-header" id="panelsStayOpen-headingOne">
                        <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#panelsStayOpen-collapseOne" aria-expanded="true" aria-controls="panelsStayOpen-collapseOne">
                          Accordion Item #1
                        </button>
                    </h2>
                    <div id="panelsStayOpen-collapseOne" class="accordion-collapse collapse show" 
                        aria-labelledby="panelsStayOpen-headingOne">
                        <div class="accordion-body">
                            <div class="d-flex flex-row form-row">
                              <div class="form-group m-1">
                                <label>Min</label>
                                <input class="form-control" placeholder="$0" type="number">
                              </div>
                              <div class="form-group text-right m-1">
                                <label>Max</label>
                                <input class="form-control" placeholder="$1,0000" type="number">
                              </div>
                            </div>
                        </div>
                    </div>
                </div>
            <div class="accordion-item">
                <h2 class="accordion-header" id="panelsStayOpen-headingTwo">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#panelsStayOpen-collapseTwo" aria-expanded="true" aria-controls="panelsStayOpen-collapseTwo">
                      Accordion Item #2
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseTwo" class="accordion-collapse collapse show" 
                    aria-labelledby="panelsStayOpen-headingTwo">
                    <div class="accordion-body">
                        <div class="filter-content">
                          <div class="card-body">
                                <check name="More cheese"></check>
                                <check name="More vegetables"></check>
                                <check name="More sauce"></check>
                          </div> 
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div> `,

}

const Check = {
    props: ['name'],
    template: `
      <label class="form-check">
      <input class="form-check-input" type="checkbox" value="">
      <span class="form-check-label">{{ name }}</span>
      </label> `
}

const Supplier = {
    props: ['id', 'title', 'logo'],
    template: `
      <div class="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
           style="border-radius: 10px; max-height: 700px; max-width: 760px">
      <div class="icon-box">
        <router-link :to="{ path : String(id) }"><img :src="logo" class="img-fluid" alt="Pizza picture"></router-link>
        <h4><a href="#">{{ title }}</a></h4>
      </div>
      </div> `
}

const Product = {
    data() {
        return {
            showItem: false,
        }
    },
    props: ['id', 'title', 'logo'],
    template: `
      <div class="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
           style="border-radius: 10px; max-height: 700px; max-width: 760px">
      <div class="icon-box">
        <img :src="logo" class="img-fluid" data-bs-toggle="modal" data-bs-target="#item-modal"
             @click="this.showItem = true" alt="menu-item"/>
        <h4>{{ title }}</h4>
        <Teleport to="#pop-portal">
          <product-item :id="id" v-if="this.showItem" @close="this.showItem = false"></product-item>
        </Teleport>
      </div>
      </div> `
}

const ProductPopUp = {
    data() {
        return {
            quantity: 1,
        };
    },
    props: ["id"],
    template: `
      <div class="modal fade" id="item-modal" tabindex="-1" aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered modal-lg">
        <div class="modal-content"
             style="
         background-color: rgba(117, 190, 218, 0.0);
         border-color:rgba(117, 190, 218, 0.0);">
          <div class="modal-body">
            <div class="d-flex flex-column item-container product-from">
              <div class="img-part">
                <img :src="item.logo" class="img-f" alt="item">
              </div>
              <div class="product-data">
                <div class="data">
                  <h1>{{ item.title }}</h1>
                  <b>Ingredients: </b>
                  <text v-for="ingredient in item.ingredients">
                    <text>{{ ingredient }}{{ ", " }}</text>
                  </text>
                </div>
              </div>
              <div class="d-flex bd-highlight item-overlay-panel">
                <div class="me-auto p-2 bd-highlight">
                  <div style="padding: 0 10px ">
                    <div class="number" data-step="1" data-min="1" data-max="100">
                      <input v-model="quantity" class="number-text" type="text" name="count" value="1" readonly>
                      <a @click="reduce()" class="number-minus">−</a>
                      <a @click="add()" class="number-plus">+</a>
                    </div>
                  </div>
                </div>
                <div class="p-2 bd-highlight">
                  <p style="margin-top: 15px !important;">Total: <strong>{{ total }}$</strong></p>
                </div>
                <div class="p-2 bd-highlight">
                  <div style="padding: 0 10px ">
                    <button @click="addToCart(id, this.quantity)" data-bs-dismiss="modal"
                            class="add-to-cart-btn">ADD TO CART
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      </div>
    `,
    computed: {
        item: function () {
            return this.$root.items.find(x => x.id == this.id)
        },
        total: function () {
            return this.item.price * parseInt(this.quantity, 10)
        }
    },
    methods: {
        addToCart: function (id, count) {
            let orderItem = {
                id: id,
                count: count
            }

            this.$root.cartList.push(orderItem)
        },
        reduce: function () {
            if (parseInt(this.quantity, 10) === 1) return;
            this.quantity--;
        },
        add: function () {
            this.quantity++;
        },
        init: function () {
            this.quantity = 1
        }
    },
    mounted() {
        this.init()
    }
}

const SuppliersList = {
    template: `
    <supplier
            v-for="restaurant in this.$root.restaurants"
            :key="restaurant.id"
            :id="restaurant.id"
            :title="restaurant.title"
            :logo="restaurant.logo"
    ></supplier> `
}

const ProductsList = {
    props: ["id"],
    template: `
      <product
          v-for="item in this.$root.items.filter(x => x.restaurantId == id)"
          :key="item.id"
          :id="item.id"
          :title="item.title"
          :logo="item.logo"
      ></product> `
}

const routes = [
    {
        path: '/',
        component: SuppliersList
    },
    {
        path: '/:id',
        component: ProductsList,
        props: true,
    },

]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHashHistory(),
    routes,
})

app.component("check", Check)
app.component("filter-panel", FiltersPanel)
app.component("product", Product)
app.component("supplier", Supplier)
app.component("login", Login)
app.component("cart", Cart)
app.component("cart-tr", CartTr)
app.component("product-item", ProductPopUp)
app.use(router)

app.mount('#main-content')
