<template>
  <ManageNav />
  <div class="tabs">
    <van-config-provider :theme-vars="themeVars">
      <van-tabs
          v-model:active="active"
          sticky
          title-active-color="#1890FF"
          animated
          swipeable
      >
        <van-tab title="待审核" name="auditing">
          <userList />
        </van-tab>
        <van-tab title="已审核" name="audited">
          <auditedList />
        </van-tab>
      </van-tabs>
    </van-config-provider>
  </div>
</template>

<script>
import { defineComponent, reactive, ref } from 'vue';
import userList from '../../components/memberManage/pendingAudit/userList.vue';
import auditedList from '../../components/memberManage/pendingAudit/auditedList.vue';
import { useRoute } from 'vue-router';
import ManageNav from '@/components/common/ManageNav.vue';
import { useRouter } from 'vue-router';
import { message } from 'ant-design-vue';

export default defineComponent({
  components: {
    ManageNav,
    userList,
    auditedList,
  },
  setup() {
    const Route = useRoute();
    const active = ref('auditing');
    const themeVars = {
      tabsBottomBarColor: '#1890FF',
    };
    return reactive({
      Route,
      active,
      themeVars,
    });
  },
  created() {
    // if (this.Route.query.active == 1) {
    //   this.active = 'audited';
    // }
    // this.$router.push({ query: {} });
  },
});
</script>

<style scoped>
@media screen and (min-width: 600px) {
  .tabs {
    width: 750px;
    margin: auto;
  }
}
</style>
