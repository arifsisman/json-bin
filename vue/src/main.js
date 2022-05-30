import { createApp } from 'vue'
import App from './App.vue'
import router from './router'


import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/lara-light-teal/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Button from 'primevue/button';
import Textarea from 'primevue/textarea';
import Toast from 'primevue/toast';
import ToastService from 'primevue/toastservice';

const app = createApp(App)

app
  .use(router)
  .use(PrimeVue)
  .use(ToastService)

app
  .component('Button', Button)
  .component('Textarea', Textarea)
  .component('Toast', Toast)
  .component('ToastService', ToastService)

app.mount('#app')
