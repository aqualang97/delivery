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
    async logout(){

      let usr = localStorage.getItem('user')
      let access = JSON.parse(usr).access_token
      console.log(access)

      let resp = await fetch("http://localhost:8080/logout",{
        method: "POST",
        body: JSON.stringify({access_token:access}),
      })

      if (resp.status===200){
        localStorage.clear()
        alert("success")
        await this.$router.push("/sign-in")
      }

    }
  },
  mounted() {
    this.logout()
  }
}
</script>

<style scoped>

</style>