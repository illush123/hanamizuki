<template>
  <div class="home">
    <Header></Header>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Header from '@/components/Header.vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component({
  components: {
    Header,
  },
})
export default class Home extends Vue {}
</script>
