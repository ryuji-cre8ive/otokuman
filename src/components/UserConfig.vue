<script lang="ts">
import { defineComponent, computed, ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from '@/store/index'

export default defineComponent({
  setup() {
    const store = useStore()
    const myInfo = computed(() => store.state)
    const myInitial = ref("")
    const newPhoto = ref("")
    myInitial.value = myInfo.value.img ? myInfo.value.img : myInfo.value.name.substring(0,1)
    const isSetCookie = ref(false)
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

    onMounted(() => {
      axios.get('/api/checkCookie').then(() => {
        isSetCookie.value = true
      }).catch((err) => {
        console.log(isSetCookie)
        console.error(err)
      })
      if (myInfo.value.id == '') {
        store.dispatch('loginWithCookie')
      }
    })

    return{ myInfo,
            links,
            myInitial,
            newPhoto,
          }
    
  },
})
</script>

<template>
  <v-card>
    <v-card-title>Profile</v-card-title>
    <v-card-text>
      <v-avatar size="36px">
        <v-img :src="myInfo.img" v-if="myInfo.img"></v-img>
        <v-icon v-if="!myInfo.img">{{myInitial}}</v-icon>
      </v-avatar>
      <p>プロフィール画像の更新</p>
      <v-file-input></v-file-input>
    </v-card-text>
  </v-card>
  
</template>