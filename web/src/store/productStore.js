import axios from "axios";

const state = {
    url : `${window.apiPrefix}/all-products`,
    posts : [],
    errors : [],
    loaded : false
}

const mutations = {
    setPosts(state, posts){
        state.posts = posts
    },
    setErrors(state, errors){
        state.errors = errors
    },
    setLoaded(state, loaded){
        state.loaded = loaded
    },
}

const actions = {
    fetch(products){
        products.commit(
            'setLoaded',
            false
        )
        axios.get(products.getters.getPostsURL).then((response) =>{
            products.commit('setPosts', response.data)
        }).catch((err)=>{
            products.commit('setErrors', [err])
        }).finally(()=>{
            products.commit('setLoaded', true)
        })
        // const response = fetch(getPostURL, {
        //     method: 'GET',}).then(
        //     products.commit('setPosts', response.data)
        //
        // )
    },
    // abc(){
    //     console.log("abc")
    // }
}

const getters = {
    getPostsURL: (state)=>{
        return state.url
    },
    getPosts: (state) =>{
        return state.posts
    },
    getPostsByID:(state)=>(id)=>{
        return state.posts.find(p =>{
            return p.id = id
        })
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