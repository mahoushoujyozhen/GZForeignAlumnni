
<style>
  h1{
    color:#E3F2FD;
  }

  h2{
    color:#BBDEFB;
  }
  h3{
    color:#90CAF9;
  }
  h4{
    color:#64B5F6;
  }
  h5{
    color:#42A5F5;
  }
  h6{
    color:#82B1FF;
  }
  
</style>

# 培训教程

## 1 开发环境搭建
全部都使用最近版本(latest)
前端: vue3
后端: goLang
dbms: postgresql
布署: docker
后端host OS: ubuntu server

### 1.1 各位请安装好以下工具链:

#### 1.1.1 安装nodejs
```
https://github.com/git-for-windows/git/releases/download/v2.34.1.windows.1/Git-2.34.1-64-bit.exe
```

#### 1.1.2 安装git
```
https://github.com/git-for-windows/git/releases/download/v2.34.1.windows.1/Git-2.34.1-64-bit.exe
```
#### 1.1.3 安装 vscode

#### 1.1.4 安装goLang 
```
https://golang.google.cn/dl/go1.17.6.windows-amd64.zip
```

### 1.2 初始化前端
```
yarn init
yarn global add @vue/cli
#或更新
#yarn global upgrade --latest @vue/cli

vue create alumni-association
#manual select features
#select all

Choose a version of Vue.js that you want to start the project with
3.x

[Vue 3] less, babel, typescript, pwa, router, vuex, eslint, unit-jest, e2e-cypress

# 在算法强权的平台上，用户行为轨迹记录着情绪和欲望，最终算法会将它们编织成数据牢笼把你困在信息茧房中，而不自知
------------------------


npm init vue@3
```

## 2 开始编码

### 2.1 清空你的App.vue并替换为以下内容
```html
<template>
  <!-- using for minimap
    http://patorjk.com/software/taag/#p=display&f=Roman&t=code

                          .                                          oooo                .             
                        .o8                                          `888              .o8             
                      .o888oo  .ooooo.  ooo. .oo.  .oo.   oo.ooooo.   888   .oooo.   .o888oo  .ooooo.  
                        888   d88' `88b `888P"Y88bP"Y88b   888' `88b  888  `P  )88b    888   d88' `88b 
                        888   888ooo888  888   888   888   888   888  888   .oP"888    888   888ooo888 
                        888 . 888    .o  888   888   888   888   888  888  d8(  888    888 . 888    .o 
                        "888" `Y8bod8P' o888o o888o o888o  888bod8P' o888o `Y888""8o   "888" `Y8bod8P' 
                                                          888                                         
                                                          o888o                                        
  ============================================================================================================

  -->
  <div class="app-entry"></div>
</template>

<style lang="less" scoped>
/* using for minimap
  ===============================================================================    
                                  .               oooo            
                                .o8               `888            
                      .oooo.o .o888oo oooo    ooo  888   .ooooo.  
                      d88(  "8   888    `88.  .8'   888  d88' `88b 
                      `"Y88b.    888     `88..8'    888  888ooo888 
                      o.  )88b   888 .    `888'     888  888    .o 
                      8oo888P'   "888"     .8'     o888o `Y8bod8P' 
                                      .o..P'                      
                                      `Y8P'      
  ===============================================================================                                      
*/
</style>

<script lang="ts">
/* using for minimap
  ===============================================================================    
                                               .o8            
                                              "888            
                      .ooooo.   .ooooo.   .oooo888   .ooooo.  
                      d88' `"Y8 d88' `88b d88' `888  d88' `88b 
                      888       888   888 888   888  888ooo888 
                      888   .o8 888   888 888   888  888    .o 
                      `Y8bod8P' `Y8bod8P' `Y8bod88P" `Y8bod8P'        

  ===============================================================================
 */

// https://www.jianshu.com/p/d3ffec2a3a0b
import {
  defineComponent,
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  onBeforeUnmount,
  onUnmounted,
  onActivated,
  onDeactivated,
  onErrorCaptured,
  onServerPrefetch,
  onScopeDispose,
  onRenderTracked,
  onRenderTriggered,
  reactive,
  watch,
  ref,
} from 'vue';

export default defineComponent({
  name: 'VueComponentSkeleton',

  components: {},

  props: { title: { type: String, default: '' } },

  computed: {
    // 计算属性
  },
  watch: {
    // 数据监听
  },
  setup() {
    console.log('setup');
    onRenderTracked((e) => {
      // onRenderTracked函数——状态追踪
      // 它会追踪页面上**所有**响应式变量和方法的状态，即我们在setup中return出去的值，
      // 一旦页面有更新，他都会进行追踪，然后生成一个event对象，我们通过event对象来查找程序的问题所在
      // - key 那边变量发生了变化
      // - newValue 更新后变量的值
      // - oldValue 更新前变量的值
      // - target 目前页面中的响应变量和函数
      console.log('onRenderTracked', e);
    });

    onRenderTriggered((e) => {
      // 它不像onRenderTracked函数，这个函数不会跟踪所有值的变化，而是给你变化值的信息，并
      // 且新值和旧值都会给你明确的展示出来
      // - key 那边变量发生了变化
      // - newValue 更新后变量的值
      // - oldValue 更新前变量的值
      // - target 目前页面中的响应变量和函数
      console.log('onRenderTriggered', e);
    });

    onBeforeMount(() => {
      // 在挂载开始之前被调用：相关的 render 函数首次被调用。
      console.log('onBeforeMount');
    });
    onMounted(() => {
      console.log('onMounted');
    });
    onBeforeUpdate(() => {
      // 数据更新时调用，发生在虚拟 DOM 重新渲染和打补丁之前。 你可以
      // 在这个钩子中进一步地更改状态，这不会触发附加的重渲染过程。
      console.log('onBeforeUpdate');
    });
    onUpdated(() => {
      // 由于数据更改导致的虚拟 DOM 重新渲染和打补丁，在这之后会调用该钩子。
      // 当这个钩子被调用时，组件 DOM 已经更新，所以你现在可以执行依赖于 DOM 的操作。
      // 然而在大多数情况下，你应该避免在此期间更改状态，因为这可能会导致更新无限循环
      console.log('onUpdated');
    });
    onBeforeUnmount(() => {
      // 实例销毁之前调用。在这一步，实例仍然完全可用。
      console.log('onBeforeUnmount');
    });
    onUnmounted(() => {
      // Vue 实例销毁后调用。调用后，Vue 实例指示的所有东西都会解绑定，所有的事
      // 件监听器会被移除，所有的子实例也会被销毁。 该钩子在服务器端渲染期间不被调用。
      console.log('onUnmounted');
    });
    onActivated(() => {
      // 激活<keep-alive>组件时
      console.log('onActivated');
    });
    onDeactivated(() => {
      // 离开<keep-alive>组件时
      console.log('onDeactivated');
    });
    onErrorCaptured(() => {
      // 当捕获一个来自子孙组件的异常时激活钩子函数
      console.log('onErrorCaptured');
    });
    onServerPrefetch(() => {
      console.log('onServerPrefetch');
    });
    onScopeDispose(() => {
      console.log('onScopeDispose');
    });

    const counter = ref(0);
    watch(counter, (newValue, oldValue) => {
      if (+newValue !== 6) {
        return;
      }

      console.log('it is 6');
    });

    return reactive({
      greeting: 'hello',
      name: 'mickey',
      gender: 'male',
      counter,
    });
  },

  beforeCreate() {
    // 在实例初始化之后，数据观测(data observer) 和 event/watcher 事件配置之前被调用。
    console.log('beforeCreate');
  },

  created() {
    // 实例已经创建完成之后被调用。在这一步，实例已完成以下的配置：数据观测(data observer)，
    //  属性和方法的运算， watch/event 事件回调。然而，挂载阶段还没开始，$el 属性目前不可见。
    console.log('created');
  },

  methods: {
    // 方法定义
    trial() {
      console.log('trial');
    },
  },
});
</script>
```


### 2.2 启动测试服务

      npm run dev


### 2.3 创建用户注册与登录页面
功能设计思路
1、基于用户名密码登录；
2、如果不存在，则创建，所以不需要用户主动注册；
3、用户登录后添加一个注销账号按钮用以删除自己；
4、用户登录后可以修改自己的名称、口令；
5、管理员列出用户列表后可以CRUD；

