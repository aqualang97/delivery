<template>
  <div>

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
        :supp-id="prod.supplierId"
        :external-supp-id="prod.externalSuppId"
        :ingredients="prod.ingredients"
    ></product-category>
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

    };
  },
  methods:{
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

    let cat = this.$route.params.cat_id
    if(this.$store.state.cart.productCart.length !==0){
      this.idList = this.listOfId()
    }
    console.log(this.$store.state.productStore.posts.length)
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {

        const response = await fetch(  `http://localhost:8080/categories/${cat}`
            , {
              method: 'GET',
            });
        this.prodCat = await response.json();
        console.log(this.prodCat)
        for (let prod in this.prodCat){
          console.log(this.prodCat[prod].name)
        }
      }
      main()
    }else {


      console.log("else")
      this.info = this.$store.state.productStore.posts
      console.log(this.info)

      for (let i in this.info) {
        console.log(typeof cat)
        if (this.info[i].categoryNum === Number(cat)){
          this.prodCat.push(this.info[i])
          console.log(this.prodCat)
        }
      }
      // console.log(this.prodCat, 123)
    }
}
}
</script>

<style scoped>

</style>