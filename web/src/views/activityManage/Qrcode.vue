<template>
  <ManageNav/>
  <MobileNav title="活动二维码"/>
  <div class="main">
    <van-image class="img" :src="picUrl" fit="fill" />
  </div>
</template>

<style lang="less" scoped>
@media screen and (max-width: 600px) {
  .main {
    width: 10rem;
    height: 17.7867rem;
    //margin: 1rem auto;
  }
  .img {
    width: 100%;
    position: absolute;
    top: 2.6667rem;
  }
}

@media screen and (min-width: 600px) {
  .main {
    width: 100%;
    height: 700px;
    //margin: 1rem auto;
  }
  .img {
    width: 450px;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
  }
}
</style>

<script lang="ts">

//import { ConsoleSqlOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { defineComponent, reactive } from "vue";
import ManageNav from "@/components/common/ManageNav.vue";
import MobileNav from "@/components/common/MobileNav.vue";
//import TopNav from "@/components/common/TopNav.vue";


export default defineComponent({
  components: {
    MobileNav,
    ManageNav
    //TopNav,
  },

  computed: {
    // 计算属性
  },
  watch: {
    // 数据监听
  },

  setup() {
    console.log("setup");

    return reactive({
      picUrl: "", //img标签显示
      imgBase: "", //存图片字节流
    });
  },

  created() {
    console.log("created");
    this.getQrcode();
  },

  methods: {
    // back() {
    //   this.$router.push("/activityManage");
    // },

    getQrcode() {
      console.log("it's function getQrcode");
      let url = "api/admin/getQrcode";
      fetch(url, {
        method: "Post",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        redirect: "follow",
        body: JSON.stringify({
          act_id: this.$route.params.act_id,
        }),
      })
        .then((v) => {
          console.log(v);
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
            message.error("null response");
            return;
          }

          if (v.status !== 0) {
            message.error(v.msg);
            console.log(v.status);
            console.log(v.msg);
            return;
          }

          this.imgBase = v.qrcode;
          this.picUrl = "data:png/jpeg;base64," + this.imgBase;
          console.log(v.msg);
          console.log("成功获取二维码");
        })
        .catch((err) => {
          console.log(err);
          message.error(err);
        });
    },
  },
});
</script>
