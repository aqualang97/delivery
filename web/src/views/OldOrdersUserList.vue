<template>
  <div>


    <old-orders-user v-for="(order) in orders"
         :key="order.ID"
         :order_id="order.order_id"
         :user_id="order.user_id"
         :status="order.status"
         :full_price="order.full_price"
                     :user_order="order.user_order"
    >
    </old-orders-user>
  </div>

</template>

<script>
export default {
  name: "OldOrdersUserList",
  components: {},
  props:{
  },
  data(){
    return{
      user_id:Number(),
      orders:null,
      order_id:Number(),
      status:String(),
      full_price:null,
      user_order:null,
      // name:String(),
      // category:String(),
      image:String(),
      isLogin:false,
      // quantity:Number(),
      // purchase_price:Number(),

    }
  },
  methods:{
    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged
    },

    async getOrders(){
      let user_id_route = this.$route.params.user_id
      this.user_id = parseInt(user_id_route)
      let resp = await fetch(`${this.$apiPrefix}/old-orders/`+user_id_route.toString(),{
        method:"GET"
      })
      this.orders = await resp.json()
      this.image = this.orders.image

    }
  },

  mounted() {
    this.checkLogin()
    this.isLogin = this.$store.state.auth.logged
    if(!this.isLogin){
      alert("oops")
      this.$router.push("/all-products")
    }else{
      let usr = localStorage.getItem("user")
      this.user_id = JSON.parse(usr).user_id
      let id = this.$route.params.user_id
      if (parseInt(id) === this.user_id){
        this.getOrders()
      }else {
        alert("oops")
        this.$router.push("/all-products")

      }
    }


  }
}
</script>

<style scoped>

</style>