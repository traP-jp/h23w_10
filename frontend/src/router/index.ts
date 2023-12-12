import { createRouter, createWebHistory } from 'vue-router'
const HomeView = () => import('../views/HomeView.vue')
const AboutView = () => import('../views/AboutView.vue')
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
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/questions/:id',
      name: 'question-detail',
      component: QuestionDetailView
    }
  ]
})

export default router
