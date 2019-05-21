<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
    <p class="fetchError" v-show="isError">{{ fetchError }}</p>
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
import {Component, Vue,Prop,Watch} from "vue-property-decorator";

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
    @Prop(Number) month!:number;
    @Prop(Number) year!:number;

    rirekiList: {
        id: number,
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    }[] = [];

    isError:boolean = false;
    fetchError:{} = {};

    @Watch('month') onMonthChanged(){
        this.fetchRireki();
    }

    created () {
        this.isError = false;
        this.fetchRireki();
    }

    fetchRireki():void{
        this.isError = false;
        let url = "http://localhost:8888/rireki/" + String(this.month) + "/63";
        fetch(url,{
            method: 'GET'
        }).then(response => {
            if (!response.ok) {
                this.isError =true;
            }
            return response.json();
        }).then(json =>{
            this.fetchError = json.error;
            this.loadJSONToRirekiList(json);
        })
    }

    loadJSONToRirekiList(json :any):void{
        this.rirekiList = json.result;
    }

}
</script>

<style scoped>
.fetchError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }
</style>
