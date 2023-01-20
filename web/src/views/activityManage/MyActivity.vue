<template>
  <MobileNav title="我的活动"/>
  <MemberNav/>
    <van-list>
      <van-cell v-for="activity in list" >
        <van-row>
          <van-col span="18" @click="toDetail(activity.act_id)">
            <div class="title">{{ activity.act_name }}</div>
            <div class="time">开始时间：{{ activity.start_time }}</div>
            <div class="time">结束时间：{{ activity.end_time }}</div>
            <div class="small-Msg">活动地点：{{ activity.act_address }}</div>
            <div class="small-Msg">报名人数：{{ activity.apply_num }}人</div>
          </van-col>
          <van-col span="6" id="apart-button">
            <van-button
                v-if="activity.isOver == 'true'"
                plain
                type="primary"
                size="small"
                @click="
                cancelActivity(activity);
              "
            >取消报名</van-button
            >
            <van-button
                v-if="activity.isOver == 'false'"
                color="#DDDDDD"
                size="small"
                @click="finishActivity"
            >活动结束</van-button
            >
          </van-col>
        </van-row>
      </van-cell>
    </van-list>
</template>

<style lang="less" scoped>
@baseFont:37.5;

@media screen and (max-width: 600px) {
.van-cell {
  padding: (15rem/@baseFont) (20rem/@baseFont);
}
.title {
  font-size: (16rem/@baseFont);
  font-weight: bolder;
  line-height: 2;
}
.time {
  line-height: 1.5;
  font-size: (14rem/@baseFont);
  color: orange;
}
.small-Msg {
  line-height: 1.5;
  font-size: (14rem/@baseFont);
}

}
@media screen and (min-width: 600px) {
  .van-list{
    width: 750px;
    margin: 50px auto;
  }
  .van-cell {
    padding: 15px 35px;
  }
  .title {
    font-size: 16px;
    font-weight: bolder;
    line-height: 2;
  }
  .time {
    line-height: 1.5;
    font-size: 14px;
    color: orange;
  }
  .small-Msg {
    line-height: 1.5;
    font-size: 14px;
  }
}
#apart-button {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import { Notify } from "vant";
import MemberNav from "@/components/common/MemberNav.vue";
import MobileNav from "@/components/common/MobileNav.vue";
import { message } from "ant-design-vue";

class activity {
  act_id?:number;
  act_name?: string;
  start_time?: string;
  end_time?: string;
  act_address?: string;
  apply_num?: number;
  publish_time?: string;
  isOver?: string;
}

export default defineComponent({
  name: "Enroll",

  components: {
    MobileNav,
    MemberNav,
    // TopNav,
    [Notify.Component.name]: Notify.Component,
  },

  computed: {
    // 计算属性
  },
  watch: {
    // 数据监听
  },
  setup() {
    console.log("setup");
    // const id = window.localStorage.getItem("userID")
    return reactive({
      list: [],
      // id,
    });
  },
  created() {
    // 实例已经创建完成之后被调用。在这一步，实例已完成以下的配置：数据观测(data observer)，
    //  属性和方法的运算， watch/event 事件回调。然而，挂载阶段还没开始，$el 属性目前不可见。
    this.getActMsg();
    console.log("created");
  },

  methods: {
    toDetail(actId){
      this.$router.push(`/activityDetail/${actId}/false`);
    },
    // 方法定义
    trial() {
      console.log("trial");
    },
    getActMsg() {
      let url = `api/user/peractmsg`;
      fetch(url, {
        method: "GET",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        redirect: "follow",
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
          this.list = v.data;
          console.log(this.list);
        })
        .catch((err)=>{
          console.log(err);
          Notify({ type: "warning", message: err });
        });
    },
    back() {
      console.log("it's function cancelActivity");
      this.$router.push("/activityHomePage")
    },
    cancelActivity(activity: any) {
      let query = [];
      query.push(`act_id=${activity.act_id}`);
      // query.push(`user_id=${userid}`);
      let q = query.join("&");
      let method = "DELETE";
      let url = `api/user/cancel?${q}`;
      fetch(url, {
        method,
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        redirect: "follow",
      })
        .then((v) => {
          console.log(v);
          return v.json();
        })
        .then((v) => {
          if (v.status !== 0) {
            Notify({ type: "warning", message: v.msg, position: "top" });
            return;
          }
          Notify({ type: "success", message: "取消成功", position: "top" });
          this.getActMsg();
        })
        .catch((err)=>{
          console.log(err);
          Notify({ type: "warning", message: err });
        });;
    },
    finishActivity() {
      console.log("it's function finishActivity");
      Notify({ type: "warning", message: "活动已结束！！！", position: "bottom" });
    },
    getQrcode() {
      console.log("it's function getQrcode");
      this.$router.push({
        path: "/qrcode",
      });
    },
  },
});
</script>
