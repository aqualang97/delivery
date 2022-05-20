<template>
  <div>
<!--    <select>-->
<!--      <option value="categoriesSelected" selected>-->
<!--        choose...-->
<!--      </option>-->
<!--      <option value="categories">-->
<!--        <categories-list/>-->
<!--      </option>-->
<!--    </select>-->
    <div class="sorted">
      <div class="categories">
        <p>Categories</p>
        <categories-list/>
      </div>
      <div class="supplier">
        <p>Suppliers</p>
        <suppliers-list/>
      </div>
    </div>
    <div>
      <div>
        <product
            v-for="(prod) in $store.state.productStore.posts"
            :key="prod.ID"
            :prod-name="prod.name"
            :id-prod="prod.ID"
            :external-prod-id="prod.id"
            :price="prod.price"
            :img-link="prod.image"
            :type="prod.type"
            :supp-id="prod.supplierId"
            :external-supp-id="prod.externalSuppId"
            :ingredients="prod.ingredients">
        </product>
      </div>
    </div>
  </div>
</template>


<script>
import categoriesList from "./CategoriesList";
import suppliersList from "./SuppliersList";
export default {
  name: "ProductList",
  elem: 'prod',
  components:{
    categoriesList,
    suppliersList,
  },
  props:{
  },
  data() {
    return {
      info: null,
      numProdCart:0,
      user_id:Number,
      access_token:"",
      refresh_token:"",
      type:String(),
      categories:null,
    };
  },
  methods:{

    async getProdList(){
      await this.$store.dispatch('productStore/fetch')
      console.log(this.$store.state.productStore.posts.length)
    },
    refreshToken(token){
      return token
    },

    plus(numProdCart){
      return numProdCart+=1
    },
    minus(numProdCart){
      return numProdCart-=1
    },


  },
  mounted() {

    if (this.$store.state.productStore.posts.length === 0){
      this.getProdList()
    }
    // if (this.$store.state.cart.productCart.length !==0){
    //   for(let p in this.$store.state.cart.productCart.length){
    //
    //   }
    // }

  }

  // mounted() {
  //   const  main = async () => {
  //     const response = await fetch("http://localhost:8080/all-products", {
  //       method: 'GET',
  //     });
  //     console.log(response)
  //
  //     const json  = await response.json();
  //     for (let prod in json){
  //       console.log(json[prod].name);
  //
  //     }
  //     this.info=json
  //
  //   }
  //   main()
  // }

}
</script>

<style scoped>
.sorted{
  display: flex;
  justify-content: space-around;
  padding-bottom: 25px;
}
.sorted .categories{
  border: #f28ec1 1px dashed;
  border-radius: 9px;
  margin: 0;
}
.sorted .supplier{
  border: #f28ec1 1px dashed;
  border-radius: 9px;
  margin: 0;
}
</style>