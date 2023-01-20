<template>
  <div>
    <ManageNav v-show="showManageNav"/>
    <MemberNav v-show="!showManageNav"/>
    <div class="imgWrapper">
      <img class="mainImg" :src="imgUrl" />
    </div>
    <div class="mainContainer">
      <div class="Label"><van-icon name="bullhorn-o" /><span>活动详情</span></div>
      <div class="detail">
        <div class="detailTitle">{{ activity.act_name }}</div>
        <ul class="detailContent">
          <li>
            <span>开始时间：{{ activity.start_time }}</span>
          </li>
          <li>
            <span>结束时间：{{ activity.end_time }}</span>
          </li>
          <li>
            <span>活动地点：{{ activity.act_address }}</span>
          </li>
          <li>
            <span>报名人数：{{ activity.apply_num }}</span>
          </li>
          <li>
            <div>活动简介：</div>
            <div>{{ activity.act_profile }}</div>
          </li>
        </ul>
        <van-button v-show="!ActIsStart && !isApplied && showApplyBtn" type="primary" @click="apply">我要报名</van-button>
        <van-button v-show="!ActIsStart && isApplied && showApplyBtn" @click="cancel">取消报名</van-button>
        <van-button v-show="ActIsStart && !ActIsEnd">活动已开始</van-button>
        <van-button v-show="ActIsEnd">活动已结束</van-button>
      </div>
      <div class="Label" v-show="recallIsShow">
        <van-icon name="paid" /><span>活动掠影</span>
      </div>
      <div class="recall" v-show="recallIsShow">
        <div class="recallContent">
          {{ activity.act_text }}
        </div>
        <div v-for="imgItem in recallImgList" class="imgItem">
          <van-image :src="imgItem" />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { Notify } from "vant";
import { reactive, defineComponent } from "vue";
import ManageNav from "@/components/common/ManageNav.vue";
import MemberNav from "@/components/common/MemberNav.vue";
import { message } from "ant-design-vue";

export default defineComponent({
  components: {
    MemberNav,
    ManageNav,
    [Notify.Component.name]: Notify.Component,
  },

  setup() {
    return reactive({
      showApplyBtn: false,
      isApplied: false,
      recallIsShow: false,
      showManageNav: false,
      ActIsStart:false,
      ActIsEnd:false,
      recallContent: "",
      imgUrl: "",
      recallImgList: [] as string[],
      activity: {
        img_url: "",
        act_name: "",
        act_profile: "",
        start_time: "",
        end_time: "",
        act_address: "",
        apply_num: 0,
        act_text: "",
        recall_img: "",
      },
    });
  },
  methods: {
    apply() {
      let query = [];
      // query.push(`user_id=${this.user_id}`);
      query.push(`act_id=${this.$route.params.act_id}`);
      let q = query.join("&");
      let method = "PUT";
      let url = `api/user/apply?${q}`;
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
            console.log("v为空");
          }
          if (v.status !== 0) {
            return;
          }
          this.isApplied = !this.isApplied;
          this.activity.apply_num = v.data.apply_num;
        })
        .catch(err=>{
        Notify({ type: "warning", message: err });
      });
    },
    cancel() {
      let query = [];
      // query.push(`user_id=${this.user_id}`);
      query.push(`act_id=${this.$route.params.act_id}`);
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
            console.log("v为空");
          }
          if (v.status !== 0) {
            console.log(v.msg);
            return;
          }
          this.isApplied = !this.isApplied;
          this.activity.apply_num = v.data.apply_num;
        })
        .catch(err=>{
        Notify({ type: "warning", message: err });
      });
    },
    //获取用户是否报名
    isApply() {
      let query = [];
      query.push(`act_id=${this.$route.params.act_id}`);
      // query.push(`user_id=${this.user_id}`);
      let q = query.join("&");
      let url = `api/user/isapply?${q}`;
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
      .then((v)=>{
        return v.json();
      })
      .then((v)=>{
        if(v.status == -999){
          message.warn("账号未登录，请先登录")
          this.$router.push('login')
        }

        if(v.status == -1999){
          message.warn("管理员账号未登录，请先登录")
          this.$router.push('login')
        }
        this.isApplied = v.data.isApply;
      })
      .catch(err=>{
        console.log("method isApply" + err);
        Notify({ type: "warning", message: err });
      })
    },
    //获取活动页面信息
    getActMsg() {
      let query = [];
      query.push(`act_id=${this.$route.params.act_id}`);
      let q = query.join("&");
      let url = `api/activity/getById?${q}`;
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
      .then((v)=>{
        console.log(v);
        return v.json();
      })
      .then((v)=>{
        if(v.status == -999){
          message.warn("账号未登录，请先登录")
          this.$router.push('login')
        }

        if(v.status == -1999){
          message.warn("管理员账号未登录，请先登录")
          this.$router.push('login')
        }
        this.activity = v;
        this.imgUrl = this.activity.img_url;
        this.recallContent = this.activity.act_text;
        if(this.activity.recall_img != null){
          this.recallImgList = this.activity.recall_img.split(",");
        }
        if (this.recallContent !== null) {
          this.recallIsShow = true;
        }

        //当前时间和活动开始时间、活动结束时间比较
        this.compareTime()
      })
      .catch(err=>{
        Notify({ type: "warning", message: err });
      })
    },

    compareTime(){
      //当前时间的时间戳
      var now = Date.now();
      //开始时间的时间戳
      var StartDate = new Date(this.activity.start_time.replace(/-/g,'/'));
      var ActStart=StartDate.getTime();
      //结束时间的时间戳
      var EndDate = new Date(this.activity.end_time.replace(/-/g,'/'));
      var ActEnd=EndDate.getTime();

      if(now >= ActStart){
        this.ActIsStart=true;
      }
      if(now >= ActEnd){
        this.ActIsEnd=true;
      }
    },

    async prepare() {
      if (this.$route.params.isManager === "false") {
        this.showApplyBtn = true;
        this.showManageNav = false
      } else {
        this.showApplyBtn = false;
        this.showManageNav = true;
      }
      this.isApply();
      this.getActMsg();
    },
  },

  created() {
    this.prepare();
  },
});
</script>

<style lang="less" scoped>
@baseFont:37.5;
#app{
  min-height: 1000px;
}
@media screen and (max-width: 600px){
  .imgWrapper{
    max-width: 1200px;
    margin: 0 auto;
  }
  .mainImg{
    //width: 10rem;
    width: 100%;
    //height: (200rem/@baseFont);
    //height: 100%;
    max-width: 1200px;
    margin: 0 auto;
  }
  .mainContainer{
    //width: 10rem;
    max-width: 1200px;
    margin: 0 auto;
    width: 100%;
    background-color: #F0F2F5;
    .Label{
      font-size: (14rem/@baseFont);
      color: #7F7F7F;
      padding: (10rem/@baseFont) (20rem/@baseFont);
      .van-icon{
        margin-right: (10rem/@baseFont);
      }
    }
    .detail{
      background-color: white;
      //width: 10rem;
      width: 100%;
      padding:(20rem/@baseFont);
      .detailTitle{
        font-size: (24rem/@baseFont);
        font-weight: bold;
        margin-bottom: (10rem/@baseFont);
      }
      .detailContent{
        font-size: (16rem/@baseFont);
        word-wrap: break-word;
      }
      .van-button{
        position: relative;
        left: 50%;
        transform: translateX(-50%);
        width: (150rem/@baseFont);
        height: (35rem/@baseFont);
        margin-top: (20rem/@baseFont);
        font-size: (14rem/@baseFont);
      }
    }
    .recall{
      position: relative;
      background-color: white;
      //width: 10rem;
      width: 100%;
      padding: (20rem/@baseFont);
      .recallContent{
        font-size: (16rem/@baseFont);
        word-wrap: break-word;
      }
      .imgItem{
        width: (335rem/@baseFont);
        margin: (20rem/@baseFont) auto 0 auto;
        .van-image{
          //max-width: 710px;
          width: (335rem/@baseFont);
        }
      }
    }
  }
}
@media screen and (min-width: 600px){
  .imgWrapper{
    max-width: 1000px;
    margin: 0 auto;
  }
  .mainImg{
    //width: 10rem;
    width: 100%;
    //height: (200rem/@baseFont);
    //height: 100%;
    max-width: 1000px;
    margin: 0 auto;
  }
  .mainContainer{
    max-width: 1000px;
    margin: 0 auto;
    .Label{
      font-size: 14px;
      color: #7F7F7F;
      padding: 10px 20px;
      .van-icon{
        margin-right: 10px;
      }
    }
    .detail{
      background-color: white;
      //width: 10rem;
      width: 100%;
      padding:35px;
      .detailTitle{
        font-size: 24px;
        font-weight: bold;
        margin-bottom: 10px;
      }
      .detailContent{
        font-size: 16px;
        word-wrap: break-word;
      }
      .van-button{
        position: relative;
        left: 50%;
        transform: translateX(-50%);
        width: 150px;
        height: 35px;
        margin-top: 20px;
        font-size: 14px;
      }
    }
    .recall{
      position: relative;
      background-color: white;
      padding: 35px;
      .recallContent{
        font-size: 16px;
        word-wrap: break-word;
        margin-bottom: 20px;
      }
      .imgItem{
        margin: 20px 0;
        .van-image{
          width: 650px;
          position: relative;
          left: 50%;
          transform: translateX(-50%);
        }
      }
    }
  }
}
</style>
