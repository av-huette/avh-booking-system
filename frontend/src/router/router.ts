import { createWebHistory, createRouter } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import BookingView from '../views/BookingView.vue'

const routes = [
  { path: '/', component: HomeView },
  { path: '/booking', component: BookingView },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router