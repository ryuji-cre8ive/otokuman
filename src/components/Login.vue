<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from "axios"
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  setup() {
    const name= ref("")
    const password = ref("")
    const router = useRouter()
    const route = useRoute()

    interface User {
      name: string,
      password: string
    }

    const addUser = () => {
      const obj:User = {
        name: name.value,
        password: password.value
      }
      axios.post("/login", obj).then(res => {
        if ( res.data.isCorrectUser ) {
          console.log(res.data)
          router.push(`/home/${res.data.Id}`)
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
