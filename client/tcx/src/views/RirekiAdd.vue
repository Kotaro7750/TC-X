<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>
    <input type="number" v-model="rireki.joid">
    <input type="number" v-model="rireki.syubetsu">
    <input type="text" v-model="rireki.about">
    <input type="time" v-model="rireki.startTime">
    <input type="time" v-model="rireki.endTime">
    {{rireki}}
    <button @click="log">button</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
})

export default class RirekiAdd extends Vue{
    rireki:{
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    } ={joid:63,syubetsu:3,about:'',startTime:'',endTime:''};


    addRireki():void{
      var url = "http://localhost:8888/rireki/" + "5";
      fetch(url,{
        method: 'POST',
        body: JSON.stringify(this.rireki)
      }).then(response => {
        return response.json();
      }).then(json => {
        console.log(json);
      })
    }

    log():void{
      console.log(this.rireki);
    }

}
</script>
