import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import PrimeVue from 'primevue/config';
import 'primevue/resources/themes/lara-light-teal/theme.css';
import 'primevue/resources/primevue.min.css';
import 'primeicons/primeicons.css';

import Button from 'primevue/button';
import Divider from 'primevue/divider';
import Textarea from 'primevue/textarea';

import VJsoneditor from 'v-jsoneditor'

const app = createApp(App)

app.use(router)
app.use(PrimeVue)
app.use(VJsoneditor)

app
.component('Button', Button)
.component('Divider', Divider)
.component('Textarea', Textarea)

app.mount('#app')
