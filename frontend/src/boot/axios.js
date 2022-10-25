import axios from 'axios'

// noinspection JSUnusedGlobalSymbols
export default async ({ Vue }) => {
  // noinspection JSUnusedGlobalSymbols
  Vue.prototype.$axios = axios
}
