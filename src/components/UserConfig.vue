<script lang="ts">
import { defineComponent, computed, ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from '@/store/index'
import {checkCookie} from './CheckCookie'

export default defineComponent({
  setup() {
    const store = useStore()
    const myInfo = computed(() => store.state)
    const myInitial = ref("")
    const newPhoto = ref("")
    myInitial.value = myInfo.value.img ? myInfo.value.img : myInfo.value.name.substring(0,1)
    const isSetCookie = ref(false)
    const newImage = ref('')

    const uploadNewProfileImage = async () => {
      var formData = new FormData()
      formData.append('image', newImage.value)
      try {
        await axios.post('/api/uploadProfileImage', formData).then((response) => {
        alert(response.data)
        if (response.ok) {
          myInfo.img = newImage.value
        }
      })
      } catch(err) {
        console.error(err)
      }
      
    }

    const registerNewImage = (e: any): void=> {
      e.preventDefault()
      newImage.value = e.target.files[0]
    }
    const links = [
    {
      title: "hi",
      content: "ju"
    },
    {
      title: "hu",
      content: "je"
    },
    {
      title: "he",
      content: "jo"
    },
    ]

    onMounted(async () => {
      const response = await checkCookie()
      console.log("response",response)
      if (response.status === 200) {
        isSetCookie.value = true
      }
      if (myInfo.value.id == '') {
        store.dispatch('loginWithCookie')
      }
    })

    return{ myInfo,
            links,
            myInitial,
            newPhoto,
            newImage,
            uploadNewProfileImage,
            registerNewImage,
          }
    
  },
})
</script>

<template>
  <div class="text-center justify-center">
  <h2>ユーザー設定</h2>
  <v-card width="500" height="800" class="text-center justify-center">
    <v-card-title>Profile</v-card-title>
    <v-card-text>
      <v-avatar size="36px">
        <v-img  :src="myInfo.img ? myInfo.img : ''"></v-img>
      </v-avatar>
      <p>プロフィール画像の更新</p>
      <v-file-input
        accept="image/*"
        label="File input"
        @change="registerNewImage"
        :clearable=true
      ></v-file-input>
    </v-card-text>
    <v-divider></v-divider>
    <v-card-actions>
      <v-btn @click="uploadNewProfileImage" color="secondary">更新する</v-btn>
    </v-card-actions>
  </v-card>
</div>
</template>