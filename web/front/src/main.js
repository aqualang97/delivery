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
import Buy from "./components/Buy";
import Checkout from "./components/Checkout";
import BackShopping from "./components/BackShopping";
import Pay from "./components/Pay";
// import CashMethod from "./components/CashMethod";
import CardMethod from "./components/CardMethod";
import OldOrders from "./components/OldOrders";
import OldOrdersUser from "./components/OldOrdersUser";
import Logout from "./components/Logout";
// import Total from "./components/Total";
// import example from "./views/example";

Vue.config.productionTip = false;

Vue.component("card", Card)
Vue.component('product', Product)
Vue.component('productSpecSupp', ProductSpecificSupplier)
Vue.component('productSupplier', ProductSupplier)
// Vue.component('total', Total)
Vue.component('productCategory', ProductCategory)
Vue.component("login", Login)
Vue.component("registration", Registration)
Vue.component("cart", Cart)
Vue.component('categories', Categories)
Vue.component('suppliers', Suppliers)
Vue.component('homehome', HomePage)
Vue.component('buy', Buy)
Vue.component('checkout', Checkout)
Vue.component('back-to-shopping', BackShopping)
Vue.component('pay', Pay)
// Vue.component('cash-method', CashMethod)
Vue.component('card-method', CardMethod)
Vue.component('old-orders', OldOrders)
Vue.component('old-orders-user', OldOrdersUser)
Vue.component('logout', Logout)
new Vue({
  router,
  store,
  render: (h) => h(App),
  // data(){
  //   return{
  //     sort:null
  //   }
  // },


}).$mount("#app");

