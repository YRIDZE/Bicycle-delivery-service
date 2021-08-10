const App = {
    data() {
        return {
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
        }
    }
}

const app = Vue.createApp(App)


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
                            <div class="form-row">
                              <div class="form-group col-md-6">
                                <label>Min</label>
                                <input class="form-control" placeholder="$0" type="number">
                              </div>
                              <div class="form-group text-right col-md-6">
                                <label>Max</label>
                                <input class="form-control" placeholder="$1,0000" type="number">
                              </div>
                            </div>
                        </div>
                    </div>
                </div>
            <div class="accordion-item">
                <h2 class="accordion-header" id="panelsStayOpen-headingTwo">
                    <button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#panelsStayOpen-collapseTwo" aria-expanded="false" aria-controls="panelsStayOpen-collapseTwo">
                      Accordion Item #2
                    </button>
                </h2>
                <div id="panelsStayOpen-collapseTwo" class="accordion-collapse collapse" 
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
