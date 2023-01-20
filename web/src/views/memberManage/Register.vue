<template>
  <van-config-provider :theme-vars="themeVars">
    <van-form @submit="onSubmit" scroll-to-error:true class="basic">
      <div class="title">基本信息</div>
      <van-cell-group inset>
        <van-field
            v-model="Must.openName"
            type="text"
            label="用户名"
            placeholder="用户名"
            :rules="[{ required: true, message: '请输入用户名' }]"
            required
        />
        <van-field
            v-model="Must.password"
            type="password"
            label="密码"
            placeholder="密码"
            :rules="[{ required: true, message: '请输入密码' }]"
            required
        />
        <van-field
            v-model="Must.name"
            type="text"
            name="姓名"
            label="姓名"
            placeholder="姓名"
            :rules="[{ required: true, message: '请输入姓名' }]"
            required
        />
        <!-- //手机 -->
        <van-field
            v-model="Must.phone"
            type="tel"
            name="手机"
            label="手机"
            placeholder="手机"
            :rules="[
            {
              required: true,
              validator: patternPhone,
              message: '请输入手机号',
              trigger: 'onBlur',
            },
          ]"
            :error-message="errMsgPhone"
            required
        />
        <!-- //性别 -->
        <van-field
            name="radio"
            label="性别"
            :rules="[{ required: true, message: '请选择性别' }]"
        >
          <template #input>
            <van-radio-group v-model="NotMust.sex">
              <van-radio name="男" icon-size="15px">男</van-radio>
              <van-radio name="女" icon-size="15px">女</van-radio>
            </van-radio-group>
          </template>
        </van-field>
        <!-- //证件类型选择器 -->
        <van-field
            v-model="result"
            is-link
            readonly
            name="picker"
            label="证件类型"
            placeholder="点击选择证件类型"
            @click="showPicker = true"
            :rules="[{ required: true, message: '请选择证件号类型' }]"
            required
        />
        <van-popup :show="showPicker" round position="bottom">
          <van-picker
              :columns="columns"
              @confirm="onConfirm"
              @cancel="showPicker = false"
          />
        </van-popup>

        <!-- //证件号 -->
        <van-field
            v-model="Must.idcard"
            type="number"
            name="证件号"
            label="证件号"
            placeholder="证件号"
            :rules="[{ required: true, message: '请填写证件号' }]"
            required
        />
        <!-- //电子邮件 -->
        <van-field
            v-model="NotMust.email"
            type="email"
            name="电子邮箱"
            label="电子邮箱"
            placeholder="电子邮箱"
            :rules="[
            {
              validator: patternEmail,
              trigger: 'onChange',
            },
          ]"
            :error-message="errMsgEmail"
        />
        <!-- //学位 -->

        <!-- //学历 -->
        <van-field
            v-model="essential_info.education"
            type="text"
            name="学历"
            label="学历"
            placeholder="学历（选填）"
        />
        <!-- //出生年月日 -->
        <van-field
            v-model="NotMust.birthDay"
            type="date"
            label="出生年月"
            placeholder="请输入出生年月（选填）"
        />
        <!-- 民族 -->
        <van-field
            v-model="essential_info.nation"
            type="text"
            name="民族"
            label="民族"
            placeholder="民族（选填）"
        />
        <!-- 籍贯 -->
        <van-field
            v-model="essential_info.native_place"
            type="text"
            name="籍贯"
            label="籍贯"
            placeholder="籍贯（选填）"
        />
        <!-- 出生地 -->
        <van-field
            v-model="essential_info.birth_place"
            type="text"
            name="出生地"
            label="出生地"
            placeholder="出生地（选填）"
        />
        <!-- 政治面貌 -->
        <van-field
            v-model="essential_info.political"
            type="text"
            name="政治面貌"
            label="政治面貌"
            placeholder="政治面貌（选填）"
        />
        <!-- 国籍 -->
        <van-field
            v-model="essential_info.nationality"
            type="text"
            name="国籍"
            label="国籍"
            placeholder="国籍（选填）"
        />
        <!--  -->
        <van-field name="resident" label="国境外永久居留权">
          <template #input>
            <van-radio-group v-model="essential_info.overeseas_permanent">
              <van-radio name="有" icon-size="15px">有</van-radio>
              <van-radio name="无" icon-size="15px">无</van-radio>
            </van-radio-group>
          </template>
        </van-field>
        <!-- 外语等级 -->
        <van-field
            v-model="essential_info.level"
            type="text"
            name="外语等级"
            label="外语等级"
            placeholder="外语等级（选填）"
        />
        <!--  -->
        <van-field name="translate" label="能否担任翻译工作">
          <template #input>
            <van-radio-group v-model="essential_info.translation">
              <van-radio name="能" icon-size="15px">能</van-radio>
              <van-radio name="不能" icon-size="15px">不能</van-radio>
            </van-radio-group>
          </template>
        </van-field>
        <!--  -->
        <van-field name="marriage" label="婚姻状况">
          <template #input>
            <van-radio-group v-model="essential_info.married_status">
              <van-radio name="已婚" icon-size="15px">已婚</van-radio>
              <van-radio name="未婚" icon-size="15px">未婚</van-radio>
            </van-radio-group>
          </template>
        </van-field>
        <van-field
            v-model="essential_info.home_phone"
            type="text"
            name="家庭电话"
            label="家庭电话"
            placeholder="家庭电话（选填）"
        />
        <van-field
            v-model="essential_info.address"
            type="text"
            name="通讯地址"
            label="通讯地址"
            placeholder="通讯地址（选填）"
        />
        <van-field
            v-model="essential_info.postal_code"
            type="text"
            name="邮政编码"
            label="邮政编码"
            placeholder="邮政编码（选填）"
        />
        <van-field
            v-model="essential_info.fax"
            type="text"
            name="传真"
            label="传真"
            placeholder="传真（选填）"
        />
        <van-field
            v-model="essential_info.hobby"
            type="text"
            name="兴趣爱好"
            label="兴趣爱好"
            placeholder="兴趣爱好（选填）"
        />
        <van-field
            v-model="essential_info.social_appointments"
            type="text"
            name="人大政协及其他社会兼职"
            label="人大政协及其他社会兼职"
            placeholder="人大政协及其他社会兼职（选填）"
        />

        <!-- 推荐人签章文件上传 -->
        <van-field name="uploader1" label="推荐人签章">
          <template #input>
            <van-uploader
                v-model="fileList_Peo"
                :after-read="afterRead_Peo"
                :before-delete="beforeDelete_Peo"
                :max-size="10000 * 1024"
                @oversize="onOversize"
            />
          </template>
        </van-field>
        <van-field
            v-model="supporting_materials.relation"
            type="text"
            name="与介绍人关系"
            label="与介绍人关系"
            placeholder="与介绍人关系（选填）"
        />
        <van-field
            v-model="supporting_materials.content"
            type="text"
            name="介绍单位意见"
            label="介绍单位意见"
            placeholder="介绍单位意见（选填）"
        />

        <!-- 介绍单位签章文件上传 -->
        <van-field name="uploader1" label="推荐单位签章">
          <template #input>
            <van-uploader
                v-model="fileList_Ini"
                :after-read="afterRead_Ini"
                :before-delete="beforeDelete_Ini"
                :max-size="10000 * 1024"
                @oversize="onOversize"
            />
          </template>
        </van-field>
      </van-cell-group>
      <div class="title">教育信息</div>
      <van-cell-group inset>
        <van-field
            v-model="education_info.degree"
            type="text"
            name="学位"
            label="学位"
            placeholder="学位（选填）"
        />
        <van-field
            v-model="education_info.major"
            type="text"
            name="学科专业"
            label="学科专业"
            placeholder="学科专业（选填）"
        />
        <van-field
            v-model="education_info.college"
            type="text"
            name="学校名称"
            label="学校名称"
            placeholder="学校名称（选填）"
        />

        <van-field
            v-model="education_info.region"
            type="text"
            name="学校所在国家和地区"
            label="学校所在国家和地区"
            placeholder="学校所在的国家和地区（选填）"
        />
        <van-field
            v-model="education_info.begin_time"
            type="date"
            name="开始在校学习年月"
            label="开始在校学习年月"
            placeholder="开始在校学习年月（选填）"
        />
        <van-field
            v-model="education_info.end_time"
            type="date"
            name="毕业年月"
            label="毕业年月"
            placeholder="毕业年月（选填）"
        />
        <van-field name="abroad" label="是否留学">
          <template #input>
            <van-radio-group v-model="NotMust.study_type">
              <van-radio name="是" icon-size="15px">是</van-radio>
              <van-radio name="否" icon-size="15px">否</van-radio>
            </van-radio-group>
          </template>
        </van-field>

        <!-- 教育经历的证明材料 -->
        <van-field
            name="uploader2"
            label="证明资料（学历学历、留学回国人员、毕业证书等证明资料）"
        >
          <template #input>
            <van-uploader
                v-model="fileList_Cer"
                style="size: 120px"
                :after-read="afterRead_Cer"
                :before-delete="beforeDelete_Cer"
                :max-size="10000 * 1024"
                @oversize="onOversize"
            />
          </template>
        </van-field>
      </van-cell-group>
      <div class="title">工作经历</div>
      <van-cell-group inset>
        <van-field
            v-model="work_info.position"
            type="text"
            name="职位"
            label="职位"
            placeholder="职位（选填）"
        />
        <van-field
            v-model="work_info.position_title"
            type="text"
            name="职称"
            label="职称"
            placeholder="职称（选填）"
        />
        <van-field
            v-model="work_info.professional"
            type="text"
            name="从事专业"
            label="从事专业"
            placeholder="从事专业（选填）"
        />
        <van-field
            v-model="work_info.organization"
            type="text"
            name="机构名称"
            label="机构名称"
            placeholder="机构名称（选填）"
        />
        <van-field
            v-model="work_info.organization_nature"
            type="text"
            name="机构性质"
            label="机构性质"
            placeholder="机构性质（选填）"
        />
        <van-field
            v-model="work_info.or_telephone"
            type="text"
            name="机构电话"
            label="机构电话"
            placeholder="机构电话（选填）"
        />
        <van-field
            v-model="work_info.region"
            type="text"
            name="机构所在国家和地区"
            label="机构所在国家和地区"
            placeholder="机构所在国家和地区（选填）"
        />
        <van-field
            v-model="work_info.begin_time"
            type="date"
            name="开始工作年月"
            label="开始工作年月"
            placeholder="开始工作年月（选填）"
        />
        <van-field
            v-model="work_info.end_time"
            type="date"
            name="结束工作年月"
            label="结束工作年月"
            placeholder="结束工作年月（选填）"
        />
      </van-cell-group>
      <van-popup :show="Must.loading" round>
        <div class="wrapper">
          <div class="block">
            <van-loading class="loading" type="spinner" size="24px" vertical
            >提交信息中...</van-loading
            >
          </div>
        </div>
      </van-popup>
    </van-form>
  </van-config-provider>
  <van-button class="button" block type="primary" native-type="submit" @click="Submit"
  >提交</van-button
  >
</template>

<script lang="ts">
import { defineComponent, ref, reactive, onMounted } from "vue";

import axios from "axios";
import SparkMD5 from "spark-md5";
import { message } from "ant-design-vue";
import MobileNav from "@/components/common/MobileNav.vue";
interface Data {
  birthDay: String;
}
export default defineComponent({
  components: { MobileNav },
  setup() {
    const onOversize = (file) => {
      console.log(file);
      message.warn("文件大小不能超过10MB");
    };
    const errMsgPhone = ref("");
    const errMsgEmail = ref("");
    const supporting_materials = reactive({
      relation: "",
      content: "",
      Id_Peo: "",
      Id_Ini: "",
      Id_Cer: "",
    });
    const education_info = reactive({
      begin_time: "",
      college: "",
      degree: "",
      end_time: "",
      major: "",
      region: "",
    });
    const work_info = reactive({
      position: "",
      position_title: "",
      professional: "",
      organization: "",
      organization_nature: "",
      or_telephone: "",
      region: "",
      begin_time: "",
      end_time: "",
    });
    const essential_info = reactive({
      nation: "",
      overeseas_permanent: "",
      translation: "",
      married_status: "",
      education: "",
      native_place: "",
      birth_place: "",
      political: "",
      nationality: "",
      level: "",
      home_phone: "",
      address: "",
      postal_code: "",
      fax: "",
      hobby: "",
      social_appointments: "",
    });
    //三个用户id用来传递图片信息
    let uid_Peo = "1";
    let uid_Ini = "2";
    let uid_Cer = "3";
    // 三个file用来获取上传图片的id
    // 三个md5获取上传图片后的md5值
    let md5Key_Peo: string = "";
    let md5Key_Ini: string = "";
    let md5Key_Cer: string = "";
    // 拿到图片数组
    const fileList_Peo = reactive([]);
    const fileList_Ini = reactive([]);
    const fileList_Cer = reactive([]);
    const deleteUrl = "api/file/delete";
    const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/";
    const getIdUrl = "api/file/generateId";
    const uploadUrl = "api/file/upload";
    //必填
    const Must = reactive({
      openName: "",
      loading:false,
      password: "",
      name: "",
      phone: "",
      idcard: "",
    });
    const columns = ["身份证", "护照", "港澳台通行证", "其他"];
    //非必填
    const NotMust = reactive({
      email: "",
      sex: "男",
      birthDay: "2020.10.03", //出生年月日
      study_type: "是",
    });
    const onSubmit = (values: any) => {
      console.log("submit", values);
    };

    const result = ref("");
    const showPicker = ref(false);

    const themeVars = {
      uploaderSize: "160px",
      fieldLabelWidth: "5em",
      fieldLabelMarginRight: "10px",
      fieldWordLimitFontSize: "17px",
      fieldErrorMessageFontSize: "17px",
      fieldClearIconSize: "30px",
    };
    const onConfirm = (value: any) => {
      result.value = value;
      showPicker.value = false;
    };
    onMounted(async () => {
      const formPeo = new FormData();
      formPeo.append("uid", uid_Peo);
      const resultPeo = await axios.post(getIdUrl, formPeo);
      //  console.log(resultPeo)
      if (resultPeo.data.statusCode === 2000) {
        supporting_materials.Id_Peo = resultPeo.data.details;
      }
      const formIni = new FormData();
      formIni.append("uid", uid_Ini);
      const resultIni = await axios.post(getIdUrl, formIni);
      if (resultIni.data.statusCode === 2000) {
        supporting_materials.Id_Ini = resultIni.data.details;
      }
      const formCer = new FormData();
      formCer.append("uid", uid_Cer);
      const resultCer = await axios.post(getIdUrl, formCer);
      if (resultCer.data.statusCode === 2000) {
        supporting_materials.Id_Cer = resultCer.data.details;
      }
    });
    const GetFileMd5 = (file) => {
      return new Promise((resolve) => {
        ``;
        const fileReader = new FileReader();
        const spark = new SparkMD5();
        fileReader.readAsBinaryString(file);
        fileReader.onload = (e) => {
          spark.appendBinary(e.target.result);
          const md5Key = spark.end();
          resolve(md5Key);
        };
      });
    };
    const afterRead_Peo = async (file) => {
      md5Key_Peo = (await GetFileMd5(file.file)) as string;
      const formPeo = new FormData();
      formPeo.append("id", supporting_materials.Id_Peo);
      formPeo.append("md5Key", md5Key_Peo);
      formPeo.append("file", file.file);

      // 显示正在上传中
      fileList_Peo.push({url: '', status: 'uploading'});

      // 发送请求
      const results = await axios.post(uploadUrl, formPeo, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      });
      // 根据上传结果判断显示成功还是失败
      if (results.data.statusCode === 2000) {
        fileList_Peo[fileList_Peo.length-1].status = 'done';
        fileList_Peo[fileList_Peo.length-1].url = `${displayUrl}${results.data.details}`;
        return;
      }
      else {
        fileList_Peo[fileList_Peo.length-1].status = 'failed';
      }
    };
    const beforeDelete_Peo = async (file, detail) => {
      // 发送删除的请求
      const form = new FormData();
      form.append("id", supporting_materials.Id_Peo);
      const results = await axios.post(deleteUrl, form);
      if (results.data.statusCode === 2000) {
        fileList_Peo.splice(detail.index, 1);
        return true;
      }
    };
    const afterRead_Ini = async (file) => {
      md5Key_Ini = (await GetFileMd5(file.file)) as string;
      const formIni = new FormData();
      formIni.append("id", supporting_materials.Id_Ini);
      formIni.append("md5Key", md5Key_Ini);
      formIni.append("file", file.file);

      // 显示正在上传中
      fileList_Peo.push({url: '', status: 'uploading'});

      // 发送请求
      const results = await axios.post(uploadUrl, formIni, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      });
      // 根据上传结果判断显示成功还是失败
      if (results.data.statusCode === 2000) {
        fileList_Ini[fileList_Ini.length-1].status = 'done';
        fileList_Ini[fileList_Ini.length-1].url = `${displayUrl}${results.data.details}`;
        return;
      }
      else {
        fileList_Ini[fileList_Ini.length-1].status = 'failed';
      }
    };
    const beforeDelete_Ini = async (file, detail) => {
      // 发送删除的请求
      const form = new FormData();
      form.append("id", supporting_materials.Id_Ini);
      const results = await axios.post(deleteUrl, form);
      if (results.data.statusCode === 2000) {
        fileList_Ini.splice(detail.index);
        return true;
      }
    };
    const afterRead_Cer = async (file) => {
      md5Key_Cer = (await GetFileMd5(file.file)) as string;
      const formCer = new FormData();
      formCer.append("id", supporting_materials.Id_Cer);
      formCer.append("md5Key", md5Key_Cer);
      formCer.append("file", file.file);

      // 显示正在上传中
      fileList_Cer.push({url: '', status: 'uploading'});

      // 发送请求
      const results = await axios.post(uploadUrl, formCer, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      });
      // 根据上传结果判断显示成功还是失败
      if (results.data.statusCode === 2000) {
        fileList_Cer[fileList_Cer.length-1].status = 'done';
        fileList_Cer[fileList_Cer.length-1].url = `${displayUrl}${results.data.details}`;
        return;
      }
      else {
        fileList_Cer[fileList_Cer.length-1].status = 'failed';
      }
    };
    const beforeDelete_Cer = async (file, detail) => {
      // 发送删除的请求
      const form = new FormData();
      form.append("id", supporting_materials.Id_Cer);
      const results = await axios.post(deleteUrl, form);
      if (results.data.statusCode === 2000) {
        fileList_Cer.splice(detail.index);
        return true;
      }
    };

    return {
      isPhoneCorrect: true,
      isEmailCorrect: true,
      errMsgEmail,
      errMsgPhone,
      themeVars,
      //图片上传
      fileList_Peo,
      fileList_Ini,
      fileList_Cer,
      Id_Peo: "",
      Id_Ini: "",
      Id_Cer: "",
      afterRead_Peo,
      beforeDelete_Peo,
      beforeDelete_Ini,
      beforeDelete_Cer,
      afterRead_Ini,
      afterRead_Cer,
      //必填
      Must,
      //json文件
      supporting_materials,
      education_info,
      essential_info,
      work_info,
      //非必填项
      NotMust,
      //其他
      result, //证件类型
      columns,
      onConfirm,
      showPicker,
      onSubmit,
      onOversize,
    };
  },
  methods: {
    patternPhone() {
      if (this.Must.phone == "") {
        this.errMsgPhone = "请输入手机号";
        this.isPhoneCorrect = false;
        return false;
      }
      let reg = /^1[3-9][0-9]{9}$/;
      if (!reg.test(this.Must.phone)) {
        this.errMsgPhone = "手机号格式错误,请重新输入";
        this.isPhoneCorrect = false;
        return false;
      }
      this.errMsgPhone = "";
      this.isPhoneCorrect = true;
      return true;
    },
    //邮箱校验
    patternEmail() {
      var reg = /^\w{3,}(\.\w+)*@[A-z0-9]+(\.[A-z]{2,5}){1,2}$/;
      if (this.NotMust.email == "") {
        this.errMsgEmail = "";
        this.isEmailCorrect = true;
        return true;
      } else if (!reg.test(this.NotMust.email)) {
        this.errMsgEmail = "邮箱地址错误，请重新输入";
        this.isEmailCorrect = false;
        return false;
      }else{
        this.errMsgEmail = "";
        this.isEmailCorrect = true;
        return true;
      }
    },

    onClickLeft() {
      history.back();
    },
    Submit() {
      if (
          this.Must.openName === "" ||
          this.Must.password === "" ||
          this.Must.phone === "" ||
          this.Must.idcard === "" ||
          this.Must.name === "" ||
          this.isPhoneCorrect === false ||
          this.isEmailCorrect === false
      ) {
        message.warning("请填写必要信息！");
        return;
      }
      let url = `api/user/submit`;
      this.Must.loading = true;
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
          work_info: this.work_info,
          education_info: this.education_info,
          essential_info: this.essential_info,
          supporting_materials: this.supporting_materials,
          //必填
          name: this.Must.name,
          openName: this.Must.openName,
          password: this.Must.password,
          phone: this.Must.phone,
          idcard_type: this.result,
          idcard: this.Must.idcard,

          email: this.NotMust.email,
          sex: this.NotMust.sex,
          birth: this.NotMust.birthDay,
          study_type: this.NotMust.study_type,
          fileId: this.supporting_materials.Id_Cer,
        }),
      })
          .then((v) => {
            console.log(v);
            this.Must.loading = false;
            return v.json();
          })

          .then((res) => {
            if (res.status == 200) {
              message.success(res.msg);
              this.$router.push("/login");
            } else {
              message.warn(res.msg);
            }
          })
          .catch((err) => {
            message.error(err);
          });
    },
  },
});
</script>

<style lang="less" scoped>
// 移动端媒体查询样式
@media screen and (max-width: 600px) {
  .basic {
    margin: 0 0.53rem;
    padding-top: 0.1rem;
    padding-bottom: 2rem;
    .title {
      font-size: 16px;
      color: #8c99a0;
      margin: 0.8rem 0 0.26rem 0.1rem;
    }
  }
  .button {
    width: 8rem;
    position: fixed;
    bottom: 0.2667rem;
    left: 50%;
    transform: translateX(-50%);
  }
}
// pc端媒体查询样式
@media screen and (min-width: 600px) {
  .basic {
    width: 750px;
    margin: 0 auto;
    padding-top: 10px;
    padding-bottom: 100px;
    .title {
      font-size: 17px;
      color: #8c99a0;
      margin: 30px 0 10px 10px;
    }
  }
  .button {
    width: 300px;
    position: fixed;
    bottom: 10px;
    left: 50%;
    transform: translateX(-50%);
  }
}
.wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.block {
  width: 120px;
  height: 120px;
  background-color: rgba(255, 255, 255, 0.432);
  .loading {
    margin-top: 40px;
  }
}
</style>

