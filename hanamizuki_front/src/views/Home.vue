<template>
  <div class="home">
    <Header></Header>
    <b-container>
      <b-row>
        <b-col cols="4">
          <QuestionSpace></QuestionSpace>
        </b-col>
        <b-col cols="8">
          <EditorSpace></EditorSpace>
        </b-col>
      </b-row>
    </b-container>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Header from '@/components/Header.vue';
import EditorSpace from '@/components/EditorSpace.vue';
import QuestionSpace from '@/components/QuestionSpace.vue';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component({
  components: {
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