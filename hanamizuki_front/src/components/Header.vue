<template>
  <header>
    <b-container class="b-header-container">
      <div class="header-image-wrapper">
        <img src="/path/to/image" class="header-image_src">
      </div>
      <div class="header-text-wrapper">
        <div class="header-text-lesson_wrapper">
          <span class="header-text-lesson">レッスン一覧</span>
        </div>
        <span class="header-text-lesson_name">{{lessonTitle}}</span>
      </div>
    </b-container>
  </header>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component
export default class Header extends Vue {
  // data
  private lessonTitle: string = 'これはダミーメッセージです';
}
</script>

<style scoped lang="scss">
  @import "../assets/style/_variables";

  header {

    .header-image-wrapper {
      float: left;
      padding-right: 20px;
    }

    .header-text-wrapper {
      width: 100%;
      // height: auto;
      text-align: left;
      border-bottom: 3px solid $header-border;
      padding-bottom: 10px;

      .header-text-lesson_wrapper {
        font-size: 15px;
        font-weight: 600;
        padding-bottom: 4px;
        padding-right: 10px;

        .header-text-lesson {
          background-color: $header-text-wrapper;
          border: solid 5px $header-text-wrapper;
          border-radius: 0.5em;
        }
      }

      .header-text-lesson {
        float: left;
        margin-left: 32px;
      }

      .header-text-lesson_name {
        padding-left: 11px;
      }
    }
  }

</style>