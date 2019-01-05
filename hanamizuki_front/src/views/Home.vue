<template>
  <div class="home">
    <Header></Header>
    <div class="container">
      <div class="row">
        <div class="col-4">
          <BinarySpace></BinarySpace>
        </div>
        <div class="col-8">
          <EditorSpace></EditorSpace>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
// import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import Header from '@/components/Header.vue';
import EditorSpace from '@/components/EditorSpace.vue';
import QuestionSpace from '@/components/QuestionSpace.vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component({
  components: {
    // HelloWorld,
    Header,
    QuestionSpace,
    EditorSpace,
  },
})
export default class Home extends Vue {}
</script>

<style scoped lang="scss">

  .container {
    padding-top: 34px;
  }

</style>