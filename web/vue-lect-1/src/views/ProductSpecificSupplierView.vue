<template>
  <div>
    <product-spec-supp

        :prod-name="info.name"
        :id-prod="info.ID"
        :external-prod-id="info.id"
        :price="info.price"
        :img-link="info.image"
        :type="info.type"
        :supp-id="info.supplierId"
        :external-supp-id="info.externalSuppId"
        :ingredients="info.ingredients"
    >
    </product-spec-supp>
  </div>
</template>

<script>
export default {
  name: "ProductSpecificSupplierView",
  components: {},
  data() {
    return {
      // suppID:Number,
      // prodID:Number,
      info: null
    };
  },

  mounted() {
    let supp = this.$route.params.supp_id
    let prod = this.$route.params.prod_id
    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {

        const response = await fetch(  `http://localhost:8080/suppliers/${supp}/products/${prod}`
            , {
              method: 'GET',
            });
        this.info = await response.json();

        this.info.forEach((item)=>{
          console.log(item)
        })
      }
      main()
    }else {
      this.info = this.$store.state.productStore.posts
      console.log(this.info)

      for (let i in this.info) {
         if (this.info[i].ID == prod && this.info[i].externalSuppId == supp){
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

</style>