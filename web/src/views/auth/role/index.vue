<template>
    <div>
        <el-card>
            <el-form :inline="true">
                <el-form-item label="一级路由">
                    <el-select v-model="routerId.oneMenu">
                        <el-option
                                v-for="item in userStore.menus"
                                :key="item.r_id"
                                :value="item.id"
                                :label="item.title"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="二级路由">
                    <el-select v-model="routerId.TwoMenu">
                        <el-option
                                v-for="item in routerData.twoRouterData"
                                :key="item.id"
                                :label="item.title"
                                :value="item.id"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="三级路由">
                    <el-select v-model="routerId.ThreeMenu">
                        <el-option
                                v-for="item in routerData.threeRouterData"
                                :key="item.id"
                                :label="item.title"
                                :value="item.id"
                        >
                        </el-option>
                    </el-select>
                </el-form-item>
            </el-form>
        </el-card>
        <el-card
                style="margin: 20px 0"
        >
            <el-dialog v-model="dialogFormVisible" title="修改路由,危险操作!!!">
                <el-form v-model="newRouter" label-width="100px">
                    <el-form-item label="路由路径">
                        <el-input disabled :placeholder="oldRouter.path"/>
                    </el-form-item>
                    <el-form-item label="路由名称">
                        <el-input disabled :placeholder="oldRouter.name"/>
                    </el-form-item>
                    <el-form-item label="路由标题">
                        <el-input v-model="newRouter.Title"/>
                    </el-form-item>
                    <el-form-item label="路由图标">
                        <el-input v-model="newRouter.Icon"/>
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
                                @click="routerUpdateTrue"
                        >
                            确定
                        </el-button>
                    </span>
                </template>
            </el-dialog>
            <el-button
                    type="primary"
                    size="small"
                    round
                    v-if="routerInfo"
            >
                添加路由
            </el-button>
            <el-table :data="routerInfo">
                <el-table-column label="路由id" prop="id"></el-table-column>
                <el-table-column label="路由路径" prop="path"></el-table-column>
                <el-table-column label="路由名称" prop="name"></el-table-column>
                <el-table-column label="路由标题" prop="title"></el-table-column>
                <el-table-column label="路由图标" prop="icon">
                    <template #default="{row,$index}">
                        <el-icon>
                            <component :is="row.icon"></component>
                        </el-icon>
                    </template>
                </el-table-column>
                <el-table-column label="操作">
                    <template #default="{row,$index}">
                        <el-button
                                size="small"
                                type="warning"
                                round
                                @click="routerUpdate(row)"
                        >
                            修改
                        </el-button>
                        <el-button
                                size="small"
                                type="danger"
                                round
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>

<script setup lang="ts">

import {rootStore} from "@/store";
import {ref, watch} from "vue";
import {MenuType} from "@/type/userApi.ts";
import {GET_TOKEN} from "@/utils/tokn.ts";
import {MenuUpdate} from "@/api/menu";
import {ElMessage} from "element-plus";

let userStore = rootStore();

let routerData = ref<{ twoRouterData: MenuType[], threeRouterData: MenuType[] }>({
    twoRouterData: [],
    threeRouterData: []
});


let routerId = ref({
    oneMenu: "",
    TwoMenu: "",
    ThreeMenu: ""
});

let routerInfo = ref<MenuType[]>();

watch(() => routerId.value.oneMenu, (value, oldValue, onCleanup) => {
    routerId.value.TwoMenu = "";
    userStore.menus.forEach(item => {
        if (item.id === Number(value)) {
            routerData.value.twoRouterData = item.children;
            routerInfo.value = routerData.value.twoRouterData;
            if (!item.children) {
                routerInfo.value = [item];
            }
        }
    });
});

watch(() => routerId.value.TwoMenu, (value, oldValue, onCleanup) => {
    routerId.value.ThreeMenu = "";
    if (routerData.value.twoRouterData) {
        routerData.value.twoRouterData.forEach(item => {
            if (item.id === Number(value)) {
                routerData.value.threeRouterData = item.children;
                routerInfo.value = routerData.value.threeRouterData;
                if (!item.children) {
                    routerInfo.value = [item];
                }
            }
        });
    }
});
watch(() => routerId.value.ThreeMenu, (value, oldValue, onCleanup) => {
    routerData.value.threeRouterData.forEach(item => {
        if (item.id === Number(value)) {
            routerInfo.value = [item];
        }
    });
});

let dialogFormVisible = ref<boolean>(false);

let newRouter = ref({
    Token: GET_TOKEN(),
    Id: 0,
    Title: "",
    Icon: ""
});

let oldRouter = ref<MenuType>({
    children: [],
    component: "",
    hidden: 0,
    icon: "",
    id: 0,
    menu_level: 0,
    name: "",
    parent_id: 0,
    path: "",
    r_id: 0,
    sort: 0,
    title: ""
});

function routerUpdate(r: MenuType) {
    oldRouter.value = r;
    dialogFormVisible.value = true;
    newRouter.value.Title = r.title;
    newRouter.value.Icon = r.icon;
    newRouter.value.Id = r.id;
}


const routerUpdateTrue = () => {
    dialogFormVisible.value = false;
    MenuUpdate(newRouter.value).then(res => {
        if (res.code === 200) {
            userStore.menus = [];
            userStore.getMenusInfo();
            ElMessage.info("修改成功");
        } else {
            ElMessage.error("修改失败");
        }
    });
};

</script>

<style scoped>

</style>