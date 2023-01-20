<template>
  <MemberNav/>
  <div class="entry">
    <div class="card">
      <div class="input">
        <van-cell-group>
          <van-form>
            <van-field
              label="职位"
              label-align="right"
              v-model="position"
              clearable
              autosize
              placeholder="请输入"
            />
            <van-field
              label="职称"
              label-align="right"
              v-model="profTitle"
              clearable
              autosize
              placeholder="请输入"
            />
            <van-field
              label="从事专业"
              label-align="right"
              v-model="professial"
              clearable
              autosize
              placeholder="请输入"
            />
            <van-field
              label="机构名称"
              label-align="right"
              v-model="orgName"
              clearable
              autosize
              placeholder="请输入"
            />
            <van-field
              label="机构性质"
              label-align="right"
              v-model="orgCharacter"
              clearable
              autosize
              placeholder="请输入"
            />
            <van-field
              label="机构电话"
              label-align="right"
              v-model="orgPhone"
              clearable
              autosize
              placeholder="请输入"
              :rules="[{ pattern, message: '请输入正确的内容' }]"
            />
            <van-field
              label="机构所在国家和地区"
              label-align="right"
              v-model="orgPlace"
              clearable
              autosize
              placeholder="请输入"
            />
          </van-form>

          <!-- 时间选择器 -->

          <div class="timeGroup">
            <div>
              <van-field
                v-model="startTime"
                label-align="right"
                placeholder="请选择"
                readonly
                label="开始工作年月"
                @click="showPopFnStart()"
              />
              <van-popup
                v-model:show="startTimeShow"
                position="bottom"
                :style="{ height: '45%' }"
              >
                <van-datetime-picker
                  v-model="startDate"
                  type="date"
                  @change="changeFnStart()"
                  @confirm="confirmFnStart()"
                  @cancel="cancelFnStart()"
                />
              </van-popup>
            </div>
            <div>
              <van-field
                v-model="endTime"
                placeholder="请选择"
                label-align="right"
                readonly
                label="结束工作年月"
                @click="showPopFnEnd()"
              />
              <van-popup
                v-model:show="endTimeShow"
                position="bottom"
                :style="{ height: '45%' }"
              >
                <van-datetime-picker
                  v-model="endDate"
                  type="date"
                  @change="changeFnEnd()"
                  @confirm="confirmFnEnd()"
                  @cancel="cancelFnEnd()"
                />
              </van-popup>
            </div>
          </div>
        </van-cell-group>
        <van-button class="btn" type="primary" @click="postMsg">提交</van-button>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
@media screen and (max-width: 600px) {
  .entry {
    margin: 0 0.53rem;
    padding-top: 0.53rem;
    .van-cell-group{
      margin: 0;
      padding: 0.26rem 0.1rem;
    }
    .timeGroup {
      display: flex;
      flex-direction: column;
    }
    .btn{
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 4rem;
      margin: 0.53rem auto;
    }
  }
}

@media screen and (min-width: 600px) {
  .entry {
    width: 750px;
    margin: 0 auto;
    padding-top: 50px;
    padding-bottom: 100px;
    .van-cell-group {
      margin: 0;
      padding: 10px;
    }
    .timeGroup {
      display: flex;
      flex-direction: column;
    }
    .btn {
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 200px;
      margin: 35px auto;
      font-size: 16px;
    }
  }
}
</style>

<script lang="ts">
import { message } from 'ant-design-vue'
import { defineComponent, reactive, } from "vue";
import MemberNav from "@/components/common/MemberNav.vue";

export default defineComponent({
  name: "VueComponentSkeleton",

  props: {},
  components: {
    MemberNav
  },

  setup() {
    // 手机号码正则表达式
    const pattern = /(^1[3456789]\d{9}$)/;

    return reactive({
      pattern,

      // userId: localStorage.getItem("userID"),
      position: "",
      profTitle: "",
      professial: "",
      orgName: "",
      orgCharacter: "",
      orgPhone: "",
      orgPlace: "",

      msg: "",
      startDate: new Date(),
      changeDateStart: new Date(),
      startTimeShow: false, // 用来显示弹出层
      startTime: "",

      endDate: new Date(),
      changeDateEnd: new Date(),
      endTimeShow: false,
      endTime: "",

      //post数据
      startTimePost: "",
      endTimePost: "",
    });
  },

  methods: {

    showPopFnStart() {
      this.startTimeShow = true;
    },

    //开始工作年月日时间
    changeFnStart() {
      // 值变化时触发
      this.changeDateStart = this.startDate; // Tue Sep 08 2020 00:00:00 GMT+0800 (中国标准时间)
    },

    confirmFnStart() {
      // 确定按钮
      this.startTime = this.timeFormat(this.startDate);
      this.startTimeShow = false;
      this.startTimePost = this.timeFormatPost(this.startDate);
    },

    cancelFnStart() {
      this.startTimeShow = true;
    },

    //结束工作年月日时间
    showPopFnEnd() {
      this.endTimeShow = true;
    },

    changeFnEnd() {
      // 值变化是触发
      this.changeDateEnd = this.endDate; // Tue Sep 08 2020 00:00:00 GMT+0800 (中国标准时间)
    },
    confirmFnEnd() {
      // 确定按钮
      this.endTime = this.timeFormat(this.endDate);
      this.endTimeShow = false;
      this.endTimePost = this.timeFormatPost(this.endDate); //拿到字符串2000-00-00形式的时间串
    },
    cancelFnEnd() {
      this.endTimeShow = true;
    },

    timeFormatPost(time) {
      let year = time.getFullYear();
      let month = time.getMonth() + 1;
      let day = time.getDate();
      return year + "-" + month + "-" + day;
    },

    timeFormat(time) {
      // 时间格式化 2019-09-08
      let year = time.getFullYear();
      let month = time.getMonth() + 1;
      let day = time.getDate();
      return year + "年" + month + "月" + day + "日";
    },

    postMsg() {
      if (
        this.position == "" ||
        this.profTitle == "" ||
        this.professial == "" ||
        this.orgName == "" ||
        this.orgCharacter == "" ||
        this.orgPhone == "" ||
        this.orgPlace == "" ||
        this.startTimePost == "" ||
        this.endTimePost == ""
      ) {
        // Toast.fail("请将信息补充完全");
        message.warning("请补全信息");
        return;
      }

      let reg_phone = /^1[0-9]{10}$/;

      if (!reg_phone.test(this.orgPhone)) {
        message.warning("手机号码格式错误");
        return;
      }

      let url = "api/user/postworkmsg";

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
          position: this.position,
          position_title: this.profTitle,
          organization: this.orgName,
          organization_nature: this.orgCharacter,
          professional: this.professial,
          begin_time: this.startTimePost,
          end_time: this.endTimePost,
          region: this.orgPlace,
          or_telephone: this.orgPhone,
        }),
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
          if (v.status != 0) {
            message.info("提交失败，请重试");
            return;
          }
          message.success("提交成功");
          this.$router.go(-1)
          this.$router.replace('/PersonalResueme')
        });
    },
  },

  mounted() {
    this.timeFormat(new Date());
  },
});
</script>
