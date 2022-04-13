<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from "axios"
import { useRoute, useRouter } from 'vue-router'
import { useStore } from '@/store/index'


export default defineComponent({
  setup() {
    const store = useStore()
    const name= ref("")
    const password = ref("")
    const router = useRouter()
    const route = useRoute()

    interface User {
      name: string,
      password: string
    }

    interface afterLoginUser extends User {
      Id: string,
    }

    const login = () => {
      const obj:User = {
        name: name.value,
        password: password.value
      }
      axios.post("/api/login", obj).then(res => {
        if ( res.data.isCorrectUser ) {
          const myprofile:afterLoginUser = res.data
          store.dispatch('login', myprofile)
          router.push(`/mypage/${myprofile.Id}`)
        } else {
          return alert(res.data.message)
        }
      })
    }

    return { 
      name,
      password,
      login
    }
  },
})
</script>

<template>
  <div>
    <h1>ログインページ</h1>
    <input type="text" name="name" v-model="name">
    <input type="text" name="password" v-model="password">
    <button @click="login">ログインする</button>
  </div>
  <div>
    <h1>新規登録はこちらから</h1>
    <router-link to='/add'>新規登録のページにいく</router-link>
  </div>
</template>
