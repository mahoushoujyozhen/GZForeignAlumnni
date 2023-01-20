<template>
  <MemberNav />
  <div class="entry">
    <div v-for="u in Work">
      <div class="card" v-if="u.position != '<nil>' && u.position != ''" @click="jumpWork(u.id)">
        <div class="content_top">
          <span class="career">{{ u.position }}</span>
          <span class="time">{{ u.begin_Time }}-{{ u.end_Time }}</span>
        </div>
        <div class="details">
          <span>从事专业：</span>
          <span>{{ u.professional }}</span>
        </div>
        <div class="details">
          <span>机构名称：</span>
          <span>{{ u.organization }}</span>
        </div>
      </div></div>

    <div v-for="n in Edu">
      <div v-if="n.degree != '' && n.degree != '<nil>'" class="card" @click="jumpEdu(n.id)">
        <div class="content_top">
          <span class="career">{{ n.degree }}</span>
          <span class="time">{{ n.begin_time }}-{{ n.end_time }}</span>
        </div>
        <div class="details">
          <span>学科专业：</span>
          <span>{{ n.major }}</span>
        </div>
        <div class="details">
          <span>学校名称：</span>
          <span>{{ n.college }}</span>
        </div>
      </div>
    </div>
    <van-config-provider :theme-vars="themeVars">
      <div class="button-wrapper">
        <van-button
          icon="manager-o"
          class="button-add"
          type="primary"
          to="/WorkResume"
          >新增工作简历</van-button
        >
        <van-button
          icon="newspaper-o"
          class="button-add"
          type="primary"
          to="/EdResume"
          >新增教育简历</van-button
        >
      </div>
    </van-config-provider>
  </div>
</template>

<style lang="less" scoped>
@media screen and (max-width: 600px) {
  .entry {
    padding-top: 0.53rem;
    .card {
      background-color: white;
      margin-top: 1vh;
      padding: 0.4rem 0.53rem;
      .content_top {
        display: flex;
        justify-content: space-between;
        .career {
          font-size: 0.48rem;
          font-weight: 600;
        }
        .time {
          color: #919090;
          margin-left: 3vw;
          font-size: 0.37rem;
        }
      }
      .details {
        font-size: 0.43rem;
      }
    }

    .button-wrapper {
      margin: auto;
      display: flex;
      flex-direction: column;
      width: 4rem;
      padding-bottom: 0.53rem;
      .van-button {
        margin-top: 0.5rem;
      }
    }
  }
}

@media screen and (min-width: 600px) {
  .entry {
    padding-top: 50px;
    width: 750px;
    margin: auto;
    .card {
      background-color: white;
      margin-top: 10px;
      padding: 15px 20px;
      .content_top {
        display: flex;
        justify-content: space-between;
        .career {
          font-size: 18px;
        }
        .time {
          color: #919090;
          margin-left: 3vw;
          font-size: 14px;
        }
      }
      .details {
        font-size: 16px;
      }
    }

    .button-wrapper {
      margin: auto;
      width: 400px;
      display: flex;
      justify-content: space-evenly;
      padding-bottom: 50px;
      .button-add {
        margin-top: 35px;
        font-size: 16px;
      }
    }
  }
}
</style>

<script>
import { defineComponent, onMounted, reactive } from "vue";

import MemberNav from "@/components/common/MemberNav.vue";
import MobileNav from "@/components/common/MobileNav.vue";
import { message } from "ant-design-vue";
import { useRouter } from "vue-router";
export default defineComponent({
  name: "VueComponentSkeleton",
  props: {},
  components: {
    MobileNav,
    MemberNav,
  },

  setup() {
    const router = useRouter();
    let Work = reactive([]);
    let Edu = reactive([]);

    //自定义设置button中的icon组件的样式
    const themeVars = {
      buttonIconSize: "16px",
    };
    onMounted(() => {
      let url = "api/user/listmsg";
      fetch(url, {
        method: "POST",
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
            router.push('login')
            return
          }

          if(v.status == -1999){
            message.warn("管理员账号未登录，请先登录")
            router.push('login')
            return
          }
          if (v.status != 0) {
            console.log("err:" + v.status + v.msg);
            message.info("信息展示失败，请重试");
            return;
          }


          Work.push(...v.work_data.data);
          Edu.push(...v.edu_data.data);
          return;
        });
    });

    return {
      Work,
      Edu,
      themeVars,
    };
  },

  methods: {
    jumpEdu(id) {
      this.$router.push({
        name: `DetailsEd`,
        query: {
          details_id: id,
        },
      });
    },
    jumpWork(id) {
      this.$router.push({
        name: `DetailsWork`,
        query: {
          details_id: id,
        },
      });
    },
    back() {
      this.$router.push("/memberCenter");
    },
  },
});
</script>
