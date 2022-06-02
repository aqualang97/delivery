<template>
  <main class="product-light-list">
    <div class="product-light-elem">
      <div>
        <h1 class="fredoka-38">
          Order Number: {{order_id}}
        </h1>
        <h3 class="fredoka-34">
          Full price: {{full_price}} $
        </h3>
        <h3 class="comfortaa-26">
          Status: {{status}}
        </h3>
      </div>

      <div class="h-prod-cont">
        <div class="h-prod">
          <div class="prod" v-for="(p) in user_order"
               :key="p.ID">
            <div class="image">
              <h3 class="fredoka-38">{{ p.name }}</h3>
              <img :src="p.image" :alt="p.prodName">
            </div>
            <div class="name-price-cart">
              <div class="name-price corinthina-25">
                <h3 class="comfortaa-22">Category: </h3>
                <h3 class="comfortaa-26">{{ replaceAndTitle(p.category) }}</h3>
                <p class="fredoka-30">Quantity: {{p.quantity}}</p>
                <p class="fredoka-38">Price: {{ p.purchase_price }}$</p>
              </div>
            </div>
          </div>
        </div>

      </div>

    </div>

  </main>
</template>

<script>
export default {
  name: "OldOrdersUser",
  props:{
    user_id:Number,
    order_id:Number,
    status:String,
    full_price:null,
    user_order:null,
    prodName:String,
    category:String,
    image:String,
    quantity:Number,
    purchase_price:Number,
  },
  methods:{
    replaceAndTitle(str){
      str=str.replace("_", " ")
      str = str[0].toUpperCase() + str.substring(1)
      return str

    }
  },

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


.product-light-list{
  display: block;
  background-color: #FEE440;
}

.product-light-elem{
  border-top: #FF865E dotted 3px;
  padding: 30px 0 30px 0;
  display: flex;
  flex-direction: column;
  /* align-content: stretch; */
  /* justify-content: space-between; */
  align-items: center;


}

.product-light-elem .h-prod-cont{

}
.product-light-elem .h-prod-cont .h-prod{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.product-light-elem .h-prod-cont .h-prod .prod{
  display: flex;
  width: 650px;
  padding-left: 20px;
  margin:10px;
  flex-direction: row;
  justify-content: space-evenly;
  padding-bottom: 30px;
  border-radius: 35px;
  background-color: #feef7d;
}
.product-light-elem .h-prod-cont .h-prod .prod .image{
  display: flex;
  flex-direction: column;
}
.product-light-elem .h-prod-cont .h-prod .prod .image h3 {
  width: 350px;
}
.product-light-elem .h-prod-cont .h-prod img{
  cursor: pointer;
  max-width: 300px;
  display: block;
  margin-left: auto;
  margin-right: auto;
  float: left;
}

.name-price-cart{
  padding-top: 40px;
  margin-right: 20px;
}

.name-price-cart .name-price{

}
.name-price-cart .name-price h6{
  margin: 15px;

}
.name-price-cart .name-price h3,.name-price-cart .name-price p{
  margin: 0;
  padding-bottom: 15px;
}

</style>