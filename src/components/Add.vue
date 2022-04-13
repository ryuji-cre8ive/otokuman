<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from "axios"

export default defineComponent({
  setup() {
    const name= ref("")
    const password = ref("")

    interface User {
      name: string,
      password: string
    }

    const addUser = async () => {
      const obj:User = {
        name: name.value,
        password: password.value
      }
      await axios.post("/api/adduser", obj).then((res) => {
        console.log("res",res)    
        console.log("res.data",res.data)
        if ( res.data.statusCode === 401 ) {
          return alert(res.data.message)
        }
        if ( res.data.statusCode === 200) {
          return alert(res.data.message)
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
    <!-- <h1 v-if=""></h1> -->
  </div>
  <div>
    <h1>新規登録ページ</h1>
    <input type="text" name="name" v-model="name">
    <input type="text" name="password" v-model="password">
    <button @click="addUser">Add User</button>
  </div>
  
</template>
