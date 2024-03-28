import request from "@/api/request.ts";
import {
    ResIconType, ResUpdateUserType, UpdateUserType,
    UserInfoResType,
    UserLoginFormType,
    UserLoginResponseType,
    VerifyTokenType
} from "@/type/userApi.ts";

enum UserApi {
    LOGIN = "/user/login",
    USERINFO = "/user/info",
    VERIFYTOKEN = "/user/verify",
    USEREXIT = "/user/exit",
    UPDATEICON = "/user/image",
    UPDATEUSERINFO = "/user/update"
}

export const Login = (userForm: UserLoginFormType) => request.post<any, UserLoginResponseType>(UserApi.LOGIN, userForm);

export const getUserInfo = () => request.get<any, UserInfoResType>(UserApi.USERINFO);

export const verifyToken = () => request.get<any, VerifyTokenType>(UserApi.VERIFYTOKEN);

export const UserExit = () => request.get(UserApi.USEREXIT);

export const UpdateIcon = (data: FormData) => request.post<any, ResIconType>(UserApi.UPDATEICON, data);

export const UpdateUserInfo = (data: UpdateUserType) => request.post<any, ResUpdateUserType>(UserApi.UPDATEUSERINFO, data);