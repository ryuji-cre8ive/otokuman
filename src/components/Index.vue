<script lang="ts">
import { defineComponent, ref } from 'vue'
import CardTop from './CardTop.vue'

export default defineComponent({
  components: {
    CardTop,
  },
  setup() {
    const cardTopDetail = ref([
      {
        title: 'セブンのお得な情報について',
        img: '/src/assets/images/y1169.png',
        content: 'プライチ情報は公式アプリを見れば確認できるけど、それ以外の情報も欲しい。'
      },
      {
        title: 'ローソンのお得な情報について',
        img: '/src/assets/images/y1170.png',
        content: 'ローソンのお得な情報もその都度確認したい！'
      },
      {
        title: 'ファミマのお得な情報について',
        img: '/src/assets/images/y0861.png',
        content: 'ファミチキが安くなってたり、セブンのプライチ的な情報を早く知りたい'
      },
      {
        title: 'カフェのお得な情報について',
        img: '/src/assets/images/y0979.png',
        content: 'スタバやタリーズなどのお得な情報や、個人店でのキャンペーンなどの情報を早くキャッチアップしたい'
      },
      {
        title: '旅行のお得な情報について',
        img: '/src/assets/images/y1156.png',
        content: 'お得に旅行したり、政府の打ち出す政策をいち早く知りたい！'
      },
    ])
    return {cardTopDetail}
  }
})
</script>

<template>
  <div align="center">
    <div class="top">
      <div class="top-content">
        <h1>Live a life with <br><span>Otoku-man</span></h1>
        <p><span>Otoku-man</span>は少しでもお得に生活したい人をサポートします。</p>
        <v-btn class="mt-8" color="yellow lighten-3">Otokuライフを始める</v-btn>
      </div>
      
    </div>

    <div class="support">
      <h2>少しでもお得に生活したい人々をサポート</h2>
      <p>毎日の生活が厳しく、近所のスーパーを慌ただしく開け抜ける日々。<br>少しでもお得に生活したい。スーパーでアイスすら我慢する生活から脱却したい。<br>そんなあなたのためにこのサービスはあります。</p>
      <div class="myProf d-flex justify-center">
        <div class="myProf-text">
          <p><strong>Ryuji Kokubu</strong></p>
          <p>海外大学でCSを学ぶ大学3年生</p>
        </div>
        <div class="myProf-img">
          <v-avatar size="64px" class="ma-4">
            <img width="56" height="56" src="../assets/images/ryuji_icon.png" alt="ryuji-pic">
          </v-avatar>
        </div>
        
      </div>
      <div class="myProf-icons">
          <v-btn icon="mdi-twitter" variant="text" href="https://twitter.com/ryuji_vlog"></v-btn>
          <v-btn icon="mdi-instagram" variant="text" href="https://instagram.com/ryuji_vlbg"></v-btn>
          <v-btn icon="mdi-github" variant="text" href="https://github.com/ryuji-cre8ive"></v-btn>
        </div>
    </div>

    <div class="experience">
      <h2>こんな経験はありませんか？</h2>
      <!-- <v-row align="center" no-gutters justify="space-around">
        <v-col cols="12" md="3" v-for="c in cardTopDetail" :key="c.title" class="ma-4">
          <CardTop :title="c.title" :img="c.img" :content="c.content"/>
        </v-col>
      </v-row> -->
      <v-carousel hide-delimiters height="400">
        <v-carousel-item v-for="c in cardTopDetail" :key="c.title">
          <CardTop :title="c.title" :content="c.content" :img="c.img" class="mt-15"/>
        </v-carousel-item>
      </v-carousel>
      <p>このような情報を常々キャッチアップしていくのはとても骨が折れる作業だと思います。<br>なのでこのサービスはお得な情報共有を目的とした<strong>お得情報特化型SNS</strong>です。</p>
      <v-btn class="mt-8" color="yellow lighten-3">Otokuライフを始める</v-btn>
    </div>

    <v-footer
      padless=true
    >
      <v-card
        flat
        tile
        width="100%"
        color="orange"
      >

        <v-divider></v-divider>

        <v-card-text class="white--text">
          {{ new Date().getFullYear() }} — <strong>Otokuman</strong>
        </v-card-text>
      </v-card>
    </v-footer>
  </div>

<link href="https://fonts.googleapis.com/css?family=Londrina+Shadow" rel="stylesheet">
</template>
<style lang="scss">
  
  .top {
    width: 100%;
    min-height: 100vh;
    position: relative;
    
    .top-content{
      margin: 0;
      position: absolute;
      top: calc(20% - 1em);
      width: 100%;
      text-align: center;
      h1{
        font-family: 'Londrina Shadow', cursive;
        font-size: 8vw;
        text-shadow: 3px 3px 0 #413f3f;
        span{ 
          color: yellowgreen
        }
      }
      p{
        font-size: 4vh;
        span{
          font-family: 'Londrina Shadow', cursive;
          color: yellow;
          text-shadow: 1px 1px 0 #000
        }
      }
    }
  }
  .support{
    h2 { 
      margin: 30px 0;
    }
    p {
      margin: 20px 0;
    }
  }

  .experience {
    width: 60vh;
    margin: 10vh 0;
  }
</style>