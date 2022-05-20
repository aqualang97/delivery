<template>
  <div>
    <product-spec-supp

        :prodName="info.name"
        :idProd="info.ID"
        :externalProdId="info.id"
        :price="info.price"
        :imgLink="info.image"
        :type="info.type"
        :suppId="info.supplierId"
        :externalSuppId="info.externalSuppId"
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
      suppID:Number,
      prodID:Number,
      info: null
    };
  },
  methods:{

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

</style>