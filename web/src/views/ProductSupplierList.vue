<template>
  <div class="main-prod-list">
    <div class="prod-list-cont">
      <product-supplier
          :list-id="idList"
          v-for="(prod) in info"
          :key="prod.name"
          :prodName="prod.name"
          :idProd="prod.ID"
          :externalProdId="prod.id"
          :price="prod.price"
          :imgLink="prod.image"
          :type="prod.type"
          :id-cat="prod.categoryNum"
          :suppId="prod.supplierId"
          :ingredients="prod.ingredients">
        :externalSuppId="prod.externalSuppId">
      </product-supplier>
    </div>
  </div>

</template>

<script>
export default {
  name: "ProductSpecificSupplier",
  components: {},

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
    this.checkLogin()
    //this.info = this.$route.params.cat_id
    let supp = this.$route.params.supp_id
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
    }
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {

        const response = await fetch(  `${this.$apiPrefix}/suppliers/${supp}/products`
            , {
              method: 'GET',
            });
        this.info = await response.json();


      }
      main()
    }else {
      for (let i in this.$store.state.productStore.posts) {
        if (this.$store.state.productStore.posts[i].supplierId === Number(supp)){
          this.prodCat.push(this.$store.state.productStore.posts[i])
        }
      }
      // for (let i of this.$store.state.productStore.posts) {
      //
      //
      // }
      this.info = this.prodCat
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