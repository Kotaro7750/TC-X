<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
    <ul>
        <li v-for="rireki in rirekiList" v-bind:key="rireki.id">
            {{ rireki.id }}
            {{ rireki.syubetsu }}
            {{ rireki.about }}
            {{ rireki.startTime |Time}}〜{{ rireki.endTime |Time}}
        </li>
    </ul>
    <button @click="fetchRireki">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
    filters: {
        //filter of DB-formatted time to User-like time  
        Time:function (DBTime:string) {
            let year: string = DBTime.split('-')[0];
            let month:string = DBTime.split('-')[1];
            let day:string = DBTime.split('-')[2].split('T')[0];
            let hour:string = DBTime.split('-')[2].split('T')[1].split(':')[0];
            let minute:string = DBTime.split('-')[2].split('T')[1].split(':')[1];

            return year + "年" + month + "月" + day + "日 " + hour + ":" + minute;
        }
    }
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
