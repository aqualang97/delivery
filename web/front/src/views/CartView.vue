<template>
<div>
  <cart
      v-for="i in cartList"
      :key="i.idProd"
      :id-prod="i.idProd"
      :prod-name="i.prodName"
      :price="i.price"
      :ingredients="i.ingredients"
      :img-link="i.imgLink"
      :type="i.type"
      :quantity="i.quantity"


  ></cart>
  <div class="total">
    <div class="total-txt corinthina-40">
      <h2>Total payment:</h2>
      <h2>{{ total.toFixed(2) }} $</h2>
    </div>
    <div class="buy">
      <button type="button" name="button"
              @click="$router.push(`/buy`)&&checkLogin()">
        Buy
      </button>
      <button type="button" name="button"
              @click="clear">
        Clear cart
      </button>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "CartView",
  components: {
  },
  methods:{
    checkLogin(){
      this.$store.dispatch('auth/isLogin')
      console.log(this.$store.state.auth.logged)
      this.isLogin = this.$store.state.auth.logged
      console.log("OOUL", this.isLogin)
    },
    async clear(){
      localStorage.removeItem("user_order")
      this.$store.commit('cart/clearCart');
      await document.location.reload()
    },

    // async isLogin(){
    //   let usr = localStorage.getItem('user')
    //   console.log(usr)
    //   if(usr===null){
    //     alert("You are not login, access denied")
    //     await this.$router.push("/sign-in")
    //
    //   }
    //   else{
    //     const obj = JSON.parse(usr)
    //     let resp = await fetch("http://localhost:8080/refresh",{
    //       method: "POST",
    //       body: JSON.stringify({refresh_token:obj.refresh_token})
    //     })
    //
    //     if (resp.status !== 200){
    //       await this.$router.push("/sign-in")
    //
    //       alert(resp.statusText)
    //       // await new Promise(r => setTimeout(r, 500));
    //       console.log(resp.status)
    //       return
    //     }
    //     let data = await resp.json()
    //     console.log(data)
    //     // alert("access is allowed)")
    //
    //     localStorage.setItem('user', JSON.stringify(data))
    //
    //
    //     console.log(obj.user_id)
    //     console.log(obj.access_token)
    //     console.log(obj.refresh_token)
    //     return true
    //   }
    // },

    ifUpdPage(){
      console.log("upd")
      let localCart = JSON.parse(localStorage.getItem("user_order"))
      console.log(localCart)
      console.log("if pd", localCart)
      if (localCart!==null){
        if(localCart.length!==0&&this.$store.state.cart.productCart.length===0){
          this.$store.commit('cart/inputLocalToState', localCart);
        }
      }
    },

    toLocal(){
      console.log("To localCart")

      console.log(this.$store.state.cart.productCart)

      localStorage.setItem('user_order', "")
      localStorage.setItem('user_order', JSON.stringify(this.$store.state.cart.productCart))
    },
    plus(numProdCart){
      return numProdCart+=1
    },
    minus(numProdCart){
      return numProdCart-=1
    },
    fromLocal(){
      let localCart = JSON.parse(localStorage.getItem("user_order"))
      console.log(localCart)
      console.log(localCart.length)
      this.$store.commit('cart/inputLocalToState', localCart);

      console.log("localCart")
      console.log("upd state", this.$store.state.cart.productCart)
    },

  },
  data() {
    return {
      cartList:null,
      total:0
    };
  },
  mounted() {
    this.ifUpdPage()
    this.toLocal()
    this.fromLocal()

    let  j = this.$store.state.cart.productCart;

    if (j.length !== 0){
      for (let i in j){
        console.log(j[i].quantity);
        this.total += (j[i].quantity * j[i].price)
      }
    }
    this.cartList = j;
  }
}
</script>

<style scoped>
.total{
  background-color: #FADADD;
  padding-top: 50px;
  padding-bottom: 50px;
  display: flex;
  text-align: center;
}
.total-txt{
  width: 50%;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
}

.total-txt h2, .total-txt h6{
  margin: 0;
}
.total .buy{
  width: 40%;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
.total .buy button{
  font-family: "Corinthia",serif;
  font-size: 25px;
  color: brown;
  border-radius: 4px;
  height: 50px;
  background-color: #c9ffb6;
  width: 50%;
}
</style>