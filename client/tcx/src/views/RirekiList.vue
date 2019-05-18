<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
    <ul>
        <li v-for="rireki in rirekiList" v-bind:key="rireki.id">
            {{ rireki.id }}
            {{ rireki.syubetsu }}
            {{ rireki.about }}
            {{ rireki.startTime }}
            {{ rireki.endTime }}
        </li>
    </ul>
    <button @click="fetchRireki">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
})

export default class RirekiList extends Vue{
    rirekiList: {
        id: number,
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    }[] = [];


    created () {
        this.fetchRireki();
    }

    fetchRireki():void{
        fetch("http://localhost:8888/rireki/5/63",{
            method: 'GET'
        }).then(response => {
            return response.json();
        }).then(json =>{
            this.loadJSONToRirekiList(json);
        })
    }

    loadJSONToRirekiList(json :any):void{
        this.rirekiList = json.result;
        console.log(this.rirekiList.length);
    }

}
</script>
