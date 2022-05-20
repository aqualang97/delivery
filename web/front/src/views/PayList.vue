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
    }
  },
  methods:{
    async successfulPaid(){
      await new Promise(r => setTimeout(r, 2000));

      let simulation = {
        status:"completed",
        last_order_id: JSON.parse(localStorage.getItem('user_last_order_id'))
      }
      let resp = await fetch("http://localhost:8080/card_pay", {
        method: "POST",
        body: JSON.stringify(simulation)
      })
      if (resp.status===200){
        this.msg = "Status: completed"
        document.location.reload()

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
    await this.checkMethod()


  }

}
</script>

<style scoped>

</style>