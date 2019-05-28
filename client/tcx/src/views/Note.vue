<template>
  <div class="note">
    <NoteAdd  @on-add="onAdd"/>
    <NoteList  ref="list"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import NoteAdd from '@/components/NoteAdd.vue';
import NoteList from '@/components/NoteList.vue';

@Component({
  components: {
    NoteAdd,
    NoteList,
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

export default class Note extends Vue{
  onAdd(){
    let noteList:any = this.$refs.list;
    noteList.getNoteList();
  }
}
</script>

<style scoped>
/* TODO: add css */ 
</style>
