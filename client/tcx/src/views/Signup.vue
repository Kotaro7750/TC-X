<template>
  <div class="signup">
    <Logo/>
    <h2>Sign up</h2>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop} from "vue-property-decorator";

import Logo from '@/components/Logo.vue';

@Component({
  components: {
    Logo,
  }
})

export default class Signup extends Vue{
  joid!: Number;
  name!: string;
  alert: string = "";

  created(){
    this.alert = "";
  }

}
</script>