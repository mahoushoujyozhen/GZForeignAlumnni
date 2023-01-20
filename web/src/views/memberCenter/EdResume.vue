<template>
  <MemberNav />
  <div class="entry">
    <div class="card">
      <div class="input">
        <van-cell-group>
          <van-field
            label="学位"
            label-align="right"
            v-model="degree"
            clearable
            autosize
            placeholder="请输入"
          />
          <van-field
            label="学科专业"
            label-align="right"
            v-model="major"
            clearable
            autosize
            placeholder="请输入"
          />
          <van-field
            label="学校名称"
            label-align="right"
            v-model="college"
            clearable
            autosize
            placeholder="请输入"
          />
          <van-field
            label="请输入学校所在国家和地区"
            label-align="right"
            v-model="region"
            clearable
            autosize
            placeholder="请输入"
          />

          <van-field
            v-model="startTime"
            placeholder="请选择"
            readonly
            label="开始工作年月"
            label-align="right"
            @click="showPopFnStart()"
          />
          <van-popup v-model:show="startTimeShow" position="bottom" :style="{ height: '45%' }">
            <van-datetime-picker
              v-model="startDate"
              type="date"
              @change="changeFnStart()"
              @confirm="confirmFnStart()"
              @cancel="cancelFnStart()"
            />
          </van-popup>

          <van-field
            v-model="endTime"
            placeholder="请选择"
            readonly
            label="结束工作年月"
            label-align="right"
            @click="showPopFnEnd()"
          />
          <van-popup v-model:show="endTimeShow" position="bottom" :style="{ height: '45%' }">
            <van-datetime-picker
              v-model="endDate"
              type="date"
              @change="changeFnEnd()"
              @confirm="confirmFnEnd()"
              @cancel="cancelFnEnd()"
            />
          </van-popup>

          <van-field label="是否留学" label-align="right">
            <template #input>
              <van-radio-group v-model="studyType">
                <van-radio class="radio" name="是">是</van-radio>
                <van-radio class="radio" name="否">否</van-radio>
              </van-radio-group>
            </template>
          </van-field>
          <fileWork @getFileId="get_file_id"></fileWork>
        </van-cell-group>
        <van-button class="btn" type="primary" @click="postedumsg">提交</van-button>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
//小于600px宽度的屏幕，移动端
@media screen and (max-width: 600px) {
  .entry {
    margin: 0 0.53rem;
    padding-top: 0.53rem;
    .van-cell-group {
      margin: 0;
      padding: 0.26rem 0.1rem;
    }
    .btn {
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 4rem;
      margin: 0.53rem auto;
    }
  }
  .radio {
    margin-bottom: 0.11rem;
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
    .btn {
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      width: 200px;
      margin: 35px auto;
      font-size: 16px;
    }
  }
   .radio {
    margin-bottom: 4px;
  }
}
</style>
<script lang="ts">
import { defineComponent, onMounted, reactive, ref } from "vue";
import fileWork from "../../components/memberCenter/fileWork.vue";
import { message } from "ant-design-vue";
import MemberNav from "@/components/common/MemberNav.vue";

export default defineComponent({
  components: {
    MemberNav,
    // TopNav,
    // Head,
    fileWork,
  },

  setup() {
    //表单内容正则表达式

    return reactive({

      studyType: "是",

      file_id: "",
      fileList: "",
      degree: "",
      major: "",
      college: "",
      region: "",

      startDate: new Date(),
      changeDateStart: new Date(),
      startTimeShow: false, // 用来显示弹出层
      startTime: "",

      endDate: new Date(),
      changeDateEnd: new Date(),
      endTimeShow: false,
      endTime: "",

      // id: localStorage.getItem("userID"),

      //post数据
      startTimePost: "",
      endTimePost: "",
    });
  },
  methods: {
    back() {
      this.$router.push("/PersonalResueme");
    },
    showPopFnStart() {
      this.startTimeShow = true;
    },

    //开始工作年月日时间
    changeFnStart() {
      // 值变化是触发
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

    //时间格式化
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

    get_file_id(fileId) {
      this.file_id = fileId;
    },
    postedumsg() {
      // let file_id = window.localStorage.getItem("file_id");

      if (
        this.degree == "" ||
        this.major == "" ||
        this.college == "" ||
        this.region == "" ||
        this.startTimePost == "" ||
        this.endTimePost == ""
      ) {
        //Toast.fail("请将信息补充完全");
        message.warning("请补全信息");
        return;
      }
      let url = "api/user/postedumsg";
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
          file_id: this.file_id,
          studyType:this.studyType,
          college: this.college,
          major: this.major,
          degree: this.degree,
          region: this.region,
          begin_time: this.startTimePost,
          end_time: this.endTimePost,
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          if (v.status == -999) {
            message.warn("账号未登录，请先登录")
            this.$router.push('login')
            return
          }

          if (v.status == -1999) {
            message.warn("管理员账号未登录，请先登录")
            this.$router.push('login')
            return
          }
          if (v.status != 0) {
            console.log("err:" + v.status + v.msg);
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
