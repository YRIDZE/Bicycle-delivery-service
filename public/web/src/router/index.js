import {createRouter, createWebHistory} from 'vue-router'
import SupplierList from "@/components/SupplierList";
import ProductsList from "@/components/ProductsList";

const routes = [
    {
        path: '/',
        component: SupplierList
    },

    {
        path: '/:id',
        component: ProductsList,
        props: true,
    },
]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
