<template>
  <div>
    <product-spec-supp

        :prod-name="info.name"
        :id-prod="info.ID"
        :external-prod-id="info.ExternalID"
        :price="123"
        :img-link="`asd`"
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
    const  main = async () => {
      console.log(123)
      console.log(123)
      console.log(123)
      console.log(123)
      let supp = this.$route.params.supp_id
      let prod = this.$route.params.prod_id

      const response = await fetch(  `http://localhost:8080/suppliers/${supp}/products/${prod}`
          , {
        method: 'GET',
      });
      console.log(response)

      const json  = await response.json();
      for (let prod in json){
        console.log(json[prod].name);

      }
      this.info=json

    }
    main()

}
}
</script>

<style scoped>

</style>