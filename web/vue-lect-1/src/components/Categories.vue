<template>
  <div class="categories-container">
    <div class="categories">
      {{categoryName}}
    </div>
  </div>
</template>

<script>
export default {
  name: "Categories",
  props:{
    categoryName:String,
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