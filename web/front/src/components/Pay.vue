<template>
  <main class="method-cont">
    <div>
      <p class="fredoka-26">
        {{msg}}
      </p>
    </div>
    <div>
      <div v-if="msg==='Card'">
        <img src="../../../pic/card-method.png">
      </div>
      <div v-else-if="msg==='Cash'">
        <img src="../../../pic/cash-method.png">
      </div>
      <div v-else-if="msg==='Status: completed'">
        <img src="../../../pic/completed.png">
      </div>
    </div>


  </main>
</template>

<script>
export default {
  name: "Pay",
  props:{
    msg:String(),
  },
  data(){
    return{

    }
  },
  methods:{

  },

}
</script>


<style scoped>
@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}

.fredoka-26{
  font-family: "Fredoka",serif;
  font-size: 26px;
}

.method-cont{
  background-color: #c6e1ff;
  padding-top: 50px;
  padding-bottom: 50px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
.method-cont img{
   opacity: 0.3;
 }
</style>