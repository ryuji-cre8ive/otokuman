<script lang="ts">
import { defineComponent, ref } from 'vue'
import CardTop from './CardTop.vue'

export default defineComponent({
  components: {
    CardTop,
  },
  setup() {
    let cardTopDetail = ref([{}])
    cardTopDetail.value = [
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
    ]
    return {cardTopDetail}
  }
})
</script>

<template>
  <div align="center">
    <div>
      <h1>Live a life with <span>Otoku-man</span></h1>
      <p><span>Otoku-man</span>は少しでも<span>お得に</span>生活したい人をサポートします。</p>
      <v-btn>Otokuライフを始める</v-btn>
    </div>

    <div>
      <h2>少しでもお得に生活したい人々をサポート</h2>
      <p>毎日の生活が厳しく、近所のスーパーを慌ただしく開け抜ける日々。<br>少しでもお得に生活したい。スーパーでアイスすら我慢する生活から脱却したい。<br>そんなあなたのためにこのサービスはあります。</p>
      <div class="myProf">
        <div class="myProf-text">
          <p><strong>Ryuji Kokubu</strong></p>
          <p>海外大学でCSを学ぶ大学3年生</p>
        </div>
        <div class="myProf-img">
          <v-icon>mdi-icon</v-icon>
        </div>
        <div class="myProf-icons">
          <v-btn icon="mdi-twitter" variant="text" href="https://twitter.com/ryuji_vlog"></v-btn>
          <v-btn icon="mdi-instagram" variant="text" href="https://instagram.com/ryuji_vlbg"></v-btn>
          <v-btn icon="mdi-github" variant="text" href="https://github.com/ryuji-cre8ive"></v-btn>
        </div>
      </div>
    </div>

    <div>
      <h2>こんな経験はありませんか？</h2>
      <v-row align="center" no-gutters justify="space-around">
        <v-col cols="12" md="5" v-for="c in cardTopDetail" :key="c.title" class="ma-4">
          <CardTop :title="c.title" :img="c.img" :content="c.content"/>
        </v-col>
      </v-row>
    </div>
  </div>
  
</template>
