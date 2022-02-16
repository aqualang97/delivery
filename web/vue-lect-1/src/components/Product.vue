<template>
  <div>
    {{prodName}}
    {{idProd}}
  </div>
</template>

<script>
export default {
  name: "Product",
  props:{
    idProd:Number,
    prodName:String,
  }
}
</script>

<style scoped>

</style>