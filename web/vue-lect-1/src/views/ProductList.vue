<template>
  <div>
    <div>
      <product
          v-for="(prod) in $store.state.productStore.posts"
          :key="prod.name"
          :prod-name="prod.name"
          :id-prod="prod.ID"
          :external-prod-id="prod.externalID"
          :price="prod.price"
          :img-link="prod.image"
          :type="prod.type"
          :supp-id="prod.supplierId"
          :external-supp-id="prod.externalSuppId"
      ></product>
    </div>
<!--    <div v-else >-->
<!--      <div>-->

<!--      </div>-->
<!--      {{$store.state.productStore.posts.length}}-->
<!--    </div>-->
  </div>
</template>

<script>

export default {
  name: "ProductList",
  elem: 'prod',

  data() {
    return {
      info: null
    };
  },
  methods:{
    getProdList(){
      this.$store.dispatch('productStore/fetch')
      console.log(this.$store.state.productStore.posts.length)
    }

  },
  mounted() {
    if (this.$store.state.productStore.posts.length === 0){
      this.getProdList()
    }

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

</style>