<template>
  <div id="main-content">
    <header id="header">
      <div class="header-bar fixed-top">
        <nav class="navbar navbar-expand-sm  navbar-dark bg-dark"
             style="background-color:#545454FF !important; padding: 1em !important;">

          <a class="navbar-brand" href="/" style="padding:0 30px">
            <font-awesome-icon :icon="['fas', 'bicycle']"/>
            <b> Bicycle</b></a>
          <button
              class="navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#toggleMenu"
              aria-controls="toggleMenu"
              aria-expanded="false"
              aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="toggleMenu">
            <ul class="navbar-nav ms-auto text-center" style="padding:0 30px">
              <li class="nav-item"><a class="nav-link" href="#">About</a></li>
              <li class="nav-item"><a class="nav-link" id="login" href="#" @click="showLogin = true" type="button"
                                      data-bs-toggle="modal"
                                      data-bs-target="#login-modal">
                <font-awesome-icon :icon="['fas', 'user']"/>
                Login</a></li>
              <li class="nav-item"><a class="nav-link" id="cart" href="#" @click="showCart = true" type="button"
                                      data-bs-toggle="modal"
                                      data-bs-target="#cart-modal">
                <font-awesome-icon :icon="['fas', 'shopping-basket']"/>
                Cart</a></li>
            </ul>
          </div>
        </nav>
      </div>

      <Login v-if="showLogin" @close="showLogin = false"></Login>
      <Cart v-if="showCart" @close="showCart = false"></Cart>

      <div class="cont-custom text-center bg-light"
           :style="{ backgroundImage: 'url(' + require(`@/assets/img/header-pizza.jpg`) + ')',
           marginBottom : '1rem',
           backgroundSize: 'cover',
           padding: '6rem 3rem 0 3rem !important',
           height: '26rem'}">

        <h1 class="mb-3"><b>Bicycle</b></h1>
        <h4 class="mb-3">Subheading</h4>

        <div class="search-container">
          <input type="text" class="input" placeholder="what are you looking for...">
          <input type="button" class="close-btn" value="Search">
        </div>
      </div>
    </header>

    <div class="container-fluid" id="content"
         style="padding-left: 70px !important; padding-right: 70px !important; font-family: 'Montserrat', sans-serif;">
      <div class="col-md-12"><h2>Food<GoBack></GoBack></h2>
        <p>...is any substance consumed to provide nutritional support for an organism. Food is usually of plant, animal
          or fungal origin, and contains essential nutrients, such as carbohydrates, fats, proteins, vitamins, or
          minerals. The substance is ingested by an organism and assimilated by the organism's cells to provide energy,
          maintain life, or stimulate growth. </p>
      </div>
      <div class="d-flex flex-row mb-5 mt-1">
        <Filter></Filter>
        <section id="services" class="services flex-shrink-1" style="padding-bottom: 10px">
          <div class="row">
            <router-view></router-view>
          </div>
        </section>
      </div>
    </div>

    <GoTop text="Back to top"></GoTop>

    <div id="pop-portal"></div>
    <footer id="footer">
      <div class="footer-top">
        <div class="container">
          <div class="row">
            <div class="col-lg-3 col-md-6 footer-links">
              <h4>Useful Links</h4>
              <ul>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Home</a></li>
                <li><i class="bx bx-chevron-right"></i> <a href="#">About us</a></li>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Services</a></li>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Terms of service</a></li>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Privacy policy</a></li>
              </ul>
            </div>
            <div class="col-lg-3 col-md-6 footer-links">
              <h4>Our Services</h4>
              <ul>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Delivery</a></li>
                <li><i class="bx bx-chevron-right"></i> <a href="#">Tasting</a></li>
              </ul>
            </div>
            <div class="col-lg-3 col-md-6 footer-contact">
              <h4>Contact Us</h4>
              <p>
                132, Akademika Pavlova Street <br>
                Kharkiv, Ukraine <br><br>
                <strong>Phone:</strong> +380988888888<br>
                <strong>Email:</strong> localyridze@gmail.com<br>
              </p>
            </div>
            <div class="col-lg-3 col-md-6 footer-info">
              <h3>About Bicycle</h3>
              <p>Cras fermentum odio eu feugiat lide par naso tierra. Justo eget nada terra videa magna derita
                valies darta donna mare fermentum iaculis eu non diam phasellus.</p>
              <div class="social-links mt-3">
                <a href="#">
                  <font-awesome-icon :icon="['fab', 'instagram']"/>
                </a>
                <a href="#">
                  <font-awesome-icon :icon="['fab', 'twitter']"/>
                </a>
                <a href="#">
                  <font-awesome-icon :icon="['fab', 'facebook-f']"/>
                </a>
                <a href="#">
                  <font-awesome-icon :icon="['fab', 'linkedin-in']"/>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="container">
        <div class="copyright">
          &copy; Copyright <strong><span>Bicycle Corporation</span></strong>
        </div>
      </div>
    </footer>
  </div>
</template>

<script>
import Login from '@/components/Login'
import Cart from '@/components/Cart'
import Filter from '@/components/FilterPanel'
import GoBack from '@/components/GoBack'
import GoTop from '@/components/GoTop'

import '../../web/public/css/main-page.css'
import '../../web/public/css/cart.css'
import '../../web/public/css/login-registration.css'
import '../../web/public/css/restaurant-menu-page.css'
import '../../web/public/css/menu-item-page.css'

export default {
  name: 'App',
  components: {
    Login,
    Cart,
    Filter,
    GoBack,
    GoTop
  },
  data() {
    return {
      showLogin: false,
      showCart: false,
    }
  },
}
</script>