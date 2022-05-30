<template>
  <main>
    <div class="promo-container corinthina-30">
      <div class="promo-txt">
        <h1 class="fredoka-38">Promo for new customers!</h1>
        <p class="comfortaa-26">Gift certificate when buying a set from 500 UAH</p>
      </div>
    </div>
  </main>
</template>

<script>
export default {
  name: "HomePage"
}
</script>

<style scoped>

@font-face {
  font-family: "Comfortaa";
  src :url("../../../fonts/comfortaa/Comfortaa-VariableFont_wght.ttf");
  font-weight: 400;
}
@font-face {
  font-family: "Fredoka";
  src :url("../../../fonts/fredoka/FredokaOne-Regular.ttf");
  font-weight: 400;
}
.fredoka-38{
  font-family: "Fredoka",serif;
  font-weight: normal;
  font-size: 38px;
}
.comfortaa-26{
  font-family: "Comfortaa",serif;

  font-size: 26px;
}
 .promo-container{
   min-height: 1000px;
   background-color: #FEE440;
 }
 .promo-container .promo-txt{
   display: flex;
   flex-direction: column;
   justify-content: space-around;
 }

</style>