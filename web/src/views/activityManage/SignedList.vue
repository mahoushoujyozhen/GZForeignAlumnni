<template>
  <ManageNav/>
  <MobileNav title="签到列表"/>
  <div class="container">
    <van-row class="list">
      <van-col span="12">
        <span>姓名</span>
        <van-list>
          <van-cell v-for="user in users" :key="user.name" :title="user.name" />
        </van-list>
      </van-col>
      <van-col span="12">
        <span>电话</span>
        <van-cell
          v-for="user in users"
          :key="user.phone"
          :title="user.phone"
      /></van-col>
    </van-row>
  </div>
</template>

<style lang="less" scoped>
@baseFont: 37.5;
@media screen and (max-width: 600px){
  .container {
    margin: 0.53rem auto;
    background-color: white;
    .list {
      text-align: center;
      padding: 0.53rem;
      .van-col{
        > span{
          color:#1890ff;
          font-size: 18px;
          margin-bottom: 10px;
        }
      }
    }
  }
}
@media screen and (min-width: 600px){
  .container {
    margin: 50px auto;
    width: 750px;
    background-color: white;
    .list {
      text-align: center;
      padding: 35px;
      .van-col{
        > span{
          color:#1890ff;
          font-size: 18px;
          margin-bottom: 10px;
        }
      }
    }
  }

}

</style>

<script lang="ts">
import { defineComponent, reactive } from 'vue';
import { Notify } from 'vant';
import ManageNav from "@/components/common/ManageNav.vue";
import MobileNav from "@/components/common/MobileNav.vue";
import { message } from 'ant-design-vue';
class user {
  name?: string;
  phone?: string;
}

export default defineComponent({
  components: {MobileNav, ManageNav},
  setup() {
    return reactive({
      users: [],
    });
  },

  created() {
    this.getUserMsg();
  },
  methods: {
    // 方法定义
    back() {
      this.$router.push('/activityManage');
    },
    getUserMsg() {
      let url = 'api/admin/signinList';
      console.log(url);
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: JSON.stringify({
          act_id: this.$route.params.act_id,
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
            if(v.status == -999){
              message.warn("账号未登录，请先登录")
              this.$router.push('login')
              return

            }

            if(v.status == -1999){
              message.warn("管理员账号未登录，请先登录")
              this.$router.push('login')
              return
            }
          if (!v) {
            Notify({ type: 'danger', message: 'null response' });
            return;
          }

          if (v.status !== 0) {
            Notify({ type: 'danger', message: v.msg });
            return;
          }
          this.users = v.data;
          console.log(this.users);
        });
    },
  },
});
</script>
