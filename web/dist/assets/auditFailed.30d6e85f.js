import{x,b as p,j as _,z as I,l as v,m as r,r as c,y as j,o as i,d as m,f as a,e as s,g as n,F as B,p as z,i as N,k as C,h as L,t as W,w as U,Y as k}from"./vendor.e1400650.js";import{_ as H,b as M}from"./index.11841eca.js";const O=x({components:{ManageNav:M},setup(){const t=new WebSocket("ws://mac.gzhu.edu.cn/alumni/api/message/ws?id=-1"),u=p("https://img.yzcdn.cn/vant/cat.jpeg"),d=p(!1),F=p(4),A=p(50),w={navBarTextColor:"#FFFFFF",navBarIconColor:"#FFFFFF",navBarTitleTextColor:"#FFFFFF",navBarBackgroundColor:"#1890FF",navBarHeight:"1.7333rem",fieldPlaceholderTextColor:"#333333",fieldWordLimitFontSize:".4533rem",fieldErrorMessageFontSize:".4533rem",fieldClearIconSize:".8rem",fieldLabelWidth:"5em",dialogHeaderLineHeight:"2rem",dialogWidth:"350px"},b="2132656565",o=p(""),V=_([]),f="\u672A\u901A\u8FC7",g="api/file/download",E="http://mac.gzhu.edu.cn/alumni/resources/",D=_([]),h=_([]),l=_([]),e=()=>history.back(),S=p(0);return I(S,(T,$)=>{if(+T!=6){console.log($);return}console.log("it is 6")}),_({jieguo:f,websocket:t,srcList:D,srcList1:h,srcList2:l,fileList:V,displayUrl:E,downloadUrl:g,onClickLeft:e,book1:"",ww:u,show:d,text:b,value:o,rate:F,slider:A,themeVars:w,datas1:[]})},beforeCreate(){console.log("beforeCreate")},created(){console.log("created"),this.onEdit()},methods:{async async1(t,u){const d=new FormData;d.append("id",t);const F=await v.post(this.downloadUrl,d);F.data.statusCode===2e3&&(this.srcList.push(F.data.details),this.async2(this.datas1[0].Id_Ini))},async async2(t){const u=new FormData;u.append("id",t);const d=await v.post(this.downloadUrl,u);d.data.statusCode===2e3&&(this.srcList1.push(d.data.details),this.async3(this.datas1[0].Id_Cer))},async async3(t){const u=new FormData;u.append("id",t);const d=await v.post(this.downloadUrl,u);d.data.statusCode===2e3&&this.srcList2.push(d.data.details)},confirm1(){fetch("api/user/userchange",{method:"POST",mode:"cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow",body:JSON.stringify({id:this.$route.params.id,sign:"2",book:this.book1})}).then(u=>(console.log(u),u.json())).then(u=>{if(u.status==-999&&(r.warn("\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login")),u.status==-1999&&(r.warn("\u7BA1\u7406\u5458\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login")),!u){r.error("\u7F51\u7EDC\u5F00\u5C0F\u5DEE\u4E86\uFF0C\u8BF7\u5237\u65B0\u91CD\u8BD5\uFF01");return}if(u.status!==0){r.error("\u7F51\u7EDC\u5F00\u5C0F\u5DEE\u4E86\uFF0C\u8BF7\u5237\u65B0\u91CD\u8BD5\uFF01");return}else this.transmit(),r.success("\u7528\u6237\u5BA1\u6838\u72B6\u6001\u66F4\u6539\u6210\u529F\uFF01"),setTimeout(()=>{this.$router.go(-1),this.$router.replace("/pendingAuditPage")},1e3)}).catch(u=>{console.log(u)})},transmit(){let t=this.$route.params.id.toString();this.book1==""&&(this.book1="\u5DF2\u901A\u8FC7");const u={ids:[t],title:"\u5DF2\u901A\u8FC7",content:this.book1,time:Date.now()};this.websocket.send(JSON.stringify(u))},onEdit(){fetch("api/user/already",{method:"POST",mode:"cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow",body:JSON.stringify({id:this.$route.params.id})}).then(u=>(console.log(u),u.json())).then(u=>{if(u.status==-999&&(r.warn("\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login")),u.status==-1999&&(r.warn("\u7BA1\u7406\u5458\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login")),!u){r.error("\u7F51\u7EDC\u5F00\u5C0F\u5DEE\u4E86\uFF0C\u8BF7\u5237\u65B0\u91CD\u8BD5\uFF01");return}if(u.status!==0){r.error("\u7F51\u7EDC\u5F00\u5C0F\u5DEE\u4E86\uFF0C\u8BF7\u5237\u65B0\u91CD\u8BD5\uFF01");return}else this.datas1=u.data,this.async1(this.datas1[0].Id_Peo,this.datas1[0].Id_Ini)}).catch(u=>{console.log(u)})}}}),y=t=>(z("data-v-25cdf87e"),t=t(),N(),t),P=y(()=>s("div",{id:"buildingg1"},null,-1)),J={class:"bground"},q=y(()=>s("div",{class:"font"},"\u57FA\u672C\u4FE1\u606F",-1)),Y={class:"imgWrapper"},G={class:"imgWrapper"},K=y(()=>s("div",{class:"font"},"\u6559\u80B2\u7ECF\u5386",-1)),Q={class:"imgWrapper"},R=y(()=>s("div",{class:"font"},"\u5DE5\u4F5C\u7ECF\u5386",-1)),X=["onUpdate:modelValue"],Z=L("\u4FEE\u6539\u5BA1\u6838\u7ED3\u679C"),ee=y(()=>s("div",{class:"dio"},"\u662F\u5426\u786E\u8BA4\u66F4\u6539\u5BA1\u6838\u7ED3\u679C\u4E3A\uFF1A\u901A\u8FC7",-1));function le(t,u,d,F,A,w){const b=c("ManageNav"),o=c("van-field"),V=c("van-image"),f=c("van-cell-group"),g=c("van-button"),E=c("van-dialog"),D=c("van-config-provider"),h=j("viewer");return i(),m(B,null,[a(b),P,s("div",J,[a(D,{"theme-vars":t.themeVars},{default:n(()=>[(i(!0),m(B,null,C(t.datas1,l=>(i(),m("div",{key:l.id,class:"one"},[q,a(f,{inset:!0},{default:n(()=>[a(o,{modelValue:l.phone,"onUpdate:modelValue":e=>l.phone=e,class:"filed",label:"\u624B\u673A\u53F7",readonly:""},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.sex,"onUpdate:modelValue":e=>l.sex=e,label:"\u6027\u522B",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.idcard,"onUpdate:modelValue":e=>l.idcard=e,class:"filed",readonly:""},{label:n(()=>[L(W(l.idcard_type),1)]),_:2},1032,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.email,"onUpdate:modelValue":e=>l.email=e,label:"\u7535\u5B50\u90AE\u4EF6",class:"filed",readonly:""},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.education,"onUpdate:modelValue":e=>l.education=e,label:"\u5B66\u5386",class:"filed",readonly:""},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.birthday,"onUpdate:modelValue":e=>l.birthday=e,class:"filed",label:"\u51FA\u751F\u5E74\u6708\u65E5",maxlength:"11",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.nation,"onUpdate:modelValue":e=>l.nation=e,label:"\u6C11\u65CF",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.native_place,"onUpdate:modelValue":e=>l.native_place=e,label:"\u7C4D\u8D2F",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.birth_place,"onUpdate:modelValue":e=>l.birth_place=e,label:"\u51FA\u751F\u5730",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.political,"onUpdate:modelValue":e=>l.political=e,label:"\u653F\u6CBB\u9762\u8C8C",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.overeseas_permanent,"onUpdate:modelValue":e=>l.overeseas_permanent=e,label:"\u56FD\uFF08\u5883\uFF09\u5916\u6C38\u4E45\u5C45\u7559\u6743",class:"filed",readonly:""},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.level,"onUpdate:modelValue":e=>l.level=e,class:"filed",label:"\u5916\u8BED\u7B49\u7EA7",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.translation,"onUpdate:modelValue":e=>l.translation=e,label:"\u662F\u5426\u80FD\u62C5\u4EFB\u7FFB\u8BD1",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.married_status,"onUpdate:modelValue":e=>l.married_status=e,label:"\u5A5A\u59FB\u72B6\u51B5",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.home_phone,"onUpdate:modelValue":e=>l.home_phone=e,class:"filed",label:"\u5BB6\u5EAD\u7535\u8BDD",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.address,"onUpdate:modelValue":e=>l.address=e,class:"filed",label:"\u901A\u8BAF\u5730\u5740",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.postal_code,"onUpdate:modelValue":e=>l.postal_code=e,class:"filed",label:"\u90AE\u653F\u7F16\u7801",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.fax,"onUpdate:modelValue":e=>l.fax=e,class:"filed",label:"\u4F20\u771F",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.hobby,"onUpdate:modelValue":e=>l.hobby=e,class:"filed",label:"\u5174\u8DA3\u7231\u597D",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.social_appointments,"onUpdate:modelValue":e=>l.social_appointments=e,label:"\u4EBA\u5927\u653F\u534F\u53CA\u5176\u4ED6\u793E\u4F1A\u517C\u804C",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{label:"\u63A8\u8350\u4EBA\u7B7E\u7AE0:",class:"filed",readonly:!0},{input:n(()=>[s("div",Y,[(i(!0),m(B,null,C(t.srcList[0],e=>U((i(),m("div",{key:e},[a(V,{fit:"cover",src:t.displayUrl+e},null,8,["src"])])),[[h,{movable:!1}]])),128))])]),_:1}),a(o,{modelValue:l.relation,"onUpdate:modelValue":e=>l.relation=e,class:"filed",label:"\u4E0E\u4ECB\u7ECD\u4EBA\u5173\u7CFB",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.content,"onUpdate:modelValue":e=>l.content=e,class:"filed",label:"\u4ECB\u7ECD\u5355\u4F4D\u610F\u89C1",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{label:"\u63A8\u8350\u5355\u4F4D\u7B7E\u7AE0:",class:"filed",readonly:!0},{input:n(()=>[s("div",G,[(i(!0),m(B,null,C(t.srcList1[0],e=>U((i(),m("div",{key:e},[a(V,{fit:"cover",src:t.displayUrl+e},null,8,["src"])])),[[h,{movable:!1}]])),128))])]),_:1})]),_:2},1024),K,a(f,{inset:!0},{default:n(()=>[a(o,{modelValue:l.degree,"onUpdate:modelValue":e=>l.degree=e,label:"\u5B66\u4F4D",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.major,"onUpdate:modelValue":e=>l.major=e,class:"filed",label:"\u5B66\u79D1/\u4E13\u4E1A",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.college,"onUpdate:modelValue":e=>l.college=e,class:"filed",label:"\u5B66\u6821\u540D\u79F0",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.region,"onUpdate:modelValue":e=>l.region=e,label:"\u5B66\u6821\u6240\u5728\u5730\u548C\u5730\u533A",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.begin_time,"onUpdate:modelValue":e=>l.begin_time=e,label:"\u8D77\u59CB\u5728\u6821\u65F6\u95F4",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.end_time,"onUpdate:modelValue":e=>l.end_time=e,class:"filed",label:"\u7ED3\u675F\u5728\u6821\u65F6\u95F4",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.study_type,"onUpdate:modelValue":e=>l.study_type=e,class:"filed",label:"\u662F\u5426\u7559\u5B66",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{label:"\u8BC1\u660E\u8D44\u6599\uFF08\u5B66\u5386\u5B66\u4F4D\u3001\u7559\u5B66\u56DE\u56FD\u4EBA\u5458\uFF0C\u6BD5\u4E1A\u8BC1\u4E66\u7B49\u8D44\u6599\uFF09:",class:"filed",readonly:!0},{input:n(()=>[s("div",Q,[(i(!0),m(B,null,C(t.srcList2[0],e=>U((i(),m("div",{key:e},[a(V,{fit:"cover",src:t.displayUrl+e},null,8,["src"])])),[[h,{movable:!1}]])),128))])]),_:1})]),_:2},1024),R,a(f,{inset:!0},{default:n(()=>[a(o,{modelValue:l.position,"onUpdate:modelValue":e=>l.position=e,label:"\u804C\u4F4D",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.position_title,"onUpdate:modelValue":e=>l.position_title=e,class:"filed",label:"\u804C\u79F0",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.professional,"onUpdate:modelValue":e=>l.professional=e,class:"filed",label:"\u4ECE\u4E8B\u4E13\u4E1A",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.organization,"onUpdate:modelValue":e=>l.organization=e,class:"filed",label:"\u673A\u6784\u540D\u79F0",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.organization_nature,"onUpdate:modelValue":e=>l.organization_nature=e,label:"\u673A\u6784\u6027\u8D28",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.or_telephone,"onUpdate:modelValue":e=>l.or_telephone=e,class:"filed",label:"\u673A\u6784\u7535\u8BDD",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.w_region,"onUpdate:modelValue":e=>l.w_region=e,label:"\u673A\u6784\u6240\u5728\u56FD\u5BB6\u548C\u5730\u533A",class:"filed",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.begin_time,"onUpdate:modelValue":e=>l.begin_time=e,label:"\u5DE5\u4F5C\u8D77\u59CB\u65F6\u95F4",class:"filed",maxlength:"11",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:l.end_time,"onUpdate:modelValue":e=>l.end_time=e,label:"\u5DE5\u4F5C\u7ED3\u675F\u65F6\u95F4",class:"filed",maxlength:"11",readonly:!0},null,8,["modelValue","onUpdate:modelValue"]),a(o,{modelValue:t.jieguo,"onUpdate:modelValue":u[0]||(u[0]=e=>t.jieguo=e),label:"\u5BA1\u6838\u7ED3\u679C",class:"filed",maxlength:"11",error:!0,readonly:!0},null,8,["modelValue"]),a(o,{label:"\u5BA1\u6838\u610F\u89C1",class:"filed"},{input:n(()=>[U(s("textarea",{"onUpdate:modelValue":e=>l.opinion=e,textarea:"",readonly:"",class:"notThroughTextarea"},null,8,X),[[k,l.opinion]])]),_:2},1024)]),_:2},1024),a(g,{type:"primary",class:"notThroughButton",onClick:u[1]||(u[1]=e=>t.show=!0)},{default:n(()=>[Z]),_:1}),a(E,{class:"dio1",show:t.show,"onUpdate:show":u[3]||(u[3]=e=>t.show=e),"show-cancel-button":"","confirm-button-color":"#1890FF","close-on-click-overlay":!0,onConfirm:t.confirm1},{default:n(()=>[ee,U(s("textarea",{"onUpdate:modelValue":u[2]||(u[2]=e=>t.book1=e),textarea:"",class:"textarea1",placeholder:"\u8F93\u5165\u5BA1\u6838\u610F\u89C1..."},null,512),[[k,t.book1]])]),_:1},8,["show","onConfirm"])]))),128))]),_:1},8,["theme-vars"])])],64)}var ue=H(O,[["render",le],["__scopeId","data-v-25cdf87e"]]);export{ue as default};