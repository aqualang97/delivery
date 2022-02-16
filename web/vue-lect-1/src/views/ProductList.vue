<template>
  <div>

    <product v-for="(prod) in info" :key="prod.Name" :prod-name="prod.Name" :id-prod="prod.ID"></product>

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
      const response = await fetch("http://localhost:8085/all-products", {
        method: 'GET',
      });
      console.log(response)

      const json  = await response.json();
      for (let prod in json){
        console.log(json[prod].Name);

      }
      this.info=json

    }
    main()
  }

}
</script>

<style scoped>

</style>