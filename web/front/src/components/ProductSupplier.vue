<template>
  <div class="product-light-list">
    <div class="product-light-elem">
      <div class="product">
        <div class="img-name-prod">
          <h3>{{ prodName }}</h3>
          <div
               @click="$router.push(`/suppliers/${suppId}/products/${idProd}`)"
          >
            <img :src="imgLink" :alt="prodName">
          </div>
        </div>

        <div class="name-price-cart">
          <div class="name-price corinthina-25">
            <h6>Category: {{ type }}</h6>
            <p>Price: {{ price }}$</p>
            <div class="btn-add-to-cart">
              <button type="button" name="add-to-cart">Add!</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductSupplier",
  props:{

    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,

    suppId:Number,
    externalSuppId:Number,

    ingredients:[]
  },

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