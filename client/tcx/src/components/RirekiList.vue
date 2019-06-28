<template>
  <div class="rirekilist">
    <h2 class="left col-lg-5 ">履歴リスト
        <span>
            <button @click="getRirekiList" class="update-button">
                <font-awesome-icon icon="sync-alt" />
            </button>
        </span>
    </h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isEditError">{{ editError }}</p>
    <p></p>
    <ul class="col-lg-12">
        <li class="col-lg-12" v-for="rireki in rirekiList" v-bind:key="rireki.id" v-show="editableID == -1 || editableID == rireki.id">
            <span class="col-lg-12 col-12 op-text row">
                <span class="op-text">
                    <font-awesome-icon icon="edit" @click="triggerUpdate(rireki.id)" v-if="editableID == -1" class="op-button"/>
                    <font-awesome-icon icon="window-close" @click="cancelUpdate(rireki.id)" v-else class="op-button"/>
                    <font-awesome-icon icon="trash-alt" @click="deleteRireki(rireki.id)" class="op-button"/>
                </span>
                {{ syubetsuNameMap[rireki.syubetsu] }}
                {{ rireki.startTime |Time(rireki.endTime)}}
            </span>

            <span class="col-lg-12 row">
                    {{ rireki.about }}
            </span>
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
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Watch} from "vue-property-decorator";
import RirekiInput from '@/components/RirekiInput.vue';

@Component({
    filters: {
        //filter of DB-formatted time to User-like time  
        Time:function (DBStartTime:string,DBEndTime:string) {
            let start_year:string = DBStartTime.split('-')[0];
            let start_month:string = DBStartTime.split('-')[1];
            let start_day:string = DBStartTime.split('-')[2].split('T')[0];
            let start_time:string = DBStartTime.split('-')[2].split('T')[1].split('+')[0];

            let end_year:string = DBEndTime.split('-')[0];
            let end_month:string = DBEndTime.split('-')[1];
            let end_day:string = DBEndTime.split('-')[2].split('T')[0];
            let end_time:string = DBEndTime.split('-')[2].split('T')[1].split('+')[0];

            if (start_month != end_month) {
                return start_month + "/" + start_day + "(" + end_month + "/" + end_day + ") " + start_time + "~" + end_time;
            }
            if (start_day != end_day) {
                return start_month + "/" + start_day + "(" + end_day + ") " + start_time + "~" + end_time;
            }
            return start_month + "/" + start_day + " " + start_time + "~" + end_time;
        },
    },
    components: {
        RirekiInput,
    }
})

export default class RirekiList extends Vue{
    @Prop(Number) month!:number;
    @Prop(Number) year!:number;

    //to debug, init list is set
    rirekiList: {
        id: number,
        joid:number,
        syubetsu:number,
        about:string,
        startTime:string,
        endTime:string,
    //}[] = [];
    }[] = [{id:23,joid:63,syubetsu:2,about:"ほげほげほげほげほげほげほげ",startTime:"2019-2-28T12:34",endTime:"2019-2-28T12:45"}];

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
    syubetsuNameMap: {[key: string]: string} = {}; 

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
        let url = process.env.VUE_APP_SERVER_URL+"/rireki/" + String(this.year) + "/"+ String(this.month) + "/" + String(this.personalInfo.joid);
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
            let url = process.env.VUE_APP_SERVER_URL+"/rireki/" + String(this.year)+ "/"+ String(this.month) + "/" + String(this.personalInfo.joid) +"/" + String(id);
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
            let url = process.env.VUE_APP_SERVER_URL+"/rireki/" + String(this.year)+ "/"+ String(this.month) + "/" + String(this.personalInfo.joid) + "/" + String(this.editableID);
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
            for(let i  in this.syubetsuList){
                this.syubetsuNameMap[this.syubetsuList[i].syubetsu] = this.syubetsuList[i].name;
            }
            this.isLoading = false;
        }).catch((error) => {
            this.listError = error;
            this.isListError = true
        });
    }


}
</script>

<style scoped>
    .listError { 
      padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
      }


    ul, ol {
      padding: 0;
      position: relative;
    }

    ul li, ol li {
      font-size: 15pt;
      text-align: left;
      color: black;
      border-left: solid 5px #2d9f91;/*左側の線*/
      background: whitesmoke;/*背景色*/
      margin-bottom: 5px;/*下のバーとの余白*/
      margin-left: 15px;
      line-height: 1.5;
      border-radius: 0 30px 30px 0;/*右側の角だけ丸く*/
      padding: 0.5em;
      list-style-type: none!important;
    }

    .op-button{
        font-size: 20pt;
        padding-left: 5px;
    }

    .op-text{
        text-align: right;
        padding-right: 10px;
    }

    .update-button{
      color: black;
      margin-left: 0.2em;
      margin-right: 0.2em;
      background: white;
      border: 1px solid white;
    }

    .left{
        text-align: left;
    }


</style>
