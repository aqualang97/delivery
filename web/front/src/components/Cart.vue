<template>
  <main class="cart-container">
    <div class="product" :key="idProd" >
      <div class="img-name-prod">
        <h3>{{ prodName }}</h3>
        <div :title="'Ingredients: ' + ingredients"
             @click="$router.push(`/suppliers/${suppId}/products/${idProd}`)">
          <img :src="imgLink" :alt="prodName" class="img-cart">
        </div>
      </div>
      <div class="plus-minus-cont">
        <div class="minus">
          <p @click="minusProd(idProd)">-</p>
        </div>
        <div class="quantity corinthina-40">
          <p>{{quantity}}</p>
        </div>
        <div class="plus">
          <p @click="plusProd(idProd)">+</p>
        </div>
      </div>
      <div class="name-price-cart">
        <div class="name-price corinthina-25">
          <h6>Category: {{ type }}</h6>
          <p>Price: {{ price }}$</p>
          <p>Sum: {{totalPosition.toFixed(2)}}</p>
        </div>
      </div>
    </div>

  </main>
</template>

<script>

export default {
  name: "Cart",
  props:{
    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,
    ingredients:[],
    quantity:Number,
  },
  data(){
    return{
      totalPosition:this.price*this.quantity
    }
  },
  methods:{
    totalEmit(){
      this.$emit(
          'calcSum'
      )
    },
    async plusProd(prodId){
      let  j = this.$store.state.cart.productCart;
      for (let i in j) {
        if (j[i].idProd === prodId){
          this.$store.commit('cart/plusNumProd', i);
          let newQuantity = this.$store.state.cart.productCart[i].quantity
          this.totalPosition = (newQuantity * this.price).toFixed(2)
          console.log(j);
          localStorage.setItem('user_order', "")
          localStorage.setItem('user_order', JSON.stringify(this.$store.state.cart.productCart))
          await new Promise(r => setTimeout(r, 100));
          document.location.reload();
        }
      }
    },
    async minusProd(prodId){
      let  j = this.$store.state.cart.productCart;
      for (let i in j) {
        if (j[i].idProd === prodId){
          this.$store.commit('cart/minusNumProd', i);
          let newQuantity = this.$store.state.cart.productCart[i].quantity
          this.totalPosition = (newQuantity * this.price).toFixed(2)
          console.log(this.$store.state.cart.productCart[i].quantity)
          console.log(j)
          localStorage.setItem('user_order', "")
          localStorage.setItem('user_order', JSON.stringify(this.$store.state.cart.productCart))
          await new Promise(r => setTimeout(r, 100));
          document.location.reload();
        }
      }
    }
  }
}
</script>

<style scoped>

.cart-container{
  background-color: #FADADD;
  padding-top: 50px;
  padding-bottom: 50px;
  display: block;
}


.cart-container .product{
  display: flex;

  padding-left: 20px;
  padding-right: 250px;
  flex-direction: row;
  justify-content: center;
}

.cart-container .product .img-name-prod{
  flex-direction: column;
  padding-right: 50px;
}
.cart-container.product .img-name-prod img{
  cursor: pointer;
  max-width: 300px;
  display: block;
  margin-left: auto;
  margin-right: auto;
  float: left;
}
.plus-minus-cont{
  display: flex;
  align-items: center;
  font-size: 30pt;
}

.plus-minus-cont .plus{
  cursor: pointer;
}
.plus-minus-cont .minus{
  cursor: pointer;

}
.plus-minus-cont .quantity{

}
.img-cart{
  max-width: 300px;
}
.your-cart{
  text-align: center;
  font-weight: bold;
}

.products-in-cart{
  display: flex;
  width: 60%;
  margin-left: auto;
  margin-right: auto;
  padding-bottom: 35px;


}
.prod{
  padding-top: 30px;
  border: 1px dashed #a9a9a9;
}
.prod img{
  width: 15%;
  display: block;
  margin-left: auto;
  margin-right: auto;

}

.prod h3{
  text-align: center;
  margin: 0;
}

.prod p{
  text-align: end;
  margin: 0;
  padding: 20px;
}

.plus, .minus{
  width: 20%;
  margin: auto;
  padding-right: 20px;
  padding-left: 20px;
  display: flex;
  justify-content: center;
}
.numbers .plus img, .numbers .minus img{
  width: 100%;
}

.total{
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