<template>
      <div class="supplier">
        <div class="img-name-supp">
          <h3>{{ supplierName }}</h3>
          <img :src="image" :alt="'supp img'">
        </div>
        <div class="type-hours-cont">
          <div class="type-hours corinthina-25">
            <h6>Category: {{ typeOfSupp }}</h6>
            <p>Working hours: </p>
            <p> {{ opening }} - {{closing}}</p>
          </div>
        </div>
      </div>
</template>

<script>
export default {
  name: "Suppliers",
  props:{
    id:Number,
    supplierName:String,
    typeOfSupp:Number, //потом изменю на стринг, когда в беке добавлю
    opening:String,
    closing:String,
    image:String,
    externalID:Number
  }

}
</script>

<style scoped>
.supplier{
  display: block;
  background-color: aquamarine;
  padding: 30px 150px 30px 0;
  display: flex;

  //padding-left: 20px;
  //padding-right: 250px;
  flex-direction: row;
  justify-content: center;
}
.supplier .img-name-supp{
  flex-direction: column;
  padding-right: 50px;
}

.supplier .img-name-supp img{
  max-width: 300px;
  display: block;
  margin-left: auto;
  margin-right: auto;
  float: left;

}
.type-hours-cont{
  padding-top: 40px;
  padding-left: 50px;
}
.type-hours-cont .type-hours{

}

</style>