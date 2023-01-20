<template>
  <MemberNav/>
  <div class="entry">
    <div class="card">
      <div class="input">
        <van-cell-group>
          <van-field
            v-model="details_edu.degree"
            name="学科专业"
            label="学科专业"
            label-align="right"
            readonly
          />

          <van-field
            v-model="details_edu.major"
            name="学位"
            label="学位"
            label-align="right"
            readonly
          />
          <van-field
            v-model="details_edu.college"
            name="学校名称"
            label="学校名称"
            label-align="right"
            readonly
          />
          <van-field
            v-model="details_edu.region"
            name="学校所在国家和地区"
            label="学校所在国家和地区"
            label-align="right"
            readonly
          />
          <van-field label="是否留学" label-align="right">
            <template #input>
              <van-radio-group disabled v-model="studyType">
                <van-radio class="radio" name="是">是</van-radio>
                <van-radio class="radio" name="否">否</van-radio>
              </van-radio-group>
            </template>
          </van-field>
          <van-field
            v-model="details_edu.begin_time"
            name="开始时间"
            label="开始时间"
            label-align="right"
            readonly
          />
          <van-field
            v-model="details_edu.end_time"
            name="结束时间"
            label="结束时间"
            label-align="right"
            readonly
          />

          <van-field name="证明材料" label="证明材料" label-align="right" readonly>
            <template #input>
              <div class="imgWrapper">
                <van-image
                    v-for="n in srcList"
                    :src="'http://mac.gzhu.edu.cn/alumni/resources/' + n"
                    v-viewer="{ movable: false }"
                    class="img"
                />
              </div>
            </template>
          </van-field>
        </van-cell-group>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.imgWrapper{
  display: flex;
  flex-wrap: wrap
}
//小于600px宽度的屏幕，移动端
@media screen and (max-width: 600px) {
  .entry {
    margin: 0 0.53rem;
    padding-top: 0.53rem;
    .van-cell-group {
      margin: 0;
      padding: 0.26rem 0.1rem;
    }
    .img {
      margin: 0.2rem;
      width: 4.8rem;
    }
    .radio {
      margin-bottom: 0.11rem;
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
    .input {
      padding: 20px 50px;
      border-radius: 10px;
    }
    .img {
      margin: 10px;
      width: 420px;
    }
    .radio {
      margin-bottom: 6px;
    }
  }
}
</style>

<script lang="ts">
import { defineComponent, reactive, onMounted,ref } from "vue";
import { message } from "ant-design-vue";
import { useRouter, useRoute } from "vue-router";
import MemberNav from "@/components/common/MemberNav.vue";
export default defineComponent({
  name: "VueComponentSkeleton",
  props: {},
  components: {
    MemberNav
    // TopNav,
    // Head,
  },

  setup() {
    const router = useRouter();
    // const getIdUrl = "api/file/generateId";
    // const uploadUrl = "api/file/upload";
    const downloadUrl = "api/file/download";
    // const deleteUrl = "api/file/delete";
    // 获取图片缩略图的地址
    // const displayUrl = "http://mac.gzhu.edu.cn/alumni/resources/";

    const details_edu: any = reactive({});
    let details_id = useRoute().query.details_id;
    let Id = details_edu.file_id;
    const srcList = reactive([]);
    let url_img = "";

    const studyType = ref('是');

    //获取图片路径
    const GetUploadFile = async function () {
      fetch(downloadUrl, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: `id=${details_edu.file_id}`,
      })
          .then((results) => {
            return results.json();
          })
          .then((results) => {
            if (results.statusCode === 2000) {
              srcList.push(...results.details);
              return;
            }

            console.log("参数不对，fileId为空或者不正确，无图片路径返回");
            message.info("图片请求失败");
          });
    };

    //获取简历信息
    const GetDetails = () => {
      fetch("api/user/detailsedu", {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ details_id }),
      })
          .then((results) => {
            return results.json();
          })
          .then((results) => {
            if (results.status == -999) {
              message.warn("账号未登录，请先登录")
              router.push('login')
              return
            }

            if (results.status == -1999) {
              message.warn("管理员账号未登录，请先登录")
              router.push('login')
              return
            }
            if (results.status != 0) {
              console.log("err:" + results.status + " " + results.msg);
              message.info("信息详情获取失败，请重试，多次失败后请联系工作人员");
              return;
            }

            Object.assign(details_edu, results.data[0]);
            studyType.value=results.data[0].studyType;

            //获取图片路径
            GetUploadFile();
          });
    };

    GetDetails();

    return {
      details_edu,
      details_id,
      url_img,
      srcList,
      Id,
      GetDetails,
      GetUploadFile,
      studyType,
    };
  },
});
</script>
