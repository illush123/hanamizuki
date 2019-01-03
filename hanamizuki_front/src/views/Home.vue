<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <Header></Header>
    <button type="button" class="btn btn-dark">Dark</button>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
// import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import Header from '@/components/Header.vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component({
  components: {
    // HelloWorld,
    Header,
  },
})
export default class Home extends Vue {}
</script>
