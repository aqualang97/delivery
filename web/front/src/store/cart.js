//import axios from "axios";

const state = () => ({
    productCart: [],
    url: "http://localhost:8080/products-in-cart",
    errors : [],
    loaded : false
})
const mutations = {

}
const actions = {}
const getters = {}
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
}