<template>
  <div class="notelist">
    <h2>ノートリスト</h2>
    <p class="listError" v-show="isListError">{{ listError }}</p>
    <p class="deleteError" v-show="isDeleteError">{{ deleteError }}</p>
    <p class="updateError" v-show="isEditError">{{ editError }}</p>
    <p class="loader" v-if="isLoading">Loading</p>
    <ul v-else>
        <li v-for="note in noteList" v-bind:key="note">
            {{ note | Note}}
            <button @click="deleteNote(note)">削除</button>
        </li>
    </ul>
    <button @click="getNoteList">click</button>
  </div>
</template>

<script lang="ts">
import {Component, Vue,Prop,Watch} from "vue-property-decorator";

@Component({
    components: {
    },
    filters: {
        Note:function(DBNote: string) {
            let year = DBNote.split('_')[0];
            let month = DBNote.split('_')[1];

            return year + "年" + month + "月ノート";
        }
    },
})

export default class NoteList extends Vue{

    isLoading = false;
    noteList: string[] = [];

    isListError:boolean = false;
    listError:{} = {};

    isDeleteError:boolean = false;
    deleteError:{} = {};

    created () {
        this.getNoteList();
    }

    getNoteList():void{
        this.isListError = false;
        this.isLoading = true;
        this.$store.dispatch('getNoteList').then(() => {
            this.noteList = this.$store.getters.getNoteList
            this.isLoading = false;
        }).catch((error) => {
            this.listError = error;
            this.isListError = true
        });
    }

    deleteNote(note:string):void{
        this.isDeleteError =false;
        let confirmDelete = confirm("消していいですか？");
        if (confirmDelete == true) {
            let year = note.split('_')[0];
            let month = note.split('_')[1];
            let url = "http://localhost:8888/note/" + String(year) + "/" + String(month);
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
                this.getNoteList();
            })
        }else{
        } 
    }
}
</script>