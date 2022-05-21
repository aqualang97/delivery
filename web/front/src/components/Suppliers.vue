<template>
  <div>
    <div v-if="$route.path===`/suppliers`"
        class="supplier">
      <div class="img-name-supp">
        <h3>{{ supplierName }}</h3>

        <div :title="'Move to : ' + replaceAndTitle(supplierName)"
             @click="$router.push(`/products&supplier=${id}`)"
        >
          <img :src="image" :alt="'supp img'">
        </div>
      </div>
      <div class="type-hours-cont">
        <div class="type-hours corinthina-25">
          <h6>Category: {{ replaceAndTitle(nameOfType) }}</h6>
          <p>Working hours: </p>
          <p> {{ opening }} - {{closing}}</p>
        </div>
      </div>
    </div>
    <div v-else @click="$router.push(`/products&supplier=${id}`)" class="supplier-sort">
      <ul>
        <li>
          {{supplierName}}
        </li>
      </ul>
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
    nameOfType:String,
    opening:String,
    closing:String,
    image:String,
    externalID:Number
  },
  methods:{
    replaceAndTitle(str){
      str=str.replace("_", " ")
      str = str[0].toUpperCase() + str.substring(1)
      return str
    },
  },

}
</script>

<style scoped>
.supplier-sort{
  display: flex;
  justify-content: space-between;
}


.supplier-sort ul{
  margin: 0;
}
.supplier-sort li{
  cursor: pointer;
  margin: 5px;
  padding: 5px;

}
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
  cursor: pointer;
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