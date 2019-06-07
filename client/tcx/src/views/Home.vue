<template>
  <div class="home">
    <Logo class="mt-5"/>
    <h1 class="mt-3 mb-3">TC-X will change your life</h1>
    <div class="row">
      <b-button pill  to="/rireki" class="button-secondary col-md-2 offset-md-4 col-xs-2 offset-xs-4">履歴ノート</b-button>
      <b-button pill variant="outline-secondary" to="/about" class=" col-md-2 col-xs-2 ">About</b-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Component,Prop, Vue } from 'vue-property-decorator';
import Logo from '@/components/Logo.vue';

@Component({
  components: {
    Logo,
  },
})
export default class Home extends Vue {
  msg!: string;

}
</script>

<style scoped>
  .btn-secondary, .btn-secondary:hover{
    color: #FFFFFF;
    background-color: #2D9F91;
    border-color: #2D9F91;
  }
</style>
