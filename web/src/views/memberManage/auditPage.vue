<template>
  <ManageNav />
  <auditForm :userId="userId"></auditForm>
</template>

<script lang="ts">

import { defineComponent, reactive, toRefs } from 'vue'
import auditForm from '../../components/memberManage/pendingAudit/auditForm.vue'
import ManageNav from "@/components/common/ManageNav.vue";
export default defineComponent({
  name: 'auditPage',
  beforeMount() {
    this.getUserId()

  },
  setup() {
    const userId = reactive({
      userId: 0
    })
    return {
      ...toRefs(userId)
    }
  },
  components: {
    ManageNav,
    auditForm,
  },
  methods: {
    getUserId() {
      this.userId = this.$route.query.userId
    }
  },
});
</script>

<style scoped></style>
