<template>

  <div >

    <div>
      <div @click="$router.push(`/products&category=${catId}`)" class="main-cat-cont">
        <div v-if="$route.path===`/categories`" class="categories-container fredoka-38">
          {{replaceAndTitle(categoryName)}}
        </div>




        <div class="categories-sort" v-else>
          <ul>
            <li>
              {{replaceAndTitle(categoryName)}}
            </li>
          </ul>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
export default {
  name: "Categories",
  props:{
    categoryName:String,
    catId:Number,
    sourceIcon:String,
  },
  methods:{
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
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}
.fredoka-38{
   font-family: "Fredoka",serif;
   font-weight: normal;
   font-size: 38px;
 }
.main-cat-cont{
  padding-top: 20px;
  padding-bottom: 20px;
}
.main-cat-cont .categories-container{
  background-color: #ffbda4;
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 30px;
  margin: 20px;
  border-radius: 50px;

}

.categories-sort{
  margin: 0;
  display: flex;
}
.categories-sort ul{
  margin: 0;
}
.categories-sort li{
  cursor: pointer;
  margin: 5px;
  padding: 5px;

}
</style>