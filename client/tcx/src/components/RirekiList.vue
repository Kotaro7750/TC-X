<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isUpdateError">{{ updateError }}</p>
    <ul>
        <li v-for="rireki in rirekiList" v-bind:key="rireki.id" v-show="updatableID == -1 || updatableID == rireki.id">
            {{ rireki.id }}
            {{ rireki.syubetsu }}
            {{ rireki.about }}
            {{ rireki.startTime |Time}}〜{{ rireki.endTime |Time}}
            <button @click="deleteRireki(rireki.id)">delete</button>
            <button @click="triggerUpdate(rireki.id)" v-if="updatableID == -1">update</button>
            <button @click="cancelUpdate(rireki.id)" v-else>cancel</button>
            <RirekiInput v-if="updatableID == rireki.id" @on-submit="updateRireki" v-bind:year="year" v-bind:month="month" v-bind:initData="{
                joid:rireki.joid,
                syubetsu:rireki.syubetsu,
                about:rireki.about,
                startDay:rireki.startTime.split('-')[2].split('T')[0],
                startTime:rireki.startTime,
                endTime:rireki.endTime,
            }"/>
        </li>
    </ul>
    <button @click="getRirekiList">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Watch} from "vue-property-decorator";
import RirekiInput from '@/components/RirekiInput.vue';

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
        },
    },
    components: {
        RirekiInput,
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

    isListError:boolean = false;
    listError:{} = {};

    isDeleteError:boolean = false;
    deleteError:{} = {};

    updatableID:number = -1;
    isUpdateError:boolean = false;
    updateError:{} = {};

    @Watch('month') onMonthChanged(){
        this.getRirekiList();
    }

    created () {
        this.isListError = false;
        this.getRirekiList();
    }

    getRirekiList():void{
        this.isListError = false;
        let url = "http://localhost:8888/rireki/" + String(this.month) + "/63";
        fetch(url,{
            method: 'GET'
        }).then(response => {
            if (!response.ok) {
                this.isListError =true;
            }
            return response.json();
        }).then(json =>{
            this.listError = json.error;
            this.loadJSONToRirekiList(json);
        })
    }

    loadJSONToRirekiList(json :any):void{
        this.rirekiList = json.result;
    }

    deleteRireki(id:number):void{
        this.isDeleteError =false;
        let confirmDelete = confirm("消していいですか？");
        if (confirmDelete == true) {
            let url = "http://localhost:8888/rireki/" + String(this.month) + "/63/" + String(id);
            fetch(url,{
                method: 'DELETE'
            }).then(response => {
                if (!response.ok) {
                    this.isDeleteError = true;
                }
                return response.json();
            }).then(json => {
                this.deleteError = json.error;
                this.getRirekiList();
            })
        }else{
        } 
    }

    triggerUpdate(id:number):void{
        this.updatableID =id;
    }

    cancelUpdate():void{
        this.updatableID = -1;
        this.getRirekiList();
    }

    updateRireki(rireki:{joid:number,syubetsu:number,about:string,startTime:string,endTime:string}):void{
        this.isUpdateError =false;
        let confirmUpdate = confirm("更新していいですか？");
        if (confirmUpdate == true) {
            let url = "http://localhost:8888/rireki/" + String(this.month) + "/63/" + String(this.updatableID);
            fetch(url,{
                method: 'PATCH',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(rireki)
            }).then(response => {
                if (!response.ok) {
                    this.isUpdateError = true;
                }
                return response.json();
            }).then(json => {
                this.updateError = json.error;
                this.getRirekiList();
            })
            this.updatableID = -1;
            this.getRirekiList();
        }else{
        } 
    }

}
</script>

<style scoped>
.listError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }
</style>
