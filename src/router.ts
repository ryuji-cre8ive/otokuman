// import Vue from 'vue'
import axios from 'axios'
import { createRouter,createWebHistory } from 'vue-router';
import TopPage from './components/Index.vue'
import Hello from './components/HelloWorld.vue'
import Login from './components/Login.vue'
import Add from './components/Add.vue'
import Mypage from './components/Mypage.vue'
import Home from './components/Home.vue'
import UserConfig from './components/UserConfig.vue'
import NewUser from './components/NewUser.vue'

const routes = [
    {
      path: '/',
      name: 'TopPage',
      component: TopPage
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
      path: '/mypage/:id',
      name: 'mypage',
      component: Mypage,
    },
    {
      path: '/config/:id',
      name: 'user-config',
      component: UserConfig,
    },
    {
      path: '/home',
      name: 'home',
      component: Home
    },
    {
      path: '/newuser',
      name: 'newuser',
      component: NewUser
    },
  ]
const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
