<template>
  <div id="app">
    <div id="nav">

      <router-link to="/">Home</router-link> |
      <router-link to="/homehome">HomeHome</router-link> |
      <router-link to="/about">About</router-link> |
      <router-link to="/example">example</router-link> |
      <router-link to="/all-products">products</router-link> |
      <router-link to="/sign-in">Sign In</router-link> |
      <router-link to="/sign-up">Sign Up</router-link> |
      <router-link to="/cart">Cart</router-link> |
      <router-link to="/categories">Categories</router-link> |
      <router-link to="/suppliers">Suppliers</router-link> |
      <router-link to="/old-orders">Old Orders</router-link> |
      <router-link to="/logout">Logout</router-link>
    </div>
    <router-view />
  </div>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
