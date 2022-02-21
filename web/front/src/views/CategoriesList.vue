<template>
  <div>
    <categories v-for="(category) in categoriesList"
                :key="category.name"
                :category-name="category.name"
                :cat-id="category.id"
    ></categories>
  </div>
</template>

<script>
export default {
  name: "CategoriesList",
  elem: 'category',
  data() {
    return {
      categoriesList: null
    };
  },
  mounted() {
    const  main = async () => {
      const response = await fetch("http://localhost:8080/categories", {
        method: 'GET',
      });
      const json  = await response.json();
      for (let category in json){
        console.log(json[category].name);
      }
      this.categoriesList=json
    }
    main()
  }
}
</script>

<style scoped>

</style>