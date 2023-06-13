import {defineStore} from "pinia";
import {MenuType, UserInfoType, UserLoginFormType} from "@/type/userApi.ts";
import {getUserInfo, Login} from "@/api/user/index.ts";
import {GET_TOKEN, SET_TOKEN} from "@/utils/tokn.ts";
import {GetMenusInfo, GetMenusList} from "@/api/menu";
import convertToRouteObjects from "@/utils/router.ts";
import router from "@/router";


export const rootStore = defineStore("root", {
    state: (): { token: string, menus: MenuType[], info: UserInfoType, list: MenuType[] } => {
        return {
            token: GET_TOKEN(),
            menus: [],
            info: {
                id: 0,
                username: "",
                avatar: "",
                identity: "",
                routers: [],
                buttons: [],
                roles: []
            },
            list: []
        };
    },
    actions: {
        async login(data: UserLoginFormType) {
            let res = await Login(data);
            console.log(res);
            if (res.data.code === 200) {
                this.token = <string>res.data.data?.token;
                SET_TOKEN(<string>res.data.data?.token);
                return "登录成功";
            } else {
                return Promise.reject(new Error(res.data?.message));
            }
        },
        async getMenusInfo() {
            let res = await GetMenusInfo();
            if (res.code === 200) {
                this.menus = res.data;
                let toRouteObjects = convertToRouteObjects(this.menus);
                toRouteObjects.forEach(item => {
                    router.addRoute(item);
                });
                return "ok";
            } else {
                console.log(res.message);
                return Promise.reject(new Error(res.message));
            }
        },
        async getUserInfo() {
            let res = await getUserInfo();
            if (res.code === 200) {
                this.info = res.data as UserInfoType;
            }
        },
        UserExit() {
            this.info = {
                id: 0,
                username: "",
                avatar: "",
                identity: "",
                routers: [],
                buttons: [],
                roles: []
            };
        },
        async getMenusIist() {
            let res = await GetMenusList();
            this.list = res;
        }
    },
    getters: {}
});