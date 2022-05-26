<template>
  <main class="card-pay-container">
    <form>
      <div class="card-num">
        <label for="card">Card number</label>
        <input v-model="card" type="number" id="card" name="u_card"
               placeholder="xxxx-xxxx-xxxx-xxxx" pattern="[1-9]|[0-9]{15}" value="" required>
      </div>
      <div class="month-year">
        <div class="month">
          <label for="month">Month</label>
          <input v-model="month" type="number" id="month" name="u_month"
                 placeholder="mm" pattern="[1-9][0-9]" value="" required>
        </div>
        <div class="year">
          <label  for="year">Year</label>
          <input v-model="year" type="number" id="year" name="u_year"
                 placeholder="yy" pattern="[0-9][0-9]"  value="" required>
        </div>

      </div>
      <div class="cvv">
        <label for="cvv">CVV</label>
        <input v-model="cvv" type="number" id="cvv" name="cvv"
               placeholder="CVV" pattern="[1-9][0-9][0-9]" maxlength="3" value="" required>
      </div>
      <button type="button" name="button"  @click="pay">Pay</button>
    </form>
  </main>
</template>

<script>
export default {
  name: "CardMethod",
  props:{

  },
  data(){
    return{
      cvv:Number,
      card:Number,
      month:Number,
      year:Number,

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
.card-pay-container{
  display: flex;
  flex-direction: column;
}
.card-pay-container .card-num{}
.card-pay-container .month-year{

}
.card-pay-container .month-year .month{}
.card-pay-container .month-year .month input{
  width:50px;
  height:30px;
}
.card-pay-container .month-year .year{}
.card-pay-container .month-year .year input{
  width:50px;
  height:30px;
}
.card-pay-container .cvv{}
</style>