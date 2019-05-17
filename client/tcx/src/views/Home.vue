<template>
  <div class="home">
    <Logo/>
    <HelloWorld msg="TC-X will change your life"/>
    <h2>{{ msg }}</h2>
    <button @click="apiPublic">Public</button>
    <button @click="apiPrivate">Private</button>
  </div>
</template>

<script lang="ts">
import { Component,Prop, Vue } from 'vue-property-decorator';
import axios from 'axios';
import Logo from '@/components/Logo.vue';
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import Button from '@/components/Button.vue';

@Component({
  components: {
    Logo,
    HelloWorld,
    Button,
  },
})
export default class Home extends Vue {
  msg!: string;

  async apiPublic(){
    var ret = await axios.get('http://localhost:8888');
    this.msg = ret.data;
  }

  async apiPrivate(){
    var ret = await axios.get('http://localhost:8888');
    this.msg = ret.data;
  }
}
</script>
