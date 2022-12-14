// noinspection NpmUsedModulesInstalled
import Vue from 'vue'
// noinspection NpmUsedModulesInstalled
import VueRouter from 'vue-router'

import routes from './routes'

Vue.use(VueRouter)

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation
 */

// noinspection JSUnusedGlobalSymbols
export default function (/* { store, ssrContext } */) {
  return new VueRouter({
    // scrollBehavior: () => ({ x: 0, y: 0 }),
    routes,

    // Leave these as is and change from quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    mode: process.env.VUE_ROUTER_MODE,
    base: process.env.VUE_ROUTER_BASE
  })
}
