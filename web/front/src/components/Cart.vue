<template>
  <main class="cart-container">
    <div class="your-cart">
      <h1 class="corinthina-40">Your cart</h1>
    </div>
    <div class="products-in-cart">
      <div class="prod corinthina-25">
<!--        <img src={{imgLink}}" alt="prod">-->
        <h3>Product Name</h3>
        <p>Price</p>
      </div>
      <div class="numbers">
        <div class="minus">
          <p>-</p>
        </div>
        <div class="num-of-products corinthina-40">
          <p></p>
        </div>
        <div class="plus">
          <p>+</p>
        </div>
      </div>
    </div>


    <div class="total">
      <div class="total-txt corinthina-40">
        <h2>Total payment:</h2>
        <h6>{{ total }} $</h6>
      </div>
      <div class="buy">
        <button type="button" name="button">Buy</button>
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
    quantity:Number,
    total:Number
  }
}
</script>

<style scoped>
.cart-container{
  background-color: #FADADD;
  padding-top: 50px;
  padding-bottom: 50px;
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

.numbers{
  display: flex;
  padding-left: 50px;
}
.num-of-products{
  display: flex;
  align-items: center;
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