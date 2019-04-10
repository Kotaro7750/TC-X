<template>
  <div class="signup">
    <Logo/>
    <h2>Sign up</h2>
      <label>Name</label>
      <input type="text" placeholder="Name">
    <button>Sign up</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import Logo from '@/components/Logo.vue';

@Component({
  components: {
    Logo,
  }
})

export default class Signup extends Vue{

}


  

</script>
