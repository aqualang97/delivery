<template>
  <div>
    <old-orders>

    </old-orders>
  </div>
</template>

<script>
export default {
  name: "OldOrdersList",
  components: {},
  methods:{
    async isLogin(){
      let usr = localStorage.getItem('user')

      if(usr===null){
        alert("You are not login, access denied")
        await this.$router.push("/sign-in")
      }
      else {

        const obj = JSON.parse(usr)
        console.log("OLD LRDER", obj)
        console.log(obj.user_id)
        console.log(obj.refresh_token)
        let resp = await fetch("http://localhost:8080/refresh", {
          method: "POST",
          body: JSON.stringify({refresh_token: obj.refresh_token})
        })
        console.log("LRDER",obj.refresh_token)
        if (resp.status !== 200) {
          await this.$router.push("/sign-in")

          alert(resp.statusText)
          // await new Promise(r => setTimeout(r, 500));
          console.log(resp.status)
          return
        }
        let data = await resp.json()
        console.log("NEW", data)
        // alert("access is allowed)")

        localStorage.setItem('user', JSON.stringify(data))
        await this.$router.push("/old-orders/"+(obj.user_id).toString())

      }
    },
  },
  mounted() {
    this.isLogin()
  }
}
</script>

<style scoped>

</style>