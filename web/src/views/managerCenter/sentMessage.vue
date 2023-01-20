<template>
  <ManageNav />
  <div style="background-color: #f8f8f8">
    <div class="headMobile">
      <div class="headCancle" @click="cancle"><span>取消</span></div>
      <div class="headMessage"><span>系统消息</span></div>
      <div class="headSent" @click="sendMessage"><span>发送</span></div>
    </div>
    <div class="app-entry">
      <div class="messagePopupMobile">
        <van-popup v-model:show="show" position="bottom" :style="{}">
          <div style="margin-top: 3px">
            <van-search
              v-model="searchUser"
              placeholder="请输入收件人..."
              @search="OnSearch"
              clearable
            />
          </div>
          <div style="background-color: #f8f8f8; margin-top: 0.5vh">
            <div style="display: flex; flex-flow: row">
              <div class="selectPage">
                <input type="checkbox" @click="checkPage" v-model="checkedPageStatus" />
                <span> 本页全选 </span>
              </div>
              <div class="selectAll">
                <input type="checkbox" @click="checkall" v-model="checkedAllStatus" />
                <span> 全选 </span>
              </div>
              <div class="cancelFont" @click="show = false">
                <span>取消</span>
              </div>
              <div class="OkFont" @click="sendReceiver">
                <span>确认</span>
              </div>
            </div>
            <div style="margin-top: 2.2vh">
              <van-checkbox-group v-model="checkModel" style="">
                <van-cell-group inset v-for="(val, key) in showUserList">
                  <van-cell :title="val.name">
                    <template #right-icon>
                      <van-checkbox
                        :name="val.id"
                        clickable
                        @click.stop
                        @click="check(val.id)"
                      />
                    </template>
                  </van-cell>
                </van-cell-group>
                <van-pagination
                  v-model="currentPage"
                  :page-count="userPageCount"
                  @change="pageChange"
                  mode="simple"
                >
                </van-pagination>
              </van-checkbox-group>
            </div>
          </div>
        </van-popup>
      </div>

      <div class="sendMessageDiv">
        <van-cell-group class="receiverDiv" inset>
          <van-field
            v-model="inputReceiver"
            label="收件人"
            placeholder="请选择收件人"
            type="textarea"
            rows="1"
            label-width="3em"
            right-icon="user-o"
            @click="showPopup"
            readonly
          />
          <van-field
            v-model="inputTheme"
            placeholder="请输入主题......"
            label="主题"
            label-width="3em"
          />
        </van-cell-group>
        <div class="ContentWrapper">
          <van-field
            type="textarea"
            v-model="inputContent"
            placeholder="请输入正文......"
            :autosize="{ minHeight: 450 }"
          />
          <a class="PCSendMessage" @click="sendMessage">发送</a>
        </div>
      </div>

      <div class="messagePopupPC">
        <div class="searchDivPC">
          <form action="/">
            <van-search
              v-model="searchUser"
              placeholder="请输入收件人..."
              @search="OnSearch"
              clearable
            />
          </form>
        </div>
        <div style="display: flex; flex-flow: row; justify-content: space-between">
          <div style="display: flex">
            <div class="selectPage">
              <input
                type="checkbox"
                @click="checkPage"
                style="outline-style: none"
                v-model="checkedPageStatus"
              />
              <span> 本页全选 </span>
            </div>
            <div class="selectAll">
              <input type="checkbox" @click="checkall" v-model="checkedAllStatus" />
              <span> 全选 </span>
            </div>
          </div>
          <div class="OKFont" @click="sendReceiver"><span>确定</span></div>
        </div>
        <div style="margin-top: 2vh">
          <van-checkbox-group v-model="checkModel">
            <van-cell-group inset v-for="(val, key) in showUserList">
              <van-cell :title="val.name" class="cell">
                <template #right-icon>
                  <van-checkbox
                    :name="val.id"
                    clickable
                    @click.stop
                    @click="check(val.id)"
                  />
                </template>
              </van-cell>
            </van-cell-group>
            <van-pagination
              v-model="currentPage"
              :page-count="userPageCount"
              @change="pageChange"
              mode="simple"
            >
            </van-pagination>
          </van-checkbox-group>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
@baseFont: 37.5;
@media screen and (min-width: 600px) {
  .headMobile {
    display: none;
  }
  .app-entry {
    overflow-y: auto;
    overflow-x: auto;
    // overflow: auto;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    .messagePopupMobile {
      display: none;
    }
    .messagePopupPC {
      margin: 50px 0 0 30px;
      background-color: #ffffff;
      padding: 20px;
      border-radius: 7px;
      height: 700px;
      .searchDivPC {
        min-width: 500px;
        .van-search {
          padding: 0;
        }
      }
      .selectPage {
        margin-left: 10px;
        margin-top: 2vh;
        font-size: 16px;
      }
      .selectAll {
        margin-left: 10px;
        margin-top: 2vh;
        font-size: 16px;
      }

      .OKFont {
        align-items: flex-end;
        margin-top: 2vh;
        margin-right: 11px;
        color: #1890ff;
        font-size: 16px;
      }
      .OKFont :hover {
        color: red;
      }
      .cell {
        padding: 12px 0;
      }
    }
    .sendMessageDiv {
      margin: 50px 0;
      width: 750px;
      .receiverDiv {
        margin: 0;
        padding: 10px;
      }
      .ContentWrapper {
        text-align: right;
        width: 750px;
        margin-top: 15px;
        background-color: white;
        border-radius: 7px;
        padding: 10px;
        a {
          font-size: 16px;
          margin: 15px 30px 10px 0;
          display: block;
          color: #1890ff;
        }
      }
    }
  }
}
@media screen and (max-width: 600px) {
  .headMobile {
    width: 100%;
    height: 1.06rem;
    display: flex;
    justify-content: space-between;
    background-color: #ffffff;
    border-bottom: 1px solid #efefef;
    .headCancle {
      font-size: 0.43rem;
      line-height: 1.06rem;
      margin-left: 0.53rem;
    }
    .headMessage {
      align-items: center;
      font-size: 0.48rem;
      line-height: 1.06rem;
    }
    .headSent {
      font-size: 0.43rem;
      color: #1890ff;
      line-height: 1.06rem;
      margin-right: 0.53rem;
    }
  }
  .app-entry {
    display: flex;
    flex-flow: column;
    margin: 0 auto;
    .messagePopupPC {
      display: none;
    }
    .sendMessageDiv {
      margin: 0.53rem;
      .receiverDiv {
        margin: 0;
        padding: 0.26rem;
      }
      .ContentWrapper {
        text-align: right;
        margin-top: 0.4rem;
        background-color: white;
        border-radius: 7px;
        padding: 0.26rem;
        a {
          display: none;
        }
      }
    }
    .selectPage {
      margin-left: 10px;
      margin-top: 2vh;
      font-size: 16px;
    }
    .selectAll {
      margin-left: 5vw;
      margin-top: 2vh;
      font-size: 16px;
    }
    .cancelFont {
      float: right;
      font-size: 16px;
      margin-top: 2vh;
      position: absolute;
      right: 50px;
    }
    .cancelFont:hover {
      color: #1890ff;
    }
    .OkFont {
      position: absolute;
      right: 9px;
      color: #1890ff;
      margin-top: 2vh;
      font-size: 16px;
    }
    .OkFont:hover {
      color: red;
    }
  }
}
</style>

<script lang="ts">
import { defineComponent, onMounted, reactive, watch, ref } from "vue";
import { message } from "ant-design-vue";
import ManageNav from "@/components/common/ManageNav.vue";
import { toRaw } from "@vue/reactivity";
import { useRouter, useRoute } from "vue-router";

export default defineComponent({
  name: "VueComponentSkeleton",

  components: { ManageNav },
  setup() {
    const router = useRouter();
    const route = useRoute();
    watch(
      () => route.path,
      () => {
        websocket.close();
        message.destroy();
      }
    );
    //ws连接
    const websocket = new WebSocket("ws://mac.gzhu.edu.cn/alumni/api/message/ws");
    const code = ref(0);
    websocket.onopen = () => {
      console.log("connected");
    };
    websocket.onmessage = (e) => {
      console.log(JSON.parse(e.data));
    };
    //若连接断开，由于nginx时间设置为3分钟，可能为长时间未操作导致的，尝试刷新重连
    websocket.onclose = (e) => {
      code.value = e.code;
      if (e.code == 1006) {
        message.destroy();
        message.info("长时间未发送消息，连接断开，请刷新重试", 0);
      }
      console.log(e);
    };

    //本页显示用户列表的id数组(用来判断是否单页全选)
    const showThisPageId: any = ref([]);
    //本页显示的用户列表---用于单页全选
    const showUserList: any = ref([]);
    //用户列表总页码
    const userPageCount: any = ref(1);
    //储存当前所有的用户列表---也可能为搜索后得到的符合条件的所有用户列表（搜索后会将其替代掉）
    const allUserList: any = ref([]);
    //真正的所有用户列表
    const recordAllUserList: any = ref([]);
    //记录用户列表的id--同allUserList
    const allUserListId: any = ref([]);
    //真正记录所有用户的id
    const recordAllUserListId: any = ref([]);

    //PC端在加载时获得第一页的数据
    const showPagePC = () => {
      let url = "api/admin/listuser";
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
          userCurrentPage: 1,
          equipment: "PC",
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          if (v.status == -999) {
            message.warn("账号未登录，请先登录");
            router.push("login");
          }

          if (v.status == -1999) {
            message.warn("管理员账号未登录，请先登录");
            router.push("login");
          }
          if (v.status != 0) {
            message.warn("用户加载失败，请刷新重试");
            return;
          } else {
            //获得总共页数
            userPageCount.value = v.pageCount;
            //接受后端传来的第一页数据
            showUserList.value.push(...v.data);
            showUserList.value.forEach((user: any) => {
              showThisPageId.value.push(user.id);
            });
          }
        });
    };
    onMounted(() => {
      let url = "api/admin/listalluser";
      fetch(url, {
        method: "GET",
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: new Headers({
          "Content-Type": "application/json",
        }),
        redirect: "follow",
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          if (v.status == -999) {
            message.warn("账号未登录，请先登录");
            router.push("login");
          }

          if (v.status == -1999) {
            message.warn("管理员账号未登录，请先登录");
            router.push("login");
          }
          if (v.status != 0) {
            message.warn("用户加载失败，请刷新后重试");
          } else {
            allUserList.value.push(...v.alluser);
            recordAllUserList.value.push(...v.alluser);
            allUserList.value.forEach((user: any) => {
              allUserListId.value.push(user.id);
            });
            recordAllUserListId.value = allUserListId.value;
          }
        });
      if (document.body.clientWidth > 600) {
        showPagePC();
      }
    });
    return reactive({
      //此时的所有用户（可能是搜索后得到的全部用户）--用以进行全选
      allUserList,
      //此时所有用户的id
      allUserListId,
      //记录所有用户---真正保存全部用户的数组
      recordAllUserList,
      //记录所有用户的id
      recordAllUserListId,
      //已选择的用户列表---已选择的用户数组
      selectUserList: [],
      //收件人，主题和正文
      inputReceiver: "",
      inputTheme: "",
      inputContent: "",
      //当前页面用户列表显示
      showUserList,
      //当前页码
      currentPage: 1,
      //总页码
      userPageCount,
      //搜索的用户
      searchUser: "",
      //全选的选择状态
      checkedAllStatus: false,
      //本页全选的选择状态
      checkedPageStatus: false,
      //数据绑定数组
      checkModel: [],
      //该页的选项id
      showThisPageId,
      //ws连接
      websocket,
      //判断此时是否是处于搜索状态
      searchStatus: false,
      //Popup的显示
      show: false,
      //Pc端显示第一页界面
      showPagePC,
      code,
    });
  },

  methods: {
    showPopup() {
      //如果是手机端才运行点击收件人栏
      if (document.body.clientWidth <= 600) {
        this.show = true;
        this.allUserList = this.recordAllUserList;
        this.allUserListId = this.recordAllUserListId;
        this.searchStatus = false;
        this.showPage();
      } else {
        return;
      }
    },
    // 显示用户页面
    showPage() {
      let equipment = "";
      if (document.body.clientWidth > 600) {
        equipment = "PC";
      } else {
        equipment = "mobile";
      }
      this.searchUser = "";
      let url = "api/admin/listuser";
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
          userCurrentPage: this.currentPage,
          equipment: equipment,
        }),
      })
        .then((v) => {
          return v.json();
        })
        .then((v) => {
          if (v.status == -999) {
            message.warn("账号未登录，请先登录");
            this.$router.push("login");
          }

          if (v.status == -1999) {
            message.warn("管理员账号未登录，请先登录");
            this.$router.push("login");
          }
          if (v.status != 0) {
            message.warn("用户加载失败，请刷新重试");
          } else {
            //获得总共页数
            this.userPageCount = v.pageCount;
            //本页显示置为空
            this.showUserList = [];
            //接受后端获得的一页数据
            this.showUserList.push(...v.data);
            //本页显示id置空
            this.showThisPageId = [];
            //遍历本页用户重新获得本页id
            this.showUserList.forEach((user: any) => {
              this.showThisPageId.push(user.id);
            });
            let checkModel = toRaw(this.checkModel);
            let showThisPageId = toRaw(this.showThisPageId);
            //翻页后判断本页全选框的状态----遍历本页显示的id是否全部处于checkModel中
            for (let i = 0; i < showThisPageId.length; i++) {
              if (checkModel.indexOf(showThisPageId[i]) > -1) {
                this.checkedPageStatus = true;
              } else {
                this.checkedPageStatus = false;
                break;
              }
            }
            //判断全选框状态----此时已选的长度与当前所有用户的长度相同时则说明全选
            if (this.checkModel.length === this.allUserListId.length) {
              this.checkedAllStatus = true;
            } else {
              this.checkedAllStatus = false;
            }
          }
        });
    },
    //用户列表翻页
    pageChange(page: any, pageSize: any) {
      this.currentPage = page;
      this.OnSearch();
    },
    //搜索
    OnSearch() {
      if (this.searchUser === "") {
        //如果查询的输入为空，则未进入搜索状态
        this.searchStatus = false;
        //将当前全部用户恢复为真正的全部用户
        this.allUserList = this.recordAllUserList;
        this.allUserListId = this.recordAllUserListId;
        //重新显示第一页
        this.showPage();
      } else {
        this.searchStatus = true;
        let equipment = "";
        if (document.body.clientWidth > 600) {
          equipment = "PC";
        } else {
          equipment = "mobile";
        }
        let url = "api/admin/searchuser";
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
            searchContent: this.searchUser,
            userCurrentPage: this.currentPage,
            equipment: equipment,
          }),
        })
          .then((v) => {
            return v.json();
          })
          .then((v) => {
            if (v.status == -999) {
              message.warn("账号未登录，请先登录");
              this.$router.push("login");
            }

            if (v.status == -1999) {
              message.warn("管理员账号未登录，请先登录");
              this.$router.push("login");
            }
            if (v.status != 0) {
              message.warn("搜索失败,请刷新后重试");
              return;
            } else {
              //搜索后当前用户情况改变，全部置空
              this.allUserList = [];
              this.showUserList = [];
              this.showThisPageId = [];
              //重新加入搜索后的用户
              this.showUserList.push(...v.data);
              this.showUserList.forEach((user: any) => {
                this.showThisPageId.push(user.id);
              });
              //获得页数
              this.userPageCount = v.pageCount;
              //获得当前所有用户的id---判断全选
              this.allUserListId = [];
              this.allUserList.push(...v.alluser);
              this.allUserList.forEach((user: any) => {
                this.allUserListId.push(user.id);
              });
              let checkModel = toRaw(this.checkModel);
              let showThisPageId = toRaw(this.showThisPageId);
              //判断本页全选状态
              for (let i = 0; i < this.showThisPageId.length; i++) {
                if (checkModel.indexOf(showThisPageId[i]) > -1) {
                  this.checkedPageStatus = true;
                } else {
                  this.checkedPageStatus = false;
                  break;
                }
              }
              //判断全选状态
              let allUserListId = toRaw(this.allUserListId);
              for (let i = 0; i < this.allUserListId.length; i++) {
                if (checkModel.indexOf(allUserListId[i]) > -1) {
                  this.checkedAllStatus = true;
                } else {
                  this.checkedAllStatus = false;
                  break;
                }
              }
            }
          });
      }
    },
    //全选功能
    checkall() {
      if (this.checkedAllStatus) {
        //取消全选
        if (this.searchStatus === false) {
          //未搜索时，将所有选择置空
          this.checkModel = [];
          this.checkedPageStatus = false;
        } else {
          //处于搜索中，则只是从之前选择的用户中将搜索后的用户去除
          this.allUserList.forEach((user) => {
            var index = this.checkModel.indexOf(user.id);
            if (index != -1) {
              this.checkModel.splice(index, 1);
            }
            this.checkedPageStatus = false;
          });
        }
      } else {
        //选择全选
        //遍历当前所有的用户，并将它们的id加入选中状态
        this.allUserList.forEach((user) => {
          //若该use.id不存在于checkModel中，则将其加入选中状态
          if (this.checkModel.indexOf(user.id) == -1) {
            this.checkModel.push(user.id);
          }
        });
        //选中全选后，将单页全选按钮置为选中状态
        this.checkedPageStatus = true;
      }
    },
    //单页全选
    checkPage() {
      if (this.checkedPageStatus) {
        //单页全选取消
        if (this.checkedAllStatus === true) {
          this.checkedAllStatus = false;
        }
        // this.checkModel.splice(0, this.checkModel.length);
        //在选中用户中去掉本页的显示
        this.showUserList.forEach((user: any) => {
          if (this.checkModel.indexOf(user.id) !== -1) {
            let index = this.checkModel.indexOf(user.id);
            this.checkModel.splice(index, 1);
          }
        });
      } else {
        //本页全选选中时，单页全选未选中--正常不会存在此情况
        if (this.checkedAllStatus === true) {
          this.checkedAllStatus = false;
          //点击单页全选时，若全选状态为选中，则将所有置为空，且将目前本页加入选中数组中
          this.checkModel = [];
          this.showUserList.forEach((user) => {
            if (this.checkModel.indexOf(user.id) == -1) {
              this.checkModel.push(user.id);
            }
          });
        } else {
          //如果此时全选状态为未选中，则正常进行本页全选功能
          this.showUserList.forEach((user: any) => {
            if (this.checkModel.indexOf(user.id) == -1) {
              this.checkModel.push(user.id);
            }
          });
          //根据是否处于搜索状态来决定全选的判断
          if (this.searchStatus === false) {
            if (this.checkModel.length === this.allUserListId.length) {
              this.checkedAllStatus = true;
            } else {
              this.checkedAllStatus = false;
            }
          } else {
            let allUserListId = toRaw(this.allUserListId);
            let checkModel = toRaw(this.checkModel);
            for (let i = 0; i < this.allUserListId.length; i++) {
              if (checkModel.indexOf(allUserListId[i]) > -1) {
                this.checkedAllStatus = true;
              } else {
                this.checkedAllStatus = false;
                break;
              }
            }
          }
        }
      }
    },
    sendReceiver() {
      //如果未选中收件人--则相当于直接退出，不发请求
      if (this.checkModel.length === 0) {
        this.inputReceiver = "";
        this.show = false;
      } else {
        this.inputReceiver = "";
        let sendId = this.checkModel.join(";");
        this.show = false;
        if (this.checkModel != []) {
          let url = "api/admin/listreceiver";
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
              searchId: sendId,
            }),
          })
            .then((v) => {
              return v.json();
            })
            .then((v) => {
              if (v.status == -999) {
                message.warn("账号未登录，请先登录");
                this.$router.push("login");
              }

              if (v.status == -1999) {
                message.warn("管理员账号未登录，请先登录");
                this.$router.push("login");
              }
              if (v.status != 0) {
                message.warn("用户加载失败,请尝试刷新");
                return;
              } else {
                this.selectUserList = [];
                this.selectUserList.push(...v.data);
                let receiverArray = [];
                this.selectUserList.forEach((user) => {
                  receiverArray.push(user.name);
                });
                this.inputReceiver = receiverArray.join(";");
              }
            });
        }
      }
    },
    sendMessage() {
      if (this.selectUserList != [] && this.inputTheme != "" && this.inputContent != "") {
        let id = [];
        this.selectUserList.forEach((user) => {
          let idString = user.id.toString();
          id.push(idString);
        });
        const data = {
          ids: id,
          title: this.inputTheme,
          content: this.inputContent,
          time: Date.now(),
        };
        this.websocket.send(JSON.stringify(data));
        this.inputTheme = "";
        this.inputContent = "";
        this.inputReceiver = "";
        this.checkModel = [];
        this.checkedPageStatus = false;
        this.checkedAllStatus = false;
        if (this.code == 0) {
          message.success("发送成功");
        } else {
          message.warn("发送失败，请刷新重试");
        }
      } else {
        message.info("缺少收件人/主题/正文");
      }
    },
    cancle() {
      this.websocket.close();
      this.$router.push("/managerCenter");
    },
    check(e: any) {
      let checkModel = toRaw(this.checkModel);
      let showThisPageId = toRaw(this.showThisPageId);
      for (let i = 0; i < showThisPageId.length; i++) {
        if (checkModel.indexOf(showThisPageId[i]) > -1) {
          this.checkedPageStatus = true;
        } else {
          this.checkedPageStatus = false;
          break;
        }
      }
      if (this.searchStatus === false) {
        if (this.checkModel.length === this.allUserListId.length) {
          this.checkedAllStatus = true;
        } else {
          this.checkedAllStatus = false;
        }
      } else {
        let allUserListId = toRaw(this.allUserListId);
        for (let i = 0; i < this.allUserListId.length; i++) {
          if (checkModel.indexOf(allUserListId[i]) > -1) {
            this.checkedAllStatus = true;
          } else {
            this.checkedAllStatus = false;
            break;
          }
        }
      }
    },
  },
});
</script>
