import request from "@/api/request.ts";
import {
    ResIconType, ResUpdateUserType, UpdateUserType,
    UserInfoResType,
    UserLoginFormType,
    UserLoginResponseType,
    VerifyTokenType
} from "@/type/userApi.ts";
import axios from "axios";

enum UserApi {
    LOGIN = "http://127.0.0.1:9666/api/user/login",
    USERINFO = "/user/info",
    VERIFYTOKEN = "/user/verify",
    USEREXIT = "/user/exit",
    UPDATEICON = "/user/image",
    UPDATEUSERINFO = "/user/update"
}

export const Login = (userForm: UserLoginFormType) => axios.post<any, UserLoginResponseType>(UserApi.LOGIN, userForm);

export const getUserInfo = () => request.get<any, UserInfoResType>(UserApi.USERINFO);

export const verifyToken = () => request.get<any, VerifyTokenType>(UserApi.VERIFYTOKEN);

export const UserExit = () => request.get(UserApi.USEREXIT);

export const UpdateIcon = (data: FormData) => request.post<any, ResIconType>(UserApi.UPDATEICON, data);

export const UpdateUserInfo = (data: UpdateUserType) => request.post<any, ResUpdateUserType>(UserApi.UPDATEUSERINFO, data);