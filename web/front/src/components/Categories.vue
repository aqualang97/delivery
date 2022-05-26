<template>

  <div >

    <div>
      <div
           @click="$router.push(`/products&category=${catId}`)"
      >
        <div v-if="$route.path===`/categories`"
             class="categories-container categories category-name"
        >{{replaceAndTitle(categoryName)}}</div>
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
.categories-container{
  background-color: #f6c101;
  display: flex;
  justify-content: center;
}
.categories-container .categories{
  font-size: 25px;
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