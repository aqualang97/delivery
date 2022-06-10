<template>
  <div>
    <pay :msg="msg"
    >

    </pay>
  </div>
</template>

<script>
export default {
  name: "PayList",
  components: {},
  data(){
    return{
      msg:String(),
      isLogin:false,
    }
  },
  methods:{
    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged
    },

    async successfulPaid(){
      await new Promise(r => setTimeout(r, 2000));

      let simulation = {
        status:"completed",
        last_order_id: JSON.parse(localStorage.getItem('user_last_order_id'))
      }
      let resp = await fetch(`${this.$apiPrefix}/card_pay`, {
        method: "POST",
        body: JSON.stringify(simulation)
      })
      if (resp.status===200){
        this.msg = "Status: completed"
        this.$store.commit('cart/clearCart')
        localStorage.removeItem("user_order")
        await new Promise(r => setTimeout(r, 2000));
        await this.$router.push("/all-products")

        // document.location.reload()

      }
      localStorage.setItem("user_pay_method", "")
    },
    async checkMethod(){
      let method = localStorage.getItem("user_pay_method")
      if (method==="Cash"){
        this.msg = "Cash"
        await this.successfulPaid()
      }
      else if (method==="Card"){
        this.msg = "Card"

        await this.successfulPaid()
      }
      else {
        this.msg = "Oops"
      }
    }
  },
  async mounted() {
    await this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    if(this.isLogin===false){
      alert("You are not login, access denied")
      await this.$router.push("/sign-in")
    }
    await this.checkMethod()

  }

}
</script>

<style scoped>

</style>