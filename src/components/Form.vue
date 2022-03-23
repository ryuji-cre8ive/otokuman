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

    const addUser = () => {
      const obj:User = {
        name: name.value,
        password: password.value
      }
      axios.post("/adduser", obj).then(res => {
        console.log(res)
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
    <input type="text" name="name" v-model="name">
    <input type="text" name="password" v-model="password">
    <button @click="addUser">Add User</button>
  </div>
</template>
