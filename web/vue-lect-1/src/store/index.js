import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import prodStore from "./productStore"
import cart from "./cart";

const modules = {
  productStore :prodStore,
  cart:cart
}

export default new Vuex.Store({
  // state: {},
  // mutations: {},
  // actions: {},
  // modules: {},
  modules

});
