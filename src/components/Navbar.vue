<script lang="ts">
import { defineComponent, computed , ref, watchEffect} from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from '@/store/index'
import axios from 'axios'

export default defineComponent({
  setup() {
    const route = useRoute()
    const router = useRouter()
    const store = useStore()
    const items = ['セブンイレブン','ファミリーマート','ローソン', 'その他']
    let dialog = ref(false)
    let choiceGenreItem = ref("")
    let isHome = ref(false)
    const title = ref("")
    const content = ref("")
    const genre = ref("")

    let myInfo = computed(() => store.state)

    const goHome = () => {
      router.push('/')
    }
    const goAbout = () => {
      router.push('/about')
    }
    const goNews = () => {
      router.push('/news')
    }
    const goLogin = () => {
      router.push('/login')
    }

    const goMypage = (id:string) => {
      router.push('/mypage/' + id)
    }

    const goHomeWithLogin = () => {
      router.push('/home')
    }

    const choiceGenre = (genre: string) => {
      choiceGenreItem.value = genre
      console.log(choiceGenreItem.value)
    }

    interface Article {
      title: string,
      content: string,
      genre: string,
      createUser: string,
    }

    const makeArticle = () => {

      const article:Article = {
        title: title.value,
        content: content.value,
        genre: choiceGenreItem.value,
        createUser: myInfo.value.id
      }

      console.log("genre: " + article.genre)


      axios.post('/api/addArticle', article).then(res => {
        alert(res.data.message)
        location.reload()
      })
    }

    watchEffect(() => {
      if ( route.path == '/home' ) {
        isHome.value = true
      } else {
        isHome.value = false
      }
    })


    return {
      goHome,
      goAbout,
      goNews,
      goLogin,
      myInfo,
      goMypage,
      goHomeWithLogin,
      items,
      dialog,
      choiceGenre,
      isHome,
      makeArticle,
      title,
      content,
      genre
    }
  },
})
</script>


<template>
  <v-app-bar
    color="yellow"
    dense
    app
    class="px-16"
  >
    <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title color="white" class="pl-4">OTOKU-MAN</v-toolbar-title>

      <v-spacer></v-spacer>
      <div class="px-10">
        <v-btn @click="goHome" v-if="myInfo.id == ''">
          <v-icon>mdi-home</v-icon>
          <p>home</p>
        </v-btn>

        <v-btn @click="goHomeWithLogin" v-else>
          <v-icon>mdi-home</v-icon>
          <p>home</p>
        </v-btn>


        <v-btn @click="goAbout">
          <v-icon>mdi-check-bold</v-icon>
          <p>about</p>
        </v-btn>

        <v-btn @click="goNews">
          <v-icon>mdi-comment-text</v-icon>
          <p>news</p>
        </v-btn>
      </div>
      
      <v-divider vertical></v-divider>

      <v-btn @click="goLogin" v-if="myInfo.id == ''">
        <p>ログイン</p>
        <v-icon>mdi-login</v-icon>
      </v-btn>

      <div v-else>
        <v-btn @click="goMypage(myInfo.id)">
          <p>マイページ</p>
          <v-icon>mdi-login</v-icon>
        </v-btn>

        <v-btn @click="dialog = !dialog" v-if="isHome">
            <p>記事を書く</p>
            <v-icon>mdi-file-document-multiple-outline</v-icon>
          </v-btn>
        <v-dialog v-if="dialog" v-model="dialog" persistent max-width="600px">
          
          <v-card width="600px">
            <v-container>
              <v-row jsutify="space-around">
                <v-col cols="5">
                  <v-card-title>Write new article</v-card-title>
                </v-col>
                <v-col>
                  <v-spacer></v-spacer>
                </v-col>
                <v-col cols="3" sm="3" md="3" class="ma-2">
                    <v-menu offset-y>
                      <template v-slot:activator="{ props }">
                        <v-btn v-bind="props">
                          Genre
                        </v-btn>
                      </template>
                      <v-list>
                        <v-list-item v-for="(item, index) in items" :key="index" :value="item" active-color="primary">
                          <v-list-item-title @click="choiceGenre(item)" v-model="genre">{{item}}</v-list-item-title>
                        </v-list-item>
                      </v-list>
                    </v-menu>
                  </v-col>
              </v-row>
            </v-container>
            <v-card-text>
              <v-container>
                <v-row>
                  <v-col cols="12" sm="12" md="12">
                    <v-text-field label="title" required v-model="title"></v-text-field>
                  </v-col>
                  <v-col cols="12" sm="12" md="12">
                    <v-textarea label="content" required v-model="content"></v-textarea>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
            <v-divider></v-divider>
            <v-card-actions class="ma-2">
              <v-btn color="red"  @click="dialog = !dialog">
                cancel
                <v-icon>mdi-cancel</v-icon>
              </v-btn>
              <v-spacer></v-spacer>
              <v-btn color="green" @click="makeArticle">
                submit
                <v-icon>mdi-comment-text</v-icon>
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        

      </div>

      
      
  </v-app-bar>
</template>