import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Card from "./components/Card";
import Product from "./components/Product";
// import example from "./views/example";

Vue.config.productionTip = false;

Vue.component("card", Card)
Vue.component('product', Product)
new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount("#app");
