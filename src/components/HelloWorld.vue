<template>
  <div>
    <button @click="sendReqToBack">結果を見る</button>
    <h1>結果確認</h1>
    <p>GET: <br/>{{consequence}} </p>
    <button @click="deleteConsequence">結果削除</button>
  </div>
</template>

<script lang="ts">
import axios from "axios";
import { defineComponent, ref } from 'vue'


export default defineComponent({
  setup() {
    let consequence = ref("")

    const sendReqToBack = async () => {
      await axios.get("/hello").then(res => {
        consequence.value = res.data
      })
    }

    const deleteConsequence = () => {
      consequence.value = ""
    }

    return { 
      consequence,
      sendReqToBack,
      deleteConsequence
    }
  },
})
</script>
