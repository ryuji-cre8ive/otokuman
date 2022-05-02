<script lang="ts">
import { defineComponent, ref, computed, onMounted, watchEffect } from 'vue'
import axios from 'axios'
import { useStore } from '@/store/index'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({
  setup() {
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

    

    // const makeArticle = () => {

    //   const article:Article = {
    //     title: title.value,
    //     content: content.value,
    //     genre: genre.value,
    //     createUser: myInfo.value.id
    //   }


    //   axios.post('/api/addArticle', article).then(res => {
    //     alert(res.data.message)
    //     location.reload()
    //   })
    // }

    const onClickMore = (article: Article) => {
      currentArticle.value = article
      dialog.value = true
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
      currentArticle
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
          <v-card
            class="mx-auto"
            max-width="344"
          >
            <v-card-text>
              
              <p class="text-h4 text--primary">
                {{list.title}}
              </p>
            </v-card-text>
            <v-card-actions>
              <v-btn
                variant="text"
                color="teal-accent-4"
                @click.stop="onClickMore(list)"
              >
                Learn More
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
      <v-dialog 
        v-model="dialog"
        v-if="currentArticle"
        max-width="600px"
      >
        <v-card max-width="600" class="pa-5">
          <v-card-title>{{currentArticle.title}}</v-card-title>
          <v-card-text>
            <v-row>
              {{currentArticle.content}}
            </v-row>
            
          </v-card-text>
          <v-card-actions>
            <v-btn @click="dialog = false" color="error">
              Close
              <v-icon>mdi-close</v-icon>
            </v-btn>
            <v-spacer></v-spacer>
            <v-chip>
              {{currentArticle.createdAt}}
              <v-icon>mdi-calendar-month</v-icon>
            </v-chip>
          </v-card-actions>
        </v-card>
      </v-dialog>
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
</style>