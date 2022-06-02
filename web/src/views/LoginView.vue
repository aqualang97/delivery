<template>
  <div>
    <login
    @send:="takeLogin"></login>
  </div>
</template>

<script>
export default {
  name: "LoginView",
  data(){
    return {email:"", password:""}
  },
  methods:{
    takeLogin(){
    }
  }
}
</script>

<style scoped>

</style>