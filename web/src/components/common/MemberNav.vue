<template>
  <div class="nav">
    <img @click="toMemberCenter" src="../../assets/img/title.png">
    <div class="funWrapper">
      <a @click="toMemberCenter">首页</a>
      <a @click="toChangeInfo">信息修改</a>
      <a @click="toResume">个人简历</a>
      <a @click="toMessage">我的消息</a>
      <a @click="toActivity">活动列表</a>
      <a @click="logout">注销</a>
    </div>
  </div>
</template>

<script>
import { useRouter } from "vue-router";
import {message} from 'ant-design-vue'

export default {
  name: "MemberNav",
  methods:{
    toMemberCenter(){
      this.$router.push('/memberCenter');
    },
    toChangeInfo(){
      this.$router.push('/informationChange');
    },
    toResume(){
      this.$router.push('/PersonalResueme');
    },
    toMessage(){
      this.$router.push('/cardList');
    },
    toActivity(){
      this.$router.push('/activityHome');
    },
  },
  setup(){
    const router = useRouter();

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
            router.push('/login')
            return
          }

        if(v.status == -1999){
          message.warn("管理员账号未登录，请先登录")
          router.push('/login')
          return
        }
        console.log(v.status)

        if (v.status == 200) {
          message.success("注销成功")
          router.push("/login");
        }
        return;
      })
      .catch((err) => {
        console.log(err);
        message.error(err);
      });
    };


    return {
      logout,
    };
  }
}
</script>

<style lang="less" scoped>
.nav{
  display: block;
  position: relative;
  width: 100%;
  height: 60px;
  border: 1px solid #efefef;
  padding: 0 60px;
  background-color: white;
  img{
    width: 250px;
    margin: auto;
    position: absolute;
    top:50%;
    transform: translateY(-50%);
  }
  .funWrapper{
    position: absolute;
    right: 60px;
    display: flex;
    text-align: center;
    a{
      font-size: 16px;
      flex: 1;
      width: 110px;
      line-height: 60px;
      text-decoration:none;
      color: #304455;
    }
  }
}
@media screen and (max-width: 1050px){
  .nav{
    .funWrapper{
      display: none;
    }
  }
}
@media screen and (max-width: 600px){
  .nav{
    display: none;
  }
}
</style>
