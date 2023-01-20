<template>
<div>
   <div id="refresh">
      <router-view v-if="isRouterAlive"></router-view>
  </div>
  <div class="activity">
    <div class="act_title" @click="toDetail">{{ msg.act_name }}</div>
    <div class="act_content">
      <van-image :src="msg.img_url" fit="cover" position="center" @click="toDetail"/>
      <div class="act_detail" @click="toDetail">{{msg.act_profile}}</div>
      <van-button v-show="isManager" class="manageBtn" type="primary" @click="showPopup">管理</van-button>
    </div>
  </div>
  <van-popup class="act_handle" v-model:show="show" round position="bottom" :style="{ height: '6.8rem' }">
    <div class="topHr"></div>
    <div class="popupWrapper">
      <div class="actEdit">
        <div class="edit" @click="toEditActivity"><span>编辑活动</span><van-icon size=".5rem" name="edit" /></div>
        <div class="delete" @click="deleteActivity"><span>删除活动</span><van-icon size=".5rem" name="delete-o" /></div>
      </div>
      <van-grid :border="false">
        <van-grid-item icon="video" icon-color="#1890ff" text="活动掠影" @click="toRecall"/>
        <van-grid-item icon="scan" icon-color="#1890ff" text="签到码" @click="signInCode"/>
        <van-grid-item icon="todo-list" icon-color="#1890ff" text="签到列表" @click="signInList"/>
      </van-grid>
    </div>
  </van-popup>
  <van-popup class="act_handle" v-model:show="showPCPopup" round :style="{ width:'450px' }">
    <div class="topHr"></div>
    <div class="popupWrapper">
      <div class="actEdit">
        <div class="edit" @click="toEditActivity"><span>编辑活动</span><van-icon size="20px" name="edit" /></div>
        <div class="delete" @click="deleteActivity"><span>删除活动</span><van-icon size="20px" name="delete-o" /></div>
      </div>
      <van-grid :border="false">
        <van-grid-item icon="video" icon-color="#1890ff" text="活动掠影" @click="toRecall"/>
        <van-grid-item icon="scan" icon-color="#1890ff" text="签到码" @click="signInCode"/>
        <van-grid-item icon="todo-list" icon-color="#1890ff" text="签到列表" @click="signInList"/>
      </van-grid>
    </div>
  </van-popup>
</div>
</template>

<script>
import {ref} from "vue";
import {message} from "ant-design-vue";

export default {
  name: "manage_activity_item",
  inject: ['reload'],
  setup() {
    const show = ref(false);
    const showPCPopup = ref(false);
    const showPopup = () => {
      if(document.body.clientWidth < 600){
        show.value = true;
      }else {
        showPCPopup.value=true;
      }
    }

    return {
      show,
      showPopup,
      showPCPopup,
    };
  },
  props:['msg','isManager'],
  methods:{
    toRecall(){
      this.$router.push({path:`/issueRecall/${this.msg.act_id}`});
    },
    toDetail(){
      console.log("this.isManager",this.isManager)
      this.$router.push({path:`/activityDetail/${this.msg.act_id}/${this.isManager}`});
    },
    toEditActivity(){
      this.$router.push({path:`/releaseActivity/${this.msg.act_id}`});
    },
    signInCode(){
      this.$router.push({path:`/qrcode/${this.msg.act_id}`});
    },
    signInList(){
      this.$router.push({path:`/signedlist/${this.msg.act_id}`});
    },
   
     deleteActivity() {
      console.log("正在删除:"+this.msg.act_id);
      let url = 'api/admin/act_delete';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        //headers是请求头：
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: JSON.stringify({
        act_id: this.msg.act_id,
        }),
      }).then((v) => {
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
              // message.error('null response');
              return;
            }

            if (v.status !== 0) {
              // message.error(v.msg);
              return;
            }
           
          })
          .catch((err) => {
            console.log(err);
          });

      // this.$router.go(0); //当前页面刷新
      this.reload();
    },
  },
  mounted() {
  }
}
</script>

<style lang="less" scoped>
@baseFont: 37.5;
@media screen and (max-width: 600px){
  .activity{
    //width: 10rem;
    width: 100%;
    height: (128rem/@baseFont);
    padding: (10rem/@baseFont) 0;
    background-color: white;
    .act_title{
      font-size: (18rem/@baseFont);
      margin: 0 (20rem/@baseFont);

      //一行省略号
      overflow:hidden;
      text-overflow:ellipsis;
      white-space:nowrap
    }
    .act_content{
      display: flex;
      margin: (10rem/@baseFont) (20rem/@baseFont);
      .van-image{
        width: (65rem/@baseFont);
        height: (65rem/@baseFont);
      }
      .act_detail{
        flex:1;
        //换行、溢出省略号
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;  /*显示的行数*/
        overflow: hidden;
        word-break:break-all; /*可择*/

        width: (156rem/@baseFont);
        font-size: (14rem/@baseFont);
        margin: 0 (10rem/@baseFont);
      }
      .van-button{
        margin-top: (36rem/@baseFont);
        width: (94rem/@baseFont);
        height: (29rem/@baseFont);
        font-size: (14rem/@baseFont);
        z-index: 50;
      }
    }
  }
  .act_handle{
    .topHr{
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
      top:(12rem/@baseFont);
      width: (40rem/@baseFont);
      height: (3rem/@baseFont);
      background-color: #919090;
      border-radius: (10rem/@baseFont);
    }
    .popupWrapper{
      margin: (40rem/@baseFont) auto 0;
      width: (320rem/@baseFont);
      .actEdit{
        display: flex;
        justify-content: space-between;
        > *{
          //flex: 1;
          width: (150rem/@baseFont);
          text-align: center;
          height: (60rem/@baseFont);
          line-height: (60rem/@baseFont);
          font-size: (16rem/@baseFont);
          border-radius: (3rem/@baseFont);
        }
        .van-icon{
          margin-left: (2rem/@baseFont);
        }
        .edit{
          background-image: linear-gradient(to bottom right,#1890ff,#389fff);
          color:white;
        }
        .delete{
          border: 1px solid #1890ff;
          color:#1890ff;
        }
      }
      .van-grid{
        margin-top: (20rem/@baseFont);
        > *{

          --van-grid-item-text-font-size:.37rem;
          --van-grid-item-icon-size:1rem;
          flex: 1;
        }
      }
    }
  }
}
@media screen and (min-width: 600px){
  .activity{
    //width: 10rem;
    width: 100%;
    //height: 128px;
    padding: 30px 45px 20px 45px;
    background-color: white;
    .act_title{
      font-size: 18px;
      //margin: 0 20px;

      //一行省略号
      overflow:hidden;
      text-overflow:ellipsis;
      white-space:nowrap
    }
    .act_content{
      display: flex;
      //width: (335rem/@baseFont);
      //width: 100%;
      margin: 10px 0;
      .van-image{
        width: 100px;
        height: 100px;
      }
      .act_detail{
        flex:1;
        //换行、溢出省略号
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 3;  /*显示的行数*/
        overflow: hidden;
        word-break:break-all; /*可择*/

        width: 156px;
        font-size: 14px;
        margin: 0 10px;
      }
      .van-button{
        margin-top: 71px;
        width: 94px;
        height: 29px;
        font-size: 14px;
        z-index: 50;
      }
    }
  }
  .act_handle{
    .topHr{
      position: absolute;
      left: 50%;
      transform: translateX(-50%);
      top:12px;
      width: 40px;
      height: 3px;
      background-color: #919090;
      border-radius: 10px;
    }
    .popupWrapper{
      margin: 50px auto 40px;
      width: 320px;
      .actEdit{
        display: flex;
        justify-content: space-between;
        > *{
          //flex: 1;
          width: 150px;
          text-align: center;
          height: 60px;
          line-height: 60px;
          font-size: 16px;
          border-radius: 3px;
        }
        .van-icon{
          margin-left: 2px;
        }
        .edit{
          background-image: linear-gradient(to bottom right,#1890ff,#389fff);
          color:white;
        }
        .delete{
          border: 1px solid #1890ff;
          color:#1890ff;
        }
      }
      .van-grid{
        margin-top: 20px;
        > *{
          --van-grid-item-text-font-size:15px;
          --van-grid-item-icon-size:50px;
          flex: 1;
        }
      }
    }
  }
}
</style>
