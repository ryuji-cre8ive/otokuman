<script lang="ts">
import { defineComponent, ref, computed, onMounted, watchEffect } from 'vue'
import axios from 'axios'
import { useStore } from '@/store/index'
import { useRoute, useRouter } from 'vue-router'
import ArticleCard from './ArticleCard.vue'
import DialogForArticleCard from './DialogForArticleCard.vue'

export default defineComponent({
  components: {
    ArticleCard,
    DialogForArticleCard
  },
  setup(props, context) {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    let isSetCookie = ref(false)
    

    interface Article {
      title: string,
      content: string,
      genre: string,
      createUser: string,
    }
    interface ArticleFromDB extends Article{
      id: string,
      createdAt: string,
    }
    
    let otokuList = ref<ArticleFromDB[]>([]);
    const myInfo = computed(() => store.state)
    let dialog = ref(false)
    let currentArticle = ref<Article>()

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

      axios.get('api/getArticles').then((res) => {
        otokuList.value = res.data.Value
      }).catch((err) => {
        alert(err)
      })
    })

    const onClickMore = (article: Article) => {
      currentArticle.value = article
      dialog.value = true
      console.log(currentArticle)
    }

    const makeDialogFalse = () => {
      dialog.value = false
    }

    

    return { 
      otokuList,
      // makeArticle,
      // title,
      // content,
      // genre,
      isSetCookie,
      dialog,
      onClickMore,
      currentArticle,
      makeDialogFalse
    }
  },
})
</script>
<template>
<v-app>
  <div v-if="isSetCookie">
    <div align=center class="pt-10">
      <h2>お得な情報一覧</h2>
      <v-row dense class="pt-10">
        <v-col
          v-for="list in otokuList"
          :key="list.id"
          class="mx-auto my-5"
          cols="12"
          sm="4"
        >
          <ArticleCard :list="list" @onClickMore="onClickMore"/>
        </v-col>
      </v-row>
      <DialogForArticleCard v-if="currentArticle" :dialog="dialog" :currentArticle="currentArticle" @makeDialogFalse="makeDialogFalse"/>

    </div>

    
  </div>

  <div v-else>
    <h1>You are not logged in</h1>
    <p>Please make sure you logged in</p>
    <router-link to="/login">ログインする</router-link>
  </div>
</v-app>
  
  
</template>

<style>
.v-card--reveal {
  bottom: 0;
  opacity: 1 !important;
  position: absolute;
  width: 100%;
}

.text-card-title {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
} 
</style>