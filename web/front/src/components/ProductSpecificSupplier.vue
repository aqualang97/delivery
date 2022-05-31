<template>
  <div class="product-light-list">
    <div class="product-light-elem" >
      <div class="product-container">
        <div class="prod-info">
          <div class="img-name-prod">
            <div class="img-ingredients comfortaa-22">
              <img :src="imgLink" :alt="prodName" >
            </div>
            <div class="name">
              <h3 class="prod-name fredoka-38">{{ prodName }}</h3>
            </div>
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
              <div class="price">
                <p class="  fredoka-38 ">Price: {{ price }}$</p>

              </div>
              <p class="ing fredoka-30">Ingredients:</p>
              <div v-for="(i) in ingredients"
                   :key="i"
              >
                <div class="comfortaa-26">
                  {{ i }}
                </div>
              </div>
              <div class="fredoka-34 add-plus-minus">
                <div class="btn-add-to-cart">
                  <div v-if="$store.state.cart.productCart.length !==0 ">
                    <div v-if="notInArray(idProd)">
                      <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)"
                              class="fredoka-26">Add to cart</button>
                    </div>
                  </div>
                  <div v-else>
                    <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)"
                            class="fredoka-26">Add to cart</button>
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

<!--<template>-->

<!--  <div class="product-light-list">-->
<!--    <div class="product-light-elem">-->
<!--      <div class="product">-->
<!--        <div class="img-name-prod">-->
<!--          <h3>{{ prodName }}</h3>-->
<!--          <div>-->
<!--            <img :src="imgLink" :alt="prodName">-->
<!--          </div>-->

<!--        </div>-->

<!--        <div class="name-price-cart">-->
<!--          <div class="name-price corinthina-25">-->
<!--            <h6>Category:</h6>-->
<!--            <h6>{{ replaceAndTitle(type) }}</h6>-->
<!--            <p>Price: {{ price }}$</p>-->
<!--            <div v-for="(i) in ingredients"-->
<!--                 :key="i"-->
<!--            >{{ i }}</div>-->
<!--            <div class="btn-add-to-cart">-->
<!--              <div v-if="$store.state.cart.productCart.length !==0 ">-->
<!--                <div v-if="notInArray(idProd)">-->
<!--                  <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)" >Add!</button>-->
<!--                </div>-->
<!--              </div>-->
<!--              <div v-else>-->
<!--                <button type="button" name="add-to-cart" :id="`add-prod${idProd}`" @click="addToCart(idProd)" >Add!</button>-->
<!--              </div>-->
<!--              <div v-for="j in $store.state.cart.productCart" :key="j.idProd">-->
<!--                <div v-if="j.idProd === idProd">-->
<!--                  &lt;!&ndash;                  {{test(j.idProd, idProd)}}&ndash;&gt;-->
<!--                  &lt;!&ndash;                  {{console.log(j.idProd)}}&ndash;&gt;-->
<!--                  <button type="button" name="plus-to-cart" :id="`plus-prod${idProd}`" @click="plusToCart(idProd)">+</button>-->
<!--                  <p>{{j.quantity}}</p>-->
<!--                  <button type="button" name="minus-from-cart" :id="`minus-prod${idProd}`" @click="minusFromCart(idProd)">-</button>-->
<!--                </div>-->

<!--              </div>-->
<!--            </div>-->
<!--          </div>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </div>-->
<!--</template>-->

<script>

export default {
  name: "ProductSpecificSupplier",

  props:{
    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,
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
  methods:{
    notInArray(idProd){
      for(let id in  this.listId){
        if(this.listId[id] === idProd){
          return false
        }
      }

      return true
    },
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
          return
        }
      }
    },
    replaceAndTitle(str){
      if (typeof str==="undefined"){
        return str
      }
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
}
.comfortaa-22{
  font-family: "Comfortaa",serif;
  font-size: 22px;
}
.product-light-list{
  display: flex;
  background-color: #FEE440;
  justify-content: space-around;
  padding-top: 40px;
  padding-bottom: 40px;
  width: 440px;
}

.product-light-elem{
  /*padding: 30px 150px 30px 0;*/
  /*border: 4px dotted #FF865E;*/
  border-radius: 35px;
  /*padding-left: 70px;*/
  /*padding-right: 70px;*/
  background-color: #feef7d;


}
.product-light-list .product-light-elem .product-container{
  display: flex;

  /*padding-left: 20px;*/
  /*padding-right: 250px;*/
  /*justify-content: end;*/
  flex-direction: column;
  min-width: 400px;

}
.product-light-list .product-light-elem .product-container .prod-info{
  min-width: 350px;
}

.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod{
  display: flex;
  flex-direction: column;
}

.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod h3{
  display: flex;
}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod .img-ingredients{
  display: flex;
  justify-content: center;

}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod .name{
  display: flex;
  justify-content: center;
}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod .name .prod-name{
  /*width: 300px;*/
  margin: 0;
  max-width: 350px;
}
.product-light-list  .product-light-elem .product-container .prod-info .img-name-prod img{
  cursor: pointer;
  max-width: 350px;
  max-height: 250px;
  display: block;
  /*margin-left: auto;*/
  /*margin-right: auto;*/
  /*float: left;*/
  /*padding-bottom: 15px;*/
  margin: 30px;
  justify-content: center;

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
  align-items: center;
  padding-bottom: 50px;
}
.name-price-cart .name-price .ing{
  color: #FF865E;
  padding-bottom: 15px;
}
.name-price-cart .name-price .price{
  padding: 20px 0;
}

.name-price-cart .name-price .add-plus-minus{
  display: flex;
  justify-content: center;
  padding-top: 15px;
}

.name-price-cart .name-price .category{
  display: flex;

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
  color: #2c3e50;
  width: 220px;
  height: 50px;

  border-width: 0;
  border-radius: 30px;
  /*border-color: black;*/
  float: right;
  cursor: pointer;
}
/*.name-price-cart .name-price .btn-add-to-cart button:hover button{*/
/*  background-color: #f28ec1;*/
/*}*/

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
  color: #2c3e50;
  width: 100px;
  height: 50px;

  border-width: 0;
  border-radius: 30px;
  /*border-color: black;*/
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


/*.product-light-list{*/
/*  display: block;*/
/*  background-color: aquamarine;*/
/*}*/

/*.product-light-elem{*/
/*  padding: 30px 150px 30px 0;*/

/*}*/
/*.product-light-elem .product{*/
/*  display: flex;*/

/*  padding-left: 20px;*/
/*  padding-right: 250px;*/
/*  flex-direction: row;*/
/*  justify-content: end;*/
/*}*/
/*.product-light-list .product .img-name-prod{*/
/*  flex-direction: column;*/
/*  padding-right: 50px;*/
/*}*/
/*.product-light-elem .product img{*/
/*  cursor: pointer;*/
/*  max-width: 300px;*/
/*  display: block;*/
/*  margin-left: auto;*/
/*  margin-right: auto;*/
/*  float: left;*/
/*}*/

/*.name-price-cart{*/
/*  padding-top: 40px;*/
/*  padding-left: 50px;*/

/*}*/
/*.name-price-cart .name-price{*/

/*}*/
/*.name-price-cart .name-price h3{*/

/*}*/
/*.name-price-cart .name-price h6{*/
/*  margin: 15px;*/

/*}*/
/*.name-price-cart .name-price h3,.name-price-cart .name-price p{*/
/*  margin: 0;*/
/*  padding-bottom: 30px;*/
/*}*/
/*.name-price-cart .name-price .btn-add-to-cart{*/
/*  padding-right: 10px;*/
/*}*/
/*.name-price-cart .name-price .btn-add-to-cart button{*/

/*  background-color: #cf1ad5;*/
/*  color: #1a1f1c;*/
/*  width: 50px;*/
/*  height: 30px;*/
/*  border-radius: 4px;*/
/*  float: right;*/

/*}*/


/*.corinthina-25{*/
/*  font-family: "Corinthia",serif;*/
/*  font-size: 25px;*/
/*}*/

</style>