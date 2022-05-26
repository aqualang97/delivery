<template>
  <div>

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
      console.log(this.$store.state.auth.logged)
      this.isLogin = this.$store.state.auth.logged
      console.log("tytyty", this.isLogin)

      if(this.isLogin){
        console.log("tytyty", this.isLogin)
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
      console.log(response.body)

      const json  = await response.json();
      for (let supp in json){
        console.log(json[supp].name);

      }
      this.info=json

    }
    main()
  }
}
</script>

<style scoped>
.supplier-sort{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
/*.supplier-sort :hover .supplier-sort{*/
/*  height: auto;*/
/*  transition: max-height 0.25s ease-in;*/
/*}*/
</style>