<script lang="ts">
import { defineComponent, computed } from 'vue'
import { useStore } from '@/store/index'

export default defineComponent({
  setup() {
    const store = useStore()
    const myInfo = computed(() => store.state)
    console.log(myInfo.value.id)
    return {
      myInfo
    }
  },
})
</script>
<template>
  <div>
    <p>this is my profile</p>
    <p>Id: {{myInfo.id}}</p>
    <p>name: {{myInfo.name}}</p>
    <p>password: {{myInfo.password}}</p>
  </div>
  <div>
    <h1>ホームに行く</h1>
    <router-link to='/home'>ホームに行く</router-link>
  </div>

  
</template>