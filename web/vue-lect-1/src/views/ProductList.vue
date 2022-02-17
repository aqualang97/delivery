<template>
  <div>
      <product v-for="(prod) in info"
             :key="prod.name"
             :prod-name="prod.name"
             :id-prod="prod.ID"
             :external-prod-id="prod.externalID"
             :price="prod.price"
             :img-link="prod.image"
             :type="prod.type"

    ></product>

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
  mounted() {
    const  main = async () => {
      const response = await fetch("http://localhost:8080/all-products", {
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