<template>
  <div class="signin">
    <Logo/>
    <h2>Sign in</h2>
      <label>Name</label>
      <input type="text" placeholder="Name">
    <button>Sign in</button>

    <p>You don't have an account?
      <router-link to="/signup">Create account now!</router-link>
    </p>
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

export default class Signin extends Vue{

}


  

</script>
