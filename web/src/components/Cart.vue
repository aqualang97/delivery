<template>
  <main class="cart-container">
    <div class="product" :key="idProd" >
      <div class="img-name-prod fredoka-38">
        <h3>{{ prodName }}</h3>
        <div :title="'Ingredients: ' + ingredients"
             @click="$router.push(`/suppliers/${suppId}/products/${idProd}`)">
          <img :src="imgLink" :alt="prodName" class="img-cart">
        </div>
      </div>

      <div class="name-price-cart">
        <div class="name-price corinthina-25">
          <h6 class="comfortaa-22">Category:</h6>
          <h6 class="category-route comfortaa-26"  @click="$router.push(`/products&category=${idCat}`)"
          >{{replaceAndTitle(type) }}</h6>
          <p class="fredoka-34">Price: {{ price }}$</p>
          <p class="fredoka-38">Sum: {{totalPosition.toFixed(2)}}</p>
          <div class="plus-minus-cont fredoka-30 ">
            <div class="minus">
              <button class="fredoka-30" @click="minusProd(idProd)">-</button>
            </div>
            <div class="quantity corinthina-40">
              <p>{{quantity}}</p>
            </div>
            <div class="plus">
              <button class="fredoka-30" @click="plusProd(idProd)">+</button>
            </div>
          </div>
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
    idCat:Number,
    suppId:Number,
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
          localStorage.setItem('user_order', "")
          localStorage.setItem('user_order', JSON.stringify(this.$store.state.cart.productCart))
          await new Promise(r => setTimeout(r, 100));
          document.location.reload();
        }
      }
    },
    replaceAndTitle(str){
      str=str.replace("_", " ")
      str = str[0].toUpperCase() + str.substring(1)
      return str
    },
    async minusProd(prodId){
      let  j = this.$store.state.cart.productCart;
      for (let i in j) {
        if (j[i].idProd === prodId){
          this.$store.commit('cart/minusNumProd', i);
          let newQuantity = this.$store.state.cart.productCart[i].quantity
          this.totalPosition = (newQuantity * this.price).toFixed(2)
          localStorage.setItem('user_order', "")
          localStorage.setItem('user_order', JSON.stringify(this.$store.state.cart.productCart))
          await new Promise(r => setTimeout(r, 100));
          document.location.reload();
        }
      }
    },
  }
}
</script>

<style scoped>

@font-face {
  font-family: "Comfortaa";
  src :url("../../fonts/comfortaa/Comfortaa-VariableFont_wght.ttf");
  font-weight: 400;
}
@font-face {
  font-family: "Fredoka";
  src :url("../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-26{
  font-family: "Fredoka",serif;
  font-size: 26px;
}

.fredoka-30{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 30px;
}
.fredoka-34{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 34px;
}.fredoka-38{
   font-family: "Fredoka",serif;
   font-weight: normal;
   font-size: 38px;
 }
.comfortaa-26{
  font-family: "Comfortaa",serif;

  font-size: 26px;
}
.comfortaa-22{
  font-family: "Comfortaa",serif;
  font-size: 22px;
}

.cart-container{
  background-color: #A2D2FF;
  padding-top: 50px;
  padding-bottom: 50px;
  display: flex;
}



.cart-container .product{
  display: flex;
  flex-direction: row;
  justify-content: center;
  border: 0;
  border-radius: 30px;
  background-color: #c6e1ff;
  width: 900px;
  height: 520px;
}
.cart-container .product .name-price-cart{}
.cart-container .product .name-price-cart .name-price{
  padding-top: 50px;
}
.cart-container .product .name-price-cart .name-price h6{
  margin: 15px;
}
.cart-container .product .name-price-cart .name-price .category-route{
  cursor: pointer;
  color: #FF865E;
}
.cart-container .product .name-price-cart .name-price .plus-minus-cont{

}
.cart-container .product .name-price-cart .name-price .plus-minus-cont .plus{}
.cart-container .product .name-price-cart .name-price .plus-minus-cont .minus{}
.cart-container .product .name-price-cart .name-price .plus-minus-cont .quantity{}


.cart-container .product .img-name-prod{

  display: flex;
  flex-direction: column;
  width: 500px;
  align-items: center;


}
.cart-container .product .img-name-prod h3{
  font-weight: normal;
}
.cart-container .product .img-name-prod img{
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
.plus-minus-cont button{
  background-color: #FEE440;
  color: #2c3e50;
  width: 100px;
  height: 50px;

  border: 4px dotted #FF865E;
  border-radius: 30px;
  /*border-color: black;*/
  float: right;
  cursor: pointer;
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