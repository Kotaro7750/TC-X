<template>
  <div class="about">
    <Logo/>
    <h1>This is an about page</h1>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Logo from '@/components/Logo.vue';

@Component({
  components: {
    Logo,
  },
})
export default class About extends Vue {}
</script>
