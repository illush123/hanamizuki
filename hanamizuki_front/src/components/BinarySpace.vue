<template>
  <div class="binary-space-wrapper">
    <div class="container">
      <div class="binary-space-lesson-title-wrapper">
        <div class="row">
          <div class="col-5">
            <span class="lesson-const">Lesson</span>
            <span class="lesson-number">{{lessonNumber}}</span>
          </div>
          <div class="col-7">
            <span class="lesson-title">{{lessonTitle}}</span>
          </div>
        </div>
      </div>
      <div class="binary-space-lesson-description-wrapper">
        <span class="lesson-description">{{lessonDescription}}</span>
      </div>
      <div class="binary-space-lesson-required-wrapper">
        <p class="required-title">要件定義</p>
        <ul>
          <li v-for="element in lessonDescriptionElements" class="required-description">
            {{element}}
          </li>
        </ul>
      </div>
      <div class="binary-space-careful-comments">
        <ul>
          <li v-for="comment in carefulComments">
            {{comment}}
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component
export default class HelloWorld extends Vue {
  @Prop({ default: '01' }) private lessonNum!: string;
  @Prop({ default: '四則演算' }) private title!: string;
  @Prop({ default: 'これはダミーメッセージです。'}) private description!: string;
  @Prop(Array) private descriptions!: string[];
  @Prop(Array) private carefuls!: string[];

  private lessonNumber: string = this.lessonNum;
  private lessonTitle: string = this.title;
  private lessonDescription: string = this.description;
  private lessonDescriptionElements: string[] = this.descriptions;
  private carefulComments: string[] = this.carefuls;
}

</script>

<style scoped lang="scss">

  .binary-space-wrapper {
    width: 100%;
    background-color: #DFF9FB;
    border: solid 1px #DFF9FB;
    border-radius: 1em;
    font-family: 'Monaco','Avenir', Helvetica, Arial, sans-serif;

    .binary-space-lesson-title-wrapper {
      padding-top: 4px;
      padding-bottom: 4px;
      text-align: center;
      background-color: cyan;
      font-weight: 600;
      border: solid 1px cyan;
      // border-right: solid 1px cyan;
      // border-left: solid 1px cyan;
      border-top-left-radius: 1em;
      border-top-right-radius: 1em;


      .lesson-const {
        padding-left: 40px;
        padding-right: 10px;
      }
    }
  }
</style>