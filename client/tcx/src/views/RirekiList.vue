<template>
  <div class="rirekilist">
    <h2>履歴リスト</h2>
  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component({
})

export default class RirekiList extends Vue{
    rirekiList!: {
        id: number,
        joid:number,
        syubetsu:number,
        startTime:string,
        endTime:string,
    }[];

    created () {
        fetch("http://localhost:8888/rireki/5/63",{
            method: 'GET'
        }).then(response => {
            return response.json();
        }).then(json =>{
            console.log(json);
        })
    }
}
</script>
