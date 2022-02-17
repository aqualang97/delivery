<template>
  <div class="example">
    Card: {{ data }} {{ title }} {{valueMax }} <br>
    <div :title="'Jump to: ' + $router.getRoutes()[0].name"
         :class="{example: value<5, clickable: click}"
         :style="{'font-size': '25px'}"
         @click="$router.push({ name: 'Home' })"

    >
      {{'кликабельно'}}
    </div>
    <div  :title="'Jump to: ' + $router.getRoutes()[1].name"
          :class="{example: value<5, clickabl2: click}"
          :style="{'font-size': '25px'}"
          @click="$router.push({ name: 'About' })"
    >
      {{ name }} {{ data2 }} {{value}}

    </div>
    <div> <button @click="n1">button</button></div>
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
  data(){
    return{
      name: "ExampleCard",
      data2:{a:1,b:2, c:3},
      clickable:false,
      value:0
    }
  },
  methods:{
    n1(){
      this.$emit(
          'calc'
      )
    }
  }

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
.clickabl2{
  cursor: progress;
  color: #ff00d5;
}
</style>