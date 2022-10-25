// noinspection NpmUsedModulesInstalled,JSUnusedGlobalSymbols
const routes = [
  {
    path: '/',
    component: () => import('layouts/Index.vue')
  }
]

// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  // noinspection NpmUsedModulesInstalled
  routes.push({
    path: '*',
    component: () => import('pages/Error404.vue')
  })
}

export default routes
