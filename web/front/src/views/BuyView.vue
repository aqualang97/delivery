<template>
  <div>
    <buy>

    </buy>
  </div>
</template>

<script>
export default {
  name: "BuyView",
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
    },
  },
  mounted() {
    this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    if(!this.isLogin){
      alert("You are not login, access denied")
      this.$router.push("/sign-in")
    }
        }
}
</script>

<style scoped>

</style>