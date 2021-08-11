const App = {
    data() {
        return {
            showLogin: false,
            showCart: false,

            restaurants: [
                {
                    id: 1,
                    title: 'My journey with Vue',
                    logo: 'https://images.unsplash.com/photo-1490717064594-3bd2c4081693?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1350&q=80'
                },
                {
                    id: 2,
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
            restMenuItems: [
                {
                    id: 1,
                    title: 'My journey with Vue',
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

            levels: ["pizza-menu", "restaurant-menu"],
            currentLevel: "pizza-menu"
        }
    },
    computed: {
        level() {
            return this.currentLevel
        },
    }
}

const app = Vue.createApp(App)

app.component('login', {
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
                  <input class="login-input" type="text" placeholder="Username"/>
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
                    <button class="ghost bt" id="signIn"  @click="signUpMode = false">SIGN IN</button>
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
})

app.component('cart', {
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
                  <th>Total</th>
                  <th scope="col"></th>
                </tr>
                </thead>
                <tbody>
                <tr>
                  <td> 1</td>
                  <td>Niam-Niam Food</td>
                  <td> Grill Bar <small class="d-block">Far far away, grill bar grill bar grill bar</small></td>
                  <td>$200.00</td>
                  <td>2</td>
                  <td>$400.00</td>
                  <td><a href="#" class="details">Details</a></td>
                </tr>
                <tr>
                  <td> 1</td>
                  <td>Niam-Niam Food</td>
                  <td> Grill Bar <small class="d-block">Far far away, grill bar grill bar grill bar</small></td>
                  <td>$200.00</td>
                  <td>2</td>
                  <td>$400.00</td>
                  <td><a href="#" class="details">Details</a></td>
                </tr>
                <tr>
                  <td> 1</td>
                  <td>Niam-Niam Food</td>
                  <td> Grill Bar <small class="d-block">Far far away, grill bar grill bar grill bar</small></td>
                  <td>$200.00</td>
                  <td>2</td>
                  <td>$400.00</td>
                  <td><a href="#" class="details">Details</a></td>
                </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="d-flex flex-row-reverse">
            <button class="cart-btn" data-bs-target="#order-data-model" data-bs-toggle="modal"
                    data-bs-dismiss="modal">Order
            </button>
            <p style="margin-top: 15px !important;">Total <strong>$1200.00</strong></p>
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
    `
})

app.component('pizza-container', {
    props: ['id', 'title', 'logo'],
    template: `
      <div class="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
           style="border-radius: 10px; max-height: 700px; max-width: 760px">
          <div class="icon-box"
               style="border-style: solid; border-width: 1px; border-color: rgba(194, 184, 184, 0.26);">
            <a class="icon" type="button" @click="this.$root.currentLevel='restaurant-menu'"><img :src="logo"
                                                                                                  class="img-fluid"
                                                                                                  alt="Pizza picture"></a>
            <h4><a href="#">{{ title }}</a></h4>
            <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore</p>
          </div>
      </div> `
})

app.component("pizza-menu", {
    template: `
    <pizza-container
            v-for="restaurant in this.$root.restaurants"
            :key="restaurant.id"
            :title="restaurant.title"
            :logo="restaurant.logo"
    ></pizza-container> `
})

app.component("restaurant-menu", {
    template: `
    <rest-menu-item
            v-for="item in this.$root.restMenuItems"
            :key="item.id"
            :title="item.title"
            :logo="item.logo"
    ></rest-menu-item> `
})

app.component("rest-menu-item", {
    props: ['id', 'title', 'logo'],
    template: `
      <div className="col-lg-3 col-md-6 d-flex align-items-stretch mt-4"
           style="border-radius: 10px; max-height: 700px; max-width: 760px">
          <div className="icon-box"
               style="border-style: solid; border-width: 1px; border-color: rgba(194, 184, 184, 0.26);">
            <a className="icon" href="#" data-bs-toggle="modal"
               data-bs-target="#item-modal"><img :src="logo" className="img-fluid" alt="menu-item"></a>
            <h4><a href="#">{{ title }}</a></h4>
            <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore</p>
          </div>
      
      </div> `
})

app.component('pizza-filter', {
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
})

app.component('check', {
    props: ['name'],
    template: `
      <label class="form-check">
      <input class="form-check-input" type="checkbox" value="">
      <span class="form-check-label">{{ name }}</span>
      </label> `
})


app.mount('#main-content')
