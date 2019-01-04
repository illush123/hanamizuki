<template>
    <div class="editor-space-container">
      <div class="container">
        <div class="editor-space-header">
          <div class="editor-space-title pt-1 pb-1">VM&nbspCode</div>
        </div>
        <div class="editor-space-body pl-4 pr-4">
          <div class="editor-space-tools pt-5 pb-2">
            <b-button id="folder-button">フォルダを開く</b-button>
            <b-dropdown id="language-dropdown-button" :text="dropDownNow">
              <b-dropdown-item-button @click="chooseLanguage('C言語')">C言語</b-dropdown-item-button>
              <b-dropdown-item-button @click="chooseLanguage('C++')">C++</b-dropdown-item-button>
              <b-dropdown-item-button @click="chooseLanguage('Java')">Java</b-dropdown-item-button>
              <b-dropdown-item-button @click="chooseLanguage('Python')">Python</b-dropdown-item-button>
            </b-dropdown>
          </div>
          <div class="editor pt-3">
            <!-- editor? -->
          </div>
          <div class="editor-space-submit-button pt-4">
            <b-button >検証スタート</b-button>
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
  private dropDownNow = "C言語";

  chooseLanguage (language) {
    this.dropDownNow = language;
  }
}

</script>

<style scoped lang="scss">

  .editor-space-container {
    font-family: 'Monaco', 'Avenir', 'HiraginoSans-W6', Helvetica, Arial, sans-serif;

    .editor-space-header {
      width: 100%;
      background-color: chartreuse;
      border: solid 1px chartreuse;
      border-top-left-radius: 1rem;
      border-top-right-radius: 1rem;
      font-weight: 600;
      text-align: left;
      padding-left: 8rem;
      box-shadow: 2px 2px 2px 2px #888888;
    }

    .editor-space-body {
      height: 780px;
      background-color: darkgreen;
      border: solid 1px darkgreen;
      border-bottom-left-radius: 1rem;
      border-bottom-right-radius: 1rem;
      box-shadow: 2px 2px 2px 2px #888888;

      .editor {
        height: 580px;
        border: solid 1px black;
        background-color: white;

      }
    }
  }

</style>