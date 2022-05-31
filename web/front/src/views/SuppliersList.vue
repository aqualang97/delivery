<template>
  <div>
    <div v-if="$route.path===`/suppliers`" class="main-prod-list">
      <div class="prod-list-cont">
        <suppliers class="supplier-sort"
                   v-for="(supp) in info"
                   :key="supp.name"
                   :id="supp.ID"
                   :supplier-name="supp.name"
                   :type-of-supp="supp.type"
                   :opening="supp.workingHours.opening"
                   :closing="supp.workingHours.closing"
                   :external-i-d="supp.externalID"
                   :image="supp.image"
                   :name-of-type="supp.categoryName"
        ></suppliers>
      </div>
    </div>
    <div v-else class="main-supp-sorted">
      <div>
        <suppliers class="supplier-sort"
                   v-for="(supp) in info"
                   :key="supp.name"
                   :id="supp.ID"
                   :supplier-name="supp.name"

        ></suppliers>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: "SuppliersList",
  elem:'supp',
  data() {
    return {
      info: null,
      isLogin:false,
    };
  },
  methods:{
    checkLogin(){
      this.$store.dispatch('auth/isLogin')
      this.isLogin = this.$store.state.auth.logged
      if(this.isLogin){
        this.$store.dispatch('auth/refresh')
        // this.$router.push("/sign-in")
      }
    },
  },
  mounted() {
    this.checkLogin()
    const  main = async () => {
      const response = await fetch("http://localhost:8080/suppliers", {
        method: 'GET',
      });
      const json  = await response.json();

      this.info=json
    }
    main()
  }
}
</script>

<style scoped>
.main-supp-sorted{
  background-color: #A2D2FF;

}
.main-supp-sorted .supplier-sort{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.main-prod-list{
  background-color: #FEE440;

}
.main-prod-list .prod-list-cont{
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

/*.supplier-sort :hover .supplier-sort{*/
/*  height: auto;*/
/*  transition: max-height 0.25s ease-in;*/
/*}*/
</style>