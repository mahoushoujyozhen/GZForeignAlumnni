<template>
  <MemberNav/>
  <div class="container">
    <div class="details-wrapper">
      <div class="details-main">
        <div class="text-item">
          <van-field
              class="text-input"
              v-model="title"
              label="主题"
              readonly
          />
        </div>
        <div class="text-item">
          <van-field
              class="text-input"
              v-model="time"
              label="时间"
              readonly
          />
        </div>
        <div class="text-item">
          <van-field
              class="text-input"
              type="textarea"
              v-model="content"
              label="正文"
              readonly
              autosize
              rows="10"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="js">
import { onMounted } from "vue";
import { useRouter, useRoute } from "vue-router";
import MemberNav from "@/components/common/MemberNav.vue";

export default {
  name: "CardDetails",
  components: {
    MemberNav
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    time: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
  },
  methods:{
    back(){
      this.$router.push('/cardList');
    }
  },
  setup() {
    return {};
  },
};
</script>

<style scoped>
@media screen and (min-width: 600px) {
  .container {
    background-color: #f7f7f7;
    display: flex;
    justify-content: center;
    align-items: center;
    padding-top: 50px;
  }
  .details-wrapper {
    width:750px;
  }
  .details-wrapper .details-main .text-item {
    margin-bottom: 10px;
    border-radius: 10px;
  }
  .details-main .text-item:last-child .text-input {
    height: 50vh;
  }
}
@media screen and (max-width: 600px) {
  .details-wrapper {
    padding-top: 0.53rem;
    min-height: 100vh;
    background-color: #f7f8fa;
  }
  .details-wrapper .details-main {
    padding: 0 0.53rem 0 0.53rem;
  }
  .details-wrapper .details-main .text-item {
    margin: 0 0 0.26rem 0;
    position: relative;
  }
  .details-wrapper .details-main .text-item .text-input {
    border-radius: 5px;
    font-size: 0.4rem;
  }
  @--van-field-label-width: 4em;
}
</style>
