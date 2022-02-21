import axios from "axios";

const state = {
    url : "http://localhost:8080/login",
    user : null,
    errors : [],
    loaded : false
}

const loginModel = {
    login: "",
    password: ""
}

const mutations = {
    setUserLogin(state, user){
        state.user = user
    },
    setErrorsLogin(state, errors){
        state.errors = errors
    },
    setLoadedLogin(state, loaded){
        state.loaded = loaded
    },
}

const actions = {
    fetch(loginModel){
        JSON.stringify(loginModel)
        // user.commit(
        //     'setLoadedLogin',
        //     false
        // )
        // axios.post(user.getters.getPostsURL).then((response)=>{
        //     user.commit('setUserLogin', state.data.push(user))
        // }).catch((err)=>{user.commit('setErrorLogin', [err])
        // }).finally(()=>{user.commit('setLoaded', true)})


        // const response = fetch(getPostURL, {
        //     method: 'GET',}).then(
        //     products.commit('setPosts', response.data)
        //
        // )
    }
}

const getters = {
    getPostsURL: (state)=>{
        return state.url
    },
    getUser: (state) =>{
        return state.user

    },
    getErrors: (state) =>{
        return state.errors
    },
    getLoaded: (state) =>{
        return state.loaded
    }
}

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
}