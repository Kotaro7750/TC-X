<template>
  <div class="syubetsulist">
    <h2>種別リスト</h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isEditError">{{ editError }}</p>
    <p class="loader" v-if="isLoading">Loading</p>
    <ul v-else>
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

    isLoading = false;
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
        this.getSyubetsuList();
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

    deleteSyubetsu(id:number):void{
        this.isDeleteError =false;
        let confirmDelete = confirm("消していいですか？");
        if (confirmDelete == true) {
            let url = process.env.VUE_APP_SERVER_URL+"/syubetsu/" + String(id);
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
            let url = process.env.VUE_APP_SERVER_URL+"/syubetsu/" + String(this.editableID);
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

.loader,
.loader:before,
.loader:after {
  background: #ffffff;
  -webkit-animation: load1 1s infinite ease-in-out;
  animation: load1 1s infinite ease-in-out;
  width: 1em;
  height: 4em;
}
.loader {
  color: #174dc2;
  text-indent: -9999em;
  margin: 88px auto;
  position: relative;
  font-size: 11px;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  position: absolute;
  top: 0;
  content: '';
}
.loader:before {
  left: -1.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 1.5em;
}
@-webkit-keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}
@keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}

</style>
