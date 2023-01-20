<template>
    <van-config-provider :theme-vars="themeVars">
        <van-nav-bar
            :title="title"
            :right-text="$route.name === 'adminpage' ? '登录' : ''"
            :left-arrow="$route.name === 'adminpage' ? false : true"
            :placeholder="true"
            :safe-area-inset-top="true"
            :fixed="true"
            @click-left="back"
            @click-right="logout"
        >
            <template #left>
                <van-icon v-if="$route.name === 'adminpage'" name="wap-home-o" size="28px" />
                <div v-else :style="{ fontSize: '20px', color: '#fff' }">
                    <van-icon name="arrow-left" />返回
                </div>
            </template>
            <template #right>
                <div
                    v-if="$route.name === 'adminpage'"
                    :style="{ fontSize: '20px', color: '#fff' }"
                >注销</div>
            </template>
        </van-nav-bar>
    </van-config-provider>
</template>

<script setup lang="js">
import { useRoute, useRouter } from 'vue-router'
import axios from "axios";
const route = useRoute()
const router = useRouter()
defineProps<{
    title: string,
}>()


const themeVars = {
    navBarBackgroundColor: '#1890FF',
    navBarIconColor: '#fff',
    navBarTextColor: '#fff',
    navBarTitleTextColor: '#fff',
    navBarHeight: '68px',
    navBarArrowSize: '23px',
    navBarTitleFontSize: '23px',
}
const login = () => {
    router.push('userLogin');
}

//注销登录
const logout = () => {
    //删除服务器sessionID
    axios.post('api/admin/logout').then((res) => {
        console.log(res)
    })
    //删除本地cookies
    let exp = new Date();
    exp.setTime(exp.getTime() - 1);

    let start = document.cookie.indexOf("name=")
    if (start == -1) {
        router.push('/adminLogin')
    }
    let end = document.cookie.indexOf(";", start + 5)
    if (end == -1) {
        end = document.cookie.length
    }
    console.log(end)

    let sessionID = document.cookie.substring(start + 5, end)
    document.cookie = "name=" + sessionID + "; expires=" + exp.toUTCString();
    router.push('adminLogin')
}


const back = () => {
    if (route.name == 'pendingAuditPage') {
        router.push('/adminpage')
    }
    else {
        router.go(-1)
    }

}
</script>

<style >
/* :root{
    --van-tabs-bottom-bar-color:#098AE8
} */
</style>


