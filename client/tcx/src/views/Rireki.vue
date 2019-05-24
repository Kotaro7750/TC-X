<template>
  <div class="rireki">

    <select v-model="selectedYear">
      <option v-for="year in years" v-bind:key="year"  v-bind:value="year">{{year}}年</option>
    </select>
    <select v-model="selectedMonth">
      <option v-for="month in months" v-bind:key="month"  v-bind:value="month">{{month}}月</option>
    </select>

    <RirekiAdd v-bind:year="selectedYear" v-bind:month="selectedMonth" @on-add="onAdd"/>
    <RirekiList v-bind:year="selectedYear" v-bind:month="selectedMonth" ref="list"/>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import RirekiAdd from '@/components/RirekiAdd.vue';
import RirekiList from '@/components/RirekiList.vue';

@Component({
  components: {
    RirekiAdd,
    RirekiList,
  }
})

export default class Rireki extends Vue{
  years:number[] = [2019];
  months:number[] = [1,2,3,4,5,6,7,8,9,10,11,12];

  selectedYear = 2019;
  selectedMonth = 5;

  onAdd(){
    let rirekiList:any = this.$refs.list;
    rirekiList.getRirekiList();
  }
}
</script>


<style scoped>
/* TODO: add css */ 
</style>
