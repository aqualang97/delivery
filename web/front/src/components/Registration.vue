<template>
  <main class="reg-form-container">
    <form class="corinthina-40" action="index.html" method="post">
      <ul class="fredoka-26">
        <li>
          <label for="emailReg">E-mail</label>
          <input v-model="emailReg" type="email" id="emailReg" name="user_mail_reg"
                 placeholder="Enter E-mail" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="loginReg">Login</label>
          <input v-model="loginReg" type="text" id="loginReg" name="user_login_reg"
                 placeholder="Enter login" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="passwordReg">Password</label>
          <input v-model="passwordReg" type="password" id="passwordReg" name="user_password_reg"
                 placeholder="Enter password" value="" class="corinthina-25" required>
        </li>

        <li>
          <label for="confirmPasswordReg">Confirm Password</label>
          <input v-model="confirmPasswordReg" type="password" id="confirmPasswordReg" name="user_confirm_password_reg"
                 placeholder="Enter password" value="" class="corinthina-25" required>
        </li>

<!--        <li>-->
<!--          <label for="address">Address</label>-->
<!--          <input v-model="address" type="text" id="address" name="user_address"-->
<!--                 placeholder="Enter address" value="" class="corinthina-25" required>-->
<!--        </li>-->
<!--        <li>-->
<!--          <label for="phone">Phone Number</label>-->
<!--          <input v-model="phone" type="text" id="phone" name="user_phone"-->
<!--                 placeholder="Enter phone" value="" class="corinthina-25" required>-->
<!--        </li>-->
      </ul>
    </form>
    <div class="reg-return-cont">
      <div class="reg-return">
        <div  class="reg">
          <button type="button" name="button" @click="registration" class="fredoka-26">Registration</button>
        </div>
        <div>
          <button type="button" name="button" @click="$router.push(`/`)" class="fredoka-26">Return to main</button>
        </div>
      </div>


    </div>
  </main>
</template>

<script>
export default {
  name: "Registration",
  props:{

  },
  data(){
    return{
      emailReg:"",
      loginReg:"",
      passwordReg:"",
      confirmPasswordReg:"",
      emptyStorage:[],
    }
  },
  methods:{
    async registration(){
      console.log(this.passwordReg)
      if ((this.loginReg || this.passwordReg || this.confirmPasswordReg|| this.emailReg) === ""){
        await alert("All fields must be completed")
        return
      }
      if (this.passwordReg !==this.confirmPasswordReg){
        await alert("The passwords are different")
        this.passwordReg = this.confirmPasswordReg = ""

        return

      }
      if (this.passwordReg.length < 8){
        await alert("The password is short")
        this.passwordReg = this.confirmPasswordReg = ""
        return

      }
      // localStorage.clear()
      // this.$store.commit('cart/clearCart');

      localStorage.setItem('user', "")
      let resp = await fetch("http://localhost:8080/registration",{
        method: "POST",
        body: JSON.stringify({email:this.emailReg, login: this.loginReg, password:this.passwordReg})
      })
      if (resp.status !== 200){
        alert((await Promise.resolve(resp.text())).toString())
        this.emailReg = this.loginReg = this.passwordReg = this.confirmPasswordReg = ""
      }
      if(this.validEmail(this.emailReg)){
        let data = await resp.json()
        console.log(data)
        alert("Successful registration")
        await this.$router.push("/cart")
      }else {
        alert("Check spelling")
      }



    },
    validEmail(email) {
      let re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@(([[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }
  }
}
</script>


<style scoped>

@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-26{
  font-family: "Fredoka",serif;
  font-size: 26px;
}

header, main, footer{
}

::-webkit-input-placeholder {
  color: #b0bfd2;
  font-family: Comfortaa;
  font-size: 22px;
  font-weight: bold;
  padding-left: 20px;
}
.reg-form-container{

  background-color: #A2D2FF;
}
/*.ubuntu-25{
  font-size: 25px;
}
.ubuntu-15{
  font-size: 15px;
}*/
.reg-form-container form{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 60px 250px 0 250px;
}

.reg-form-container form ul{
  padding-left: 0;
  width: 700px;

}
.reg-form-container form ul li{
  list-style-type: none;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  color: #2c3e50;
}
.reg-form-container form ul li label{
  padding-bottom: 20px;
}

.reg-form-container form ul li input{
  height: 50px;
  border-radius: 30px;
  border: 4px solid #FF865E;
}
.reg-return-cont{
  display: flex;
  justify-content: space-around;
}
.reg-return{
  padding-bottom: 50px;
}
.reg-return-cont .reg-return .reg{
  padding-bottom: 100px;
}
.reg-return-cont .reg-return button{
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

input[type~=email],input[type~=text], input[type~=password]{
  color: #2c3e50;
  font-family: Comfortaa;
  font-size: 22px;
  font-weight: bold;
  padding-left: 20px;

}
</style>