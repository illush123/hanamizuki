<template>
    <div class="editor-space-container">
      <b-container>
        <div class="editor-space-header">
          <div class="editor-space-title pt-1 pb-1">VM&nbsp;Code</div>
        </div>
        <div class="editor-space-body pl-4 pr-4">
          <b-row>
            <b-col cols="9"></b-col>
            <b-col cols="3">
              <div class="editor-space-tools pt-3 pb-2">
                <b-button class="folder-button">&nbsp;&nbsp;&nbsp;&nbsp;フォルダを開く</b-button>
              </div>
            </b-col>
          </b-row>
          <div class="editor">
            <!-- editor? -->
            <div id="editor-space" class="editor" style="{height: 100%}"></div>
          </div>
          <div class="editor-space-submit-button pt-4">
            <b-button class="submit-button" size="lg">検証スタート</b-button>
          </div>
        </div>
      </b-container>
    </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import * as monaco from 'monaco-editor';
import IStandaloneCodeEditor = monaco.editor.IStandaloneCodeEditor;
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';

@Component
export default class EditorSpace extends Vue {
  private editorData = '';

  async mounted() {
    await this.$store.dispatch('DataStore/uploadUserProgram');
    const target = document.getElementById('editor-space');
    console.log(this.$store.state.DataStore.test);
    if (target !== null) {
      const editor: IStandaloneCodeEditor = monaco.editor.create(target, {
        value: this.$store.state.DataStore.test.status,
        language: 'c',
        minimap: {
          enabled: false,
        },
        roundedSelection: false,
        readOnly: false,
      });
    }
  }

  // monaco-editorはeditor.getValue()で内容を取得できそう。
}

</script>

<style scoped lang="scss">
  @import "../assets/style/_variables";

  .editor-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .editor-space-header {
      width: 100%;
      background-color: $editor-header;
      border: solid 1px $editor-header;
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
      box-shadow: 2px 2px 2px 2px $box-shadow;

      .editor-space-tools {

        .folder-button {
          font-size: 0.3rem;
          background-color: $editor-button;
          color: $h-black;
          border-color: $editor-button;
        }
      }

      .editor {
        text-align: left;
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