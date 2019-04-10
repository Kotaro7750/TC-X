<template>
  <b-button class=Button>{{ title }}</b-button>
</template>

<script lang="ts">
import { Component,Prop, Vue } from 'vue-property-decorator';

@Component
export default class Button extends Vue{
  @Prop() 
  private title!: string; 
}
</script>
