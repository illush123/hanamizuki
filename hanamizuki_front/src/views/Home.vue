<template>
  <div class="home">
    <Header messageProp="this is propmessage"></Header>
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
