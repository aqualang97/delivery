<template>
  <div id="app">
    <div id="nav" class="comfortaa-22">

<!--      <router-link to="/">Home</router-link> |-->
      <router-link  to="/">Home
      </router-link>
      <router-link to="/about">About</router-link>
<!--      <router-link to="/example">example</router-link> |-->
      <router-link to="/all-products">Products</router-link>

      <router-link to="/cart">Cart</router-link>
      <router-link to="/categories">Categories</router-link>
      <router-link to="/suppliers">Suppliers</router-link>
      <router-link v-if="this.isLogin"  to="/old-orders">Old Orders</router-link>
      <router-link v-if="this.isLogin"  to="/logout">Logout</router-link>
      <router-link v-if="!this.isLogin" to="/sign-in">Sign In</router-link>
      <router-link v-if="!this.isLogin" to="/sign-up">Sign Up</router-link>
    </div>
    <router-view />
  </div>
</template>
<script>
// import CartView from "./views/CartView";
export default {
  data(){
    return{
      isLogin:false,
      isReload:false,
    }
  },
  methods:{
    refreshToken(){
      this.isLogin = this.$store.dispatch('auth/isLogin')
      if (this.isLogin){
        this.$store.dispatch('auth/refresh')
      }
    },
    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      console.log(this.$store.state.auth.logged)
      this.isLogin = this.$store.state.auth.logged
      console.log("app", this.isLogin)
    },

  },
  mounted() {
    // this.logged = CartView.methods.isLogin()
    this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    let isr = localStorage.getItem('isReload')
    console.log('reload',isr===null)
    if(isr===null){
      localStorage.setItem('isReload', 'true')
      document.location.reload()
    }else if(Boolean(isr)===false){
      localStorage.setItem('isReload', 'true')
      document.location.reload()

    }

    console.log("APP", this.isLogin)
    console.log("APP", this.isLogin)
    console.log("APP", this.isLogin)

  }
}

</script>
<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
@font-face {
  font-family: "Comfortaa";
  src :url("../../fonts/comfortaa/Comfortaa-VariableFont_wght.ttf");
  font-weight: 400;
}
@font-face {
  font-family: "Fredoka";
  src :url("../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.comfortaa-22{
  font-family: "Comfortaa",serif;
  font-size: 22px;
  text-decoration: unset;
}

#nav {
  padding: 30px;
  display:flex;
  justify-content: space-evenly;
  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {

       margin-top: 3px;
       /*text-decoration: underline dotted;*/
       border-bottom: 3px #FF865E dotted;

       color: #FF865E;
     }
  }
}
#nav a{

}
a:-webkit-any-link{
  text-decoration: none;
}
</style>
