//import axios from "axios";

const state = () => ({
    productCart: [],
    url: "http://localhost:8080/",
    errors : [],
    loaded : false
})
const mutations = {
    set
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