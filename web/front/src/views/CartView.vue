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
      <h2>{{ total }} $</h2>
    </div>
    <div class="buy">
      <button type="button" name="button">Buy</button>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "CartView",
  components: {},
  methods:{
    plus(numProdCart){
      return numProdCart+=1
    },
    minus(numProdCart){
      return numProdCart-=1
    }
  },
  data() {
    return {
      cartList:null,
      total:0
    };
  },
  mounted() {
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