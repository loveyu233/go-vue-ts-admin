// 登录表单类型
export interface UserLoginFormType {
    username: string,
    password: string
}

// 登录返回类型
export interface UserLoginResponseType {
    code?: number;
    message?: string;
    data?: {
        token?: string;
    };
}

// 路由类型
export interface MenuType {
    id: number,
    r_id: number,
    menu_level: number,
    parent_id: number,
    path: string,
    name: string,
    hidden: number,
    component: string,
    sort: number,
    title: string,
    icon: string,
    children: MenuType[]
}

// 动态路由返回类型
export interface GeneralResponseType {
    code?: number;
    message?: string,
    data: MenuType[]
}

// 用户相关数据类型
export interface UserInfoType {
    id: number,
    username: string,
    avatar: string,
    identity: string,
    routers: {
        name: string
    }[],
    buttons: {
        name: string
    }[],
    roles: {
        name: string
    }[]
}

// 返回用户数据类型
export interface UserInfoResType {
    code?: number,
    message?: string,
    data?: UserInfoType
}

// 返回验证token类型
export interface VerifyTokenType {
    code?: number,
    message?: string,
    data: {
        verify?: boolean
    }
}

// 用户退出返回类型
export interface UserExitType {
    code?: number,
    data: {
        message?: string
    }
}

export interface LimitType {
    page: number,
    pageSize: number
}

export interface ResLimitType {
    code: number,
    data: {
        menus: MenuType[],
        total: number
    }
}

export interface ResIconType {
    code: number,
    data: {
        url: string
    }
}


export interface UpdateUserType {
    token: string
    username: string,
    password: string,
    icon: string
}

export interface ResUpdateUserType {
    code: number,
    message: string,
    data: {
        "msg": string
    }
}


export interface MenuUpdateType {
    Id: number,
    Title: string,
    Icon: string,
    Token: string
}

export interface ResMenuUpdateType {
    code: number,
    message: string,
    data: {
        msg: string
    }
}

export interface CarouselType {
    id: number,
    url: string
    isShow: boolean
}

export interface ResCarouselType {
    code: number,
    message: string,
    data: CarouselType[]
}

export interface AddCarouselType {
    code: number,
    data: number,
    message: number
}