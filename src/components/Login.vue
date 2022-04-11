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

    const addUser = () => {
      const obj:User = {
        name: name.value,
        password: password.value
      }
      axios.post("/login", obj).then(res => {
        if ( res.data.isCorrectUser ) {
          const myprofile:afterLoginUser = res.data
          console.log(myprofile)
          console.log(store.dispatch('login', myprofile))
          router.push(`/mypage/${myprofile.Id}`)
        }
      })
    }

    return { 
      name,
      password,
      addUser
    }
  },
})
</script>

<template>
  <div>
    <h1>ログインページ</h1>
    <input type="text" name="name" v-model="name">
    <input type="text" name="password" v-model="password">
    <button @click="addUser">ログインする</button>
  </div>
  <div>
    <h1>新規登録はこちらから</h1>
    <router-link to='/add'>新規登録のページにいく</router-link>
  </div>
</template>
