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

      let resp = await fetch("http://localhost:8080/logout",{
        method: "GET",
        mode:"cors",
        headers: {"Authorization": `Bearer ${access}`},
      })
      console.log(access)
      console.log(access)
      console.log(access)
      console.log(access)
      console.log(access)
      console.log(resp.status)
      console.log(resp.status)
      console.log(resp.status)
      console.log(resp.status)
      console.log(resp.status)
      if (resp.status===200){
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