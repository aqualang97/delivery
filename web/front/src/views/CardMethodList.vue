<template>
  <div>
    <card-method>

    </card-method>
  </div>
</template>

<script>
export default {
  name: "CardMethodList",
  components:{},
  data(){
    return{
      isLogin:false,
    }
  },

  methods:{
    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      console.log(this.$store.state.auth.logged)
      this.isLogin = this.$store.state.auth.logged
      console.log("BYY", this.isLogin)
    },
  },
  async mounted() {
    await this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    if(this.isLogin===false) {
      alert("You are not login, access denied")
      await this.$router.push("/sign-in")
    }
  }
}
</script>

<style scoped>

</style>