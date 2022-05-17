<template>
  <v-card width="400" height="240" class="pa-2">
    <v-card-title>{{title}}</v-card-title>
    <v-card-content>
      <v-row>
        <v-col cols="5"><v-img :src="img" max-height="160"></v-img></v-col>
        <v-col cols="7" align="left">{{content}}</v-col>
      </v-row>
    </v-card-content>
  </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    title: String,
    img: String,
    content: String,
  },
})
</script>
