<template>
  <div class="home">
    <Logo/>
    <HelloWorld msg="TC-X will change your life"/>
    <h2>{{ msg }}</h2>
    <router-link to="/rireki/add">Add Rireki</router-link>
    <router-link to="/rireki/list">List Rireki</router-link>
  </div>
</template>

<script lang="ts">
import { Component,Prop, Vue } from 'vue-property-decorator';
import axios from 'axios';
import Logo from '@/components/Logo.vue';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src

@Component({
  components: {
    Logo,
    HelloWorld,
  },
})
export default class Home extends Vue {
  msg!: string;
}
</script>
