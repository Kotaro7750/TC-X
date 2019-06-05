<template>
  <div class="container-fluid">
    <img class="logo" alt="Vue logo" src="../assets/logo.png">
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';

@Component({
})
export default class Logo extends Vue {}
</script>

<style scoped>
  .logo{
    width: 400px;
    max-width: 90%;
    height: auto;
  }

</style>
