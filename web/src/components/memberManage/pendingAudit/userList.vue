<template>
  <div class="bg">
    <div class="list">
      <div class="listBg">
        <van-list
          v-model:loading="loading"
          :finished="finished"
          finished-text="没有更多了"
          @load="onLoad"
        >
          <div class="userList">
            <van-config-provider :theme-vars="themeVars">
              <van-cell-group>
                <div v-if="list[0].id">
                  <van-cell
                    :is-link="true"
                    :clickable="true"
                    :center="true"
                    v-for="user in list"
                    :key="user.id"
                    @click="goToAuditPage(user.id)"
                    class="cell"
                  >
                    <template #title>
                      <div class="userName">
                        <a class="name">姓名</a>
                        {{ user.name }}
                      </div>
                    </template>
                    <template #label>
                      <div class="time">申请时间：{{ dateTime(user.date) }}</div>
                    </template>
                  </van-cell>
                </div>
              </van-cell-group>
            </van-config-provider>
          </div>
        </van-list>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { message } from 'ant-design-vue';
import { defineComponent, reactive, ref } from 'vue';
import dayjs from "dayjs";
import utc from 'dayjs/plugin/utc'
export default defineComponent({
  mounted() {
    this.getUsers();
  },
  setup() {
    const users = ref([]);
    const list: any = ref([[]]);
    const loading = ref(false);
    const finished = ref(false);
    const refreshing = ref(false);
    const themeVars = {
      cellFontSize: '17px',
      cellLabelFontSize: '14px',
    };
    return reactive({
      list,
      loading,
      finished,
      refreshing,
      themeVars,
      users,
    });
  },
  methods: {
    dateTime: (date: any) => {
      dayjs.extend(utc)
      return dayjs(date).utc().format('YYYY-MM-DD HH:mm:ss');
    },
    onLoad: function () {
      setTimeout(() => {
        if (this.refreshing) {
          this.list = [[]];
          this.refreshing = false;
          this.getUsers();
        }
        this.loading = false;
        if (this.list.length >= this.users.length) {
          this.finished = true;
        }
      }, 1000);
    },
    onRefresh: function () {
      // 清空列表数据
      this.finished = false;
      // 重新加载数据
      // 将 loading 设置为 true，表示处于加载状态
      this.loading = true;
      this.onLoad();
    },
    goToAuditPage: function (user: any) {
      this.$router.push({
        name: 'auditPage',
        query: {
          userId: user
        }
      })
    },
    getUsers: function () {
      let url = 'api/admin/getUser';
      fetch(url, {
        method: 'GET',
        mode: 'cors',
        cache: 'no-cache',

        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
      })
        .then((v) => {
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
          v.data = v.data.sort((a: any, b: any) => {
            let x = this.dateTime(a['date']);
            let y = this.dateTime(b['date']);
            let aTime = new Date(x).getTime();
            let bTime = new Date(y).getTime();
            return bTime - aTime;
          });

          this.users = v.data;
          console.log('users: ', this.users);
          for (let i = 0; i < this.users.length; i++) {
            this.list[i] = this.users[i];
          }


        });
    },
  },
});
</script>

<style scoped>
.bg {
  background: #f7f7f7;
  width: 100%;
}
.listBg {
  margin-top: 20px;
  min-height: 100vh;
}
@media screen and (max-width: 600px) {
  .cell {
    font-size: 0.4267rem;
  }
  .userList {
    width: auto;
  }
  .userName {
    margin: 0.2667rem;
  }
  .name {
    margin-right: 0.32rem;
  }
  .time {
    font-size: 0.3733rem;
    margin: 0.4rem 0 0 0.2667rem;
  }
}
@media screen and (min-width: 600px) {
  .cell {
    font-size: 16px;
  }
  .userList {
    margin: auto;
    width: 750px;
  }
  .userName {
    margin: 10px;
  }
  .name {
    margin-right: 12px;
  }
  .time {
    margin: 15px 0 0 10px;
    font-size: 14px;
  }
}
</style>
