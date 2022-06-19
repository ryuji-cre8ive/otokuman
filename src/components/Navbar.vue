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
    let drawer = ref(false)
    const title = ref("")
    const content = ref("")
    const genre = ref("")
    const links = [{
      title: "ログアウト",
      link: '/logout',
      icon: 'mdi-logout'
    }]

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

    const goConfig = (id:string) => {
      router.push('/config/' + id)
    }

    // const choiceGenre = (genre: string) => {
    //   choiceGenreItem.value = genre
    //   console.log(choiceGenreItem.value)
    // }

    interface Article {
      title: string,
      content: string,
      genre: string,
      createUser: string,
    }

    interface User {
      id: string,
      name: string,
      password: string
    }

    const makeArticle = () => {

      const article:Article = {
        title: title.value,
        content: content.value,
        genre: genre.value,
        createUser: myInfo.value.id
      }

      console.log("genre: " + article.genre)


      axios.post('/api/addArticle', article).then(res => {
        alert(res.data.message)
        location.reload()
      })
    }

    const logout = () => {
      const logoutParam:User = {
        id: myInfo.value.id,
        name: myInfo.value.name,
        password: myInfo.value.password
      }
      axios.post('/api/logout', logoutParam).then(res => {
        alert(res.data)
        return location.href = "/"
      }).catch(err => {
        alert(err)
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
      // choiceGenre,
      isHome,
      makeArticle,
      title,
      content,
      genre,
      drawer,
      links,
      logout,
      goConfig
    }
  },
})
</script>


<template>
<v-navigation-drawer
  v-model="drawer"
  temporary
  class="d-none d-md-flex"
  >
    <v-list-item
      prepend-avatar="https://randomuser.me/api/portraits/men/78.jpg"
      :title="myInfo.name"
    ></v-list-item>

      <v-divider></v-divider>
      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-view-dashboard" title="ホーム" value="home" @click="goHomeWithLogin" class="py-3"></v-list-item>
        <v-list-item prepend-icon="mdi-check-bold" title="Otokumanについて" value="about" @click="goAbout" class="py-3"></v-list-item>
        <v-list-item prepend-icon="mdi-comment-text" title="新着情報" value="news" @click="goNews" class="py-3"></v-list-item>
        <v-list-item prepend-icon="mdi-login" title="マイページ" value="mypage" @click="goMypage(myInfo.id)" class="py-3"></v-list-item>
        <v-list-item prepend-icon="mdi-file-document-multiple-outline" title="記事を書く" value="mypage" @click="dialog = !dialog" class="py-3"></v-list-item>
        <v-list-item prepend-icon="mdi-cog" title="設定" value="config" @click="goConfig(myInfo.id)" class="py-3"></v-list-item>
        <v-divider></v-divider>
        <v-btn color="error" class="ma-6" @click="logout">ログアウトする</v-btn>
      </v-list>
    </v-navigation-drawer>
  <v-app-bar
    color="yellow"
    dense
    app
  >
  
    <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title color="white" class="pl-4 d-none d-sm-flex">OTOKU-MAN</v-toolbar-title>

      <v-spacer></v-spacer>
      <div class="px-10 d-sm-flex">
        <v-btn @click="goHome" v-if="myInfo.id == ''">
          <v-icon>mdi-home</v-icon>
          <p>home</p>
        </v-btn>

        <v-btn @click="goHomeWithLogin" v-else>
          <v-icon>mdi-home</v-icon>
          <p>home</p>
        </v-btn>


        <v-btn @click="goAbout" class="d-none d-sm-flex">
          <v-icon>mdi-check-bold</v-icon>
          <p>about</p>
        </v-btn>

        <v-btn @click="goNews" class="d-none d-sm-flex">
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
          
          <v-card max-width="700px">
            <v-container>
              <v-row jsutify="space-around">
                <v-col cols="12" sm="5">
                  <v-card-title>Write new article</v-card-title>
                </v-col>
                <v-col>
                  <v-spacer></v-spacer>
                </v-col>
                <v-col cols="12" sm="7" class="ma-2">
                  <v-select v-model="genre" :items="items" label="ジャンル">
                  </v-select>
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