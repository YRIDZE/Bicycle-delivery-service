import Vue from "vue";
import VueRouter from "vue-router";
import SupplierList from "@/components/SupplierList";
import ProductsList from "@/components/ProductsList";
import AllProducts from "@/components/AllProducts";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: SupplierList,
  },
  {
    path: "/:id",
    component: ProductsList,
    props: true,
  },
  {
    path: "/all",
    component: AllProducts,
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
