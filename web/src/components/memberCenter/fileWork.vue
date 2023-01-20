<template>
  <van-field name="uploader" label="证明材料" label-align="right">
    <template #input>
      <div>
        <div>
          <van-uploader v-model="fileList" :after-read="afterRead" :before-delete="beforeDelete" />
          <!-- <van-button type="primary" @click="GetUploadFile">获取图片的路径</van-button> -->
          <div v-for="src in srcList">
            <div>{{ displayUrl + src }}</div>
          </div>
        </div>
      </div>
    </template>
  </van-field>
</template>

<script>
import SparkMD5 from "spark-md5";
import { reactive, onMounted } from "vue";
import { message } from "ant-design-vue";
import axios from "axios";

export default {
  name: "Vantload",
  emits: ["getFileId"],
  setup(props, context) {
    // 文件列表
    const fileList = reactive([]);
    const getIdUrl = "api/file/generateId";
    const uploadUrl = "api/file/upload";
    const downloadUrl = "api/file/download";
    const deleteUrl = "api/file/delete";
    const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/"; // 获取图片缩略图的地址

    let fileId = "";
    let md5Key = "";
    const srcList = reactive([]);
    // 页面刷新时,发送请求
    onMounted(async () => {
      const results = await axios.post(getIdUrl);

      if (results.data.statusCode === 2000) {
        fileId = results.data.details.toString();
        context.emit("getFileId", fileId);
        return;
      }

      message.info("网络繁忙，请刷新重试");
      console.log("generateId failed");
      // 有可能有网络问题导致拿不到Id,这里可以处理一下,比如不显示页面什么的
    });

    // 获取文件的MD5值
    const GetFileMd5 = (file) => {
      return new Promise((resolve) => {
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
      });
    };

    // 用户上传了文件
    const afterRead = async (file) => {
      md5Key = await GetFileMd5(file.file);
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

      if (results.data.statusCode === 2000) {
        fileList.splice(detail.index, 1);
        return;
      }

      message.info("图片删除失败，请重试");
      console.log("deleteFile failed");
    };

    const GetUploadFile = async () => {
      const form = new FormData();
      form.append("id", fileId);
      const results = await axios.post(downloadUrl, form);

      if (results.data.astatusCode === 2000) {
        srcList.push(...results.data.details);
        return;
      }

      console.log("getFileUrl failed");

    };

    return {
      srcList,
      fileList,
      displayUrl,
      afterRead,
      beforeDelete,
      GetUploadFile,
    };
  },
};
</script>

<style scoped>
.container {
  width: 100%;
  min-height: 10vh;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 4vh;
}
/* 上传表单 */
.container .upload-wrapper {
  width: 30vw;
  padding: 20px;
  background: rgba(255, 255, 255, 0.1);
  box-shadow: 0 25px 45px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.5);
  border-right: 1px solid rgba(255, 255, 255, 0.2);
  border-bottom: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 10px;
}
.container .upload-wrapper {
  text-align: center;
}
</style>

