<template>
<div>
  <ManageNav/>
  <div style="min-height: 1000px">
    <div class="addAct">
      <van-button class="addAct" plain type="primary" @click="toIssueActivity">发布活动</van-button>
    </div>
    <div class="activities">
      <div v-for="item in activities">
        <div class="itemContainer">
          <manage_activity_item :msg="item" :isManager="isManager"></manage_activity_item>
        </div>
      </div>
    </div>
  </div>

</div>
</template>

<script>
import Manage_activity_item from "@/components/activityManage/manage_activity_item.vue";
import axios from "axios";
import SparkMD5 from "spark-md5";
import {reactive} from "vue";
import ManageNav from "@/components/common/ManageNav.vue";
import {message} from "ant-design-vue";
export default {
  name: 'ActivityManage',
  components: {ManageNav, Manage_activity_item},
  data(){
    return{
      activities:[
        {data: 0},
      ],
      isManager:true,
    }
  },
  methods:{
    toIssueActivity(){
      this.$router.push({path:`/releaseActivity/0`});
    },
    clickLeft(){
      this.$router.push('/managerCenter');
    }
  },
  mounted() {
    let url = 'api/activity/getAll';
    fetch(url, {
      method: "GET",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      headers: new Headers({
        "Content-Type": "application/json",
      }),
      redirect: "follow",
    }).then((res) => {
      return res.json();

    }).then((res) => {
      if(res.status == -999){
        message.warn("账号未登录，请先登录")
        this.$router.push('login')
      }

      if(res.status == -1999){
        message.warn("管理员账号未登录，请先登录")
        this.$router.push('login')
      }

      this.activities=res;
    });
  },
  setup() {

    // 文件列表
    const fileList = reactive([]);

    const uploadUrl = "api/file/upload";
    const downloadUrl = "api/file/download";
    const deleteUrl = "api/file/delete";
    const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/"  // 获取图片缩略图的地址

    let id = '1';  // id是事项的id，比如简历的id
    let md5Key = '';
    const srcList = reactive([]);

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
      form.append("id", id);
      form.append("md5Key", md5Key);
      form.append("file", file.file);
      // 发送请求
      const results = await axios.post(uploadUrl, form, {
        headers: {
          "Content-Type": "multipart/form-data"
        }
      });
      // 上传成功展示缩略图
      if(results.data.statusCode === 2000) {
        fileList.push({url: `${displayUrl}${results.data.details}`})
      }
    };

    // 用户不想上传了
    const beforeDelete = async (file, detail) => {
      // 发送删除的请求
      const form = new FormData();
      form.append("id", id);
      const results = await axios.post(deleteUrl, form);
      if(results.data.statusCode === 2000) {
        fileList.splice(detail.index, 1);
      }
    };

    const GetUploadFile = async () => {
      const form = new FormData();
      form.append("id", id);
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
      GetUploadFile
    }
  }
}
</script>

<style lang="less" scoped>
@baseFont:37.5;
body {
  min-width: 320px;
  overflow: hidden;
  height: 100%;
  div {
    //width: 15rem;
    //width: 10rem;
    width: 100%;
  }
}
@media screen and (max-width: 600px) {
  .addAct {
    width: 10rem;
    margin: 0 auto;
    //height: (70rem/@baseFont);
    //margin-top: (7rem/@baseFont);
    .van-button {
      float: left;
      height: (33rem/@baseFont);
      width: (335rem/@baseFont);
      margin: (10rem/@baseFont) (20rem/@baseFont);
      font-size: (16rem/@baseFont);
      border-radius: (2rem/@baseFont);
    }
  }
  .activities {
    overflow: hidden;
    max-width: 1200px;
    margin: 0 auto;
  }
  .itemContainer{
    overflow: hidden;
    margin: (5rem/@baseFont) auto;
  }
}
@media screen and (min-width: 600px) {
  .addAct {
    .van-button {
      position: relative;
      left: 50%;
      transform: translateX(-50%);
      height: 45px;
      width: 1000px;
      margin: 10px auto;
      font-size: 16px;
      border-radius: 2px;
    }
  }
  .activities {
    overflow: hidden;
    max-width: 1000px;
    margin: 0 auto;
  }
  .itemContainer{
    overflow: hidden;
    margin: 5px auto;
    width: 1000px;
  }
}
</style>
