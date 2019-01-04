<template>
  <header>
    <div class="container">
      <div class="header-image-wrapper pr-4">
        <img src="/path/to/image" class="header-image_src">
      </div>
      <div class="header-text-wrapper">
        <div class="header-text-lesson-wrapper pr-2">
          <div class="header-text-lesson ml-5">レッスン一覧</div>
        </div>
        <div class="header-text-lesson_name pt-1 pl-3">{{lessonTitle}}</div>
      </div>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component
export default class Header extends Vue {

  // data
  lessonTitle = 'これはダミータイトルです';
}
</script>

<style scoped lang="scss">

  header {

    .header-image-wrapper {
      float: left;
      padding-right: 20px;
    }

    .header-text-wrapper {
      width: 80%;
      height: 100%;
      text-align: left;

      .header-text-lesson-wrapper {
        font-size: 15px;
        font-weight: 600;

        .header-text-lesson {
          background-color: #B8E994;
          border: solid 5px #B8E994;
          border-radius: 0.5em;
          float: left;
        }
      }


      .header-text-lesson_name {
        float: left;
      }
    }
  }

</style>