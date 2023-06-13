import request from "@/api/request.ts";
import {
    GeneralResponseType,
    LimitType,
    MenuType,
    MenuUpdateType,
    ResLimitType,
    ResMenuUpdateType
} from "@/type/userApi.ts";

enum MenuApi {
    MENUINFO = "/menu/info",
    MENULIST = "/menu/list",
    MENULIMIT = "/menu/limit",
    MENUUPDATE = "/menu/update"
}

export const GetMenusInfo = () => request.get<any, GeneralResponseType>(MenuApi.MENUINFO);

export const GetMenusList = () => request.get<any, MenuType[]>(MenuApi.MENULIST);

export const GetMenusLimit = (data: LimitType) => request.get<any, ResLimitType>(MenuApi.MENULIMIT, {params: data});

export const MenuUpdate = (data: MenuUpdateType) => request.post<any, ResMenuUpdateType>(MenuApi.MENUUPDATE, data);