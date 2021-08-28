const App = {
    data() {
        return {
            showLogin: false,
            showCart: false,

            restaurants: [
                {
                    id: 1,
                    items: [
                        {
                            id: 1,
                            title: 'My journey with Vue',
                            description: "auf",
                            price: 122,
                            logo: '../img/pizza-menu-item.jpg'
                        },
                        {
                            id: 2,
                            title: 'My journey with Vue',
                            logo: '../img/pizza-menu-item.jpg'
                        },
                        {
                            id: 3,
                            title: 'My journey with Vue',
                            logo: '../img/pizza-menu-item.jpg'
                        }
                    ],
                    title: 'My journey with Vue',
                    logo: 'https://images.unsplash.com/photo-1490717064594-3bd2c4081693?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80'
                },
                {
                    id: 2,
                    items: [
                        {
                            id: 1,
                            title: '2',
                            logo: '../img/pizza-menu-item.jpg'
                        },
                        {
                            id: 2,
                            title: '2',
                            logo: '../img/pizza-menu-item.jpg'
                        },
                        {
                            id: 3,
                            title: '2',
                            logo: '../img/pizza-menu-item.jpg'
                        }
                    ],
                    title: 'Blogging with Vue',
                    logo: 'https://images.unsplash.com/photo-1506354666786-959d6d497f1a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80'
                },
                {
                    id: 3,
                    title: 'Why Vue is so fun',
                    logo: 'https://images.unsplash.com/photo-1590947132387-155cc02f3212?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8MTN8fHxlbnwwfHx8fA%3D%3D&w=1000&q=80'
                },
                {
                    id: 4,
                    title: 'Blogging with Vue',
                    logo: 'https://images.unsplash.com/photo-1617219474432-2e0b3ba569b8?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1050&q=80'
                },
                {
                    id: 5,
                    title: 'Blogging with Vue',
                    logo: 'https://images.unsplash.com/photo-1620374645466-dc3ff1558148?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80'
                },
                {
                    id: 6,
                    title: 'Blogging with Vue',
                    logo: 'https://images.unsplash.com/photo-1513104890138-7c749659a591?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxleHBsb3JlLWZlZWR8MXx8fGVufDB8fHx8&w=1000&q=80'
                }
            ],
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

const Cart = {
    data() {
        return {
            cartList: [
                {
                    id: 1,
                    name: 'Iphone xs',
                    from: "IRA",
                    price: 10000,
                    count: 1
                },
                {
                    id: 2,
                    name: 'Ipad Pro',
                    from: "IRA",
                    price: 6666,
                    count: 1
                },
                {
                    id: 3,
                    name: 'MacBook Pro',
                    from: "IRA",
                    price: 25000,
                    count: 1
                },
            ]

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
                <tr v-for="(item, index) in cartList">
                  <td>{{ index + 1 }}</td>
                  <td>{{ item.name }}</td>
                  <td>{{ item.from }}</td>
                  <td>{{ item.price }}</td>
                  <td>
                    <a style="color: #3b3b3b" @click="reduce(index)" :disable="item.count === 1"><i
                        class="fas fa-minus-circle"></i></a>
                    {{ item.count }}
                    <a style="color: #3b3b3b" @click="add(index)"><i class="fas fa-plus-circle"></i></a>
                  </td>
                  <td></td>
                  <td>
                    <a @click="del(index)"><i class="far fa-trash-alt"></i></a>
                  </td>
                </tr>
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
    methods: {
        reduce: function (index) {
            if (this.cartList[index].count === 1) return;
            this.cartList[index].count--;
        },
        add: function (index) {
            this.cartList[index].count++;
        },
        del: function (index) {
            this.cartList.splice(index, 1);
        }
    },
    computed: {
        total: function () {
            let total = 0;
            for (let i = 0; i < this.cartList.length; i++) {
                total += this.cartList[i].price * this.cartList[i].count;
            }
            return total.toString().replace(/\B(?=(\d{3})+$)/g, ',');
        }
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
                                <check name="More meat"></check>
                                <check name="More cheese"></check>
                                <check name="More vegetables"></check>
                                <check name="More sauce"></check>
                          </div> 
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div> `
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
      <div class="icon-box"
           style="border-style: solid; border-width: 1px; border-color: rgba(194, 184, 184, 0.26);">
        <router-link :to="{ path : title }"><img :src="logo" class="img-fluid" alt="Pizza picture"></router-link>
        <h4><a href="#">{{ title }}</a></h4>
        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore</p>
      </div>
      </div> `
}

const Product = {
    props: ['id', 'title', 'logo'],
    template: `
      <div class="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
           style="border-radius: 10px; max-height: 700px; max-width: 760px">
      <div class="icon-box"
           style="border-style: solid; border-width: 1px; border-color: rgba(194, 184, 184, 0.26);">
        <a class="icon" href="#" data-bs-toggle="modal" data-bs-target="#item-modal"><img :src="logo" class="img-fluid"
                                                                                          alt="menu-item"></a>
        <h4><a href="#">{{ title }}</a></h4>
        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore</p>

        <router-view></router-view>

      </div>
      </div> `
}

const ProductPopUp = {
    props: ["title", "itemName"],
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
                  <p>{{ item.description }}</p>
                </div>
              </div>
              <div class="d-flex bd-highlight item-overlay-panel">
                <div class="me-auto p-2 bd-highlight">
                  <div style="padding: 0 10px ">
                    <div class="number" data-step="1" data-min="1" data-max="100">
                      <input class="number-text" type="text" name="count" value="1" readonly>
                      <a href="#" class="number-minus">−</a>
                      <a href="#" class="number-plus">+</a>
                    </div>
                  </div>
                </div>
                <div class="p-2 bd-highlight">
                  <input class="price" type="text" name="price" v-bind:value=item.price readonly>
                </div>
                <div class="p-2 bd-highlight">
                  <div style="padding: 0 10px ">
                    <button class="add-to-cart-btn">ADD TO CART</button>
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
        item() {
            return this.$root.restaurants.find(x => x.title === title).items.find(y => y.title === itemName)
        }

    }
}

const SuppliersList = {
    template: `
    <supplier
            v-for="restaurant in this.$root.restaurants"
            :key="restaurant.id"
            :title="restaurant.title"
            :logo="restaurant.logo"
    ></supplier> `
}

const ProductsList = {
    props: ["title"],
    template: `
      <product
          v-for="item in this.$root.restaurants.find(x => x.title === title).items"
          :key="item.id"
          :title="item.title"
          :logo="item.logo"
      ></product> `
}

const routes = [
    {path: '/', component: SuppliersList},
    {
        path: '/:title',
        component: ProductsList,
        props: true,
        children: [
            {
                path: ':itemName',
                components: ProductPopUp,
                props: true
            },
        ]
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
app.use(router)

app.mount('#main-content')
