import{_ as c,M as i,b as l}from"./index.11841eca.js";import{x as d,j as g,m as o,r as s,o as p,d as m,f as t,e as _,F as h}from"./vendor.e1400650.js";const f=d({components:{MobileNav:i,ManageNav:l},computed:{},watch:{},setup(){return console.log("setup"),g({picUrl:"",imgBase:""})},created(){console.log("created"),this.getQrcode()},methods:{getQrcode(){console.log("it's function getQrcode"),fetch("api/admin/getQrcode",{method:"Post",mode:"cors",cache:"no-cache",credentials:"same-origin",headers:new Headers({"Content-Type":"application/json"}),redirect:"follow",body:JSON.stringify({act_id:this.$route.params.act_id})}).then(e=>(console.log(e),e.json())).then(e=>{if(e.status==-999){o.warn("\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login");return}if(e.status==-1999){o.warn("\u7BA1\u7406\u5458\u8D26\u53F7\u672A\u767B\u5F55\uFF0C\u8BF7\u5148\u767B\u5F55"),this.$router.push("login");return}if(!e){o.error("null response");return}if(e.status!==0){o.error(e.msg),console.log(e.status),console.log(e.msg);return}this.imgBase=e.qrcode,this.picUrl="data:png/jpeg;base64,"+this.imgBase,console.log(e.msg),console.log("\u6210\u529F\u83B7\u53D6\u4E8C\u7EF4\u7801")}).catch(e=>{console.log(e),o.error(e)})}}}),F={class:"main"};function B(a,e,N,v,C,M){const n=s("ManageNav"),r=s("MobileNav"),u=s("van-image");return p(),m(h,null,[t(n),t(r,{title:"\u6D3B\u52A8\u4E8C\u7EF4\u7801"}),_("div",F,[t(u,{class:"img",src:a.picUrl,fit:"fill"},null,8,["src"])])],64)}var Q=c(f,[["render",B],["__scopeId","data-v-366291a8"]]);export{Q as default};