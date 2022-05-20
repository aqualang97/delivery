<template>
  <div class="product-light-list">
    <div class="product-light-elem">
      <div class="product">
        <div class="img-name-prod">
          <h3>{{ prodName }}</h3>
          <div :title="'Ingredients: ' + ingredients"
               @click="$router.push(`/suppliers/${suppId}/products/${idProd}`)"
          >
            <img :src="imgLink" :alt="prodName">
          </div>
        </div>

        <div class="name-price-cart">
          <div class="name-price corinthina-25">
            <h6>Category: {{ replaceAndTitle(type) }}</h6>
            <p>Price: {{ price }}$</p>
            <div class="btn-add-to-cart">
              <button type="button" name="add-to-cart" :id="`prod${idProd}`" @click="addToCart(idProd)" >Add!</button>
              <div v-for="j in $store.state.cart.productCart" :key="j.idProd">
                <div v-if="j.idProd === idProd" >
                  <button type="button" name="plus-to-cart" :id="`prod${idProd}`" @click="plusToCart(idProd)">+</button>
                  <p>{{j.quantity}}</p>
                  <button type="button" name="minus-from-cart" :id="`prod${idProd}`" @click="minusFromCart(idProd)">-</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductCategory",
  props:{
    // title:{
    //   type: String,
    //   required: true,
    // },
    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,

    suppId:Number,
    externalSuppId:Number,
    ingredients:[],
    quantityInCart:Number
  },
  data(){
    return{
      quantity:0,
    }
  },
  methods:{
    addToCart(prodIdAdd) {
      const prod = {
        idProd: this.idProd,
        externalProdId: this.externalProdId,
        prodName: this.prodName,
        price: this.price,
        imgLink: this.imgLink,
        type: this.type,

        suppId: this.suppId,
        externalSuppId: this.externalSuppId,
        ingredients: this.ingredients,

        quantity: this.quantity,
      }
      let elem = document.getElementById(`prod${this.idProd}`)
      elem.parentNode.removeChild(elem);

      let j = this.$store.state.cart.productCart
      console.log(j.length)
      if (j.length === 0) {
        this.$store.commit('cart/addProdToCart', prod)
        console.log(j)

      } else {
        for (let i in j) {
          if (j[i].idProd === prodIdAdd) {
            this.$store.commit('cart/plusNumProd', i)
            console.log(j)
            return
          }
        }
        this.$store.commit('cart/addProdToCart', prod)
        console.log(j)

      }


      // if(prod in  this.$store.state.cart.productCart){
      //   console.log(this.$store.state.cart.quantity)
      //   // this.$store.commit('cart/updateNumProd',  )
      //   console.log("+")
      //
      // }else {
      //   this.$store.commit('cart/addProdToCart', prod)
      //   console.log("-")
      // }
      // console.log(this.$store.state.cart)
    },
    plusToCart(prodIdAdd) {
      let  j = this.$store.state.cart.productCart;
      for (let i in j) {
        if (j[i].idProd === prodIdAdd){
          this.$store.commit('cart/plusNumProd', i);
          let newQuantity = this.$store.state.cart.productCart[i].quantity
          this.totalPosition = (newQuantity * this.price).toFixed(2)
          console.log(j);
          return
        }
      }
    },
    minusFromCart(prodIdMinus) {
      let  j = this.$store.state.cart.productCart;
      for (let i in j) {
        if (j[i].idProd === prodIdMinus){
          this.$store.commit('cart/minusNumProd', i);
          let newQuantity = this.$store.state.cart.productCart[i].quantity
          this.totalPosition = (newQuantity * this.price).toFixed(2)
          console.log(this.$store.state.cart.productCart[i].quantity)
          console.log(j)
          return
        }
      }
    },
    replaceAndTitle(str){
      str=str.replace("_", " ")
      str = str[0].toUpperCase() + str.substring(1)
      return str
    },
  }

};
</script>

<style scoped>

@font-face {
  font-family: "Corinthia";
//src: url("/home/yurii/delivery/web/shop/fonts/Corinthia-Regular.ttf");
  font-weight: 400;
}
.product-light-list{
  display: block;
  background-color: aquamarine;
}

.product-light-elem{
  padding: 30px 150px 30px 0;

}
.product-light-elem .product{
  display: flex;

  padding-left: 20px;
  padding-right: 250px;
  flex-direction: row;
  justify-content: center;
}
.product-light-list .product .img-name-prod{
  flex-direction: column;
  padding-right: 50px;
}
.product-light-elem .product img{
  cursor: pointer;
  max-width: 300px;
  display: block;
  margin-left: auto;
  margin-right: auto;
  float: left;
}

.name-price-cart{
  padding-top: 40px;
  padding-left: 50px;

}
.name-price-cart .name-price{

}
.name-price-cart .name-price h3{

}
.name-price-cart .name-price h6{
  margin: 15px;

}
.name-price-cart .name-price h3,.name-price-cart .name-price p{
  margin: 0;
  padding-bottom: 30px;
}
.name-price-cart .name-price .btn-add-to-cart{
  padding-right: 10px;
}
.name-price-cart .name-price button{

  background-color: #cf1ad5;
  color: #1a1f1c;
  width: 50px;
  height: 30px;
  border-radius: 4px;
  float: right;

}

.corinthina-25{
  font-family: "Corinthia",serif;
  font-size: 25px;
}
</style>