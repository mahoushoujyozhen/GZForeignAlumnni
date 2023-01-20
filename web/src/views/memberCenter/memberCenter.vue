<template>
  <MemberNav />
  <div class="headMobile">
    <div class="headTitle"><span>个人中心</span></div>
    <div class="headFunc" @click="logout"><span>注销</span></div>
  </div>
  <div class="bg">
    <div class="whole" style="background-color: #f7f7f7">
      <Card></Card>

      <div class="wrapper">
        <Function></Function>
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
.bg {
  background-color: #f7f7f7;
  min-height: 100vh;
}

.whole {
  padding-top: 0.6rem;
}
.bar {
  height: 1.0667rem;
}

@media screen and (min-width: 600px) {
  .headMobile {
    display: none;
  }
  .bar {
    display: none;
    height: 40px;
  }
  .whole {
    margin: auto;
    margin-top: 0px;
    padding-top: 22.5px;
  }
}
</style>

<script lang="ts">
import Function from "../../components/memberCenter/Function.vue";
import { defineComponent } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";

import MemberNav from "@/components/common/MemberNav.vue";
import Card from "@/components/memberCenter/Card.vue";
import { message } from "ant-design-vue";
export default defineComponent({
  components: {
    Card,
    MemberNav,
    Function,
  },
  setup() {
    const router = useRouter();
    const back = function () {
      router.push("/login");
    };

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
      back,
      logout,
    };
  },
  data() {
    return {
      name: "xxx",
      m_url: "",
    };
  },
});
</script>
