<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <router-link to="/about">About</router-link>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';

@Component
export default class HelloWorld extends Vue {
  @Prop() private msg!: string;
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  align: center;
  margin: 40px 0 0;
}
</style>
