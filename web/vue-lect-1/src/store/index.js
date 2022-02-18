import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);
import prodStore from "./productStore"

const modules = {
  productStore :prodStore
}

export default new Vuex.Store({
  // state: {},
  // mutations: {},
  // actions: {},
  // modules: {},
  modules

});
