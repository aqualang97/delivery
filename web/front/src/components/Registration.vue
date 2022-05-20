<template>
  <main class="login-form-container">
    <form class="corinthina-40" action="index.html" method="post">
      <ul>
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
      <button type="button" name="button" @click="registration">Registration</button>
      <div class="return-forgot">
        <button type="button" name="button" class="btn-bg-return-colour">Return to main</button>
      </div>
    </form>
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

      let data = await resp.json()
      console.log(data)
      alert("Successful registration")
      // localStorage.setItem('user', JSON.stringify(data))
      // localStorage.setItem("user_order", JSON.stringify([]))

      await this.$router.push("/cart")
      // console.log((await Promise.resolve(resp.text())).toString()


    }
  }
}
</script>


<style scoped>
header, main, footer{
  max-width: 1440px;
}


.login-form-container{

  background-color: #FADADD;
}
/*.ubuntu-25{
  font-size: 25px;
}
.ubuntu-15{
  font-size: 15px;
}*/
.login-form-container form{
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 100px 250px 100px 250px;
}

.login-form-container form ul{
  padding-left: 0;
  width: 100%;

}
.login-form-container form ul li{
  list-style-type: none;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  color: #26938e;
}

.login-form-container form ul li input{
  height: 30px;
  border-radius: 8px;
}

.login-form-container button{
  background-color: #c9ffb6;
  color: black;
  border-radius: 8px;
  width: 100%;
  height: 50px;
  font-family: "Corinthia";
  font-size: 30px;

}

.login-form-container .return-forgot{
  display: flex;
  justify-content: space-between;
  margin: 0;
  width: 100%;
  padding-top: 50px;

}
.return-forgot button{
  height: 40px;
  width: 100%;

}
.return-forgot .btn-bg-return-colour{
}
.return-forgot .btn-bg-forgot-colour{
}

</style>