import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      redirect: '/home'
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/convert',
      name: 'convert',
      component: () => import('../views/ConvertView.vue')
    },
    {
      path: '/edit',
      name: 'edit',
      component: () => import('../views/EditView.vue')
    },
    {
      path: '/ai',
      name: 'ai',
      component: () => import('../views/AIView.vue')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
