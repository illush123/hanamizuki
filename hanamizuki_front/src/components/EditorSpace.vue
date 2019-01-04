<template>
    <div class="editor-space-container">
      <div class="container">
        <div class="editor-space-header">
          <div class="editor-space-title pt-1 pb-1">VM&nbspCode</div>
        </div>
        <div class="editor-space-body pl-4 pr-4">
          <div class="editor-space-tools pt-5 pb-2">
            <b-button class="folder-button btn btn-default pull-right">&nbsp&nbsp&nbsp&nbspフォルダを開く</b-button>
          </div>
          <div class="editor pt-3">
            <!-- editor? -->
          </div>
          <div class="editor-space-submit-button pt-4">
            <b-button class="submit-button" size="lg">検証スタート</b-button>
          </div>
        </div>
      </div>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component
export default class EditorSpace extends Vue {
}

</script>

<style scoped lang="scss">

  .editor-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .editor-space-header {
      width: 100%;
      background-color: #AFE686;
      border: solid 1px #AFE686;
      border-top-left-radius: 1rem;
      border-top-right-radius: 1rem;
      font-weight: 600;
      text-align: left;
      padding-left: 8rem;
      box-shadow: 2px 2px 2px 2px #888888;
    }

    .editor-space-body {
      height: 780px;
      background-color: #C1ECA2;
      border: solid 1px #C1ECA2;
      border-bottom-left-radius: 1rem;
      border-bottom-right-radius: 1rem;
      box-shadow: 2px 2px 2px 2px #888888;

      .editor-space-tools {

        .folder-button {
          font-size: 0.3rem;
          background-color: #93DD5C;
          color: black;
          border-color: #93DD5C;
        }
      }

      .editor {
        height: 580px;
        border: solid 1px black;
        background-color: white;
      }

      .submit-button {
        font-size: 1.5rem;
        background-color: #93DD5C;
        color: #FFF;
        border-color: #93DD5C;
      }
    }
  }

</style>