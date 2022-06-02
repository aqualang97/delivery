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
      <div class="comfortaa-26">
        <p v-if="showCat===false" class="categories" @click="showCat=true">Categories</p>
        <p v-else @click="showCat=false" class="categories">Categories</p>
        <categories-list v-if="showCat" class="categories-cont"/>
      </div>
<!--      <div class="supplier">-->
      <div class="comfortaa-26">
        <p v-if="showSupp===false" @click="showSupp = true" class="supplier">Suppliers</p>
        <p v-else @click="showSupp = false" class="supplier">Suppliers</p>
        <suppliers-list v-if="showSupp" class="supplier-cont"/>
      </div>
    </div>
    <div class="main-prod-list">
      <div class="prod-list-cont">
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
            :id-cat="prod.categoryNum"
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
      isLogin:false

    };
  },
  methods:{

    async getProdList(){
      await this.$store.dispatch('productStore/fetch')

    },
    async checkLogin(){
      await this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged

      if(this.isLogin){
        await this.$store.dispatch('auth/refresh')
      }
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
    this.checkLogin()
    if (this.$store.state.productStore.posts.length === 0){
      this.getProdList()
    }
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
    }
  }
}
</script>

<style scoped>
@font-face {
  font-family: "Fredoka";
  src :url("../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}
@font-face {
  font-family: "Comfortaa";
  src :url("../../fonts/comfortaa/Comfortaa-VariableFont_wght.ttf");
  font-weight: 400;
}
.comfortaa-26{
  font-family: "Comfortaa",serif;
  font-weight: bold;
  font-size: 26px;
}
.sorted{
  display: flex;
  justify-content: space-between;

  background-color: #A2D2FF;
  text-align: center;
  padding: 20px 70px;
}
.sorted .categories{
  border-radius: 30px;
  margin: 0;
  background-color: #c6e1ff;
  width: 400px;
  height: 40px;
  padding-top: 15px;
}
.sorted .supplier{
  border-radius: 30px;
  margin: 0;
  background-color: #c6e1ff;
  width: 400px;
  height: 40px;
  padding-top: 15px;
}
.sorted .categories .categories-cont{

}
.sorted .supplier .supplier-cont{

}
.main-prod-list{
  background-color: #FEE440;

}
.main-prod-list .prod-list-cont{
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}


</style>