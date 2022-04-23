<script lang="ts">
import { defineComponent, computed, onMounted, ref} from 'vue'
import { useStore } from '@/store/index'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

export default defineComponent({
  setup() {
    const store = useStore()
    const router = useRouter()
    const myInfo = computed(() => store.state)
     

    let isSetCookie = ref(false)
    console.log(myInfo.value.id)

    onMounted(() => {
      axios.get('/api/checkCookie').then(() => {
        isSetCookie.value = true
      }).catch((err) => {
        return
      })
    })

    const goHome = () => {
      axios.get('/api/checkCookie').then(() => {
        router.push('/home')
      }).catch((err) => {
        alert(err)
      })
    }
    return {
      myInfo,
      goHome,
      isSetCookie
    }
  },
})
</script>
<template>
  <div v-if="isSetCookie">
    <div>
      <p>this is my profile</p>
      <p>Id: {{myInfo.id}}</p>
      <p>name: {{myInfo.name}}</p>
      <p>password: {{myInfo.password}}</p>
    </div>
    <div>
      <h1>ホームに行く</h1>
      <button @click="goHome">ホームにいく</button>
    </div>
  </div>
  <div v-else>
    <h1>You are not logged in</h1>
    <p>Please make sure you logged in</p>
    <router-link to="/login">ログインする</router-link>
  </div>

  
</template>