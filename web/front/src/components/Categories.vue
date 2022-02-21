<template>
  <div class="categories-container">
    <div class="categories">
      <div :title="'Ingredients: ' + ingredients"
           @click="$router.push(`/products&category=${catId}`)"
      >
        {{categoryName}}
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

  }
}
</script>

<style scoped>
.categories-container{
  background-color: #f6c101;
}
.categories-container .categories{
  font-size: 25px;
}
</style>