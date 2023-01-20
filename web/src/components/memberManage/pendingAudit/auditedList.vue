<template>
  <div class="bground">
    <van-list v-model:loading="loading" :finished="finished" finished-text="没有更多了" @load="onLoad">
      <div class="div1">
        <van-config-provider :theme-vars="themeVars">
          <van-cell-group v-for="user in list" :key="user.id">
            <div>
              <van-cell
                :is-link="true"
                :clickable="true"
                :center="true"
                :to="'/auditPage'"
                @click="run(user.id, user.status)"
                v-if="user.status == 1 || user.status == 2"
              >
                <template #title>
                  <div class="div2">
                    <a class="name">姓名</a>
                    {{ user.name }}
                  </div>
                </template>
                <template #label>
                  <div class="div3">申请时间：{{ dateTime(user.created_at) }}</div>
                </template>
                <template #right-icon>
                  <van-tag
                    color="#bfffbe"
                    text-color="#007420"
                    v-if="user.status == 1"
                    size="large"
                  >已通过</van-tag>
                  <van-tag
                    color="#ffe1e1"
                    text-color="#ad0000"
                    v-if="user.status == 2"
                    size="large"
                  >未通过</van-tag>
                </template>
              </van-cell>
            </div>
          </van-cell-group>
        </van-config-provider>
      </div>
    </van-list>
  </div>
</template>

<style scoped>
@media screen and (max-width: 600px) {
  .div1 {
    min-height: 100vh;
    margin-top: 0.5333rem;
  }
  .bground {
    background: #f7f7f7;
    width: 100%;
  }
  .div3 {
    font-size: 0.3733rem;
    margin: 0.4rem 0 0 0.2667rem;
  }
  .div2 {
    margin: 0.2667rem;
    font-size: 0.4267rem;
  }
  .name {
    margin-right: 0.32rem;
  }
  .text3 {
    margin-left: 0.7467rem;
  }
  .cell {
    background: rgb(16, 108, 219);
  }
  .one {
    background: rgb(41, 122, 245);
  }
}
@media screen and (min-width: 600px) {
  .bground {
    background: #f7f7f7;
    width: 100%;
  }
  .div3 {
    font-size: 13.9988px;
    margin: 15px 0 0 10.0012px;
  }
  .div2 {
    margin: 10.0012px;
    font-size: 16px;
  }
  .div1 {
    min-height: 100vh;
    margin-top: 20px;
  }
  .name {
    margin-right: 12px;
  }

  .text3 {
    margin-left: 28.0013px;
  }
  .cell {
    background: rgb(16, 108, 219);
  }
  .one {
    background: rgb(41, 122, 245);
  }
}
</style>

<script lang="ts">
import dayjs from "dayjs";
import utc from 'dayjs/plugin/utc'
import { ref, reactive, defineComponent } from 'vue';
import { Toast } from 'vant';
import { message } from 'ant-design-vue';
import { RouterLink } from 'vue-router';

export default defineComponent({
  setup() {
    const users = ref([]);
    const list: any = ref([[]]);
    const rate = ref(4);
    const slider = ref(50);
    const active1 = ref(0);
    const loading = ref(false);
    const finished = ref(false);
    const refreshing = ref(false);
    const onClickTab1 = ({ title }) => Toast(title);

    // themeVars 内的值会被转换成对应 CSS 变量
    // 比如 sliderBarHeight 会转换成 `--van-slider-bar-height`
    const themeVars = {
      navBarTextColor: '#FFFFFF',
      navBarIconColor: '#FFFFFF',
      navBarTitleTextColor: '#FFFFFF',
      navBarBackgroundColor: '#1890FF',
      navBarHeight: '1.6rem',
      cellBackgroundColor: '#FFFFF55',
      cellFontSize: '.4533rem',
      cellLabelFontSize: '.3733rem',
      cellBorderColor: '#1890FF',
    };

    return reactive({
      list,
      users,
      loading,
      finished,
      refreshing,
      active1,
      onClickTab1,
      rate,
      slider,
      themeVars,
      datas1: [],
    });
  },
  created() {
    this.onEdit();
  },
  methods: {
    setItem: function (user :any) {
      localStorage.setItem('userInfo', JSON.stringify(user));
    },
    onRefresh: function () {
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;

      this.onLoad();
    },
    onLoad: function () {
      setTimeout(() => {
        if (this.refreshing) {
          this.list = [];
          this.refreshing = false;
          this.onEdit();
        }

        this.loading = false;
        if (this.list.length >= this.users.length) {
          this.finished = true;
        }
      }, 1000);
    },
    dateTime: (date) => {
      dayjs.extend(utc)
      return dayjs(date).utc().format('YYYY-MM-DD HH:mm:ss');
    },
    onEdit() {
      let url = 'api/user/nouser';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: JSON.stringify({}),
      })
        .then((v) => {
          console.log(v);
          return v.json();
        })
        .then((v) => {
          if (v.status == -999) {
            message.warn("账号未登录，请先登录")
            this.$router.push('login')
            return
          }

          if (v.status == -1999) {
            message.warn("管理员账号未登录，请先登录")
            this.$router.push('login')
            return
          }
          if (!v) {
            message.error('网络开小差了，请刷新重试!');
            return;
          }

          if (v.status !== 0) {
            message.error('网络开小差了，请刷新重试！');
            return;
          } else {
            v.data = v.data.sort((a:any, b:any) => {
              let x = this.dateTime(a['created_at']);
              let y = this.dateTime(b['created_at']);
              let aTime = new Date(x).getTime();
              let bTime = new Date(y).getTime();
              return bTime - aTime;
            });
            this.users = v.data;
            console.log('users: ', this.users);
            for (let i = 0; i < this.users.length; i++) {
              this.list[i] = this.users[i];
            }
            console.log('list: ', this.list);
          }
        });
    },
    run(id, sign) {
      if (sign == 2) {
        this.$router.push('/auditFailed/' + id);
      } else if (sign == 1) {
        this.$router.push('/approved/' + id);
      }
    },
  },
});
</script>
