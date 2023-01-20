import { createApp } from 'vue';
import { createPinia } from 'pinia';
import piniaPersist from 'pinia-plugin-persist'

import App from './App.vue';
import router from './router';

import antDesign from 'ant-design-vue';
import 'ant-design-vue/dist/antd.css';
import * as antIcons from '@ant-design/icons-vue';
import { Dialog } from 'vant';

import Vant from 'vant';
import 'vant/lib/index.css';
import './utils/js/lib-flexible/index';

import axios from 'axios';
import VueAxios from 'vue-axios';

import 'viewerjs/dist/viewer.css';
import VueViewer from 'v-viewer';

const app = createApp(App);

const pinia = createPinia();
pinia.use(piniaPersist);

// app.provide('$message', message);
// app.config.globalProperties.$message = message;

app.use(VueViewer);
app.use(VueAxios, axios);
app.use(Vant);
app.use(Dialog);

app.use(createPinia());
app.use(pinia)

app.use(router);
app.use(antDesign);
const icons: any = antIcons;
for (const i in icons) {
  app.component(i, icons[i]);
}

app.mount('#app');
