<template>
  <ManageNav />
  <div id="buildingg1"></div>
  <div class="bground">
    <van-config-provider :theme-vars="themeVars">
      <div v-for="u in datas1" :key="u.id" class="one">
        <div class="font">基本信息</div>
        <van-cell-group :inset="true">
          <van-field v-model="u.phone" label="手机号" class="filed" readonly />
          <van-field v-model="u.sex" label="性别" class="filed" :readonly="true" />
          <van-field v-model="u.idcard" class="filed" readonly >
            <template #label>
              {{u.idcard_type}}
            </template>
          </van-field>
          <van-field v-model="u.email" label="电子邮件" class="filed" readonly />

          <van-field v-model="u.education" label="学历" class="filed" readonly />
          <van-field
            v-model="u.birthday"
            label="出生年月日"
            maxlength="11"
            class="filed"
            :readonly="true"
          />
          <van-field v-model="u.nation" label="民族" class="filed" :readonly="true" />
          <van-field v-model="u.native_place" label="籍贯" class="filed" :readonly="true" />
          <van-field v-model="u.birth_place" label="出生地" class="filed" :readonly="true" />
          <van-field v-model="u.political" label="政治面貌" class="filed" :readonly="true" />

          <van-field v-model="u.overeseas_permanent" label="国（境）外永久居留权" class="filed" readonly />
          <van-field v-model="u.level" label="外语等级" class="filed" :readonly="true" />
          <van-field v-model="u.translation" label="是否能担任翻译" class="filed" :readonly="true" />
          <van-field v-model="u.married_status" label="婚姻状况" class="filed" :readonly="true" />

          <van-field v-model="u.home_phone" label="家庭电话" class="filed" :readonly="true" />
          <van-field v-model="u.address" label="通讯地址" class="filed" :readonly="true" />
          <van-field v-model="u.postal_code" label="邮政编码" class="filed" :readonly="true" />
          <van-field v-model="u.fax" label="传真" class="filed" :readonly="true" />
          <van-field v-model="u.hobby" label="兴趣爱好" class="filed" :readonly="true" />
          <van-field
            v-model="u.social_appointments"
            label="人大政协及其他社会兼职"
            class="filed"
            :readonly="true"
          />
          <van-field label="推荐人签章:" class="filed" :readonly="true">
              <template #input>
                <div class="imgWrapper">
                  <div
                      v-for="src in srcList[0]"
                      :key="src"
                      v-viewer="{ movable: false }"
                  >
                    <van-image fit="cover" :src="displayUrl + src" />
                  </div>
                </div>
            </template>
          </van-field>


          <van-field v-model="u.relation" label="与介绍人关系" :readonly="true" class="filed" />
          <van-field v-model="u.content" label="介绍单位意见" :readonly="true" class="filed" />
          <van-field label="推荐单位签章:" class="filed" :readonly="true">
            <template #input>
              <div class="imgWrapper">
                <div
                    v-for="src in srcList1[0]"
                    :key="src"
                    v-viewer="{ movable: false }"
                >
                  <van-image fit="cover" :src="displayUrl + src" />
                </div>
              </div>
            </template>
          </van-field>


        </van-cell-group>
        <div class="font">教育经历</div>
        <van-cell-group :inset="true">
          <van-field v-model="u.degree" label="学位" class="filed" :readonly="true" />
          <van-field v-model="u.major" label="学科/专业" class="filed" :readonly="true" />
          <van-field v-model="u.college" label="学校名称" class="filed" :readonly="true" />
          <van-field v-model="u.region" label="学校所在地和地区" :readonly="true" class="filed" />
          <van-field v-model="u.begin_time" label="起始在校时间" :readonly="true" class="filed" />
          <van-field v-model="u.end_time" label="结束在校时间" :readonly="true" class="filed" />
           <van-field
            v-model="u.study_type"
            class="filed"
            label="是否留学"
            :readonly="true"
          />
          <van-field class="filed" label="证明资料（学历学位、留学回国人员，毕业证书等资料）:" :readonly="true">
            <template #input>
              <div class="imgWrapper">
                <div
                    v-for="src in srcList2[0]"
                    :key="src"
                    v-viewer="{ movable: false }"
                >
                  <van-image fit="cover" :src="displayUrl + src" />
                </div>
              </div>
            </template>
          </van-field>

        </van-cell-group>
        <div class="font">工作经历</div>
        <van-cell-group :inset="true">
          <van-field v-model="u.position" label="职位" class="filed" :readonly="true" />
          <van-field v-model="u.position_title" label="职称" class="filed" :readonly="true" />
          <van-field v-model="u.professional" label="从事专业" :readonly="true" class="filed" />
          <van-field v-model="u.organization" label="机构名称" :readonly="true" class="filed" />
          <van-field v-model="u.organization_nature" label="机构性质" :readonly="true" class="filed" />
          <van-field v-model="u.or_telephone" label="机构电话" :readonly="true" class="filed" />
          <van-field v-model="u.w_region" label="机构所在国家和地区" :readonly="true" class="filed" />
          <van-field
            v-model="u.begin_time"
            label="工作起始时间"
            maxlength="11"
            class="filed"
            :readonly="true"
          />
          <van-field
            v-model="u.end_time"
            label="工作结束时间"
            class="filed"
            maxlength="11"
            :readonly="true"
          />
          <van-field
            v-model="jieguo"
            label="审核结果"
            class="filed"
            maxlength="11"
            :error="true"
            :readonly="true"
          />
          <van-field label="审核意见" class="filed">
            <template #input>
              <textarea v-model="u.opinion" textarea readonly class="throughTextarea"></textarea>
            </template>
          </van-field>
        </van-cell-group>
        <van-button type="primary" class="throughButton" @click="show = true">修改审核结果</van-button>
        <van-dialog
          class="dio1"
          v-model:show="show"
          show-cancel-button
          confirm-button-color="#1890FF"
          :close-on-click-overlay="true"
          @confirm="confirm1"
        >
          <div class="dio">是否确认更改审核结果为：未通过</div>
          <textarea v-model="book1" textarea class="textarea1" placeholder="输入审核意见..."></textarea>
        </van-dialog>
      </div>
    </van-config-provider>
  </div>
</template>
<style scoped>
.imgWrapper{
  display: flex;
  flex-wrap: wrap
}
@media screen and (max-width: 600px) {
  .dio1 {
    width: 320px;
    margin: auto;
  }
  .dio {
    width: 500px;
    margin-top: 30px;
    margin-bottom: 30px;
    margin-left: 45px;
  }
  .van-image{
    width: 100px;
    height: 100px;
    margin: 2px;
  }
  .van-cell-group {
    margin: 0;
    padding: 0.26rem 0.1rem;
  }
  .font {
    font-size: 16px;
    color: #8c99a0;
    margin: 0.8rem 0 0.26rem 0.1rem;
  }
  #buildingg1 {
    background: #f7f7f7;
    width: 100%;
    height: 100%;
    position: fixed;
    background-size: 100% 100%;
  }
  .throughTextarea {
    border: 0.0267rem solid rgb(202, 200, 200);
    margin-left: 0;
  }

  textarea {
    font-size: 0.4267rem;
    width: 80%;
    border-color: rgb(90, 175, 245);
    margin-left: 10%;
  }

  .textarea1 {
    margin-top: -0.2rem;
    margin-bottom: 0.3rem;
    border-color: #c0bfbf;
    height: 3rem;
  }
  .throughButton {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    width: 4rem;
    margin: 0.53rem auto;
  }
  .one {
    position: absolute;
    margin: 0 0.53rem;
  }
}
@media screen and (min-width: 600px) {
  .dio {
    width: 500px;
    margin-top: 30px;
    margin-bottom: 50px;
    margin-left: 45px;
  }
  .van-cell-group {
    margin: 0;
    padding: 10px;
  }
  .van-image{
    width: 120px;
    height: 120px;
    margin: 5px;
  }
  .font {
    font-size: 17px;
    color: #8c99a0;
    margin: 30px 0 10px 10px;
  }
  .bground {
    max-width: 750px;
    margin: auto;
  }
  #buildingg1 {
    margin-top: -20px;
    background: #f7f7f7;
    width: 100%;
    height: 100%;
    position: fixed;
    background-size: 100% 100%;
    z-index: -100;
  }
  .throughTextarea {
    border: 1.0012px solid rgb(202, 200, 200);
    margin-left: 0;
  }

  textarea {
    font-size: 16px;
    width: 80%;
    border-color: rgb(90, 175, 245);
    margin-left: 10%;
    margin-bottom: 10px;
    height: 80px;
  }

  .textarea1 {
    margin-top: -30px;
    border-color: #c0bfbf;
    height: 112.5px;
  }
  .throughButton {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    width: 200px;
    margin: 35px auto;
    font-size: 16px;
  }
  .one {
    padding-top: 10px;
    width: 100%;
  }
}
</style>

<script lang="ts">
import {
  defineComponent,
  ref,
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  onBeforeUnmount,
  onUnmounted,
  onActivated,
  onDeactivated,
  onErrorCaptured,
  onServerPrefetch,
  onScopeDispose,
  onRenderTracked,
  onRenderTriggered,
  reactive,
  watch,
} from 'vue';
import axios from 'axios';
import { Toast } from 'vant';
import { message } from 'ant-design-vue';
import ManageNav from "@/components/common/ManageNav.vue";
export default defineComponent({
  components: { ManageNav },
  setup() {
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

    const show = ref(false);
    const rate = ref(4);
    const slider = ref(50);
    // themeVars 内的值会被转换成对应 CSS 变量
    // 比如 sliderBarHeight 会转换成 `--van-slider-bar-height`
    const themeVars = {
      // cellGroupTitleFontSize: '.5333rem',
      navBarTextColor: '#FFFFFF',
      navBarIconColor: '#FFFFFF',
      navBarTitleTextColor: '#FFFFFF',
      navBarBackgroundColor: '#1890FF',
      navBarHeight: '1.7333rem',
      fieldPlaceholderTextColor: '#333333',
      fieldLabelWidth: '5em',
      dialogHeaderLineHeight: '2.5rem',
      dialogWidth: '350px',
    };
    const text = '2132656565';
    const value = ref('');

    const fileList = reactive([]);
    const getIdUrl = 'api/file/generateId';
    const downloadUrl = 'api/file/download';
    const displayUrl = 'http://mac.gzhu.edu.cn/alumni/resources/'; // 获取图片缩略图的地址

    let uid = '1'; // "uid"
    let fileId = '';
    const srcList = reactive<any[]>([]);
    const srcList1 = reactive<any[]>([]);
    const jieguo = '已通过';

    const srcList2 = reactive<any[]>([]);

    const counter = ref(0);
    watch(counter, (newValue, oldValue) => {
      if (+newValue !== 6) {
        console.log(oldValue);
        return;
      }
      console.log('it is 6');
    });
    const onClickLeft = () => history.back();
    return reactive({
      book1: '',
      srcList,
      srcList1,

      srcList2,
      jieguo,
      onClickLeft,
      websocket,

      fileList,
      displayUrl,
      fileId: 0,
      downloadUrl,
      show,
      getIdUrl,
      text,
      value,
      rate,
      slider,
      themeVars,
      datas1: [],
    });
  },
  beforeCreate() {
    // 在实例初始化之后，数据观测(data observer) 和 event/watcher 事件配置之前被调用。
    console.log('beforeCreate');
  },
  created() {
    // 实例已经创建完成之后被调用。在这一步，实例已完成以下的配置：数据观测(data observer)，
    //  属性和方法的运算， watch/event 事件回调。然而，挂载阶段还没开始，$el 属性目前不可见。
    console.log('created');
    this.onEdit();
  },
  methods: {
    async async1(rp: any, ru: any) {
      const form = new FormData();
      form.append('id', rp);
      const results = await axios.post(this.downloadUrl, form);
      if (results.data.statusCode === 2000) {
        this.srcList.push(results.data.details);
        this.async2(this.datas1[0].Id_Ini);
      }
    },
    async async2(t: any) {
      const form = new FormData();
      form.append('id', t);
      const results = await axios.post(this.downloadUrl, form);
      if (results.data.statusCode === 2000) {
        this.srcList1.push(results.data.details);
        this.async3(this.datas1[0].Id_Cer);
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

    confirm1() {
      let url = 'api/user/userchange';
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
          id: this.$route.params.id,
          sign: '1',
          book: this.book1,
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
          if (!v) {
            message.error('网络开小差了，请刷新重试！');
            return;
          }

          if (v.status !== 0) {
            message.error('网络开小差了，请刷新重试！');
            return;
          } else {
            this.transmit();
            message.success('用户审核状态更改成功');
            setTimeout(() => {
              this.$router.go(-1)
              this.$router.replace('/pendingAuditPage')
            }, 1000); //延时1000毫秒执行this.$router.push("/");
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },

    transmit() {
      let q = this.$route.params.id.toString();
      if (this.book1 == '') {
        this.book1 = '未通过';
      }
      const data = {
        ids: [q],
        title: '未通过',
        content: this.book1,
        time: Date.now(),
      };
      this.websocket.send(JSON.stringify(data));
    },
    onEdit() {
      let url = 'api/user/already';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: JSON.stringify({ id: this.$route.params.id }),
      })
        .then((v) => {
          console.log(v);
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
          if (!v) {
            message.error('网络开小差了，请刷新重试！');
            return;
          }

          if (v.status !== 0) {
            message.error('网络开小差了，请刷新重试！');
            return;
          } else {
            this.datas1 = v.data;
            console.log("data1----", this.datas1)
            this.async1(this.datas1[0].Id_Peo, this.datas1[0].Id_Ini);
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
});
</script>
