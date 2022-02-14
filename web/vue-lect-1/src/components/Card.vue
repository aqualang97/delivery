<template>
  <div class="example">
    Card: {{ data }} {{ title }} {{valueMax}}
    <div :title="'Jump to: ' + $router.getRoutes()[0].name"
         :class="{example: true, clickable: click}"
         :style="{'font-size': '25px'}"
         @click="$router.push({ name: 'Home' })"
    >
      {{'кликабельно'}}
    </div>

  </div>
</template>

<script>
export default {
  name: "Card",
  props:{
    title:{
      type: String,
      required: true,
    },
    valueMax:{
      type:Number,
      default:0,
      validator:(value) =>{
        return value >=0;
      },
    },
    data:{
      type: Object,
    },
    click:{
      type:Boolean,
    }
  },
};
</script>

<style scoped>
.example{
  color: #42b983;
}
.clickable{
  cursor: pointer;
  color: red;
}
</style>