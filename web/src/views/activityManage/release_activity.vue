<template>
  <div class="div1">
    <ManageNav/>
    <van-form class="activityForm">
      <van-cell-group inset>
        <van-field v-model="act_name" label="活动名称" placeholder="请输入活动名称"  :rules="[{ required: true, message: '请填写活动名' }]"/>
        <van-field v-model="act_address" label="活动地址" placeholder="请输入活动地址"  :rules="[{ required: true, message: '请填写活动地址' }]"/>
        <van-field v-model="startTime" label="开始时间" placeholder="请选择活动开始时间" @click="showPicker_Start = true" :rules="[{ required: true, message: '请填写活动开始时间' }]" />
        <van-popup v-model:show="showPicker_Start" position="bottom">
          <van-datetime-picker
              v-model="currentDate"
              type="datetime"
              v-show="showPicker_Start"
              :min-date="minDate"
              :max-date="maxDate"
              @cancel="showPicker_Start = false"
              @confirm="confirmPicker_Start"
          />
        </van-popup>
        <van-field v-model="endTime" label="结束时间" placeholder="请选择活动结束时间" @click="showPicker_End = true" :rules="[{ required: true, message: '请填写活动结束时间' }]"/>
        <van-popup v-model:show="showPicker_End" position="bottom">
          <van-datetime-picker
              v-model="currentDate"
              type="datetime"
              v-show="showPicker_End"
              :min-date="minDate"
              :max-date="maxDate"
              @cancel="showPicker_End = false"
              @confirm="confirmPicker_End"
          />
        </van-popup>
        <van-field v-model="act_profile" label="活动简介" placeholder="请输入活动简介" :rules="[{ required: true, message: '请填写活动简介' }]"/>
        <div class="upload-wrapper">
          <van-uploader  v-model="fileList" :after-read="afterRead" :before-delete="beforeDelete" :max-count="1"/>
        </div>
      </van-cell-group>
      <van-button type="primary" native-type="submit" @click="test" >
        提交
      </van-button>
    </van-form>
  </div>
</template>

<script lang="js">
import axios from "axios";
import SparkMD5 from "spark-md5";
import {onMounted, reactive} from "vue";
import { message } from 'ant-design-vue';
import { ref } from 'vue';
import { DatetimePicker } from "vant";
import {useRoute, useRouter} from "vue-router";
import ManageNav from "@/components/common/ManageNav.vue";

export default {
  setup() {

    // 文件列表
    const fileList = reactive([]);

    const getIdUrl = "api/file/generateId"
    const uploadUrl = "api/file/upload";
    const downloadUrl = "api/file/download";
    const deleteUrl = "api/file/delete";
    const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/"  // 获取图片缩略图的地址

    let uid = '1';  // "uid"
    let fileId = '';
    let md5Key = '';
    const srcList = reactive([]);

    const route = useRoute();
    const router = useRouter();

    let actId = 0;
    let act_name = ref('');
    let act_address = ref('');
    let startTime = ref('');
    let endTime = ref('');
    let act_profile = ref('');


    // 页面刷新时,发送请求
    onMounted( async ()=>{
      //获取活动信息
      actId = parseInt(route.params.act_id)
      if(actId!==0){
        let url =`api/activity/getById?act_id=${actId}`
        fetch(url, {
          method: "GET",
          mode: "cors",
          cache: "no-cache",
          credentials: "same-origin",
          headers: new Headers({
            "Content-Type": "application/json",
          }),
          redirect: "follow",
        }).then((res)=>{
          console.log(res);
          return res.json();
        }).then((res)=>{
          if(res.status == -999){
            message.warn("账号未登录，请先登录")
            this.$router.push('login')
          }

          if(res.status == -1999){
            message.warn("管理员账号未登录，请先登录")
            this.$router.push('login')
          }

          act_name.value =res.act_name
          act_address.value =res.act_address
          startTime.value =res.start_time
          endTime.value = res.end_time
          act_profile.value =res.act_profile

          console.log("res.recall_img=",res.recall_img)
          console.log("res.img_url=",res.img_url)
          if(res.img_url!=null){
            let imgUrl=res.img_url;
            fileList.push({url: imgUrl});
          }
          console.log("fileList:",fileList)

        })
      }

      //图片功能
      const form = new FormData();
      form.append("uid", uid)
      const results = await axios.post(getIdUrl, form);
      if(results.data.statusCode === 2000) {
        fileId = results.data.details;
      }
    });

    //点击发布活动，先校验表单
    const test = function() {
      actId = parseInt(route.params.act_id);
      if(act_name.value.length!==0&&act_address.value.length!==0&&startTime.value!==0&&endTime.value!==0&&act_profile.value.length!==0&&fileList.length!==0)
      {
        if(actId==0){ //actId==0，则该活动尚未发布
          act_release();
        }else{ //对已发布的活动进行修改
          act_change(actId);
        }
        router.go(-1);
        router.replace('/activityManage');
      }
    }

    const act_release = function() {
      let url = 'api/admin/act_release';
      console.log("发送：",fileList[0].url)
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin', 
        //headers是请求头：
        headers: new Headers({
          'Content-Type': 'application/json',   //后台发送数据的格式必须为json字符串
        }),
        redirect: 'follow',  //redirect重定向
        //body是请求体：要向服务器发送的数据
        body: JSON.stringify({  //JSON.stringify：将值转换为相应的 JSON 格式字符串
          name:act_name.value,
          address:act_address.value,
          start_Time:startTime.value,
          end_Time:endTime.value,
          profile:act_profile.value,
          img_url:fileList[0].url
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
            if (!v) {
              message.error('null response');
              return;
            }

            if (v.status !== 0) {
              message.error(v.msg);
              return;
            }
            message.success('发布成功');
          })
          .catch((err) => {
            message.error(err);
          });
    }

    const act_change = function(actId) {
      let url = 'api/admin/act_change';
     
      console.log("发送：",fileList[0].url)
      let u=fileList[0].url
      fetch(url, {
        method: 'POST',
        mode: 'cors', 
        cache: 'no-cache',
        credentials: 'same-origin', 
        //headers是请求头：
        headers: new Headers({
          'Content-Type': 'application/json',   //后台发送数据的格式必须为json字符串
        }),
        redirect: 'follow',
        body: JSON.stringify({
          name:act_name.value,
          address:act_address.value,
          start_Time:startTime.value,
          end_Time:endTime.value,
          profile:act_profile.value,
          img_url:fileList[0].url,
          cur_actId:actId,
        }),
      })
          .then((v) => {
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
              message.error('null response');
              return;
            }

            if (v.status !== 0) {
              message.error(v.msg);
              return;
            }
            message.success('修改成功');
          })
          .catch((err) => {
            console.log(err);
            message.error(err);
          });
    }

    // 获取文件的MD5值
    const GetFileMd5 = (file) => {
      return new Promise((resolve => {
        const fileReader = new FileReader();
        const spark = new SparkMD5();
        // 读取文件
        fileReader.readAsBinaryString(file);
        // 读取文件完毕的回调函数
        fileReader.onload = (e) => {
          spark.appendBinary(e.target.result);
          const md5Key = spark.end();
          resolve(md5Key);
        };
      }));
    };

    // 用户上传了文件
    const afterRead = async (file) => {
      md5Key =  await GetFileMd5(file.file);
      const form = new FormData();
      form.append("id", fileId);
      form.append("md5Key", md5Key);
      form.append("file", file.file);
      // 显示正在上传中
      fileList.push({url: '', status: 'uploading'});

      // 发送请求
      const results = await axios.post(uploadUrl, form, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      });
     // 根据上传结果判断显示成功还是失败
      if (results.data.statusCode === 2000) {
        fileList[fileList.length-1].status = 'done';
        fileList[fileList.length-1].url = `${displayUrl}${results.data.details}`;
        return;
      }
      else {
        fileList[fileList.length-1].status = 'failed';
      }
    };

    // 用户不想上传了
    const beforeDelete = async (file, detail) => {
      // 发送删除的请求
      const form = new FormData();
      form.append("id", fileId);
      const results = await axios.post(deleteUrl, form);

      if(results.data.statusCode === 2000) {
        fileList.splice(detail.index, 1);
      }
    };

    const GetUploadFile = async () => {
      const form = new FormData();
      form.append("id", fileId);
      const results = await axios.post(downloadUrl, form);
      if(results.data.statusCode === 2000) {
        srcList.push(...results.data.details);
      }
    };
    return {
      srcList,
      fileList,
      displayUrl,
      afterRead,
      beforeDelete,
      GetUploadFile,
      act_name,
      act_address,
      startTime,
      endTime,
      act_profile,
      test
    }
  },
  data() {
    const onClickLeft = () => history.back();
    const currentDate = ref(new Date());
    const show = ref(false);
    const showPopup = () => {
      show.value = true;
    };

    const result = ref('');
    const showPicker = ref(false);
    const onConfirm = (value) => {
      endTime.value = value;
      showPicker.value = false;
    };

    return {
      onClickLeft,
      minDate: currentDate.value,
      maxDate: new Date(2025, 10, 1),
      currentDate,
      show,
      showPopup,
      result,
      onConfirm,
      showPicker: false,
      username:'',
      showPicker_Start: false,
      showPicker_End: false,
      isShow : false,//出生日期组件显示或隐藏
    };
  },
  components: {
    ManageNav,
    vanDatetimePicker : DatetimePicker
  },
  methods: {
    //确定选择时间
    confirmPicker_Start (val) {
      // console.log(val)
      let year = val.getFullYear()//年
      let month = val.getMonth() + 1//月
      let day = val.getDate()//日
      let hour = val.getHours()//时
      let minute = val.getMinutes()//分
      if (month >= 1 && month <= 9) { month = `0${month}` }
      if (day >= 1 && day <= 9) { day = `0${day}` }
      if (hour >= 0 && hour <= 9) { hour = `0${hour}` }
      if (minute >= 0 && minute <= 9) { minute = `0${minute}` }
      this.startTime = `${year}-${month}-${day} ${hour}:${minute}`
      this.showPicker_Start = false
      console.log(typeof(this.startTime));
    },
    //确定选择时间
    confirmPicker_End (val) {
      // console.log(val)
      let year = val.getFullYear()//年
      let month = val.getMonth() + 1//月
      let day = val.getDate()//日
      let hour = val.getHours()//时
      let minute = val.getMinutes()//分
      if (month >= 1 && month <= 9) { month = `0${month}` }
      if (day >= 1 && day <= 9) { day = `0${day}` }
      if (hour >= 0 && hour <= 9) { hour = `0${hour}` }
      if (minute >= 0 && minute <= 9) { minute = `0${minute}` }
      this.endTime = `${year}-${month}-${day} ${hour}:${minute}`
      this.showPicker_End = false
    },

  }
};

</script>

<style lang="less" scoped>
@baseFont:37.5;

//PC
@media screen and (min-width: 600px){
.activityForm{
  width: 750px;
  padding-top: 50px;
  margin: 0 auto;
  .van-cell-group{
    margin:0;
    padding: 20px;
  }
  .van-field{
    // height:37.5px;
    font-size: (562.5px/@baseFont);
    line-height: (900px/@baseFont);
    //font-size: .3733rem;
    //line-height: .64rem;
  }
}
#div1{
  background-color: #f7f8fa;
  // display:none
}
  .van-button {
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    width: 200px;
    margin: 35px auto;
    font-size: 16px;
  }
.upload-wrapper{
  margin: (750px/@baseFont);
  .van-uploader{
  }
}
}
//移动端
@media screen and (max-width: 600px){
.activityForm{
  width: 10rem;
  margin: 0 auto;
  padding-top: (35rem/@baseFont);
  .van-cell-group{
    padding: (20rem/@baseFont);
  }
  .van-field{
    font-size: (15rem/@baseFont);
    line-height: (24rem/@baseFont);
  }
}
#div1{
  background-color: #f7f8fa;
}
  .van-button{
    position: relative;
    left: 50%;
    transform: translateX(-50%);
    width: 4rem;
    margin: 0.53rem auto;
  }
.upload-wrapper{
  margin: (20rem/@baseFont);
  .van-uploader{
  }
}

}



</style>

