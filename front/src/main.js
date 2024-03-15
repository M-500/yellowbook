import Vue from 'vue'
import App from './App.vue'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

import router from './router';
import './assets/theme/index.css'
import './assets/css/style.css'

import { Loading } from 'element-ui';
Vue.use(Loading.directive);

Vue.config.productionTip = false

// 将全局的echarts对象挂载到Vue原型对象上,在别的组件上使用echarts只需使用this.$echarts即可
Vue.prototype.$myecharts = window.echarts


Vue.use(ElementUI);
new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
