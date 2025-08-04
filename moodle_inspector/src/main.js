import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi'

// Font
import "@fontsource/pt-sans";

// Icons
import "@mdi/font/css/materialdesignicons.css";

// Theme
const UofGTheme = {
  dark: false,
  colors: {
    background: '#FFFFFF',
    surface: '#FFFFFF',
    'surface-bright': '#FFFFFF',
    'surface-light': '#EEEEEE',
    'surface-variant': '#424242',
    'on-surface-variant': '#EEEEEE',
    primary: '#005398',
    'primary-darken-1': '#023966',
    secondary: '#006630',
    'secondary-darken-1': '#004520',
    error: '#B30C00',
    info: '#005398',
    success: '#006630',
    warning: '#be4d00',
    'university-blue': '#011451',
    'burgundy': '#7d2239',
    'lavender': '#5b4d94',
    'leaf': '#006630',
    'moss': '#385a4f',
    'pillarbox': '#b30c00',
    'rust': '#bE4d00',
    'sandstone': '#605643',
    'skyblue': '#005398',
    'slate': '#4f5961',
    'thistle': '#951272',
    'reversed': '#f7f7f7',
  },
  variables: {
    'border-color': '#000000',
    'border-opacity': 0.12,
    'high-emphasis-opacity': 0.87,
    'medium-emphasis-opacity': 0.60,
    'disabled-opacity': 0.38,
    'idle-opacity': 0.04,
    'hover-opacity': 0.04,
    'focus-opacity': 0.12,
    'selected-opacity': 0.08,
    'activated-opacity': 0.12,
    'pressed-opacity': 0.12,
    'dragged-opacity': 0.08,
    'theme-kbd': '#212529',
    'theme-on-kbd': '#FFFFFF',
    'theme-code': '#F5F5F5',
    'theme-on-code': '#000000',
  }
}

const app = createApp(App)

const vuetify = createVuetify({
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    },
  },
  theme: {
    defaultTheme: 'UofGTheme',
    themes: {
      UofGTheme,
    },
  },
})

app.use(createPinia())
app.use(router)
app.use(vuetify)

app.mount('#app')
