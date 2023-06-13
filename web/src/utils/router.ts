import {RouteRecordRaw} from "vue-router";
import {MenuType} from "@/type/userApi.ts";

function convertToRouteObjects(data: MenuType[]): RouteRecordRaw[] {
    return data.map((item: MenuType) => {
        const route: RouteRecordRaw = {
            path: item.path,
            name: item.name,
            component: () => import(".." + `${item.component}`),
            meta: {
                title: item.title,
                icon: item.icon
            },
            children: item.children ? convertToRouteObjects(item.children) : []
        };
        return route;
    });
}


export default convertToRouteObjects;
