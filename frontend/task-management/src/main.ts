import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { createVuetify } from 'vuetify'
import 'vuetify/styles'
import '@mdi/font/css/materialdesignicons.css'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import VueApexCharts from 'vue3-apexcharts'
const app = createApp(App)
const vuetify = createVuetify({
  components,
  directives,
})

app.use(createPinia())

app.component('apexchart', VueApexCharts)
app.use(router)
app.use(vuetify)
app.mount('#app')
