<template>
  <div>
    <product-supplier
          v-for="(prod) in info"
          :key="prod.name"
          :prod-name="prod.name"
          :id-prod="prod.ID"
          :external-prod-id="prod.id"
          :price="prod.price"
          :img-link="prod.image"
          :type="prod.type"
          :supp-id="prod.supplierId"
          :external-supp-id="prod.externalSuppId">
    </product-supplier>

  </div>

</template>

<script>
export default {
  name: "ProductSpecificSupplier",
  components: {},

  data() {
    return {
      info: null,
      prodCat: []
    };
  },
  mounted() {

    //this.info = this.$route.params.cat_id
    let supp = this.$route.params.supp_id
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {

        const response = await fetch(  `http://localhost:8080/suppliers/${supp}/products`
            , {
              method: 'GET',
            });
        this.info = await response.json();
        console.log(this.info)
        console.log(this.info)
        console.log(this.info)
        console.log(this.info)
        for (let prod in this.info){
          console.log(this.info[prod].name)
        }
      }
      main()
    }else {
      console.log("else")
      console.log(supp)
      console.log(this.$store.state.productStore.posts)

      for (let i in this.$store.state.productStore.posts) {
        console.log(this.$store.state.productStore.posts[i].supplierId)
        if (this.$store.state.productStore.posts[i].supplierId == supp){
          console.log(this.$store.state.productStore.posts[i])
          this.prodCat.push(this.$store.state.productStore.posts[i])

        }
      }
      // for (let i of this.$store.state.productStore.posts) {
      //   console.log(i)
      //
      //
      // }
      this.info = this.prodCat

      console.log(this.prodCat, 123)
      // console.log(this.prodCat, 123)


    }

  }
}
</script>

<style scoped>

</style>