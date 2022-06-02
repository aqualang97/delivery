<template>
  <main class="pay-form-container">
    <form class="fredoka-26">
      <div class="inputs">
        <label for="card">Card number</label>
        <input minlength="16" maxlength="16" v-model="card" type="text" id="card" name="u_card"
               placeholder="xxxx-xxxx-xxxx-xxxx" pattern="[1-9]|[0-9]{15}" value="" required>
      </div>
      <div>
        <div class="inputs">
          <label for="month">Month</label>
          <input v-model="month" type="text" id="month" name="u_month"
                 placeholder="mm" pattern="[1-9][0-9]" minlength="2" maxlength="2" value="" required>
        </div>
        <div class="inputs">
          <label  for="year">Year</label>
          <input v-model="year" type="text" id="year" name="u_year"
                 placeholder="yy" pattern="[0-9][0-9]" minlength="2" maxlength="2" value="" required>
        </div>

      </div>
      <div class="inputs">
        <label for="cvv">CVV</label>
        <input v-model="cvv" type="text" id="cvv" name="cvv"
               placeholder="CVV" pattern="[1-9][0-9][0-9]" minlength="3" maxlength="3" value="" required>
      </div>
    </form>
    <div class="pay-btn">
      <div class="pay">
        <button type="button" name="button"  @click="pay" class="fredoka-26">Pay</button>

      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: "CardMethod",
  props:{

  },
  data(){
    return{
      cvv:String(),
      card:String(),
      month:String(),
      year:String(),

    }
  },
  methods:{
    async pay(){
      console.log(this.card)
      let simulation = {
        card:this.card,
        status:"paid",
        last_order_id: JSON.parse(localStorage.getItem('user_last_order_id'))
      }
      let resp = await fetch("http://localhost:8080/card_pay", {
        method: "POST",
        body: JSON.stringify(simulation)
      })
      if (resp.status === 200) {
        alert("Ok")
        await this.$router.push("/pay")
      }else{
        alert(resp.status)
      }
    },
  }
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

.pay-form-container{
  display: flex;
  align-items: center;
  flex-direction: column;
  background-color: #A2D2FF;
}

.pay-form-container form{
  padding: 60px 250px;
}


.pay-form-container label{
  padding-bottom: 20px;
  padding-right: 20px;
}

.pay-form-container input{
  height: 50px;
  border-radius: 30px;
  border: 4px solid #FF865E;

}
.pay-form-container .inputs{
  padding-top: 10px;
  padding-bottom: 10px;
  display: flex;
  justify-content: flex-end;
}

.pay-btn{
  padding-bottom: 50px;
}
.pay-btn .pay{
  display: flex;
  justify-content: center;
}
.pay-btn button{
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

input[type~=text]{
  color: #2c3e50;
  font-family: Comfortaa;
  font-size: 22px;
  font-weight: bold;
  padding-left: 20px;

}
</style>