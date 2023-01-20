<template>
  <div class="card">
    <div class="cardtitle">
      <img src="../../assets/img/广州欧美同学会.png" class="cardtitle" />
    </div>
    <hr
      style="
        height: 1px;
        border: none;
        border-top: 1px solid white;
        width: 80%;
        margin: 0 auto;
      "
    />
    <div>
      <div class="cardmsg">
        <span>姓名 {{ username }}</span>
<!--        <span>编号 {{ userid }}</span>-->
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.card {
  margin: 0rem auto;
  //height: auto;
  width: 8.7rem;
  box-sizing: border-box;
  border-radius: 0.1333rem;
  background-size: cover;
  text-align: center;
  box-shadow: 0px 2px 7px #b9bcbe;
  background-image: linear-gradient(to bottom, #4481eb 10%, #04affec9 90%);
  padding-top: 0.4rem;
}
.cardtitle {
  //font-size: 1.0667rem;
  color: whitesmoke;
  //font-family: 'STXingkai';
  //line-height: 1rem;
  img {
    width: 100%;
  }
}
.cardmsg {
  text-align: left;
  display: flex;
  justify-content: space-between;

  line-height: 0.5333rem;
  color: white;
  font-size: 0.3733rem;
  padding-bottom: 0.5rem;
  margin: 0.4rem 1rem;
}

@media screen and (min-width: 600px) {
  .card {
    margin: 0 auto 15px;
    max-width: 670px;
    box-sizing: border-box;
    border-radius: 5px;
    background-size: cover;
    text-align: center;
    box-shadow: 0px 2px 7px #b9bcbe;
    background-image: linear-gradient(to bottom, #4481eb 10%, #04affec9 90%);
    padding: 15px 0;
  }

  .cardtitle {
    color: whitesmoke;
    img {
      width: 100%;
    }
  }
  .cardmsg {
    text-align: left;
    display: flex;
    justify-content: space-between;
    line-height: 20px;
    color: white;
    font-size:20px;
    padding-bottom: 0;
    margin: 40px 60px 15px 60px;
  }
}
</style>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from 'vue';
import {message} from "ant-design-vue";
import { useRouter } from "vue-router";

export default defineComponent({
  setup() {
    const router = useRouter();
    let username = ref('');
    const respondmsg = ref('');
    onMounted(() => {
      const url = 'api/user/showmsg';
      fetch(url, {
        method: 'POST',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        body: JSON.stringify({
        }),
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
          respondmsg.value = v.msg;
          //从后端拿用户名
          username.value = v.username;
          // userid.value = v.userid;
          // opinion.value= v.useropinion;
        });
    });

    return {
      username,

    };
  },
});
</script>
