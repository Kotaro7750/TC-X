<template>
  <div class="syubetsulist">
    <h2>種別リスト</h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isEditError">{{ editError }}</p>
    <ul>
        <li v-for="syubetsu in syubetsuList" v-bind:key="syubetsu.syubetsu" v-show="editableID == -1 || editableID == syubetsu.syubetsu">
            {{ syubetsu.syubetsu }}
            {{ syubetsu.name }}
            {{ syubetsu.salary }}
            <button @click="deleteSyubetsu(syubetsu.syubetsu)">削除</button>
            <button @click="triggerUpdate(syubetsu.syubetsu)" v-if="editableID == -1">編集</button>
            <button @click="cancelUpdate(syubetsu.syubetsu)" v-else>キャンセル</button>
            <SyubetsuInput v-if="editableID == syubetsu.syubetsu" @on-submit="editSyubetsu" v-bind:initData="{
                syubetsu:syubetsu.syubetsu,
                name:syubetsu.name,
                salary:syubetsu.salary,
            }"/>
        </li>
    </ul>
    <button @click="getSyubetsuList">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Watch} from "vue-property-decorator";
import SyubetsuInput from '@/components/SyubetsuInput.vue';

@Component({
    components: {
        SyubetsuInput,
    }
})

export default class SyubetsuList extends Vue{

    syubetsuList: {
        syubetsu:number,
        name: string,
        salary:number,
    }[] = [];

    isListError:boolean = false;
    listError:{} = {};

    isDeleteError:boolean = false;
    deleteError:{} = {};

    editableID:number = -1;
    isEditError:boolean = false;
    editError:{} = {};


    created () {
        this.isListError = false;
        this.getSyubetsuList();
    }

    getSyubetsuList():void{
        this.isListError = false;
        let url = "http://localhost:8888/syubetsu";
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
            this.loadJSONToSyubetsuList(json);
        })
    }

    loadJSONToSyubetsuList(json :any):void{
        this.syubetsuList = json.result;
    }

    deleteSyubetsu(id:number):void{
        this.isDeleteError =false;
        let confirmDelete = confirm("消していいですか？");
        if (confirmDelete == true) {
            let url = "http://localhost:8888/syubetsu/" + String(id);
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
                this.getSyubetsuList();
            })
        }else{
        } 
    }

    triggerUpdate(id:number):void{
        this.editableID =id;
    }

    cancelUpdate():void{
        this.editableID = -1;
        this.getSyubetsuList();
    }

    editSyubetsu(syubetsu:{syubetsu:number,name:string,salary:number}):void{
        this.isEditError =false;
        let confirmEdit = confirm("更新していいですか？");
        if (confirmEdit == true) {
            let url = "http://localhost:8888/syubetsu/" + String(this.editableID);
            fetch(url,{
                method: 'PATCH',
                headers: {'Content-Type': 'application/json','Authorization': 'Bearer ' + this.$store.getters.userInfo.token,},
                body: JSON.stringify(syubetsu)
            }).then(response => {
                if (!response.ok) {
                    this.isEditError = true;
                }
                return response.json();
            }).then(json => {
                this.editError = json.error;
                this.getSyubetsuList();
            })
            this.editableID = -1;
            this.getSyubetsuList();
        }else{
        } 
    }

}
</script>

<style scoped>
/* TODO: add css */
.listError { 
  padding:12px; font-weight:850; color:#262626; background:#FFEBE8; border:2px solid #990000; 
  }
</style>
