import axios from 'axios'
import router from '@/router'

const createHttpRequest = (apiUrl: string) => {
  const instance = axios.create({
    baseURL: apiUrl,
    timeout: 12000,
  })

  instance.defaults.headers.common['Content-Type'] = 'application/json'

  instance.interceptors.request.use(
    (config) => {
      return config
    },
    (error) => {
      return Promise.reject(error)
    },
  )

  instance.interceptors.response.use(
    (response) => {
      return response
    },
    (error) => {
      if (error.response && error.response.status === 401) {
        router.push({ name: 'Login' })
      }
      return Promise.reject(error)
    },
  )

  return instance
}

export default createHttpRequest