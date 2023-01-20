<template>
  <MemberNav />
  <div class="headMobile">
    <div class="headCancle" @click="back"><span>取消</span></div>
    <div class="headTitle"><span>信息修改</span></div>
    <div class="headChange" @click="changemsg"><span>修改</span></div>
  </div>
  <div class="whole">
    <div class="totalmsg">
      <van-cell-group inset>
        <van-field
          label="姓名"
          v-model="username"
          @blur="checkmsg"
          :error-message="nameError"
          :disabled="userStatus === 1"
          clearable
          placeholder="请输入姓名"
        />
        <van-field label="会员编号" :model-value="userID" disabled />
        <van-field
          label="用户名"
          :model-value="nickname"
          disabled
          placeholder="请输入用户名"
        />
        <!-- 性别 -->
        <van-field name="radio" label="性别">
          <template #input>
            <van-radio-group v-model="userSex">
              <van-radio name="男" icon-size="15px" @click="checkmsg">男</van-radio>
              <van-radio name="女" icon-size="15px" @click="checkmsg">女</van-radio>
            </van-radio-group>
          </template>
        </van-field>

        <!-- //证件类型选择器 -->
        <!-- click函数：判断屏幕宽度，如果大于600px，pop="true"，否则picker=true，要多写一个弹出框  -->

        <van-field
          v-model="IdCardType"
          is-link
          :disabled="userStatus === 1"
          name="picker"
          label="证件类型"
          @click="ShowType"
          :rules="[{ required: true, message: '请选择证件号类型' }]"
        />
        <van-popup v-model:show="showPicker" round position="bottom">
          <van-picker
            :columns="columns"
            @confirm="onConfirm"
            @cancel="showPicker = false"
          />
        </van-popup>
        <van-popup
          v-model:show="showPcPicker"
          round
          :style="{ width: '450px', height: '300px' }"
        >
          <van-picker
            :visible-item-count="4"
            :columns="columns"
            @confirm="onConfirm"
            @cancel="showPcPicker = false"
          />
        </van-popup>

        <!-- //证件号 -->
        <van-field
          v-model="IdCard"
          type="number"
          name="证件号"
          label="证件号"
          :rules="[{ required: true, message: '请填写证件号' }]"
          clearable
          @blur="checkmsg"
          :error-message="idcardError"
          :disabled="userStatus === 1"
          placeholder="请输入证件号"
        />
        <!-- 手机号 -->
        <van-field
          v-model="userphone"
          type="tel"
          label="手机号"
          clearable
          @blur="checkmsg"
          :error-message="perror"
          placeholder="请输入手机号"
        />
        <!-- 电子邮箱 -->
        <van-field
          v-model="useremail"
          label="电子邮箱"
          clearable
          @blur="checkmsg"
          :error-message="eerror"
          placeholder="请输入邮箱（选填）"
        />
      </van-cell-group>
    </div>
    <div class="buttonWrapper">
      <van-button type="primary" @click="changemsg">提交修改</van-button>
      <van-button type="primary" plain @click="back">取消</van-button>
    </div>
  </div>
</template>

<style lang="less" scoped>
/deep/.van-field--disabled .van-field__label {
  color: var(--van-field-text-color);
}
.whole {
  background-color: #f7f7f7;
  min-height: 100vh;
  margin: auto;
  .van-cell-group {
    margin: 0;
    padding: 0.26rem 0.1rem;
  }
}
.headMobile {
  width: 100%;
  height: 1.06rem;
  display: flex;
  justify-content: space-between;
  background-color: #ffffff;
  border-bottom: 1px solid #efefef;
  .headCancle {
    font-size: 0.43rem;
    line-height: 1.06rem;
    margin-left: 0.53rem;
  }
  .headTitle {
    align-items: center;
    font-size: 0.48rem;
    line-height: 1.06rem;
  }
  .headChange {
    font-size: 0.43rem;
    color: #1890ff;
    line-height: 1.06rem;
    margin-right: 0.53rem;
  }
}
.buttonWrapper {
  display: none;
}
.bar {
  height: 1.0667rem;
}

.msg {
  box-sizing: border-box;
  border: 1px solid #f2f2f2;
  height: 40px;
  width: auto;
  display: flex;
  flex-flow: row;
  align-items: center;
}

.lb {
  display: flex;
  flex-direction: row;
  padding-left: 5%;
}
.changeFontColor {
  color: white;
}
.totalmsg {
  padding-top: 0.53rem;
  margin: 0 0.53rem;
}

@media screen and (min-width: 600px) {
  .headMobile {
    display: none;
  }

  .buttonWrapper {
    margin: auto;
    width: 730px;
    display: flex;
    flex-direction: row-reverse;
    justify-content: space-between;
    padding-bottom: 50px;
    .van-button:nth-child(1) {
      width: 150px;
      margin-top: 35px;
      font-size: 16px;
    }
    .van-button:nth-child(2) {
      width: 150px;
      margin-top: 35px;
      font-size: 16px;
    }
  }
  .bar {
    height: 40px;
  }
  .msg {
    box-sizing: border-box;
    border: 1px solid #f2f2f2;
    height: 40px;
    width: auto;
    display: flex;
    flex-flow: row;
    align-items: center;
  }

  .lb {
    display: flex;
    flex-direction: row;
    padding-left: 5%;
  }
  .changeFontColor {
    color: white;
  }
  .changeFontColor:hover {
    font-weight: 700;
  }
  .totalmsg {
    padding-top: 50px;
    width: 750px;
    margin: auto;
  }

  .whole {
    background-color: #f7f7f7;
    .van-cell-group {
      margin: 0;
      padding: 10px;
    }
  }
}
</style>

<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from "vue";
import MemberNav from "@/components/common/MemberNav.vue";
import { message } from "ant-design-vue";
import router from "@/router";
export default defineComponent({
  inject: ["reload"],
  components: {
    MemberNav,
  },
  data() {
    return {
      perror: "",
      eerror: "",
      nameError: "",
      idcardError: "",
      phoneChange: false,
      emailChange: false,
      idcardChange: false,
      idtypeChange: false,
      sexChange: false,
      nameChange: false,
      if_read: true,
      m_url: "memberCenter",
    };
  },

  setup() {
    const IdCardType = ref("");
    const showPicker = ref(false); //底部弹出
    const showPcPicker = ref(false); //弹窗
    const columns = ["身份证", "护照", "港澳台通行证", "其他"];
    //实时用户信息
    const username = ref("");
    const userSex = ref("");
    const nickname = ref("");
    const IdCard = ref("");
    const userStatus = ref(0);
    let userphone = ref("");
    let useremail = ref("");
    const userID = ref();
    const respondmsg = ref("");
    // 记录原本的信息
    const oldphone = ref("");
    const oldemail = ref("");
    const oldIdCard = ref("");
    const oldIdCardType = ref("");
    const oldSex = ref("");
    const oldName = ref("");
    // 正则表达判断手机号和邮箱
    const phonePattern = /^1[3-9][0-9]{9}$/;
    const emailPattern = /^\w{3,}(\.\w+)*@[A-z0-9]+(\.[A-z]{2,5}){1,2}$/;

    const onConfirm = (value: any) => {
      IdCardType.value = value;
      showPicker.value = false;
      showPcPicker.value = false;
    };

    const ShowType = () => {
      if (userStatus.value === 1) {
        return;
      }
      if (document.body.clientWidth < 600) {
        showPicker.value = true;
      } else {
        showPcPicker.value = true;
      }
    };

    const changeTips = () => {
      if (userStatus.value === 1) {
        message.destroy();
        message.info("审核通过，不可修改此信息");
      }
    };
    const showmsg = () => {
      let url = "api/user/showmsg";
      fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        body: JSON.stringify({
          // id: userid,
        }),
        redirect: "follow",
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          //从后端接收信息
          respondmsg.value = v.msg;
          username.value = v.username;
          userphone.value = v.userphone;
          useremail.value = v.useremail;
          userSex.value = v.usersex;
          nickname.value = v.nickname;
          IdCard.value = v.id_card;
          IdCardType.value = v.id_card_type;
          userStatus.value = v.user_status;
          userID.value = v.userid;
          //记录原始信息
          oldphone.value = v.userphone;
          oldemail.value = v.useremail;
          oldIdCard.value = v.id_card;
          oldIdCardType.value = v.id_card_type;
          oldSex.value = v.usersex;
          oldName.value = v.username;
        });
    };

    onMounted(() => {
      showmsg();
    });
    return {
      username,
      userID,
      userphone,
      useremail,
      userSex,
      nickname,
      IdCard,
      userStatus,
      phonePattern,
      emailPattern,
      IdCardType,
      columns,
      onConfirm,
      showPicker,
      showPcPicker,
      ShowType,
      oldphone,
      oldemail,
      oldIdCard,
      oldIdCardType,
      oldSex,
      oldName,
      showmsg,
      changeTips,
    };
  },

  methods: {
    back() {
      this.$router.push("/memberCenter");
    },
    changemsg() {
      const respondmsg = ref("");
      let url = "api/user/changemsg";
      if (!(this.perror === "" && this.eerror === "")) {
        message.destroy();
        message.warn("填写信息有误");
      } else if (
        this.username === "" ||
        this.userID === "" ||
        this.userphone === "" ||
        this.nickname === "" ||
        this.IdCard === ""
      ) {
        message.destroy();
        message.warn("信息填写不完全");
      } else if (
        (this.emailChange ||
          this.phoneChange ||
          this.idcardChange ||
          this.IdCardType != this.oldIdCardType ||
          this.sexChange ||
          this.nameChange) &&
        this.perror === "" &&
        this.eerror === ""
      ) {
        fetch(url, {
          method: "POST",
          headers: new Headers({
            "Content-Type": "application/json",
          }),
          redirect: "follow",
          body: JSON.stringify({
            id: this.userid,
            phone: this.userphone,
            email: this.useremail,
            idcard: this.IdCard,
            idcardtype: this.IdCardType,
            sex: this.userSex,
            user_status: this.userStatus,
            name: this.username,
          }),
        })
          .then((v) => {
            return v.json();
          })
          .then((v) => {
            if (v.status == -999) {
              message.warn("账号未登录，请先登录");
              router.push("login");
              return;
            }

            if (v.status == -1999) {
              message.warn("管理员账号未登录，请先登录");
              router.push("login");
              return;
            }
            respondmsg.value = v.msg;
          });
        this.showmsg();
        message.success("修改成功");
        this.$router.push("/memberCenter");
      } else {
        message.destroy();
        message.info("信息未修改");
      }
    },
    checkmsg() {
      // 检测手机号
      if (!this.phonePattern.test(this.userphone)) {
        this.perror = "请输入正确的手机号";
      } else this.perror = "";
      // 检测邮箱
      if (!this.emailPattern.test(this.useremail) && this.useremail != "") {
        this.eerror = "请输入正确的电子邮箱";
      } else this.eerror = "";
      // 记录手机邮箱是否修改
      if (this.phonePattern.test(this.userphone) && this.oldphone != this.userphone) {
        this.phoneChange = true;
      } else {
        this.phoneChange = false;
      }
      if (this.emailPattern.test(this.useremail) && this.oldemail != this.useremail) {
        this.emailChange = true;
      } else {
        this.emailChange = false;
      }
      //记录其他信息是否修改
      if (this.oldIdCard != this.IdCard || this.oldName != this.username) {
        this.idcardChange = true;
        this.nameChange = true;
      } else {
        this.idcardChange = false;
        this.nameChange = false;
        this.nameError = "";
        this.idcardError = "";
      }
      //测试性别是否修改
      if (this.oldSex != this.userSex) {
        this.sexChange = true;
      } else {
        this.sexChange = false;
      }
    },
  },
});
</script>
