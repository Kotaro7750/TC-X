<template>
  <div class="rirekiadd">
    <h2>履歴追加</h2>
    <input type="number" v-bind:value="rireki.joid">
    <input type="number" v-bind:value="rireki.syubetsu">
    <input type="text" v-bind:value="rireki.about">
    <button @click="addRireki">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
})

export default class RirekiAdd extends Vue{
    rireki!: {
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    };

    created(){
      this.rireki = {
        joid:63,
        syubetsu:12,
        about:"abababa",
        startTime:"2019-05-01T12:58:00Z",
        endTime:"2019-05-01T12:56:00Z"
      }
    }

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
