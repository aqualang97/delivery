<template>

  <div class="product-light-list">
    <div class="product-light-elem">
      <div class="product">
        <div class="img-name-prod">
          <h3>{{ prodName }}</h3>
          <div
          >

          </div>
          <img :src="imgLink" :alt="prodName">
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
  name: "ProductSpecificSupplier",

  props:{
    idProd:Number,
    externalProdId:Number,
    prodName:String,
    price:Number,
    imgLink:String,
    type:String,

  },
}
</script>

<style scoped>

</style>