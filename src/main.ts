import { createApp } from 'vue'
import App from './App.vue'
import {store,key} from './store'
import vuetify from './plugins/vuetify'
import { loadFonts } from './plugins/webfontloader'
import router from './router'

loadFonts()

createApp(App)
  .use(store, key)
  .use(vuetify)
  .use(router)
  .mount('#app')
