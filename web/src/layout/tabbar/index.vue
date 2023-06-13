<template>
    <div class="main">
        <div class="main-left">
            <el-icon style="margin-right: 20px;cursor: pointer" @click="change">
                <component :is="fold?'Expand':'Fold'"></component>
            </el-icon>
            <el-breadcrumb separator-icon="ArrowRight">
                <el-breadcrumb-item
                        v-for="item in route.matched"
                        :key="item.path"
                        v-show="item.meta.title"
                >
                    <el-icon>
                        <component :is="item.meta.icon"></component>
                    </el-icon>
                    {{ item.meta.title }}
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="main-right">
            <el-button icon="Refresh" circle @click="refresh"/>
            <!--            Close-->
            <el-button :icon="isOpen?'FullScreen':'Close'" circle @click="fullScreen"/>


            <el-popover
                    placement="bottom"
                    title="全局设置"
                    :width="200"
                    trigger="hover"
                    content="this is content, this is content, this is content"
            >
                <el-form>
                    <el-form-item label="主题样式">
                        <el-color-picker
                                v-model="color"
                                @change="changeColor"
                                show-alpha
                        />
                    </el-form-item>
                    <el-form-item label="暗黑模式">
                        <el-switch
                                v-model="switchBlack"
                                class="mt-2"
                                style="margin-left: 24px"
                                inline-prompt
                                active-icon="MoonNight"
                                inactive-icon="Sunrise"
                                @change="changeSwitchBlack"
                        />
                    </el-form-item>
                </el-form>
                <template #reference>
                    <el-button icon="Setting" circle/>
                </template>
            </el-popover>


            <img
                    :src="userStore.info.avatar"
                    alt="头像"
                    style="width: 24px;height: 24px; margin: 0 10px;border-radius: 50%"
            >
            <el-dropdown>
                <span class="el-dropdown-link">
                  <span>{{ userStore.info.username }}</span>
                  <el-icon class="el-icon--right">
                    <arrow-down/>
                  </el-icon>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item @click="updateUserInfoShowDialog(userStore.info)">
                            修改个人信息
                        </el-dropdown-item>
                        <el-dropdown-item @click="userExit">
                            退出登录
                        </el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
        <el-dialog v-model="dialogFormVisible" title="修改个人信息,点击头像即可重新上传新头像">
            <el-form
                    :model="newUserInfo"
                    :rules="rules"
                    label-width="90px"
                    ref="formData"
            >
                <el-form-item label="用户名" prop="username">
                    <el-input
                            placeholder="输入新的用户名,空值则不修改"

                            v-model="newUserInfo.username"
                    />
                </el-form-item>
                <el-form-item label="密码" prop="password">
                    <el-input
                            placeholder="输入新的密码,空值则不修改"
                            v-model="newUserInfo.password"
                            show-password
                    />
                </el-form-item>
                <el-form-item label="头像" prop="icon">
                    <el-upload
                            class="avatar-uploader"
                            action=""
                            :http-request="upload"
                            :show-file-list="false"
                            :on-success="handleAvatarSuccess"
                            :before-upload="beforeAvatarUpload"
                            :headers="handler"
                    >
                        <img
                                v-if="newUserInfo.icon"
                                :src="newUserInfo.icon"
                                class="avatar"
                                alt="重复上传以最后一个为主"
                        />
                        <el-icon v-else class="avatar-uploader-icon">
                            <Plus/>
                        </el-icon>
                    </el-upload>
                </el-form-item>
            </el-form>
            <template #footer>
                  <span class="dialog-footer">
                    <el-button
                            type="primary"
                            @click="dialogFormVisible=false"
                    >
                        取消
                    </el-button>
                    <el-button
                            type="danger"
                            @click="updateUserInfo"
                    >
                      确定
                    </el-button>
                  </span>
            </template>
        </el-dialog>
    </div>
</template>

<script setup lang="ts">

import {reactive, ref} from "vue";
import {layoutStore} from "@/store/layout.ts";
import {useRoute, useRouter} from "vue-router";
import {rootStore} from "@/store";
import {UpdateIcon, UpdateUserInfo, UserExit} from "@/api/user";
import {Del_TOKEN} from "@/utils/tokn.ts";
import {ElMessage, FormRules, UploadFile, UploadProps} from "element-plus";
import {ResIconType} from "@/type/userApi.ts";

let route = useRoute();

let router = useRouter();

let store = layoutStore();

let userStore = rootStore();

let fold = ref(false);

let handler = "token:" + userStore.token;

let formData = ref();

const change = () => {
    fold.value = !fold.value;
    store.checkFold();
};

const refresh = () => {
    store.checkRefresh();
};

// 控制全屏按钮icon的切换
let isOpen = ref(true);
// 开启全屏和关闭全屏
const fullScreen = () => {
    // 获取当前是否为全屏
    let is = document.fullscreenElement;
    if (!is) {
        // 开启全屏
        document.documentElement.requestFullscreen();
        isOpen.value = false;
    } else {
        // 关闭全屏
        document.exitFullscreen();
        isOpen.value = true;
    }
};

const userExit = () => {
    UserExit();
    Del_TOKEN();
    userStore.UserExit();
    router.replace("/login");
};

let newUserInfo = reactive({
    token: userStore.token,
    username: "",
    password: "",
    icon: ""
});

let dialogFormVisible = ref<boolean>(false);

const updateUserInfoShowDialog = (user) => {
    console.log(user);
    dialogFormVisible.value = true;
    newUserInfo.icon = user.avatar;
    newUserInfo.password = "";
    newUserInfo.username = user.username;
};
// 上传成功后的回调
const handleAvatarSuccess: UploadProps["onSuccess"] = (
    response: ResIconType,
    uploadFile
) => {
    console.log(response);
    newUserInfo.icon = response.data.url;
};

// 上传前的回调
const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
    if (rawFile.type === "image/jpeg" || rawFile.type === "image/png" || rawFile.type === "image/gif") {
        if (rawFile.size / 1024 / 1024 > 2) {
            ElMessage.error("图片大小超出限制,大小请在2mb内!");
            return false;
        }
        return true;
    } else {
        ElMessage.error("图片类型不符合");
        return false;
    }
};


const upload = (param) => {
    const formData = new FormData();
    formData.append("file", param.file);
    UpdateIcon(formData).then((res) => {
        console.log(res);
        newUserInfo.icon = res.data.url;
    });
};

const updateUserInfo = async () => {
    try {
        await formData.value.validate();
        dialogFormVisible.value = false;
        UpdateUserInfo(newUserInfo).then((res) => {
            console.log(res);
            if (res.code === 200) {
                ElMessage.info("信息修改成功,请重新登陆");
                Del_TOKEN();
                userStore.token = "";
                router.replace("/login");
            }
        });
    } catch (error) {
        ElMessage.error("用户名和头像不能为空");
    }
};


const rules = reactive<FormRules>({
    username: [
        {required: true, min: 6, max: 10, message: "账号长度至少六位,最多十位", trigger: "change"}
        // { trigger: 'change', validator: validatorUserName }
    ]
});

let switchBlack = ref<boolean>(false);


const changeSwitchBlack = (v) => {
    let html = document.documentElement;
    store.checkDark();
    if (v) {
        html.className = "dark";
    } else {
        html.className = "";
    }
};

let color = ref("");

const changeColor = () => {
    document.documentElement.style.setProperty("--el-color-primary", color.value);
};

</script>

<style scoped lang="scss">
.main {
  display: flex;
  justify-content: space-between;
  //background-image: linear-gradient(to right, #2f2c2c, #238888, #005d5d);

  .main-left {
    display: flex;
    align-items: center;
    margin-left: 30px;
  }

  .main-right {
    display: flex;
    align-items: center;
  }
}

.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>