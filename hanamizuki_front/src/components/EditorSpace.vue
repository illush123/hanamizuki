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
  @import "../assets/style/_variables";

  .editor-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .editor-space-header {
      width: 100%;
      background-color: $editor-back;
      border: solid 1px $editor-back;
      border-top-left-radius: 1rem;
      border-top-right-radius: 1rem;
      font-weight: 600;
      text-align: left;
      padding-left: 8rem;
      box-shadow: 2px 2px 2px 2px $box-shadow;
    }

    .editor-space-body {
      height: 780px;
      background-color: $editor-back;
      border: solid 1px $editor-back;
      border-bottom-left-radius: 1rem;
      border-bottom-right-radius: 1rem;
      box-shadow: 2px 2px 2px 2px ;

      .editor-space-tools {

        .folder-button {
          font-size: 0.3rem;
          background-color: $editor-button;
          color: $h-black;
          border-color: $editor-button;
        }
      }

      .editor {
        height: 580px;
        border: solid 1px $h-black;
        background-color: white;
      }

      .submit-button {
        font-size: 1.5rem;
        background-color: $editor-button;
        color: #FFF;
        border-color: $editor-button;
      }
    }
  }

</style>