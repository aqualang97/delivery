<template>
<div>
  <div class="cart-prod-cont">
    <div class="cart-responsive">
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
          :id-cat="i.idCat"
          :supp-id="i.suppId"



      ></cart>
    </div>
  </div>

  <div class="total">
    <div v-if="this.$store.state.cart.productCart.length===0">
      <img src="../../../pic/empty-cart.png">
    </div>
    <div v-else class="total-not-empty" >
      <div class="total-txt ">
        <h2 class="fredoka-28">Total payment:</h2>
        <h2 class="fredoka-44">{{ total.toFixed(2) }} $</h2>
      </div>
      <div class="buy-clear" >
        <div class="btn-buy">
          <button class="fredoka-34" type="button" name="button"
                  @click="$router.push(`/buy`)&&checkLogin()">
            Buy
          </button>
        </div>

        <div class="btn-clear">
          <button class="fredoka-34" type="button" name="button"
                  @click="clear">
            Clear cart
          </button>
        </div>


      </div>
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



@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-44{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 44px;
}
.fredoka-34{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 34px;
}
.fredoka-28{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 29px;
}
.total{
  background-color: #c6e1ff;
  padding-top: 50px;
  padding-bottom: 50px;
  display: flex;
  justify-content: center;
}
.total img{
  opacity: 0.3;
}
.total .total-not-empty{

}
.total .total-not-empty .total-txt{
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
}

.total .total-not-empty .total-txt h2, .total-not-empty .total-txt h6{
  margin: 0;
  font-weight: normal;
  padding-bottom: 25px;
}

.total .buy-clear{
  display: flex;
  justify-content: flex-end;
  align-items: center;

}

.total .buy-clear .btn-clear{
  padding-right: 25px;
  padding-left: 25px;
}
.total .buy-clear .btn-buy{
  padding-right: 25px;
  padding-right: 25px;
}
.total .buy-clear button{
  background-color: #FEE440;
  color: #2c3e50;
  width: 400px;
  height: 50px;
  border: 4px dotted #FF865E;
  border-radius: 30px;
  /*border-color: black;*/
  float: right;
  cursor: pointer;
}
.cart-prod-cont{
  background-color: #A2D2FF;
}
.cart-prod-cont .cart-responsive{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}
</style>