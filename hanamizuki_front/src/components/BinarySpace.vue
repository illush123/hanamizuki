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
      <div class="binary-space-body">
        <div class="lesson-description">{{lessonDescription}}</div>
        <div class="binary-space-lesson-required-box pr-3 pl-3">
          <div class="required-title pt-4 pb-2 pl-5">要件定義</div>
          <ul class="required-descriptions">
            <li v-for="element in lessonDescriptionElements" class="required-description">
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
  export default class HelloWorld extends Vue {
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
  // 1rem = 16px

  ul {
    list-style: none;
  }

  .binary-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .binary-space-lesson-title {
      width: 100%;
      background-color: #C7ECEE;
      border: solid 1px #C7ECEE;
      border-top-left-radius: 1rem;
      border-top-right-radius: 1rem;
      font-weight: 600;
      box-shadow: 2px 2px 2px 2px #888888;
    }

    .binary-space-body {
      width: 100%;
      height: 780px;
      padding-left: 25px;
      padding-right: 25px;
      background-color: #DFF9FB;
      border: solid 1px #DFF9FB;
      border-bottom-left-radius: 1rem;
      border-bottom-right-radius: 1rem;
      box-shadow: 2px 2px 2px 2px #888888;

      .lesson-description {
        font-size: 1.2rem;
        text-align: left;
        padding-top: 2rem;
        padding-bottom: 2rem;
      }

      .binary-space-lesson-required-box {
        border: solid 1px black;
        background-color: white;
        border-radius: 0.1rem;
        text-align: left;
        padding-bottom: 2.5rem;

        .required-title {
          border-bottom: solid 1px darkgray;
        }

        .required-descriptions {
          padding-top: 1.8rem;
          font-size: 0.8rem;

          .required-description {
            padding-bottom: 0.1rem;
          }
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