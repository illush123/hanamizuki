<template>
  <header>
    <div class="header-image-wrapper">
      <img src="/path/to/image" class="header-image_src">
    </div>
    <div class="header-text-wrapper">
      <div class="header-text-lesson_wrapper">
        <span class="header-text-lesson">レッスン一覧</span>
      </div>
      <span class="header-text-lesson_name">{{lessonTitle}}</span>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
  @Prop({ default: 'dummy message' }) private messageProp!: string;

  // data
  private lessonTitle: string = this.messageProp;
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
      // height: auto;
      text-align: left;
      border-bottom: 3px solid #a9eda9;
      padding-bottom: 10px;

      .header-text-lesson_wrapper {
        font-size: 15px;
        font-weight: 600;
        padding-bottom: 4px;
        padding-right: 10px;

        .header-text-lesson {
          background-color: #B8E994;
          border: solid 5px #B8E994;
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