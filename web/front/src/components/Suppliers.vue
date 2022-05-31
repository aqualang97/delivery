<template>
  <div  v-if="$route.path===`/suppliers`"  class="supplier-light-list">
    <div class="supplier-light-elem" >
      <div class="supplier-container">
        <div class="supplier-info">
          <div class="img-name-supplier">
            <div class="img-supplier comfortaa-22"
                 @click="$router.push(`/products&supplier=${id}`)">
              <img :src="image" :alt="supplierName" >
            </div>
            <div class="name">
              <h3 class="fredoka-38 supplier-name">{{ supplierName }}</h3>
            </div>
          </div>
        </div>
        <div class="category-hours-cont">
          <div class="category-hours fredoka-26">
            <div>
              <h3 class="comfortaa-22">Category: </h3>
            </div>
            <div class="comfortaa-24">
              <h3 class="color-category">{{ replaceAndTitle(nameOfType) }}</h3>
            </div>
            <p>Working hours: </p>
            <p class="fredoka-30"> {{ opening }} - {{closing}}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else @click="$router.push(`/products&supplier=${id}`)" class="supplier-sort-list">
    <div class="supplier-sort">
      <ul>
        <li>
          {{supplierName}}
        </li>
      </ul>
    </div>
  </div>
</template>

<!--<template>-->
<!--  <div class="supplier-cont">-->
<!--    <div v-if="$route.path===`/suppliers`"-->
<!--        class="supplier-wrap">-->
<!--      <div class="supplier">-->
<!--        <div class="img-name-supp">-->
<!--          <h3>{{ supplierName }}</h3>-->

<!--          <div :title="'Move to : ' + replaceAndTitle(supplierName)"-->
<!--               @click="$router.push(`/products&supplier=${id}`)"-->
<!--          >-->
<!--            <img :src="image" :alt="'supp img'">-->
<!--          </div>-->
<!--        </div>-->
<!--        <div class="type-hours-cont">-->
<!--          <div class="type-hours corinthina-25">-->
<!--            <h6>Category: {{ replaceAndTitle(nameOfType) }}</h6>-->
<!--            <p>Working hours: </p>-->
<!--            <p> {{ opening }} - {{closing}}</p>-->
<!--          </div>-->
<!--        </div>-->

<!--      </div>-->

<!--    </div>-->
<!--    <div v-else @click="$router.push(`/products&supplier=${id}`)" class="supplier-sort">-->
<!--      <ul>-->
<!--        <li>-->
<!--          {{supplierName}}-->
<!--        </li>-->
<!--      </ul>-->
<!--    </div>-->
<!--  </div>-->
<!--</template>-->

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
  font-size: 30px;

}
.fredoka-38{
   font-family: "Fredoka",serif;
   font-weight: normal;
   font-size: 38px;
 }

.comfortaa-24{
  font-family: "Comfortaa",serif;

  font-size: 24px;
}
.comfortaa-22{
  font-family: "Comfortaa",serif;
  font-size: 22px;
}
.supplier-light-list{
  display: flex;
  background-color: #FEE440;
  justify-content: space-around;
  padding-top: 40px;
  width: 440px;
}

.supplier-light-elem{
  /*padding: 30px 150px 30px 0;*/
  /*border: 4px dotted #FF865E;*/
  border-radius: 35px;
  /*padding-left: 70px;*/
  /*padding-right: 70px;*/
  background-color: #feef7d;


}
.supplier-light-list .supplier-light-elem .supplier-container{
  display: flex;

  /*padding-left: 20px;*/
  /*padding-right: 250px;*/
  /*justify-content: end;*/
  flex-direction: column;
  min-width: 400px;

}
.supplier-light-list .supplier-light-elem .supplier-container .supplier-info{
  min-width: 350px;
}

.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier{
  display: flex;
  flex-direction: column;
  min-height: 360px;
}

.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier h3{
  display: flex;
}
.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier .img-supplier{
  display: flex;
  justify-content: center;

}
.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier .name{
  display: flex;
  justify-content: center;
}
.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier .name .supplier-name{
  /*width: 300px;*/
  margin: 0;
  max-width: 350px;
}
.supplier-light-list  .supplier-light-elem .supplier-container .supplier-info .img-name-supplier img{
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

.supplier-light-list  .supplier-light-elem .supplier-container .category-hours-cont{
  display: flex;
  flex-direction: column;
}

.supplier-light-list  .supplier-light-elem .supplier-container .category-hours-cont .category-hours{
  display: flex;
  flex-direction: column;
  margin-left: 15px;
  margin-top: 20px;
  padding-bottom: 50px;
}
.supplier-light-list  .supplier-light-elem .supplier-container .category-hours-cont .category-hours .color-category{
  color: #FF865E;
}
.supplier-light-list  .supplier-light-elem .supplier-container .category-hours-cont .category-hours h3{
  margin: 0;
  padding-bottom: 20px;

}
.supplier-light-list  .supplier-light-elem .supplier-container .category-hours-cont .category-hours p{
  margin: 0;
  padding-top: 20px;
}
.supplier-sort-list{

}
.supplier-sort-list .supplier-sort{
  display: flex;
  justify-content: space-between;
}

.supplier-sort-list .supplier-sort ul{
  margin: 0;
}
.supplier-sort-list .supplier-sort li{
  cursor: pointer;
  margin: 5px;
  padding: 5px;
}
</style>