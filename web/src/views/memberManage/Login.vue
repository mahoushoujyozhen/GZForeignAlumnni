<template>
  <img src="../../assets/img/title.png">
  <van-form @submit="onFinish()" class="form">
    <van-cell-group inset>
      <van-field
          v-model="postData.username"
          name="用户名"
          label="用户名"
          placeholder="用户名"
          :rules="[{ required: true, message: '请填写用户名' }]"
      />
      <van-field
          v-model="postData.password"
          type="password"
          name="密码"
          label="密码"
          placeholder="密码"
          :rules="[{ required: true, message: '请填写密码' }]"
      />
    </van-cell-group>

    <div class="button">
      <van-button block type="primary" native-type="submit">用户登录</van-button>
      <van-button block plain type="primary" to="register">账号注册</van-button>
    </div>
  </van-form>
</template>
<script lang="ts">
import { defineComponent, reactive,onMounted} from "vue";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import { useRouter, useRoute } from "vue-router";


interface FormState {
  username: string;
  password: string;
}
export default defineComponent({
  components: {
    UserOutlined,
    LockOutlined,
  },
  setup() {
    const router = useRouter();
    const route = useRoute();
    const postData = reactive<FormState>({
      username: "",
      password: "",
    });   
     onMounted(()=>{
        let url = 'api/user/is_login';
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
            console.log(v);
            return v.json();
          })
          .then((res) => {
            console.log(res);
            if (!res) {
              message.error('null response');
              return;
            }

            if (route.query.act_id !== undefined) {
              Signin(); //如果携带活动id,实行扫码签到
              return;
            }

            if(res.status == 201){
              message.success("管理员账号已登录")
              router.push('managerCenter')
              return
            }

            if(res.status == 202){
              message.success("账号已登录")
              router.push('memberCenter')
              return
            }

          })
          .catch((err) => {
            console.log(err);
            message.error(err);
          });

    })




    const onFinish = () => {
      //去掉用户名和密码前后所有空格
      let username = postData.username.replace(/(^\s*)|(\s*$)/g, "");
      let password = postData.password.replace(/(^\s*)|(\s*$)/g, "");

      let url = 'api/user/login';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: JSON.stringify({
          username,
          password
        }),
      })
          .then((v) => {
            console.log(v);
            return v.json();
          })
          .then((res) => {
            console.log(res);

            if (!res) {
              message.error('null response');
              return;
            }

            if(res.status == -20004){
              message.error("密码错误，请重新输入！")
              return
            }


            if (route.query.act_id !== undefined) {
              Signin(); //如果携带活动id,实行扫码签到
              return
            }


            if(res.data.username == "admin"){
              message.success("管理员账号登录成功")
              router.push('/managerCenter')
            }else{
              message.success("用户账号登录成功")
              router.push('/memberCenter');
            }


          })
          .catch((err) => {
            console.log(err);
            message.error(err);
          });

    };

    const Signin = () => {
      let url = "api/user/signIn";
      fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        redirect: "follow",
        body: JSON.stringify({
          act_id: route.query.act_id,
        }),
      })
          .then((v) => {
            console.log(v);
            return v.json();
          })
          .then((v) => {
            console.log("v=",v);
            console.log("v.data=",v.data);
            if (!v) {
              message.error("null response");
              console.log(v.status);
              return;
            }

            if (v.status !== 0) {
              message.error(v.msg);
              console.log(v.msg);
              console.log(v.status);
              return;
            }

            if (v.data.isApplied === false) {
              console.log("用户未报名");
              message.warning("未报名该活动");
              router.push("/memberCenter");
              return;
            }

            if (v.data.isOver === true) {
              console.log("活动已结束");
              message.warning("活动已结束");
              router.push("/memberCenter");
              return;
            }

            if (v.data.isSignIn === true) {
              console.log("用户已签到");
              message.success("已签到");
              router.push("/memberCenter");
              return;
            }
            console.log(v.status, v.msg);
            message.success("签到成功");
            router.push("/memberCenter");

            //跳转到签到成功页面
          })
          .catch((err) => {
            console.log(err);
            message.error(err);
          });
    };


    return {
      postData,
      onFinish,
      Signin,
    };
  },

});
</script>
<style lang="less" scoped>
@media screen and (max-width: 600px){
  img{
    position: relative;
    width: 8.5rem;
    left: 50%;
    transform: translateX(-50%);
    margin: 2rem 0 .8rem;
  }
  .form{
    width: 9.5rem;
    margin:auto;
    .van-cell-group{
      padding: 15px;
      margin-bottom: 40px;
    }
    .button{
      margin:20px auto;
      width:300px;
      >*{
        margin-bottom: .3rem;
      }
    }
  }


}
//大于600px宽度的屏幕，pc端
@media screen and (min-width: 600px){
  img{
    position: relative;
    width: 550px;
    left: 50%;
    transform: translateX(-50%);
    margin: 150px 0 70px;
  }
  .form{
    width:560px;
    margin:auto;
    .van-cell-group{
      padding: 15px;
      margin-bottom: 50px;
    }
    .button{
      display: flex;
      justify-content: space-evenly;
      >*{
        width: 200px;
      }
    }
  }
}

</style>
