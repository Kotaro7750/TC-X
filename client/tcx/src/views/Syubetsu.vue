<template>
  <div class="syubetsu">

    <SyubetsuAdd  @on-add="onAdd"/>
    <SyubetsuList  ref="list"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import SyubetsuAdd from '@/components/SyubetsuAdd.vue';
import SyubetsuList from '@/components/SyubetsuList.vue';

@Component({
  components: {
    SyubetsuAdd,
    SyubetsuList,
  },
  beforeRouteEnter: function(to, from, next) {
    next(component => {
      if (component.$store.getters.isSignIn == false) {
        next('/signin');
      }else{
        next();
      }
    })
  },
})

export default class Syubetsu extends Vue{
  onAdd(){
    let syubetsuList:any = this.$refs.list;
    syubetsuList.getSyubetsuList();
  }
}
</script>


<style scoped>
/* TODO: add css */ 
</style>
