<template>
  <div v-if="$route.path===`/categories`" class="main-cat-list">
    <div class="cat-cont">
      <categories
          v-for="(category) in categoriesList"
          :key="category.name"
          :category-name="category.name"
          :cat-id="category.id"
          :source-icon="`../../../pic/categories/`+ category.id + `.png`"
      >
      </categories>
    </div>
  </div>
  <div v-else class="main-cat-sorted">
    <div>
      <categories class="cat-sort"
          v-for="(category) in categoriesList"
          :key="category.name"
          :category-name="category.name"
          :cat-id="category.id"
      >
      </categories>
    </div>

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
  methods:{
    checkLogin(){
      this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged

      if(this.isLogin){
        this.$store.dispatch('auth/refresh')
        // this.$router.push("/sign-in")
      }
    },

  },
  mounted() {
    this.checkLogin()
    const  main = async () => {
      const response = await fetch("http://localhost:8080/categories", {
        method: 'GET',
      });
      const json  = await response.json();
      this.categoriesList=json
      return json
    }
    main()
  }
}
</script>

<style scoped>
.main-cat-list{
  background-color: #FF865E;
}
.main-cat-list .cat-cont{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
.main-cat-sorted{
  background-color: #A2D2FF;

}
.main-cat-sorted .cat-sort{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

</style>