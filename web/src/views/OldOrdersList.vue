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
  data(){
    return{
      isLogin:false,
    }
  },
  methods:{

    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged
      if(this.isLogin){
        let usr = localStorage.getItem('user')
        const obj = JSON.parse(usr)
        await this.$router.push("/old-orders/"+(obj.user_id).toString())
      }

    },

  },
  mounted() {
    this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    if (this.isLogin===false){
      alert("You are not login, access denied")
      this.$router.push("/sign-in")

    }
  }
}
</script>

<style scoped>

</style>