<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isEditError">{{ editError }}</p>
    <ul>
        <li v-for="rireki in rirekiList" v-bind:key="rireki.id" v-show="editableID == -1 || editableID == rireki.id">
            {{ rireki.id }}
            {{ syubetsuToAbout(rireki.syubetsu) }}
            {{ rireki.about }}
            {{ rireki.startTime |Time}}〜{{ rireki.endTime |Time}}
            <button @click="deleteRireki(rireki.id)">削除</button>
            <button @click="triggerUpdate(rireki.id)" v-if="editableID == -1">編集</button>
            <button @click="cancelUpdate(rireki.id)" v-else>キャンセル</button>
            <RirekiInput v-if="editableID == rireki.id" @on-submit="editRireki" v-bind:year="year" v-bind:month="month" v-bind:initData="{
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

    editableID:number = -1;
    isEditError:boolean = false;
    editError:{} = {};

    personalInfo :{
      joid: number,
      name: string,
      token: string,
    } = this.$store.getters.userInfo;

    isLoading = false;
    syubetsuList:{syubetsu:number,name:string,salary:number}[] = [];

    @Watch('userInfo') onUserInfoChanged(){
      this.personalInfo = this.userInfo;
    }

    get userInfo(): {joid:number,name:string,token:string}{
      let userInfo = this.$store.getters.userInfo;
      return userInfo;
    }

    @Watch('month') onMonthChanged(){
        this.getRirekiList();
    }

    created () {
        this.isListError = false;
        this.getRirekiList();
        this.getSyubetsuList();
    }

    getRirekiList():void{
        this.isListError = false;
        let url = "http://localhost:8888/rireki/" + String(this.month) + "/" + String(this.personalInfo.joid);
        fetch(url,{
            method: 'GET',
            headers: {'Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
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
            let url = "http://localhost:8888/rireki/" + String(this.month) + "/" + String(this.personalInfo.joid) +"/" + String(id);
            fetch(url,{
                method: 'DELETE',
                headers: {'Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
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
        this.editableID =id;
    }

    cancelUpdate():void{
        this.editableID = -1;
        this.getRirekiList();
    }

    editRireki(rireki:{joid:number,syubetsu:number,about:string,startTime:string,endTime:string}):void{
        this.isEditError =false;
        let confirmEdit = confirm("更新していいですか？");
        if (confirmEdit == true) {
            let url = "http://localhost:8888/rireki/" + String(this.month) + "/" + String(this.personalInfo.joid) + "/" + String(this.editableID);
            fetch(url,{
                method: 'PATCH',
                headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
                body: JSON.stringify(rireki)
            }).then(response => {
                if (!response.ok) {
                    this.isEditError = true;
                }
                return response.json();
            }).then(json => {
                this.editError = json.error;
                this.getRirekiList();
            })
            this.editableID = -1;
            this.getRirekiList();
        }else{
        } 
    }

    getSyubetsuList():void{
        this.isListError = false;
        this.isLoading = true;
        this.$store.dispatch('getSyubetsuList').then(() => {
            this.syubetsuList = this.$store.getters.getSyubetsuList
            this.isLoading = false;
        }).catch((error) => {
            this.listError = error;
            this.isListError = true
        });
    }

    syubetsuToAbout(key:string):string{
        let syubetsuMap: {[key: string]: string} = {}; 
        for(let i  in this.syubetsuList){
            syubetsuMap[this.syubetsuList[i].syubetsu] = this.syubetsuList[i].name;
        }
        return syubetsuMap[key];

    }
}
</script>

<style scoped>
/* TODO: add css */
.listError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }
</style>
