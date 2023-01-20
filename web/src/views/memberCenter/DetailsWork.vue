<template>
  <MemberNav/>
  <div class="entry">
    <van-cell-group inset>
      <van-field
        v-model="details_work.position"
        name="职位"
        label="职位"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.position_title"
        name="职称"
        label="职称"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.professional"
        name="从事专业"
        label="从事专业"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.organization"
        name="机构名称"
        label="机构名称"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.organization_nature"
        name="机构性质"
        label="机构性质"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.org_phone"
        name="机构电话"
        label="机构电话"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.region"
        name="机构所在国家和地区"
        label="机构所在国家和地区"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.begin_Time"
        name="开始时间"
        label="开始时间"
        label-align="right"
        readonly
      />
      <van-field
        v-model="details_work.end_Time"
        name="结束时间"
        label="结束时间"
        label-align="right"
        readonly
      />
    </van-cell-group>
  </div>
</template>

<style lang="less" scoped>
//小于600px宽度的屏幕，移动端
@media screen and (max-width: 600px) {
  .entry {
    padding-top: 0.53rem;
    margin: 0 0.53rem;
    .van-cell-group {
      margin: 0;
      padding: 0.26rem 0.1rem;
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
  }
}
</style>

<script lang="ts">
import { defineComponent, reactive, onMounted, isReactive } from "vue";
import MemberNav from "@/components/common/MemberNav.vue";
import { message } from "ant-design-vue";
import { useRouter, useRoute } from "vue-router";

export default defineComponent({
  name: "VueComponentSkeleton",
  props: {},
  components: {
    MemberNav
  },

  setup() {
    const router = useRouter();
    let details_work: any = reactive({});
    let details_id = useRoute().query.details_id;

    onMounted(() => {
      let url = "api/user/detailswork";
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
          details_id: details_id,
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          if(v.status == -999){
            message.warn("账号未登录，请先登录")
            router.push('login')
            return
          }

          if(v.status == -1999){
            message.warn("管理员账号未登录，请先登录")
            router.push('login')
            return
          }
          if (v.status != 0) {
            console.log("err:" + v.status + " " + v.msg);
            message.info("信息详情获取失败，请重试，多次失败后请联系工作人员");
            return;
          }

          Object.assign(details_work, v.data[0]);
        });
    });

    return {
      details_work,
    };
  },
});
</script>
