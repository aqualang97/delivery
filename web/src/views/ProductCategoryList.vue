<template>
  <div class="main-prod-list">
    <div class="prod-list-cont">
      <product-category
          :list-id="idList"
          v-for="(prod) in prodCat"
          :key="prod.ID"
          :prodName="prod.name"
          :id-prod="prod.ID"
          :external-prod-id="prod.id"
          :price="prod.price"
          :img-link="prod.image"
          :type="prod.type"
          :id-cat="prod.categoryNum"
          :supp-id="prod.supplierId"
          :external-supp-id="prod.externalSuppId"
          :ingredients="prod.ingredients"
      ></product-category>
    </div>
  </div>

</template>

<script>
export default {
  name: "ProductCategoryList",
  components: {},
  props:{},
  data() {
    return {
      info: null,
      prodCat: [],
      idList:[],
      isLogin:false,
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
    listOfId(){
      let lst = []
      if (this.$store.state.cart.productCart.length!==0){
        for (let j in this.$store.state.cart.productCart){
          lst.push(this.$store.state.cart.productCart[j].idProd)
        }

        return lst
      }else {
        return []
      }
    },
  },
  mounted() {
    //this.info = this.$route.params.cat_id
    this.checkLogin()
    let cat = this.$route.params.cat_id
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
    }
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {

        const response = await fetch(  `${this.$apiPrefix}/categories/${cat}`
            , {
              method: 'GET',
            });
        this.prodCat = await response.json();

      }
      main()
    }else {
      this.info = this.$store.state.productStore.posts
      for (let i in this.info) {
        if (this.info[i].categoryNum === Number(cat)){
          this.prodCat.push(this.info[i])
        }
      }
    }
}
}
</script>

<style scoped>
.main-prod-list{
  background-color: #FEE440;

}
.main-prod-list .prod-list-cont{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

</style>