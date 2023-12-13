import { createRouter, createWebHistory } from 'vue-router'
const HomeView = () => import('../views/HomeView.vue')
const QuestionDetailView = () => import('../views/QuestionDetailView.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/questions/:id',
      name: 'question-detail',
      component: QuestionDetailView
    }
  ]
})

export default router
