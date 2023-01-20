<template>
  <div>
    <ManageNav/>
    <div class="headMobile">
      <div class="headCancle"><span @click="toActivityManage">取消</span></div>
      <div class="headTitle"><span>发布掠影</span></div>
      <div class="headFunc"><span @click="deleteRecall">删除</span><span @click="toDetail">发布</span></div>
    </div>
    <div class="textArea">
      <van-field
          v-model="recallContent"
          autosize=" { minHeight: 45rem }"
          type="textarea"
          border="false"
          placeholder="请输入正文"
      />
      <div class="upload-wrapper">
        <van-uploader v-model="fileList" :after-read="afterRead" :before-delete="beforeDelete" multiple :max-count="100"/>
      </div>
    </div>
    <div class="btnGroup">
      <van-button id="cancleBtn" type="primary" plain @click="toActivityManage">取消</van-button>
      <van-button id="deleteBtn" type="danger" plain @click="deleteRecall">删除</van-button>
      <van-button id="issueBtn" type="primary" @click="toDetail">发布</van-button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import {onMounted, ref, reactive} from "vue";
import SparkMD5 from "spark-md5";
import {useRoute, useRouter} from "vue-router";
import ManageNav from "@/components/common/ManageNav.vue";
import {message} from "ant-design-vue";

export default {
  name: "IssueRecall",
  components: {ManageNav},
  methods:{
    toActivityManage(){
      this.$router.push("/activityManage");
    }
  },
  setup() {

    // 文件列表
    const fileList = reactive([]);

    const getIdUrl = "api/file/generateId"
    const uploadUrl = "api/file/upload";
    const downloadUrl = "api/file/download";
    const deleteUrl = "api/file/delete";
    const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/"  // 获取图片缩略图的地址

    const route = new useRoute();
    const router = useRouter();

    let uid = '1';  // "uid"
    let fileId = '';
    let md5Key = '';
    let srcList = reactive(["1.png, 2.png"]);

    let actId = 0;
    let recallContent = ref ('');
    let imgSrcList = reactive([]);

    // 页面刷新时,发送请求
    onMounted( async ()=>{

      //获取掠影信息
      actId = parseInt(route.params.act_id)
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
        recallContent.value=res.act_text
        console.log("res.recall_img=",res.recall_img)
        if(res.recall_img!=null){
          let recallImgList=res.recall_img.split(',')
          for(var i=0;i<recallImgList.length;i++){
            fileList.push({url: recallImgList[i]});
          }
          console.log("fileList:",fileList)
        }
      })

      //图片
      const form = new FormData();
      form.append("uid", uid)
      const results = await axios.post(getIdUrl, form);
      if(results.data.statusCode === 2000) {
        fileId = results.data.details;
      }
    });

    //发布掠影
    const toDetail=function (){
      // const router = useRouter()
      for(var i=0;i<fileList.length;i++){
        imgSrcList[i]=fileList[i].url;
      }
      console.log("发送：",imgSrcList)
      let imgArr=imgSrcList.toString()
      let recallData=JSON.stringify({
        id:actId,
        act_text:recallContent.value,
        recall_img:imgArr
      });
      let url='api/admin/createRecall';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: recallData,
      }).then((v)=>{
        return v.json();
      }).then((v)=>{
        if(v.status == -999){
          message.warn("账号未登录，请先登录")
          this.$router.push('login')
        }

        if(v.status == -1999){
          message.warn("管理员账号未登录，请先登录")
          this.$router.push('login')
        }

        router.replace(`/activityDetail/${actId}/true`)
      })
    };

    //删除掠影
    const deleteRecall=function(){
      var recallData=JSON.stringify({
        id:actId,
        act_text:null,
        recall_img:null
      });
      let url='api/admin/deleteRecall';
      fetch(url, {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        headers: new Headers({
          'Content-Type': 'application/json',
        }),
        redirect: 'follow',
        body: recallData,
      }).then((v)=>{
        return v.json();
      }).then((v)=>{
        if(v.status == -999){
          message.warn("账号未登录，请先登录")
          this.$router.push('login')
        }

        if(v.status == -1999){
          message.warn("管理员账号未登录，请先登录")
          this.$router.push('login')
        }
        router.replace(`/activityDetail/${actId}/true`)
      })
    };

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
      console.log(file, detail);
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
      actId,
      recallContent,
      imgSrcList,
      toDetail,
      deleteRecall,
    }
  },
}
</script>

<style lang="less" scoped>
@baseFont:37.5;
#app{
  min-height: 1000px;
}
@media screen and (max-width: 600px) {
  .headMobile {
    width: 100%;
    height: 1.06rem;
    position:relative;
    background-color: #ffffff;
    border-bottom: 1px solid #efefef;
    .headCancle {
      position: absolute;
      font-size: 0.43rem;
      line-height: 1.06rem;
      left: 0.53rem;
      color:#5f5f5f;
    }
    .headTitle {
      position: absolute;
      left:50%;
      transform: translateX(-50%);
      font-size: 0.48rem;
      line-height: 1.06rem;
    }
    .headFunc {
      position: absolute;
      font-size: 0.43rem;
      color: #1890ff;
      line-height: 1.06rem;
      right: 0.53rem;
      :nth-child(1){
        color:#5f5f5f;
      }
      :nth-child(2){
        margin-left: .15rem;
      }
    }
  }
  .operate{
    span{
      font-size: (16rem/@baseFont);
      color: white;

    }
    span:nth-child(2){
      margin-left: (10rem/@baseFont);
    }
  }
  .textArea{
    background-color: white;
    padding: (20rem/@baseFont);
    .van-field{
      //width: 10rem;
      //padding: 0;
      min-height: (200rem/@baseFont);
      width: 100%;
      font-size: (16rem/@baseFont);
    }
  }
  .upload-wrapper{
    margin: (20rem/@baseFont) (20rem/@baseFont) 0 (20rem/@baseFont);
  }
  .btnGroup{
    display: none;
  }
}
@media screen and (min-width: 600px) {
  .headMobile{
    display: none;
  }
  .operate{
    span{
      font-size: 16px;
      color: white;

    }
    span:nth-child(2){
      margin-left: 10px;
    }
  }
  .textArea{
    width: 1000px;
    margin: 45px auto;
    background-color: white;
    padding: 35px;
    .van-field{
      min-height: 300px;
      font-size: 16px;
    }
  }
  .upload-wrapper{
    background-color: white;
    margin: 20px 20px 0 20px;
  }
  .btnGroup{
    width: 1000px;
    margin: 0 auto;
    .van-button{
      width: 100px;
    }
    #cancleBtn{
    }
    #issueBtn{
      float:right;
    }
    #deleteBtn{
      float:right;
      margin-left: 15px;
    }
  }
}

</style>
