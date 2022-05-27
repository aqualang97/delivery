// import axios from "axios";

const state = {
    url : "http://localhost:8080",
    user : {
        user_id:Number,
        access:"",
        refresh:"",
    },
    logged:false,
    errors : [],
    loaded : false
}
//
// const loginModel = {
//     login: "",
//     password: ""
// }

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
    async isLogin(){
        let usr = localStorage.getItem('user')
        if(usr===null){
            state.logged = false
        }
        else{
            const obj = JSON.parse(usr)
            console.log(
                ` ${obj.access_token}`, "profile")
            let resp = await fetch("http://localhost:8080/profile",{
                method: "GET",
                headers:{
                    // Authorization:`Bearer ${JSON.stringify({access_token:obj.access_token})}`
                    Authorization:(`Bearer ${obj.access_token}`),

                }
            })
            if (resp.status !== 200){
                // alert(resp.statusText)
                state.logged = false
                // document.location.reload()
            }else{
                console.log("login pass")
                state.logged = true
                console.log(state.logged)
            }

        }
    },
    async refresh(){
        let usr = localStorage.getItem('user')

        console.log(usr)
        const obj = JSON.parse(usr)

        console.log(`Bearer`, obj.refresh_token)

        if(usr===null){
            // await this.$router.push("/sign-in")
        }
        else{
            const obj = JSON.parse(usr)
            let resp = await fetch("http://localhost:8080/refresh",{
                method: "POST",
                body: JSON.stringify({refresh_token:`Bearer ` +obj.refresh_token,
                    access_token:""})
            })

            if (resp.status !== 200){
                // await this.$router.push("/sign-in")
                alert(resp.statusText)

            }else {
                let data = await resp.json()
                console.log(data)
                state.user = {
                    user_id:obj.user_id,
                    access: obj.access_token,
                    refresh: obj.refresh_token
                }
                localStorage.setItem('user', JSON.stringify(data))
            }

        }
    },

    async logout(){

        let usr = localStorage.getItem('user')
        const obj = JSON.parse(usr)

        let resp = await fetch("http://localhost:8080/logout",{
            method: "POST",
            body: JSON.stringify({access_token:`Bearer ` +obj.access_token,
                refresh_token:""})
        })

        if (resp.status===200){
            localStorage.clear()
            state.logged = false
            alert("success")
            document.location.reload()

        }

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