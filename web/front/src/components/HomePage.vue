<template>
  <main>
    <div class="promo-container corinthina-30">
      <div class="promo-txt">
        <h1>Promo for new customers!</h1>
        <p>Gift certificate when buying a set from 500 UAH</p>
      </div>
    </div>
    <div class="about-container">
      <div class="about-txt corinthina-25">
        <h3>Welcome to the CraftShack family. We’ve been waiting for you.</h3>
        <p> You’re about to join hundreds of thousands of craft beer, wine, and spirit enthusiasts from around the globe; people who have a genuine passion for all things booze. </p>
        <p> Here in our family, we do things a little differently. </p>
        <p> Once you order from us, you’re in the family for life. Seriously. And when family members have problems, we solve them. Something wrong with your order? Did a new brew not live up to your expectations? Call us. Email us. Bang on our door. A real person will respond immediately and make things right. At CraftShack, family matters, and we haven't spent a decade building relationships for this to change anytime soon.</p>
      </div>
    </div>
    <div class="about-container reverce">
      <div class="about-txt corinthina-25">
        <h3>Welcome to the CraftShack family. We’ve been waiting for you.</h3>
        <p> You’re about to join hundreds of thousands of craft beer, wine, and spirit enthusiasts from around the globe; people who have a genuine passion for all things booze. </p>
        <p> Here in our family, we do things a little differently. </p>
        <p> Once you order from us, you’re in the family for life. Seriously. And when family members have problems, we solve them. Something wrong with your order? Did a new brew not live up to your expectations? Call us. Email us. Bang on our door. A real person will respond immediately and make things right. At CraftShack, family matters, and we haven't spent a decade building relationships for this to change anytime soon.</p>
      </div>
    </div>
    <div class="learn-more-container">
      <div class="learn-more corinthina-25">

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

.logo-about-prod-container{
  display: flex;
  width: 30%;
  justify-content: inherit;
  padding-top: 20px;
  padding-left: 20px;
}

.logo-about-prod-container img{
  width: 160%;
  margin: 0;
}

.login-cart-container{
  display: flex;
  width: 30%;
  justify-content: inherit;
  padding-top: 20px;
  padding-right: 20px;
}


.logo-about-prod-container a, .login-cart-container a {
  margin: 0;
  text-decoration: none;
  color: black;

}

.corinthina-30{
  font-family: "Corinthia";
  font-size: 30px;
}
.corinthina-40{
  font-family: "Corinthia";
  font-size: 40px;
}
.corinthina-25{
  font-family: "Corinthia";
  font-size: 25px;
}

.corinthina-15{
  font-family: "Corinthia";
  font-size: 15px;
}


.promo-container, .about-container{
  display: flex;
  width: 100%;
  background-color: #FADADD;
  margin: 0;
  padding-top: 50px;
}
.promo-container .promo-txt, .about-container .about-txt{
  width: 100%;
  padding: 50px 0 0 25px;

}
.promo-container .promo-img img, .about-container .about-img img{
  width: 100%;
  position: static;
  margin: 0;
}

.about-txt h3, .about-txt p{
  padding: 0 50px 0 50px;
}

.reverce{
  flex-direction: row-reverse;
}

.learn-more-container{
  display: flex;
  flex-direction: column;
  padding: 30px 300px;
  background-color: #FADADD;

}
.learn-more{
  justify-content: center;
  text-align: center;
}
.different-beers-btn {
  display: flex;
  justify-content: space-between;
}
.different-beers-btn button{
  font-family: "Corinthia";
  font-size: 25px;
  color: white;
  border-radius: 8px;
  height: 50px;
  width: 15%;
}

.light{
  background-color: #f6c101;
}

.dark{
  background-color: #290e05;
}

.unfiltered{
  background-color: #f28ec1;
}

.craft{
  background-color: #b84b3b;
}

footer{
  background-color: #c9ffb6;
  color: black;
  display: flex;
  flex-direction: row;
  padding-top: 25px;
}
.sotial-networks{
  display: flex;
  flex-direction: column;
  text-align: center;
  width: 50%;
}
.networks-img{
  display: flex;
  justify-content: space-evenly;
}
.copyrights{
  display: flex;
  align-items: self-end;
  width: 100%;
  justify-content: end;
  padding-right: 25px;
}

</style>