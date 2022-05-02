<template>
  <v-app>
    <v-main>
      <Navbar />
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import HelloWorld from './components/HelloWorld.vue'
import Navbar from './components/Navbar.vue'

export default defineComponent({
  name: 'App',

  components: {
    HelloWorld,
    Navbar
  },

  data () {
    return {
      //
    }
  },
})
</script>
