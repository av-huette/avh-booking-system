import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import TestView from '../views/TestView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/test', component: TestView },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router