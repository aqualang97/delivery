import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import example from "../views/example";
import products from "../views/ProductList"
import axios from "axios"

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
  {
    path: "/example",
    name: "example",
    component: example,
  },
  {
    path: "/example/:id",
    name: "example",
    component: example,
  },
  {
    path: "/all-products",
    name: "products",
    component: products,
  },

];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
