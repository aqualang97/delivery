<template>
  <div>
    <logout></logout>
  </div>
</template>

<script>
// import cartView from "./CartView";
export default {
  name: "LogoutVue",
  methods:{
    logout(){
      this.$store.dispatch('auth/logout')
    }
  },
  mounted() {
    this.logout()
    let logout = this.$store.state.auth.logged
    if (!logout){
      alert("You are not login, access denied")
      this.$router.push("/sign-in")


    }
  }
}
</script>

<style scoped>

</style>