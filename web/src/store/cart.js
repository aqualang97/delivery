//import axios from "axios";

const state = {
    productCart: [],
    url: `${window.apiPrefix}/products-in-cart`,
    errors : [],
    loaded : false
}
const mutations = {
    // addProdFromLocal(state, product){
    //     state.productCart.push(product)
    // },
    addProdToCart(state, product){
        product.quantity += 1
        state.productCart.push(product);
    },
    delProdFromCart(state, product){
        state.productCart = state.productCart.splice(state.productCart.indexOf(product), 1)
    },
    clearCart(state){
        state.productCart = []
    },
    inputLocalToState(state, local){
        let newState = []
        for (let product in local){
            newState.push(local[product])
        }
        state.productCart = newState

    },
    plusNumProd(state, product){
        //state().productCart[product].quantity = product[newQuantity]
        console.log(state.productCart[product].quantity)
        state.productCart[product].quantity +=1
    },
    minusNumProd(state, product){
        //state().productCart[product].quantity = product[newQuantity]
        console.log(state.productCart[product].quantity)
        if (state.productCart[product].quantity === 0){
            state.productCart[product].quantity = 0
        }
        else {
            state.productCart[product].quantity -=1
        }
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