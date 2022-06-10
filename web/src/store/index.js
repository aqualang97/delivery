import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import prodStore from "./productStore"
import cart from "./cart";
import auth from "./auth";

const modules = {
  productStore:prodStore,
  cart:cart,
  auth:auth,
}

export default new Vuex.Store({

  modules

});
