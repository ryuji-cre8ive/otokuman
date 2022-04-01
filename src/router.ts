import Vue from 'vue'
import { createRouter,createWebHistory } from 'vue-router';
import Hello from './components/HelloWorld.vue'
import Login from './components/Login.vue'
import Add from './components/Add.vue'
import Home from './components/Home.vue'

const routes = [
    {
      path: '/',
      name: 'hello',
      component: Hello
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/add',
      name: 'addUser',
      component: Add
    },
    {
      path: '/home/:id',
      name: 'home',
      component: Home
    }
  ]
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router