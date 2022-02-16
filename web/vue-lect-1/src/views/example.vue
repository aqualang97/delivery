<template xmlns:div="http://www.w3.org/1999/xlink">
  <div>
    <div class="example">
      Example1 {{ $route.params.id }}
      <card :title="'Title'" :valueMax="5" :data="{a:1, b:2}" :click="true" @calc="increment"></card>
    </div>
    <div>
      {{val}}
    </div>
  </div>
</template>

<script>
export default {
  name: "example1",
  components: {
  },
  data(){
    return{
      val:0
    }
  },
  methods:{
    increment(){
      this.val+=1
    }

  }
}
</script>

<style scoped>
.example{
  color: #42b983;
}


</style>