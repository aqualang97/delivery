import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import example from "../views/example";
import products from "../views/ProductList"
import login from "../views/LoginView"
import registration from "../views/RegistrationView"
import cart from "../views/CartView";
import productSpec from "../views/ProductSpecificSupplierView";
import categoriesList from "../views/CategoriesList";
import suppliersList from "../views/SuppliersList";
import HomePageView from "../views/HomePageView";
// import axios from "axios"

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/homehome",
    name: "HomeHome",
    component: HomePageView,
  },
  {
    path: "/about",
    name: "About",
    // router level code-splitting
    // this generates a separate chunk (about.[hash].js) for this router
    // which is lazy-loaded when the router is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/example",
    name: "example",
    component: example,
  },
  {
    path: "/example/:id/prod/:idd",
    name: "example",
    component: example,
  },
  {
    path: "/all-products",
    name: "products",
    component: products,
  },
  {
    path: "/categories",
    name: "categories",
    component: categoriesList,
  },
  {
    path: "/suppliers",
    name: "suppliers",
    component: suppliersList,
  },

  {
    path: "/suppliers/:supp_id/products/:prod_id",
    name: "productsSpecificSupp",
    component: productSpec,
  },
  {
    path: "/sign-in",
    name: "login",
    component: login,
  },
  {
    path: "/sign-up",
    name: "registration",
    component: registration,
  },
  {
    path: "/cart",
    name: "cart",
    component: cart,
  },

];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
