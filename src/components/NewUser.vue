<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default defineComponent({
  setup() {
    const router = useRouter()
    interface AddUser {
      name: string
      password: string
    }
    const isSetCookie = ref(false)
    const valid = ref(true)
    const email = ref('')
    const password = ref('')
    const emailRule = [
      (v: any) => !!v || 'email is required',
      (v: string) => /.+@.+\..+/.test(v) || 'Email must be valid'
    ]
    const loading = ref(false)
    const sendEmail = () => {
      loading.value = true
      const params: AddUser = {
        name: email.value,
        password: password.value
      }
      axios.post('/api/adduser', params).then(res => {
        loading.value = false
        alert("登録に成功しました！ログインしてください")
        console.log(res)
        router.push('/login')
      }).catch(err => {
        alert(err)
      })
    }
    return {
      valid,
      email,
      emailRule,
      sendEmail,
      loading,
      password
    }
  },
})
</script>


<template>
  <div align="center">
    <v-card tonal width="600" class="mt-8">
      <v-card-title>新規会員登録</v-card-title>
      <v-card-text>
        <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field v-model="email" :rules="emailRule" label="email"></v-text-field>
          <v-text-field v-model="password" label="password"></v-text-field>
        </v-form>
      </v-card-text>
      <v-divider></v-divider>
      <v-card-actions >
        <v-btn color="success" @click="sendEmail" outlined :loading="loading">登録する</v-btn>
      </v-card-actions>
    </v-card>
  </div>
  
</template>