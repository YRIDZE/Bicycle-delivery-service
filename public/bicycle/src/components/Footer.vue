<template>
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
</template>

<script>
export default {
  name: "Footer"
}
</script>

<style scoped>

</style>