<template>
  <div class="rireki">

    <select v-model="selected">
      <option v-for="note in noteList" v-bind:key="note"  v-bind:value="note">{{note | Note}}</option>
    </select>
    <RirekiAdd v-bind:year="Number(selected.split('_')[0])" v-bind:month="Number(selected.split('_')[1])" @on-add="onAdd"/>
    <RirekiList v-bind:year="Number(selected.split('_')[0])" v-bind:month="Number(selected.split('_')[1])" ref="list"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import moment from "moment-timezone"
import RirekiAdd from '@/components/RirekiAdd.vue';
import RirekiList from '@/components/RirekiList.vue';

@Component({
  components: {
    RirekiAdd,
    RirekiList,
  },
  filters: {
      Note:function(DBNote: string) {
          let year = DBNote.split('_')[0];
          let month = DBNote.split('_')[1];

          return year + "年" + month + "月ノート";
      }
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

export default class Rireki extends Vue{
  selected = String(moment().year()) + "_" + String(moment().month()+1) + "_rireki";

  noteList: string[] = [];

  isListError:boolean = false;
  listError:{} = {};

  created(){
    this.isListError = false;
    this.$store.dispatch('getNoteList').then(() => {
        this.noteList = this.$store.getters.getNoteList
    }).catch((error) => {
        this.listError = error;
        this.isListError = true
    });
  }

  onAdd(){
    let rirekiList:any = this.$refs.list;
    rirekiList.getRirekiList();
  }

}
</script>


<style scoped>
/* TODO: add css */ 
</style>
