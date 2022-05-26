<template>
  <div>
    <categories
        v-for="(category) in categoriesList"
        :key="category.name"
        :category-name="category.name"
        :cat-id="category.id"
    >
    </categories>
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
      console.log(this.$store.state.auth.logged)
      this.isLogin = this.$store.state.auth.logged
      console.log("tytyty", this.isLogin)

      if(this.isLogin){
        console.log("tytyty", this.isLogin)
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

</style>