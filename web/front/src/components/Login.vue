<template>

  <main class="login-form-container">
    <form class="fredoka-26" action="index.html" method="post">
      <ul>
        <li>
          <label for="email">E-mail</label>
          <input v-model="emailInput" type="email" id="email" name="user_mail"
                 placeholder="Enter E-mail" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="password">Password</label>
          <input v-model="passwordInput" type="password" id="password" name="user_password"
                 placeholder="Enter password" value="" class="corinthina-25" required>
        </li>
      </ul>

    </form>
    <div class="login-return-cont">
      <div class="login-return">
        <div class="login">
          <button type="button" name="button" @click="login" class="fredoka-26">Login</button>

        </div>
        <div class="return-forgot">
          <div class="return">
            <button type="button" name="button" class="fredoka-26" @click="returnToMain">Return to main</button>
          </div>
          <div class="forgot">
            <button type="button" name="button" class="fredoka-26">Forgot password?</button>
          </div>

        </div>
      </div>
    </div>

  </main>
</template>

<script>
export default {
  name: "Login",
  props:{
  },
  data(){
    return{
      body:null,
      emailInput:"",
      passwordInput:"",
    }
  },
  methods:{
    returnToMain(){
      this.$router.push("/")
    },
    async login(){
      localStorage.setItem('user', "")
      // localStorage.clear()
      // this.$store.commit('cart/clearCart');

      let resp = await fetch("http://localhost:8080/login",{
        method: "POST",



        body: JSON.stringify({email:this.emailInput, password:this.passwordInput})
      })
      if (resp.status !== 200){
        alert(resp.statusText)
        console.log(resp.status)
        return
      }
      let data = await resp.json()
      console.log(data)
      alert("Successful authorization")
      localStorage.setItem('user', JSON.stringify(data))
      await this.$router.push("/cart")
      await document.location.reload()

      // for (let index in await resp.json()){
      //   console.log(index)
      // }

      // .then(()=> this.$router.push('Home')
      // ).catch(err=>console.log(err))

}}}
</script>

<style scoped>
/*header, main, footer{*/
/*  max-width: 1440px;*/
/*}*/

@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-26{
  font-family: "Fredoka",serif;
  font-size: 26px;
}



::-webkit-input-placeholder {
  color: #b0bfd2;
  font-family: Comfortaa;
  font-size: 22px;
  font-weight: bold;
  padding-left: 20px;
}
.login-form-container{

  background-color: #A2D2FF;
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
  padding: 60px 250px 0 250px;
}

.login-form-container form ul{
  padding-left: 0;
  width: 700px;

}
.login-form-container form ul li{
  list-style-type: none;
  display: flex;
  flex-direction: column;
  padding-bottom: 20px;
  color: #2c3e50;
}
.login-form-container form ul li label{
  padding-bottom: 20px;
}

.login-form-container form ul li input{
  height: 50px;
  border-radius: 30px;
  border: 4px solid #FF865E;
}
.login-return-cont{
  display: flex;
  justify-content: space-around;
}
.login-return{
  padding-bottom: 50px;
}
.login-return-cont .login-return .login{
  display: flex;
  justify-content: center;
}
.login-return-cont .login-return button{
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
.login-form-container .login-return-cont .return-forgot{
  display: flex;
  justify-content: space-around;
  margin: 0;
  padding-top: 50px;
}
.login-form-container .login-return-cont .return-forgot button{
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

.login-form-container .login-return-cont .return-forgot .return, .login-form-container .login-return-cont .return-forgot .forgot{
  padding-left: 20px;
  padding-right: 20px;
}

.return-forgot button{
  height: 40px;

}

</style>