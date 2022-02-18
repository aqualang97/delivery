<template>
  <div>
    <product-spec-supp

        :prod-name="info.name"
        :id-prod="info.ID"
        :external-prod-id="info.ExternalID"
        :price="123"
        :img-link="info.name"
        :type="info.Category"
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

    if (this.$store.state.productStore.posts.length === 0){
      const  main = async () => {
        let supp = this.$route.params.supp_id
        let prod = this.$route.params.prod_id
        const response = await fetch(  `http://localhost:8080/suppliers/${supp}/products/${prod}`
            , {
              method: 'GET',
            });
        this.info = await response.json();
      }
      main()
    }else {
      this.info = this.$store.state.productStore.posts
    }

}
}
</script>

<style scoped>

</style>