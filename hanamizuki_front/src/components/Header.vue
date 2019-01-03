<template>
  <header>
    <div class="header-image-wrapper">
      <img src="/path/to/image" class="header-image_src">
    </div>
    <div class="header-test-wrapper">
      <span class="header-text-lesson">レッスン一覧</span>
      <span class="header-text-lesson_name">{{lessonTitle}}</span>
    </div>
  </header>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
  // @Prop() private msg!: string;
}
</script>

<style scoped>

</style>