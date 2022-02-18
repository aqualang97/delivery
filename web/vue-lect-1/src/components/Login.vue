<template>

  <main class="login-form-container">
    <form class="corinthina-40" action="index.html" method="post">
      <ul>
        <li>
          <label for="email">E-mail</label>
          <input v-model="email" type="email" id="email" name="user_mail"
                 placeholder="Enter E-mail" class="corinthina-25" value="" required>
        </li>
        <li>
          <label for="password">Password</label>
          <input v-model="password" type="password" id="password" name="user_password"
                 placeholder="Enter password" value="" class="corinthina-25" required>
        </li>
      </ul>
      <button type="button" name="button" @click="login">Login</button>
      <div class="return-forgot">
        <button type="button" name="button" class="btn-bg-return-colour">Return to main</button>
        <button type="button" name="button" class="btn-bg-forgot-colour">Forgot password?</button>
      </div>
    </form>
  </main>
</template>

<script>
export default {
  name: "Login",
  props:{

  },
  data(){
    return {email:String, password:String}
  },
  methods:{
    login(){
      let email = this.email
      let password = this.password
      this.$store.dispatch(
          'login',
          {email, password}
      )
      .then(()=> this.$router.push('/')
      ).catch(err=>console.log(err))

}}}
</script>

<style scoped>
/*header, main, footer{*/
/*  max-width: 1440px;*/
/*}*/


.login-form-container{
  display: block;

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
  padding-left: 0px;
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
  width: 40%;
}
.return-forgot .btn-bg-return-colour{
}
.return-forgot .btn-bg-forgot-colour{
}

</style>