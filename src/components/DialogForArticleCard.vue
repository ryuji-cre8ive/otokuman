<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    dialog: Boolean,
    currentArticle: Object,
  },
  emits: ['makeDialogFalse'],
  setup(){
    
  }
  
})
</script>


<template>
  <v-dialog 
    v-model="dialog"
    
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
          <v-btn @click="$emit('makeDialogFalse')" color="error">
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
</template>