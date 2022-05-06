<script lang="ts">
import { defineComponent, computed, onMounted, ref, watchEffect } from 'vue'
import { useStore } from '@/store/index'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

export default defineComponent({
  setup() {
    const store = useStore()
    const router = useRouter()
    const myInfo = computed(() => store.state)
    const articleWrittenByMe = ref<ArticleFromDB[]>([]);
    let currentArticle = ref<Article>()
    let dialog = ref(false)
     

    let isSetCookie = ref(false)
    console.log(myInfo.value.id)

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

    onMounted(() => {
      axios.get('/api/checkCookie').then(() => {
        isSetCookie.value = true
      }).catch((err) => {
        return
      })
      axios.get('/api/getMyArticle').then((res) => {
        console.log(res.data)
        articleWrittenByMe.value = res.data
        console.log(articleWrittenByMe)
      }).catch((err) => {
        console.error(err)
      })

      if (myInfo.value.id == "") {
        store.dispatch('loginWithCookie')
      }
    })

    const goHome = () => {
      axios.get('/api/checkCookie').then(() => {
        router.push('/home')
      }).catch((err) => {
        alert(err)
      })
    }

    const onClickMore = (article: Article) => {
      currentArticle.value = article
      dialog.value = true
    }

    return {
      myInfo,
      goHome,
      isSetCookie,
      articleWrittenByMe,
      currentArticle,
      onClickMore,
      dialog
    }
  },
})
</script>
<template>
  <div v-if="isSetCookie">
    <div>
      <p>this is my profile</p>
      <p>Id: {{myInfo.id}}</p>
      <p>name: {{myInfo.name}}</p>
      <p>password: {{myInfo.password}}</p>
    </div>
    <div>
      <h1>ホームに行く</h1>
      <button @click="goHome">ホームにいく</button>
    </div>

    <div>
      <h1>自分が書いた記事一覧</h1>
      <v-row dense class="pt-10">
        <v-col
          v-for="list in articleWrittenByMe"
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

  
</template>