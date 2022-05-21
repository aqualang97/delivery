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
import cartView from "./CartView";
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
      // quantity:Number(),
      // purchase_price:Number(),

    }
  },
  methods:{
    async getOrders(){
      let user_id_route = this.$route.params.user_id
      this.user_id = parseInt(user_id_route)
      let resp = await fetch("http://localhost:8080/old-orders/"+user_id_route.toString(),{
        method:"GET"
      })
      console.log(resp.status)
      this.orders = await resp.json()
      this.image = this.orders.image
      // for (let order in ordersData){
      //   console.log(order)
      // }
    }
  },
  mounted() {
    cartView.methods.isLogin()
    let usr = localStorage.getItem("user")
    this.user_id = JSON.parse(usr).user_id
    let id = this.$route.params.user_id
    if (parseInt(id) === this.user_id){
      this.getOrders()
    }else {
      alert("oops")
      this.$router.push("/all-productsv")

    }

    }
}
</script>

<style scoped>

</style>