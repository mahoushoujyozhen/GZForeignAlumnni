<template>
  <div >
    <van-grid :gutter="10" :column-num="3" square :border="false">
      <van-grid-item  to="pendingAuditPage" class="changeFontColor">
        <van-icon name="manager-o" color="#faa616"/>
        <span>会员管理</span>
      </van-grid-item>
      <van-grid-item to="sentMessage" class="changeFontColor">
        <van-icon name="envelop-o"  color="#1890ff"/>
        <span>消息发送</span>
      </van-grid-item>
      <van-grid-item to="activityManage" class="changeFontColor">
        <van-icon name="flag-o"   color="rgb(92 182 90)"/>
        <span>活动管理</span>
      </van-grid-item>
      
    </van-grid>


  </div>
</template>

<style scoped>
.van-icon-manager-o{
  font-size: 1rem;
}
.van-icon-envelop-o{
  font-size: 1rem;
}
.van-icon-flag-o{
  font-size: 1rem;
}
.van-row {
  box-sizing: border-box;
  border-bottom: 0.0267rem solid rgba(235, 235, 235, 1);
  height: 1.8667rem;
  display: flex;
  flex-flow: row;
  align-items: center;
  justify-content: space-between;
  margin: auto 5%;
}
.van-col {
  display: flex;
  box-sizing: border-box;
  align-items: center;
  justify-content: space-between;
}
.go {
  width: 0.32rem;
  height: 0.5333rem;
  display: flex;
  text-align: left;
}
a {
  color: black;
}
.changeFontColor:hover {
  color: #1890ff;
}
.context {
  margin: 0.5333rem;
  margin-top: 0rem;
}
span {
  font-size: 0.3733rem;
}
.van-icon{
  color: #1890ff;
  margin-bottom: 0.2rem;
}

@media screen and (min-width: 600px) {
  .van-icon-manager-o{
  font-size: 37.5px;
}
  .van-icon-envelop-o{
  font-size: 37.5px;
}
  .van-icon-flag-o{
  font-size: 37.5px;
}
  .van-grid {
  display: flex;
  align-items: baseline;
  justify-content: space-around;
  /* margin: 0 0.5rem; */
  margin: auto;
  /* max-width: 900px; */
    max-width: 670px;
}

span{
  font-size: 16px;
}
}
</style>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'MFunction',
  setup() {},
});
</script>
