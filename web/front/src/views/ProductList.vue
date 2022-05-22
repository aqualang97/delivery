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
      <div>
        <p v-if="showCat===false" class="categories" @click="showCat=true">Categories</p>
        <p v-else @click="showCat=false" class="categories">Categories</p>
        <categories-list v-if="showCat" class="categories-cont"/>
      </div>
<!--      <div class="supplier">-->
      <div>
        <p v-if="showSupp===false" @click="showSupp = true" class="supplier">Suppliers</p>
        <p v-else @click="showSupp = false" class="supplier">Suppliers</p>
        <suppliers-list v-if="showSupp" class="supplier-cont"/>
      </div>
    </div>
    <div>
      <div>
        <product
            :list-id="idList"
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
      idList:[],
      showSupp:false,
      showCat:false,

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
    if (this.$store.state.productStore.posts.length === 0){
      this.getProdList()
    }
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
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
  background-color: #A2D2FF;
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
.sorted .categories .categories-cont{

}
.sorted .supplier .supplier-cont{

}


</style>