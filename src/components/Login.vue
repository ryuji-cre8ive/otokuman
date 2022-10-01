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
        const myprofile:afterLoginUser = res.data
        store.dispatch('login', myprofile)
        router.push(`/mypage/${myprofile.Id}`)
      }).catch(err => {
        alert(err)
      })
    }

    const showPassword = ref(false)

    return { 
      name,
      password,
      login,
      showPassword
    }
  },
})
</script>

<template>
    
    <v-main class="ma-5">
      <v-card max-width="600"  class="ma-auto">
        <v-card-title class="pt-10">
          <h1 class="display-1">ログイン</h1>
        </v-card-title>
        <v-card-text>
          <v-form class="pt-7">
            <v-text-field label="メールアドレス" prepend-icon="mdi-account-circle" v-model="name" />
            <v-text-field :type="showPassword ? 'text' : 'password'" label="パスワード" prepend-icon="mdi-lock" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'" v-model="password" @click:append="showPassword = !showPassword" />
          </v-form>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-btn color="success" @click="login">ログイン</v-btn>
        </v-card-actions>
      </v-card>
      <div align="center">
        <v-row>
          <v-col>
            <router-link to="/newuser">新規会員登録はこちら</router-link>
          </v-col>
        </v-row>
      </div>
    </v-main>
</template>