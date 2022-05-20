<template>
  <main class="buy-form-container">
    <form class="corinthina-40" action="index.html" method="post">
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
          <div>
            <input type="radio" id="cash_method" name="user_payment_method"
                   class="corinthina-25" v-model="paymentMethod" value="Cash">
            <label for="cash_method">Cash</label>
            <input type="radio" id="online_method" name="user_payment_method"
                   class="corinthina-25" v-model="paymentMethod" value="Online payment">
            <label for="online_method">Online payment</label>
          </div>


        </li>
        <button type="button" name="button" @click="confirm">Confirm</button>

      </ul>
    </form>
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
      // console.log(typeof this.phone )
      // console.log(typeof (this.firstName) !== "string" )
      // console.log( typeof(this.lastName) !== "string")
      // console.log( typeof(this.address) !== "string")
      // console.log(this.firstName.length <= 2 )
      // console.log( this.lastName.length <= 2 )
      // console.log(this.address.length < 10)
      // console.log(this.phone.toString().length<10 )
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
        console.log(user_id)
        console.log("user_id", user_id, "firstName", this.firstName, "lastName", this.lastName,
            "address", this.address, "phone", this.phone, "paymentMethod", this.paymentMethod)
        let cartParse = JSON.parse(localStorage.getItem('user_order'))
        let cart = []
        console.log("cartParse", cartParse)
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
        console.log(orderInfo)
        let resp = await fetch("http://localhost:8080/confirm", {
          method: "POST",
          body: JSON.stringify(orderInfo)
        })
        console.log(resp.status)
        if (resp.status === 200) {
          let orderId = await resp.json()
          if (this.paymentMethod === "Cash") {
            await this.$router.push("/pay")
            localStorage.setItem('user_pay_method', "Cash")
            localStorage.setItem('user_last_order_id', JSON.stringify(orderId))
          }
          else if (this.paymentMethod === "Online payment") {
            await this.$router.push("/card-method")
            localStorage.setItem('user_pay_method', "Card")
            localStorage.setItem('user_last_order_id', JSON.stringify(orderId))

          }
        }
      }
    }
  },
}
</script>

<style scoped>
.buy-form-container{
  display: block;

  background-color: #FADADD;
}
.buy-form-container form{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 100px 250px 100px 250px;
}

.buy-form-container form ul{
  padding-left: 0px;
  width: 100%;

}
.buy-form-container form ul li{
  list-style-type: none;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  color: #26938e;
}

.buy-form-container form ul li input{
  height: 30px;
  border-radius: 8px;
}

.buy-form-container button{
  background-color: #c9ffb6;
  color: black;
  border-radius: 8px;
  width: 100%;
  height: 50px;
  font-family: "Corinthia";
  font-size: 30px;

}

</style>