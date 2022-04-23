<script lang="ts">
import { defineComponent, ref, computed, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from '@/store/index'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  setup() {
    const store = useStore()
    let isSetCookie = ref(false)

    interface Article {
      title: string
      content: string
      genre: string
      createUser: string
    }

    const title = ref("")
    const content = ref("")
    const genre = ref("")
    const otokuList = [
      {
        title: 'マックのお得',
        content: 'マックのお得な情報を発信しています'
      },
      {
        title: 'セブンのお得',
        content: 'セブンのお得な情報を発信しています'
      },
      {
        title: 'ローソンのお得',
        content: 'ローソンのお得な情報を発信しています'
      },
    ]
    const myInfo = computed(() => store.state)

    onMounted(() => {
      axios.get('/api/checkCookie').then(() => {
        isSetCookie.value = true
      }).catch((err) => {
        console.log(isSetCookie)
        console.error(err)
      })
    })


    const makeArticle = () => {

      const article:Article = {
        title: title.value,
        content: content.value,
        genre: genre.value,
        createUser: myInfo.value.id
      }


      axios.post('/api/addArticle', article).then(res => {

        alert(res.data)
      })
    }

    return { otokuList, makeArticle, title, content, genre, isSetCookie}
  },
})
</script>
<template>
  <div v-if="isSetCookie">
    <h1>this is home page which is posted a lot of articles about otoku</h1>
    <ul>
      <li v-for="o in otokuList" :key="o.title">
        <p>{{o.title}}</p>
        <p>{{o.content}}</p>
      </li>
    </ul>

    <h1>make articles</h1>
    <input type="text" name="title" placeholder="title" v-model="title">
    <textarea name="content" v-model="content"></textarea>
    <select name="genre" v-model="genre">
      <option value="">選択してください</option>
      <option value="セブンイレブン">セブンイレブン</option>
      <option value="ローソン">ローソン</option>
      <option value="ファミリーマート">ファミリーマート</option>
    </select>
    <button @click="makeArticle">submit</button>
  </div>

  <div v-else>
    <h1>You are not logged in</h1>
    <p>Please make sure you logged in</p>
    <router-link to="/login">ログインする</router-link>
  </div>
  
</template>