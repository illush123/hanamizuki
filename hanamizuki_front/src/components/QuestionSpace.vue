<template>
  <div class="binary-space-container">
    <div class="container">
      <div class="binary-space-lesson-title">
        <div class="row">
          <div class="col-5 pt-1 pb-1 pl-5">
            <span class="lesson-const">Lesson</span>
            <span class="lesson-number">&nbsp&nbsp{{lessonNumber}}</span>
          </div>
          <div class="col-7 pt-1 pb-1">
            <span class="lesson-title">{{lessonTitle}}</span>
          </div>
        </div>
      </div>
      <div class="binary-space-body pl-4 pr-4">
        <div class="lesson-description pt-4 pb-4">{{lessonDescription}}</div>
        <div class="binary-space-lesson-required-box pr-3 pl-3 pb-4">
          <div class="required-title pt-4 pb-2 pl-5">要件定義</div>
          <ul class="required-descriptions pt-4">
            <li v-for="element in lessonDescriptionElements" class="required-description pb-1">
              {{element}}
            </li>
          </ul>
        </div>
        <div class="binary-space-careful-comments pt-5 pb-5">
          <ul class="carefuls-list">
            <li v-for="comment in carefulComments" class="carefuls-list-item">
              {{comment}}
            </li>
          </ul>
        </div>
        <div class="last pb-5">
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
  import {Component, Prop, Vue} from 'vue-property-decorator';
  import 'bootstrap/dist/css/bootstrap.css';
  import 'bootstrap-vue/dist/bootstrap-vue.css';

  @Component
  export default class QuestionSpace extends Vue {
    @Prop({default: '01'}) private lessonNum!: string;
    @Prop({default: '四則演算'}) private title!: string;
    @Prop({default: '四則演算を実装しよう'}) private description!: string;
    @Prop(Array) private descriptions!: string[];
    @Prop(Array) private carefuls!: string[];

    private lessonNumber: string = this.lessonNum;
    private lessonTitle: string = this.title;
    private lessonDescription: string = this.description;
    // とりあえず確認のため
    private lessonDescriptionElements: string[] = ['要件定義の中身です', '要件定義の中身です', '要件定義の中身です', '要件定義の中身です'];
    // とりあえず確認のため
    private carefulComments: string[] = ['気をつけることなどコメント', '気をつけることなどコメント', '気をつけることなどコメント', '気をつけることなどコメント'];
  }

</script>

<style scoped lang="scss">
  @import "../assets/style/_variables";
  // 1rem = 16px

  ul {
    list-style: none;
  }

  .binary-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .binary-space-lesson-title {
      width: 100%;
      background-color: $question-header;
      border: solid 1px $question-header;
      border-top-left-radius: 1rem;
      border-top-right-radius: 1rem;
      font-weight: 600;
      box-shadow: 2px 2px 2px 2px $box-shadow;
    }

    .binary-space-body {
      width: 100%;
      height: 780px;
      background-color: $question-body;
      border: solid 1px $question-body;
      border-bottom-left-radius: 1rem;
      border-bottom-right-radius: 1rem;
      box-shadow: 2px 2px 2px 2px $box-shadow;

      .lesson-description {
        font-size: 1.2rem;
        text-align: left;
      }

      .binary-space-lesson-required-box {
        border: solid 1px $h-black;
        background-color: white;
        border-radius: 0.1rem;
        text-align: left;

        .required-title {
          border-bottom: solid 1px darkgray;
        }

        .required-descriptions {
          font-size: 0.8rem;
        }
      }
    }

    .binary-space-careful-comments {
      text-align: left;
      font-size: 0.7rem;

      .carefuls-list-item {
        padding-bottom: 0.2rem;
        padding-left: 0px;
      }
    }
  }
</style>