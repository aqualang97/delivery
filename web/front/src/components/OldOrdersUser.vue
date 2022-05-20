<template>
  <main class="product-light-list">
    <div class="product-light-elem">
      <h1>
        Order Number:{{order_id}}
      </h1>
      <h3>
        Full price: {{full_price}} $
      </h3>
      <h3>
        Status: {{status}}
      </h3>
      <div class="h-prod-cont" v-for="(p) in user_order"
           :key="p.ID">
        <div>
          <h3>{{ p.name }}</h3>

          <img :src="p.image" :alt="p.prodName">
        </div>
        <div class="name-price-cart">
          <div class="name-price corinthina-25">
            <h3>Category: {{ replaceAndTitle(p.category) }}</h3>
            <p>Quantity: {{p.quantity}}</p>
            <p>Price: {{ p.purchase_price }}$</p>
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
.product-light-list{
  display: block;
  background-color: aquamarine;
}

.product-light-elem{
  border-top: #1a1f1c solid 2px;
  padding: 30px 150px 30px 0;

}
.product-light-elem .h-prod-cont{
  display: flex;

  padding-left: 20px;
  padding-right: 250px;
  flex-direction: row;
  justify-content: flex-end;
  padding-bottom: 30px;
}

.product-light-elem .h-prod-cont img{
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
</style>