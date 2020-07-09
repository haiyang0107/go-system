<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>
<style>
  *{
      -webkit-box-sizing: border-box;
      -moz-box-sizing: border-box;
      box-sizing: border-box;
      margin:0px;
      padding:0px;
  }
  html,body {
    width: 100%;
    height: 100%;
  }
  ol,ul {
    list-style: none;
  }
  #app {
    width: 100%;
    height: 100%;
  }
</style>
