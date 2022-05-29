<template>
  <div class="main-prod-list">
    <div class="prod-list-cont">
      <product-spec-supp
          :list-id="idList"
          :prodName="info.name"
          :idProd="info.ID"
          :externalProdId="info.id"
          :price="info.price"
          :imgLink="info.image"
          :type="info.type"
          :id-cat="info.categoryNum"
          :suppId="info.supplierId"
          :externalSuppId="info.externalSuppId"
          :ingredients="info.ingredients"
      >
      </product-spec-supp>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductSpecificSupplierView",
  components: {},
  data() {
    return {
      suppID:Number,
      prodID:Number,
      info: [],
      idList:[],
      isLogin:false,
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
    let supp = this.$route.params.supp_id
    let prod = this.$route.params.prod_id
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
    }
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {
        const response = await fetch(  `http://localhost:8080/suppliers/${supp}/products/${prod}`
            , {
              method: 'GET',
            });
        this.info = await response.json();
      }
      main()
    }else {
      this.info = this.$store.state.productStore.posts
      console.log(this.info)

      for (let i in this.info) {
         if (this.info[i].ID === Number(prod) && this.info[i].externalSuppId === Number(supp)){
            this.info = this.info[i]
           break
         }
          // if (this.info[i].ID === prod && this.info[i].externalSuppId ===supp){
          //   console.log(this.info[i])
          // }
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
  justify-content: center;
}

</style>