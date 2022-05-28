<template>

  <div class="product-light-list">
    <div class="product-light-elem" >
      <div class="product-container">
        <div class="prod-info">
          <div class="img-name-prod">
            <div :title="'Ingredients: ' + ingredients"
                 @click="$router.push(`/suppliers/${suppId}/products/${idProd}`)">
              <img :src="imgLink" :alt="prodName" class="comfortaa-22">
            </div>
            <h3 class="fredoka-38 prod-name">{{ prodName }}</h3>

          </div>
        </div>
        <div class="category-price-cont">
          <div class="name-price-cart">
            <div class="name-price fredoka-26">
              <div class="category">
                <h6 class="comfortaa-22">Category:</h6>
                <h6 class="category-route comfortaa-26"  @click="$router.push(`/products&category=${idCat}`)"
                >{{replaceAndTitle(type) }}</h6>
              </div>
              <p class="fredoka-38">Price: {{ price }}$</p>
              <div class="fredoka-34 add-plus-minus">
                <div class="btn-add-to-cart">
                  <div v-if="$store.state.cart.productCart.length !==0 ">
                    <div v-if="notInArray(idProd)">
                      <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)"
                              class="fredoka-26">Add!</button>
                    </div>
                  </div>
                  <div v-else>
                    <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)"
                            class="fredoka-26">Add!</button>
                  </div>
                </div>
                <div class="plus-minus-cont fredoka-30">
                  <div v-for="j in $store.state.cart.productCart" :key="j.idProd" class="plus-minus-wrap">
                    <div v-if="j.idProd === idProd">
                      <button type="button" name="plus-to-cart" :id="`plus-prod${idProd}`" @click="plusToCart(idProd)"
                      class="fredoka-30 fix-btn-padding" >+</button>
                      <p>{{j.quantity}}</p>
                      <button type="button" name="minus-from-cart" :id="`minus-prod${idProd}`"
                              @click="minusFromCart(idProd)" class="fredoka-30 fix-btn-padding">-</button>
                    </div>
                  </div>

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

  name: "Product",
  title:{
    type: String,
    required: true,

  },
  components:{

  },
  props:{
    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,//name of category
    idCat:Number,
    suppId:Number,
    externalSuppId:Number,
    ingredients:[],
    quantityInCart:Number,
    listId:[],
  },

  data(){
    return{
      quantity:0,
    }
  },
  mounted() {
  },
  methods: {
    // test(j, idProd){
    //   console.log("j.idProd", j,"\nidProd", idProd)
    // },

    notInArray(idProd){
      for(let id in  this.listId){
        if(this.listId[id] === idProd){
          return false
        }
      }
      return true
    },
    // add() {
    //   this.$emit(
    //       'plus', 'minus'
    //   )
    //   let b = document.getElementById(`prod${this.idProd}`)
    //   b.parentNode.removeChild(b)
    //   let minus = document.createElement("minusbtn")
    //   minus.setAttribute("id", `products${this.idProd}`);
    //   minus.innerHTML = "1111111"
    //   document.body.append(minus);
    // },

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
      let elem = document.getElementById(`add-prod${this.idProd}`)
      elem.parentNode.removeChild(elem);

      let j = this.$store.state.cart.productCart
      if (j.length === 0) {
        this.$store.commit('cart/addProdToCart', prod)

      } else {
        for (let i in j) {
          if (j[i].idProd === prodIdAdd) {
            this.$store.commit('cart/plusNumProd', i)
            return
          }
        }
        this.$store.commit('cart/addProdToCart', prod)
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
          // console.log(this.$store.state.cart.productCart[i].quantity)
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


}
</script>

<style scoped>

@font-face {
  font-family: "Corinthia";
  //src: url("/home/yurii/delivery/web/shop/fonts/Corinthia-Regular.ttf");
  font-weight: 400;
}

@font-face {
  font-family: "Comfortaa";
  src :url("../../../fonts/comfortaa/Comfortaa-VariableFont_wght.ttf");
  font-weight: 400;
}
@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
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
}.comfortaa-22{
  font-family: "Comfortaa",serif;
  font-size: 22px;
}
.product-light-list{
  display: flex;
  background-color: #FEE440;
  justify-content: space-around;
  padding-top: 40px;
  min-width: 700px;
  max-width: 1440px;

}

.product-light-elem{
  /*padding: 30px 150px 30px 0;*/
  border: 4px dotted #FF865E;
  border-radius: 35px;
  /*padding-left: 70px;*/
  /*padding-right: 70px;*/
  background-color: #FCEA60;
  min-height: 800px;


}
.product-light-list .product-light-elem .product-container{
  display: flex;

  /*padding-left: 20px;*/
  /*padding-right: 250px;*/
  /*justify-content: end;*/
  flex-direction: column;
  min-width: 680px;
  min-height: 715px;

}
.product-light-list .product-light-elem .product-container .prod-info{
  min-width: 350px;
}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod{
  display: flex;
  flex-direction: column;
  min-height: 520px;
}

.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod h3{
  display: flex;
}

.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod .prod-name{
  /*width: 300px;*/
  margin: 0;
  justify-content: center;
  max-width: 680px;
}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod img{
  cursor: pointer;
  max-width: 450px;
  /*max-height: 415px;*/
  display: block;
  /*margin-left: auto;*/
  /*margin-right: auto;*/
  float: left;
  /*padding-bottom: 15px;*/
  margin: 30px;

}

.product-light-list  .product-light-elem .product-container .category-price-cont{
}

.name-price-cart{
  /*padding-top: 40px;*/
  /*padding-left: 50px;*/

}
.name-price-cart .name-price{
  display: flex;
  flex-direction: column;
}
.name-price-cart .name-price .add-plus-minus{
  display: flex;
  justify-content: center;
}

.name-price-cart .name-price .category{
  display: flex;
  margin-left: 15px;
  margin-top: 20px;

}
.name-price-cart .name-price .category .category-route{
  cursor: pointer;
  color: #FF865E;
}

/*.name-price-cart .name-price .category .category-route :hover .category-route {*/
/*  color: red;*/
/*}*/
.name-price-cart .name-price h3{

}
.name-price-cart .name-price h6{
  margin: 15px;

}
.name-price-cart .name-price h3,.name-price-cart .name-price p{
  margin: 0;
}
.name-price-cart .name-price .btn-add-to-cart{
  /*padding-right: 10px;*/
  display: flex;
  /*justify-content: space-around;*/
  padding-top: 20px;

}
.name-price-cart .name-price .btn-add-to-cart button{

  background-color: #FF865E;
  color: #1a1f1c;
  width: 100px;
  height: 50px;

  border-width: 3px;
  border-radius: 16px;
  border-color: black;
  float: right;
  cursor: pointer;
}
.name-price-cart .name-price .btn-add-to-cart button:hover button{
  background-color: #f28ec1;
}

.name-price-cart .name-price .plus-minus-cont {
  display: flex;
  /*justify-content: center;*/
  /*position: static;*/
}
.name-price-cart .name-price .plus-minus-cont .plus-minus-wrap{
  display: flex;
  flex-wrap: wrap;
}

.name-price-cart .name-price .plus-minus-cont p{
  /*padding: 0;*/
}
.name-price-cart .name-price .plus-minus-cont .fix-btn-padding{
}
.name-price-cart .name-price .plus-minus-cont button{
  background-color: #FF865E;
  color: #1a1f1c;
  width: 100px;
  height: 50px;

  border-width: 3px;
  border-radius: 16px;
  border-color: black;
  float: right;
  cursor: pointer;

}


.name-price-cart .name-price .plus-minus-cont button p{
  padding: 0;
  margin: 0;
}
/*.name-price-cart .cart{*/
/*}*/

/*.name-price-cart .cart img{*/
/*  width: 40%;*/

/*}*/

.corinthina-25{
  font-family: "Corinthia",serif;
  font-size: 25px;
}

</style>