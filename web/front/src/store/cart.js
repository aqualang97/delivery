//import axios from "axios";

const state = () => ({
    productCart: [],
    url: "http://localhost:8080/products-in-cart",
    errors : [],
    loaded : false
})
const mutations = {
    addProdToCart(state, product){
        state().productCart.push(product)
    },
    delProdFromCart(state, product){
        state().productCart = state().productCart.splice(state().productCart.indexOf(product), 1)
    },
    clearCart(state){
        state().productCart = []
    },
    updateNumProd(state, product, newQuantity){
        state().productCart[product].quantity = product[newQuantity]
    }

}
const actions = {}
const getters = {

}
export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
}