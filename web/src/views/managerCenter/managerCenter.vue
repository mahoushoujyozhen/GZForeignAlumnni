<template>
  <div class="headMobile">
    <div class="headTitle"><span>管理中心</span></div>
    <div class="headFunc" @click="logout"><span>注销</span></div>
  </div>
  <ManageNav />
  <div class="whole">
    <div class="viewsize">
      <div class="card">
        <div class="cardtitle">
          <img src="../../assets/img/广州欧美同学会.png" class="cardtitle" />
        </div>
      </div>
      <div class="margin">
        <MFunction></MFunction>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.headMobile {
  width: 100%;
  height: 1.06rem;
  position: relative;
  justify-content: space-between;
  background-color: #ffffff;
  border-bottom: 1px solid #efefef;
  .headTitle {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
    font-size: 0.48rem;
    line-height: 1.06rem;
  }
  .headFunc {
    position: absolute;
    right: 0.53rem;
    font-size: 0.43rem;
    color: #1890ff;
    line-height: 1.06rem;
  }
}
.margin {
  margin: 0.5333rem;
  margin-top: 0px;
}
.card {
  margin: 0px auto 0.6rem auto;
  /* height: auto; */
  width: 8.7rem;
  box-sizing: border-box;
  border-radius: 0.1333rem;
  background-size: cover;
  text-align: center;
  box-shadow: 0px 2px 7px #b9bcbe;
  background-image: linear-gradient(to bottom, #4481eb 10%, #04affec9 90%);
  padding: 0.4rem 0;
}

.cardtitle {
  /* font-size: 1.0667rem; */
  /* color: whitesmoke; */
  /* font-family: 'STXingkai'; */
  /* line-height: 1rem; */
  img {
    width: 100%;
  }
}
.whole {
  background-color: #f7f7f7;
  padding-top: 0.6rem;

  margin-top: 0px;
  min-height: 100vh;
  overflow: hidden;
}

@media screen and (min-width: 600px) {
  .headMobile {
    display: none;
  }
  .viewsize {
    max-width: 900px;
    margin: auto;
  }
  .margin {
    margin: 20px;
    margin-top: 0px;
  }
  .card {
    margin: 0px auto 22.5px auto;
    /* height: auto; */
    max-width: 670px;
    box-sizing: border-box;
    border-radius: 5px;
    background-size: cover;
    text-align: center;
    box-shadow: 0px 2px 7px #b9bcbe;
    background-image: linear-gradient(to bottom, #4481eb 10%, #04affec9 90%);
    padding: 15px 0;
    max-height: 700px;
  }

  .cardtitle {
    /* font-size: 1.0667rem; */
    /* color: whitesmoke; */
    /* font-family: 'STXingkai'; */
    /* line-height: 1rem; */
    img {
      width: 100%;
    }
  }

  .whole {
    background-color: #f7f7f7;
    padding-top: 22.5px;

    margin: auto;
    margin-top: 0px;
    min-height: 100vh;
    overflow: hidden;
  }
}
</style>

<script lang="ts">
import { defineComponent, reactive } from "vue";

import MFunction from "../../components/managerCenter/managerFunction.vue";
import { useRouter } from "vue-router";
import axios from "axios";
import { message } from "ant-design-vue";
import { useUserStore } from "@/stores";
import { storeToRefs } from "pinia";
import ManageNav from "@/components/common/ManageNav.vue";

interface FormState {
  username: string;
  password: string;
}

export default defineComponent({
  components: {
    ManageNav,
    MFunction,
  },
  data() {
    return {
      m_url: "",
    };
  },

  setup() {
    const router = useRouter();
    const postData = reactive<FormState>({
      username: "",
      password: "",
    });


    const logout = () => {
      let url = 'api/user/logout';
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
         if (!v) {
            message.error('null response');
            return;
          }

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
      

        if (v.status == 200) {
          message.success("注销成功")
          router.push("login");
        }
        return;
      })
      .catch((err) => {
        message.error(err);
      });
    };

    return {
      postData,
      logout,
    };
  },
});
</script>
