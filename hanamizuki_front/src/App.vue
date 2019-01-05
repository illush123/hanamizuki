<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style lang="scss">
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding-bottom: 100px;
  padding-top: 50px;

  // 中央寄せ
  margin-right: auto;
  margin-left: auto;
  max-width: 1200px;
  padding-left: 10px;
  padding-right: 10px;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
