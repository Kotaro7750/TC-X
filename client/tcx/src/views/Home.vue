<template>
  <div class="home">
    <Logo/>
    <h1>TC-X will change your life</h1>
    <div class="row">
      <b-button pill variant="info" to="/rireki" class="col-md-2 offset-md-4 col-xs-2 offset-xs-4">履歴ノート</b-button>
      <b-button pill variant="outline-secondary" to="/about" class="col-md-2 col-xs-2 ">About</b-button>
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
/* TODO: add css */

</style>
