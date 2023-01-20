import {
  createRouter,
  createWebHashHistory,
  createWebHistory,
} from 'vue-router';
import ActivityManage from '../views/activityManage/ActivityManage.vue';
import IssueRecall from '../views/activityManage/IssueRecall.vue';
import ActivityDetail from '../views/activityManage/ActivityDetail.vue';
import ActivityHome from '../views/activityManage/ActivityHome.vue';
import releaseActivity from '../views/activityManage/release_activity.vue';

import memberCenter from '../views/memberCenter/memberCenter.vue';
import managerCenter from '../views/managerCenter/managerCenter.vue';
//个人简历功能
import WorkResume from '../views/memberCenter/WorkResume.vue';
import PersonalResueme from '../views/memberCenter/PersonalResueme.vue';
import DetailsEd from '../views/memberCenter/DetailsEd.vue';
import DetailsWork from '../views/memberCenter/DetailsWork.vue';
import EdResume from '../views/memberCenter/EdResume.vue';
//消息发送功能
import sentMessage from '../views/managerCenter/sentMessage.vue';
//系统消息功能
import CardList from '../views/memberCenter/CardList.vue';
import CardDetails from '../views/memberCenter/CardDetails.vue';

import informationChange from "@/views/memberCenter/informationChange.vue";

const router = createRouter({
  // history: createWebHistory(import.meta.env.BASE_URL),
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      redirect: '/login',
    },
    //*********************第一组*********************
    //#region
    {
      //登录
      path: '/login',
      name: 'login',
      component: () => import('../views/memberManage/Login.vue'),
    },
    {
      //用户注册页
      path: '/register',
      name: 'register',
      component: () => import('../views/memberManage/Register.vue'),
    },
    {
      //审核信息页
      path: '/auditPage',
      name: 'auditPage',
      component: () => import('../views/memberManage/auditPage.vue'),
    },
    {
      //管理员审核中心
      path: '/pendingAuditPage',
      name: 'pendingAuditPage',
      meta: { requestAdminAuth: true },
      component: () =>
        import('../views/memberManage/pendingAuditPage.vue'),
    },
    {
      //审核未通过
      path: '/auditFailed/:id',
      name: 'auditFailed',
      component: () => import('../views/memberManage/auditFailed.vue'),
      meta: { requestAdminAuth: true },
    },
    {
      //审核通过
      path: '/approved/:id',
      name: 'approved',
      component: () => import('../views/memberManage/approved.vue'),
      meta: { requestAdminAuth: true },
    },
    //#endregion
    //*********************第三组*********************
    //#region
    {
      //会员个人中心
      path: '/memberCenter',
      name: 'memberCenter',
      component: memberCenter,
    },
    {
      //管理总后台
      path: '/managerCenter',
      name: 'managerCenter',
      component: managerCenter,
      meta: { requestAdminAuth: true },
    },
    {
      //信息修改
      path: '/informationChange',
      name: 'informationChange',
      component: informationChange,
    },
    {
      //个人简历
      path: '/PersonalResueme',
      name: 'PersonalResueme',
      component: PersonalResueme,
    },
    {
      //新增工作经历
      path: '/WorkResume',
      name: 'WorkResume',
      component: WorkResume,
    },
    {
      //新增教育经历
      path: '/EdResume',
      name: 'EdResume',
      component: EdResume,
    },
    {
      //教育经历详情
      path: '/DetailsEd',
      name: 'DetailsEd',
      component: DetailsEd,
    },
    {
      //工作经历详情
      path: '/DetailsWork',
      name: 'DetailsWork',
      component: DetailsWork,
    },
    {
      //消息发送功能
      path: '/sentMessage',
      name: 'sentMessage',
      component: sentMessage,
    },
    {
      //系统消息
      path: '/cardList',
      name: 'CardList',
      component: CardList,
    },
    {
      //信息详情
      path: '/cardDetails/:title/:time/:content',
      name: 'CardDetails',
      component: CardDetails,
      props: true,
    },
    //#endregion
    //*********************第二组*********************
    //#region
    {
      //活动管理后台
      path: '/activityManage',
      name: 'activityManage',
      component: ActivityManage,
      meta: { requestAdminAuth: true },
    },
    {
      //发表掠影
      path: '/issueRecall/:act_id',
      name: 'issueRecall',
      component: IssueRecall,
      meta: { requestAdminAuth: true },
    },
    {
      //活动详情
      path: '/activityDetail/:act_id/:isManager',
      name: 'activityDetail',
      component: ActivityDetail,
    },
    {
      //活动首页（供用户看）
      path: '/activityHome',
      name: 'activityHome',
      component: ActivityHome,
    },
    {
      //发表活动
      path: '/releaseActivity/:act_id',
      name: 'releaseActivity',
      component: releaseActivity,
    },
    {
      //个人活动列表
      path: '/myActivity',
      name: 'myActivity',
      component: () => import('../views/activityManage/MyActivity.vue'),
    },
    {
      //活动二维码
      path: '/qrcode/:act_id',
      name: 'qrcode',
      component: () => import('../views/activityManage/Qrcode.vue'),
    },
    {
      //活动签到列表
      path: '/signedlist/:act_id',
      name: 'signedlist',
      component: () => import('../views/activityManage/SignedList.vue'),
    },
    //#endregion
  ],
});
export default router;
