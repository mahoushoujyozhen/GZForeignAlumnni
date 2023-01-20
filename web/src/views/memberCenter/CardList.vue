<template>
  <MemberNav />
  <div class="container">
    <div class="list-wrapper">
      <div class="list-item" v-for="msg in messageList" @click="OnShowDetails(msg)">
        <MessageCard :title="msg.title" :time="msg.time" :content="msg.content">
        </MessageCard>
      </div>
    </div>
  </div>
</template>

<script lang="js">
import { ref, onMounted } from "vue";
import MessageCard from "../../components/memberCenter/MessageCard.vue";
import { useRouter } from "vue-router";
import dayjs from "dayjs";
import MemberNav from "@/components/common/MemberNav.vue";
import MobileNav from "@/components/common/MobileNav.vue";
export default {
  name: "CardList",
  components: {
    MobileNav,
    MemberNav,
    MessageCard,
  },
  setup() {
    // const messageList = ref([{
    //   title:ref([]),
    //   time: ref([])_,
    //   content: ref([]),
    // }]);
    const messageList=ref([])
    const id = 0;
    // 开启websocket
    onMounted(() => {
      // const websocket = new WebSocket(`ws://172.22.72.229:9090/api/message/ws`);
      const websocket = new WebSocket(`ws://mac.gzhu.edu.cn/alumni/api/message/ws`);

      websocket.onopen = () => {
        let storageList = window.localStorage.getItem("messageList");
        if (storageList != null) {
          messageList.value.push(...JSON.parse(storageList));
        }
      };

      websocket.onmessage = (e) => {
        const responseList = JSON.parse(e.data).details;
        if (responseList !== undefined) {
          responseList.push(...messageList.value);
          // 根据时间排序
          responseList.sort(function (front, back) {
            return back.time - front.time;
          });

          responseList.map((item) => {
            item.time = dayjs(item.time).format("YYYY-MM-DD HH:mm:ss");
          });
          messageList.value.splice(0, messageList.value.length);
          messageList.value.push(...responseList);
          // 缓存信息
          window.localStorage.setItem("messageList", JSON.stringify(messageList.value));
        }
      }

      websocket.onclose = () => {
        // websocket.send(JSON.stringify({name: "zhangsan"}));
        websocket.close(1000, "close");
      };
    });

    const router = useRouter();

    const OnShowDetails = (msg) => {
      router.push({
        name: "CardDetails",
        path: "/cardDetails",
        params: {
          title: msg.title,
          time: msg.time,
          content: msg.content,
        },
      });
    };

    return {
      messageList,
      OnShowDetails,
    };
  },
  methods:{
    back(){
      this.$router.push('/memberCenter')
    }
  }
};
</script>

<style scoped>
@media screen and (min-width: 600px) {
  .container {
    background-color: #f7f7f7;
  }
  .list-wrapper {
    width: 750px;
    margin: 50px auto;
  }
  .list-item {
    border-radius: 10px;
    background-color: white;
    margin-bottom: 12px;
  }
}
@media screen and (max-width: 600px) {
  .container {
    min-height: 100vh;
    background-color: #f7f7f7;
    padding: 0.53rem;
  }
  .list-item {
    margin-bottom: 0.26rem;
  }
}
</style>
