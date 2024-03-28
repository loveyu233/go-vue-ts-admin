import {createRouter, createWebHistory} from "vue-router";

const router = createRouter({
    history: createWebHistory("/"),
    routes: [],
    scrollBehavior() {
        return {
            left: 0,
            top: 0
        };
    }
});

import {rootStore} from "@/store";
import {verifyToken} from "@/api/user";
// @ts-ignore
import {ElNotification} from "element-plus";
import nprogress from "nprogress";
import "nprogress/nprogress.css";

router.beforeEach((to, _, next) => {
    nprogress.start();
    if (to.path === "/" || to.path === "/home") {
        router.replace("/home/index");
    }
    if (to.path.includes("/home")) {
        verifyToken().then((res) => {
            console.log(res);
            if (res.code === 200) {
                next();
            } else {
                console.log("登录.....");
                router.replace("/login");
                ElNotification({
                    title: "Error",
                    message: "请登录"
                });
            }
        });
    }
    let store = rootStore();
    // TODO 动态获取路由
    if (store.menus.length === 0) {
        store.getMenusInfo().then(() => {
            router.replace(to.path);
        });
    }
    // @ts-ignore
    document.title = to.meta.title ? to.meta.title : "";
    next();
});


router.afterEach(() => {
    nprogress.done();
});
export default router;