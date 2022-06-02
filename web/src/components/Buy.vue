<template>
  <main class="buy-form-container">
    <form class="fredoka-26" action="index.html" method="post">
      <ul>
        <li>
          <label for="first-name">First Name</label>
          <input v-model="firstName" type="text" id="first-name" name="user_name"
                 placeholder="Enter First name " class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="last-name">Last Name</label>
          <input v-model="lastName" type="text" id="last-name" name="user_last_name"
                 placeholder="Enter Last name" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="address">Address</label>
          <input v-model="address" type="text" id="address" name="user_address"
                 placeholder="Enter your address" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="phone">Phone number</label>
          <input v-model="phone" type="number" id="phone" name="user_phone"
                 placeholder="Enter your phone number" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="paymentMethod">Payment method</label>
          <div class="radio-cont">
            <div class="radio">
              <input type="radio" id="cash_method" name="user_payment_method"
                     class="corinthina-25" v-model="paymentMethod" value="Cash">
              <label for="cash_method">Cash</label>
            </div>
            <div class="radio">
              <input type="radio" id="online_method" name="user_payment_method"
                     class="corinthina-25" v-model="paymentMethod" value="Online payment">
              <label for="online_method">Online payment</label>
            </div>

          </div>


        </li>

      </ul>
    </form>
    <div class="buy-btn-cont">
      <div class="buy-btn">
        <div class="buy">
          <button type="button" name="button" @click="confirm" class="fredoka-26">Confirm</button>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: "Buy",
  props:{

  },
  data(){
    return{
      firstName:"",
      lastName:"",
      address:"",
      phone:Number,
      paymentMethod:"",
      user_id:0,
    }
  },
  methods:{
    async confirm() {

      if (typeof (this.firstName) !== "string" || typeof (this.lastName) !== "string" ||
          typeof (this.address) !== "string"){
        await alert("Please, enter correct data")
      }

      else if (this.firstName.length <= 2 || this.lastName.length <= 2 || this.address.length < 10 ||
          this.phone.toString().length<10) {
        await alert("Please, enter correct data")
      } else if (this.paymentMethod ===""){
        await alert("Please, choose payment method")

      }
        else {
        let user = JSON.parse(localStorage.getItem('user'))
        let user_id = user.user_id
        let cartParse = JSON.parse(localStorage.getItem('user_order'))
        let cart = []
        for (let p in cartParse) {

          let product = {
            product_id: cartParse[p].idProd,
            purchase_price: cartParse[p].price,
            quantity: cartParse[p].quantity,
          }

          cart.push(product)
        }
        let orderInfo = {
          contact_data: {
            user_id: user_id,
            first_name: this.firstName,
            last_name: this.lastName,
            address: this.address,
            phone_number: this.phone,
          },
          payment_method: this.paymentMethod,
          cart: cart,
        }
        let resp = await fetch("http://localhost:8080/confirm", {
          method: "POST",
          body: JSON.stringify(orderInfo)
        })
        if (resp.status === 200) {
          let orderId = await resp.json()
          if (this.paymentMethod === "Cash") {
            localStorage.setItem('user_pay_method', "Cash")
            localStorage.setItem('user_last_order_id', JSON.stringify(orderId))
            await this.$router.push("/pay")

          }
          else if (this.paymentMethod === "Online payment") {
            localStorage.setItem('user_pay_method', "Card")
            localStorage.setItem('user_last_order_id', JSON.stringify(orderId))
            await this.$router.push("/card-method")


          }
        }
      }
    }
  },
}
</script>

<style scoped>
@font-face {
  font-family: "Fredoka";
  src :url("../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-26{
  font-family: "Fredoka",serif;
  font-size: 26px;
}

.buy-form-container{
  display: block;

  background-color: #A2D2FF;
}

.buy-form-container form{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 250px 0 250px;
}

.buy-form-container form ul{
  padding-left: 0;
  width: 700px;

}
.buy-form-container form ul li{
  list-style-type: none;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  color: #2c3e50;
}
.buy-form-container form ul li label{
  padding-bottom: 20px;
}

.buy-form-container form ul li input{
  height: 50px;
  border-radius: 30px;
  border: 4px solid #FF865E;
}
.buy-btn-cont{
  display: flex;
  justify-content: space-around;
}
.buy-btn-cont .buy-btn{
  padding-bottom: 50px;
}
.buy-btn-cont .buy-btn .buy{
  display: flex;
  justify-content: center;
}
.buy-btn-cont .buy-btn button{
  background-color: #FEE440;
  color: #2c3e50;
  width: 300px;
  height: 50px;
  border: 4px dotted #FF865E;
  border-radius: 30px;
  /*border-color: black;*/
  float: right;
  cursor: pointer;

}

input[type~=number],input[type~=text]{
  color: #2c3e50;
  font-family: Comfortaa;
  font-size: 22px;
  font-weight: bold;
  padding-left: 20px;

}
.buy-form-container .radio-cont{
  display: flex;
  justify-content: space-between;
}
.buy-form-container .radio-cont .radio{
}
input[type~=radio]{
  width: 100px;

}
</style>