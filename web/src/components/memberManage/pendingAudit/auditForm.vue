<template>
  <div class="bg">
    <div class="infoList">
      <van-config-provider :theme-vars="themeVars">
        <div class="groupTitle">基本信息</div>
        <van-cell-group inset>
          <van-field v-model="formInfo.phone" label="手机号" readonly class="fieldFont" />
          <van-field v-model="formInfo.sex" label="性别" :readonly="true" class="fieldFont" />
          <van-field v-model="formInfo.idCard" readonly class="fieldFont" >
            <template #label>{{formInfo.idcard_type}}</template>
          </van-field>
          <van-field v-model="formInfo.email" label="电子邮件" readonly class="fieldFont" />

          <van-field v-model="essentialInfo.education" label="学历" readonly class="fieldFont" />
          <van-field
            v-model="formInfo.birthday"
            label="出生年月日"
            maxlength="10"
            :readonly="true"
            class="fieldFont"
          />
          <van-field v-model="essentialInfo.nation" label="民族" :readonly="true" class="fieldFont" />
          <van-field
            v-model="essentialInfo.native_place"
            label="籍贯"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.birth_place"
            label="出生地"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.political"
            label="政治面貌"
            :readonly="true"
            class="fieldFont"
          />

          <van-field
            v-model="essentialInfo.overseas_permanent"
            label="国（境）外永久居留权"
            readonly
            class="fieldFont"
          />
          <van-field v-model="essentialInfo.level" label="外语等级" :readonly="true" class="fieldFont" />
          <van-field
            v-model="essentialInfo.translation"
            label="是否能担任翻译"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.married_status"
            label="婚姻状况"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.home_phone"
            label="家庭电话"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.address"
            label="通讯地址"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="essentialInfo.postal_code"
            label="邮政编码"
            :readonly="true"
            class="fieldFont"
          />
          <van-field v-model="essentialInfo.fax" label="传真" :readonly="true" class="fieldFont" />
          <van-field v-model="essentialInfo.hobby" label="兴趣爱好" :readonly="true" class="fieldFont" />
          <van-field
            v-model="essentialInfo.social_appointments"
            label="人大政协及其他社会兼职"
            :readonly="true"
            class="fieldFont"
          />
          <van-field label="推荐人签章:" :readonly="true" class="fieldFont" >
            <template #input>
              <div class="imgWrapper">
                <div v-for="src in srcList[0]" :key="src" v-viewer="{ movable: false }">
                  <van-image fit="cover" :src="displayUrl + src" />
                </div>
              </div>
            </template>
          </van-field>

          <van-field
            v-model="supportialMaterial.relation"
            label="与介绍人关系"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="supportialMaterial.content"
            label="介绍单位意见"
            :readonly="true"
            class="fieldFont"
          />
          <van-field label="推荐单位签章:" :readonly="true" class="fieldFont">
            <template #input>
              <div class="imgWrapper">
                <div v-for="src in srcList1[0]" :key="src" v-viewer="{ movable: false }">
                  <van-image fit="cover" :src="displayUrl + src" />
                </div>
              </div>
            </template>
          </van-field>
        </van-cell-group>
        <div class="groupTitle">教育经历</div>
        <van-cell-group inset style="margin: 0">
          <van-field v-model="educationInfo.degree" label="学位" :readonly="true" class="fieldFont" />
          <van-field
            v-model="educationInfo.major"
            label="学科/专业"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="educationInfo.college"
            label="学校名称"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="educationInfo.region"
            label="学校所在地和地区"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="educationInfo.begin_time"
            label="起始在校时间"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="educationInfo.end_time"
            label="结束在校时间"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="formInfo.study_type"
            label="是否留学"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
              label="证明资料（学历学位、留学回国人员，毕业证书等资料）:"
              :readonly="true"
              class="fieldFont">
            <template #input>
              <div class="imgWrapper">
                <div v-for="src in srcList2[0]" :key="src" v-viewer="{ movable: false }" class="img">
                  <van-image fit="cover" :src="displayUrl + src" />
                </div>
              </div>
            </template>
          </van-field>

        </van-cell-group>
        <div class="groupTitle">工作经历</div>
        <van-cell-group inset style="margin: 0">
          <van-field v-model="workInfo.position" label="职位" :readonly="true" class="fieldFont" />
          <van-field
            v-model="workInfo.position_title"
            label="职称"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.professional"
            label="从事专业"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.organization"
            label="机构名称"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.organization_nature"
            label="机构性质"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.or_telephone"
            label="机构电话"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.region"
            label="机构所在国家和地区"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.begin_time"
            label="工作起始时间"
            maxlength="11"
            :readonly="true"
            class="fieldFont"
          />
          <van-field
            v-model="workInfo.end_time"
            label="工作结束时间"
            maxlength="11"
            :readonly="true"
            class="fieldFont"
          />
        </van-cell-group>
        <div class="groupTitle">审核意见</div>
        <van-cell-group inset style="margin: 0">
          <van-field
            v-model="comment"
            rows="5"
            :autosize="true"
            label="审核意见"
            type="textarea"
            placeholder="请输入审核意见"
            @update:model-value="input"
            class="fieldFont"
          />
        </van-cell-group>
        <div class="btn">
          <van-button type="primary" class="passBtn" @click="passed">通过</van-button>
          <van-button type="danger" class="unpassBtn" @click="failPassed" plain>退回修改</van-button>
        </div>
      </van-config-provider>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, toRefs } from 'vue';
// import moment from 'moment';
import axios from 'axios';
import { message } from 'ant-design-vue';
import { Toast } from 'vant';
import dayjs from "dayjs";
export default defineComponent({

  mounted() {
    this.getUserInfo();
  },
  props: {
    userId: Number
  },
  setup(props) {

    const userId = props.userId
    console.log("userID", userId)
    const websocket = new WebSocket(
      'ws://mac.gzhu.edu.cn/alumni/api/message/ws?id=-1'
    );
    const wsconnect = () => {
      // websocket = new WebSocket("ws://localhost:9090/api/message/ws?id=1");
      websocket.onopen = () => {
        console.log('connected');
      };
      websocket.onmessage = (e) => {
        // console.log(e);
        console.log(JSON.parse(e.data));
      };
      websocket.onclose = (e) => {
        console.log(e);
      };
    };
    const themeVars = {
      fieldPlaceholderTextColor: '#333333',
      cellGroupInsetTitlePadding: '0',
      cellGroupTitleFontSize: '.4533rem',
      fieldWordLimitFontSize: '.4267rem',
      fieldLabelWidth: '5em',
    };
    const comment = ref('审核通过');
    const focus = ref(false);
    const status = ref(3);
    const downloadUrl = 'api/file/download';
    const displayUrl = 'http://mac.gzhu.edu.cn/alumni/resources/'; // 获取图片缩略图的地址

    //推荐人签章图片
    const filePeople = reactive([]);
    const srcListPeople = reactive([]);
    //推荐单位签章图片
    const fileUnit = reactive([]);
    const srcListUnit = reactive([]);
    //教育信息证明资料图片
    const fileMaterial = reactive([]);
    const srcListMaterial = reactive([]);

    const dateTime = (date: any) => {
      // moment(date).utc().format('YYYY/MM/DD');
      return dayjs(date).format('YYYY-MM-DD');
    };

    const srcList = reactive<any[]>([]);
    const srcList1 = reactive<any[]>([]);
    const srcList2 = reactive<any[]>([]);
    const formInfo = reactive<any>([])
    const essentialInfo = reactive<any>([])
    const educationInfo = reactive<any>([])
    const supportialMaterial = reactive<any>([])
    const workInfo = reactive<any>([])
    return reactive({
      ...toRefs(props),
      userId,
      // formInfo: [] as any,
      formInfo,
      essentialInfo,
      educationInfo,
      supportialMaterial,
      workInfo,
      // essentialInfo: [] as any,
      // educationInfo: [] as any,
      // supportialMaterial: [] as any,
      // workInfo: [] as any,
      themeVars,
      comment,
      focus,
      dateTime,
      status,
      filePeople,
      srcListPeople,
      fileUnit,
      srcListUnit,
      fileMaterial,
      srcListMaterial,
      displayUrl,
      downloadUrl,
      srcList,
      srcList1,
      srcList2,
      datas1: [],
      websocket,
    });
  },
  methods: {
    async async1(rp: any, ru: any) {
      const form = new FormData();
      form.append('id', rp);
      console.log('form:', form.getAll('id'));
      const results = await axios.post(this.downloadUrl, form);
      console.log("图片结果", results)
      if (results.data.statusCode === 2000) {
        this.srcList.push(results.data.details);
        console.log('srcList:', this.srcList);
        this.async2(ru);
      }
    },
    async async2(t: any) {
      const form = new FormData();
      form.append('id', t);
      const results = await axios.post(this.downloadUrl, form);
      if (results.data.statusCode === 2000) {
        this.srcList1.push(results.data.details);
        this.async3(this.fileMaterial);
      }
    },
    async async3(t: any) {
      const form = new FormData();
      form.append('id', t);
      const results = await axios.post(this.downloadUrl, form);
      if (results.data.statusCode === 2000) {
        this.srcList2.push(results.data.details);
      }
    },

    showToast: function () {
      let second = 3;
      const toast = Toast.loading({
        duration: 0,
        forbidClick: true,
        message: '提交成功\n3秒后返回',
      });
      const timer = setInterval(() => {
        second--;
        if (second) {
          toast.message = `提交成功\n${second}秒后返回`;
        } else {
          clearInterval(timer);
          toast.clear();
        }
      }, 1000);
      let q = this.userId.toString();
      const data = {
        ids: [q],
        title: '审核结果',
        content: this.comment,
        time: Date.now(),
      };

      this.websocket.send(JSON.stringify(data));
    },
    failPassed: function () {
      if (!this.comment) {
        message.warning('退回修改意见不能为空！');
        return;
      }
      if (this.comment == '审核通过') {
        message.warning('请输入退回修改意见');
        this.comment = '';
        return;
      }

      this.status = 2;
      this.submit();
      setTimeout(() => {
        this.$router.go(-1);
        this.$router.replace('/pendingAuditPage')
      }, 3000);
    },

    input: function () {
      if (this.comment) {
        this.errorMsg = '';
      }
    },

    passed: function () {
      if (this.comment == '') {
        message.warning('审核意见不能为空');
        return;
      }
      this.status = 1;
      this.submit();
      setTimeout(() => {
        this.$router.go(-1);
        this.$router.replace('/pendingAuditPage')
      }, 1000);
    },
    getUserInfo: function asyn() {
      console.log("用户Id", this.userId)
      let q = `id=${this.userId}`;

      let url = `api/admin/userInfo?${q}`;
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
          console.log(v)

          console.log(v)
          this.formInfo = v.data;
          console.log("formInfo", this.formInfo)
          this.essentialInfo = v.data['essential_info'];
          this.educationInfo = v.data['education_info'];
          this.supportialMaterial = v.data['supportial_material'];
          this.workInfo = v.data['work_info'];

          this.filePeople = v.data['supportial_material'].Id_Peo;
          this.fileUnit = v.data['supportial_material'].Id_Ini;
          this.fileMaterial = v.data.file_id;

          this.async1(this.filePeople, this.fileUnit);
        });
    },
    submit: function () {
      let query = [];
      query.push(`id=${this.userId}`);
      query.push(`opinion=${this.comment}`);
      query.push(`status=${this.status}`);
      let q = query.join('&');
      let url = `api/admin/opinionSubmit?${q}`;
      fetch(url, {
        method: 'PUT',
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

          if (v.status == 200) {
            //this.showToast();
            message.success('审核结果提交成功');
            let q = this.userId.toString();
            const data = {
              ids: [q],
              title: '审核结果',
              content: this.comment,
              time: Date.now(),
            };

            this.websocket.send(JSON.stringify(data));

            return;
          } else {
            message.error('审核结果提交失败，请刷新重试');
            return;
          }
        });
    },
  },
});
</script>

<style lang="less" scoped>
.bg {
  background: #f7f7f7;
  display: flex;
  justify-content: center;
  width: 100%;
}
.imgWrapper{
  display: flex;
  flex-wrap: wrap
}
/* 移动端 */
@media screen and (max-width: 600px) {
  .van-cell-group {
    margin: 0;
    padding: 0.26rem 0.1rem;
  }
  .van-image{
    width: 100px;
    height: 100px;
    margin: 2px;
  }
  .groupTitle {
    font-size: 16px;
    color: #8c99a0;
    margin: 0.8rem 0 0.26rem 0.1rem;
  }
  .infoList {
    padding: 0 0.5333rem;
  }
  .btn {
    margin: auto;
    display: flex;
    flex-direction: column;
    width: 4rem;
    padding-bottom: 0.53rem;
    .van-button {
      margin-top: 0.5rem;
    }
  }
}
/* pc端 */
@media screen and (min-width: 600px) {
  .groupTitle {
    font-size: 17px;
    color: #8c99a0;
    margin: 30px 0 10px 10px;
  }
  .van-cell-group {
    margin: 0;
    padding: 10px;
  }
  .infoList {
    margin: 0 20px;
    width: 750px;
    .van-cell-group {
      margin: 0;
      padding: 10px;
    }
  }
  .van-image{
    width: 120px;
    height: 120px;
    margin: 5px;
  }
  .btn {
    margin: auto;
    width: 400px;
    display: flex;
    justify-content: space-evenly;
    padding-bottom: 50px;
    .van-button {
      margin-top: 35px;
      font-size: 16px;
      width: 150px;
    }
  }
}
</style>

