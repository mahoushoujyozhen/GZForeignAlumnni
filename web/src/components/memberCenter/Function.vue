<template >
  <van-notice-bar
    v-if="userStatus !== 1"
    mode="closeable"
    left-icon="volume-o"
    :text="noticeText"
    class="notice"
    
  />

  <van-grid :gutter="5" :column-num="4" square class="func" :border="false">
    <van-grid-item to="informationChange">
      <van-icon name="records" color="rgb(241,143,96)" />
      <div class="changeBackgroundColor">信息修改</div>
      <!-- <van-grid-item id="msg" icon="records" text="信息修改" to="informationChange" /> -->
      <!-- v-if="userOption === '审核通过'" -->
    </van-grid-item>

    <van-grid-item @click="toResueme">
      <van-icon name="description" color="#40a9ff" />
      <div class="funcTitle">个人简历</div>
    </van-grid-item>

    <van-grid-item to="CardList">
      <van-icon name="envelop-o" color="rgb(250,192,61)" />
      <div class="funcTitle">我的消息</div>
    </van-grid-item>
    <!-- v-if="userStatus ==  1" -->
    <van-grid-item @click="toActivity">
      <van-icon name="flag-o" color="rgb(104 213 144)" />
      <div class="funcTitle">活动列表</div>
    </van-grid-item>
  </van-grid>
</template>

<script lang="ts">
import router from '@/router';
import { defineComponent, onBeforeMount, onMounted, onUpdated, ref } from 'vue';
import { message } from 'ant-design-vue';

export default defineComponent({
  name: 'Function',
  setup() {
    // let userid = localStorage.getItem('userID');
    let userStatus = ref(0);
    const noticeText = ref('');
    const noticeChange = () => {
      if (userStatus.value === 2) {
        noticeText.value =
          '审核未通过，请在信息修改页面重新提交信息,暂时只可使用信息修改与消息通知功能，不便之处，敬请原谅';
      } else if (userStatus.value === 3) {
        noticeText.value =
          '信息审核中，暂时只可使用信息修改与消息通知功能，不便之处，敬请原谅';
      }
  
    };
    const toResueme = () => {
      if (userStatus.value !== 1) return;
      else {
        router.push('/PersonalResueme');
      }
    };
    const toActivity = () => {
      if (userStatus.value !== 1) return;
      else {
        router.push('/activityHome');
      }
    };
    onMounted(() => {
      const url = 'api/user/getopinion';
      fetch(url, {
        method: 'POST',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        body: JSON.stringify({
          // id: userid,
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          //只用从后端拿审核状态
          userStatus.value = v.user_status;
          noticeChange();
        });
    });

    return {
      noticeText,
      userStatus,
      toResueme,
      toActivity,
    };
  },
});
</script>

<style lang="less" scoped>
.van-icon-records {
  font-size: 0.8rem;
}
.van-icon-description {
  font-size: 0.8rem;
}

.van-icon-envelop-o {
  font-size: 0.8rem;
}
.van-icon-flag-o {
  font-size: 0.8rem;
}
.notice {
  margin: 0 0.5rem 0.2667rem 0.5rem;
}
.func {
  text-align: center;
  font-size: 0.3733rem;
}
.van-grid {
  display: flex;
  align-items: baseline;
  justify-content: space-around;
  margin: 0 0.5rem;
}
.funcTitle {
  color: #304455;
}
.van-icon {
  color: #1890ff;
  margin-bottom: 0.2rem;
}

@media screen and (min-width: 600px) {
  .van-icon-records {
    font-size: 37.5px;
  }
  .van-icon-description {
    font-size: 37.5px;
  }
  .van-icon-envelop-o {
    font-size: 37.5px;
  }
  .van-icon-flag-o {
    font-size: 37.5px;
  }
  .notice {
    margin: 0 auto 15px;
    width: 670px;
  }
  .func {
    text-align: center;
    font-size: 16px;
  }
  .van-grid {
    display: flex;
    align-items: baseline;
    justify-content: space-around;
    margin: auto;
    max-width: 670px;
  }
  .van-icon {
    color: #1890ff;
    margin-bottom: 0.1rem;
  }
}
</style>
