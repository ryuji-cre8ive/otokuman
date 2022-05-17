<script lang="ts">
import { defineComponent, ref } from 'vue'


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

export default defineComponent({
  emits: ['onClickMore'],
  setup(props, context) {
    let currentArticle = ref<Article>()
    let dialog = ref(false)
    

    const onClickMore = (article: Article) => {
      context.emit('onClickMore', article)
    }

    const setDialogToFalse = () => {
      dialog.value = false
    }

    return {
      onClickMore,
      currentArticle,
      dialog,
      setDialogToFalse
    }
  },
  props: {
    list: Object
  }
})
</script>


<template>
  <v-card
    class="pt-4"
    max-width="320"
    height="150"
  >
    <v-card-text>
      <p class="text-h5 text--primary text-card-title">
        {{list.title}}
      </p>
    </v-card-text>
    <v-card-actions>
      <v-btn
        class="mt-4"
        variant="text"
        color="teal-accent-4"
        @click.stop="onClickMore(list)"
      >
        Learn More
      </v-btn>
    </v-card-actions>
  </v-card>
  <DialogForArticleCard :currentArticle="currentArticle" :dialog="dialog"/>

</template>