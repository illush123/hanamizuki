<template>
  <div class="binary-space-wrapper">
    <div class="binary-space-lesson-title-wrapper">
      <span class="lesson-number">{{lessonNumber}}</span>
      <span class="lesson-title">{{lessonTitle}}</span>
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
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
  @Prop({ default: '01' }) private lessonNum!: string;
  @Prop({ default: 'dummyLesson' }) private title!: string;
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

  
</style>