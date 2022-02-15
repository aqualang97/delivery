<template xmlns:div="http://www.w3.org/1999/xlink">
  <div class="example">
    Example {{ $route.params.id }}
    <card :title="'Title'" :valueMax="5" :data="{a:1, b:2}" :click="true"></card>
<!--    <div :title="'Jump to: ' + $router.getRoutes()[0].name"-->
<!--         @click="$router.push({ name: 'Home' })">-->
<!--      {{'кликабельно'}}-->

<!--    </div>-->
  </div>
</template>

<script>
export default {
  name: "example",
  components: {}
}
</script>

<style scoped>
.example{
  color: #42b983;
}


</style>