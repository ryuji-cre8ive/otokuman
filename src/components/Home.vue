<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  setup() {
    const otokuList = [
      {
        title: 'マックのお得',
        content: 'マックのお得な情報を発信しています'
      },
      {
        title: 'セブンのお得',
        content: 'セブンのお得な情報を発信しています'
      },
      {
        title: 'ローソンのお得',
        content: 'ローソンのお得な情報を発信しています'
      },
    ]

    return { otokuList}
  },
})
</script>
<template>
  <h1>this is home page which is posted a lot of articles about otoku</h1>
  <ul>
    <li v-for="o in otokuList" :key="o.title">
      <p>{{o.title}}</p>
      <p>{{o.content}}</p>
    </li>
  </ul>
</template>