<template>
    <div class="box">
        <el-row>
            <el-col>
                <el-form
                        :model="loginForm"
                        class="login_form"
                        :rules="rules"
                        ref="form"
                >
                    <el-form-item
                            prop="username"
                    >
                        <el-input
                                class="input"
                                placeholder="请输入账号"
                                prefix-icon="UserFilled"
                                v-model="loginForm.username"
                                autofocus
                        ></el-input>
                    </el-form-item>
                    <el-form-item prop="password">
                        <el-input
                                class="input"
                                placeholder="请输入密码"
                                prefix-icon="Lock"
                                show-password
                                v-model="loginForm.password"
                        ></el-input>
                    </el-form-item>
                    <el-form-item>
                        <el-button
                                :loading="loading"
                                class="button"
                                @click="login">
                            Login
                        </el-button>
                    </el-form-item>
                </el-form>
            </el-col>
        </el-row>
    </div>
</template>

<script setup lang="ts">
import {reactive, ref} from "vue";
import {rootStore} from "@/store/index.ts";
import {ElNotification, FormRules} from "element-plus";
import {useRouter} from "vue-router";
import {GetTime} from "@/utils/time.ts";

let router = useRouter();

let store = rootStore();

let loginForm = reactive({
    username: "",
    password: ""
});

let form = ref();

const rules = reactive<FormRules>({
    username: [
        {required: true, min: 5, max: 10, message: "账号长度至少六位,最多十位", trigger: "change"}
        // { trigger: 'change', validator: validatorUserName }
    ],
    password: [
        {required: true, min: 5, max: 15, message: "密码长度至少6位,最多十五位", trigger: "change"}
        // { trigger: 'change', validator: validatorPassword }
    ]
});


let loading = ref(false);

const login = async () => {
    await form.value.validate();

    loading.value = true;
    debugger
    try {
        await store.login(loginForm);
        ElNotification({
            type: "success",
            title: "Hi, 欢迎回来",
            message: `${GetTime()}好`
        });
        loading.value = false;
        await router.push("/home");
    } catch (error: Error) {
        ElNotification({
            type: "error",
            message: error.message
        });
        loading.value = false;
    }
};
</script>

<style scoped lang="scss">
.box {
  width: 100%;
  height: 100vh;
  background: #75e5e5;
  background-size: contain;
  overflow: hidden;
  display: flex;
  justify-content: center;

  .login_form {
    width: 400px;
    margin-top: 100px;
    opacity: 0.9;
    padding: 50px;
    backdrop-filter: blur(20px);
    color: #fff;
    border-radius: 20px;
    box-shadow: 0 0 30px 10px rgba(0, 0, 0, .4);

    .button {
      width: 300px;
      margin: 0 auto;
    }
  }
}
</style>