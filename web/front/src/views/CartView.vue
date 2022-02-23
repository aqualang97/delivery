<template>
  <cart></cart>
</template>

<script>
export default {
  name: "CartView",
  components: {},
  methods:{
    plus(numProdCart){
      return numProdCart+=1
    },
    minus(numProdCart){
      return numProdCart-=1
    }
  }
}
</script>

<style scoped>

</style>