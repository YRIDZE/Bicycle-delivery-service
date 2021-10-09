import Vue from "vue";
import VueRouter from "vue-router";
import SupplierList from "@/components/SupplierList";
import ProductsList from "@/components/ProductsList";
import SupplierFilter from "@/components/SupplierFilter";
import ProductFilter from "@/components/ProductFilter";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    components: {
      filter: SupplierFilter,
      content: SupplierList,
    },
  },
  {
    path: "/:id",
    components: {
      filter: ProductFilter,
      content: ProductsList,
    },
    props: {
      filter: false,
      content: true
    },
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
