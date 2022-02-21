import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Card from "./components/Card";
import Product from "./components/Product";
import Login from "./components/Login";
import Registration from "./components/Registration";
import Cart from "./components/Cart";
import ProductSpecificSupplier from "./components/ProductSpecificSupplier";
import Categories from "./components/Categories";
import Suppliers from "./components/Suppliers";
import HomePage from "./components/HomePage";
import ProductCategory from "./components/ProductCategory";
import ProductSupplier from "./components/ProductSupplier";

// import example from "./views/example";

Vue.config.productionTip = false;

Vue.component("card", Card)
Vue.component('product', Product)
Vue.component('productSpecSupp', ProductSpecificSupplier)
Vue.component('productSupplier', ProductSupplier)

Vue.component('productCategory', ProductCategory)
Vue.component("login", Login)
Vue.component("registration", Registration)
Vue.component("cart", Cart)
Vue.component('categories', Categories)
Vue.component('suppliers', Suppliers)
Vue.component('homehome', HomePage)

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
