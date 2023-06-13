<template>
    <div class="main">
        <el-card
                class="box-card"
                shadow="hover"
                style="height: 100vh"
        >
            <el-dialog v-model="dialogFormVisible" :title="dialogTitle">
                <el-form label-width="80px">
                    <el-form-item label="r_id">
                        <el-input
                                placeholder="路由唯一id,不可重复"
                                v-model="newMenu.r_id"
                        />
                    </el-form-item>
                    <el-form-item label="menu_level">
                        <el-input
                                placeholder="路由层级"
                                v-model="newMenu.menu_level"
                        />
                    </el-form-item>
                    <el-form-item label="parent_id">
                        <el-input
                                placeholder="路由的父路由id"
                                v-model="newMenu.parent_id"
                        />
                    </el-form-item>
                    <el-form-item label="path">
                        <el-input
                                placeholder="路由path"
                                v-model="newMenu.parent_id"
                        />
                    </el-form-item>
                    <el-form-item label="name">
                        <el-input
                                placeholder="路由name"
                                v-model="newMenu.name"
                        />
                    </el-form-item>
                    <el-form-item label="hidden">
                        <el-input
                                placeholder="路由是否隐藏"
                                v-model="newMenu.hidden"
                        />
                    </el-form-item>
                    <el-form-item label="component">
                        <el-input
                                placeholder="路由对应vue页面所在位置"
                                v-model="newMenu.component"
                        />
                    </el-form-item>
                    <el-form-item label="sort">
                        <el-input
                                placeholder="路由在自己层级的显示顺序"
                                v-model="newMenu.sort"
                        />
                    </el-form-item>
                    <el-form-item label="title">
                        <el-input
                                placeholder="路由对应的网页标题"
                                v-model="newMenu.title"
                        />
                    </el-form-item>
                    <el-form-item label="icon">
                        <el-input
                                placeholder="路由显示时前边的图标"
                                v-model="newMenu.icon"
                        />
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
                                @click="soreMenu"
                        >
                            确定
                        </el-button>
                    </span>
                </template>
            </el-dialog>


            <el-dialog
                    v-model="deleteMenuDialog"
                    title="删除路由"
                    width="30%"
            >
                <span style="color: red">确定要删除吗? (危险操作)</span>
                <template #footer>
                  <span class="dialog-footer">
                    <el-button
                            type="primary"
                            @click="cancelDeleteMenu"
                    >
                        取消
                    </el-button>
                    <el-button
                            type="danger"
                            @click="soreDeleteMenu"
                    >
                      确定
                    </el-button>
                  </span>
                </template>
            </el-dialog>


            <el-button
                    type="primary"
                    icon="Plus"
                    style="margin: 20px 0"
                    @click="addRouter"
            >
                添加路由
            </el-button>
            <el-table border :data="data">
                <el-table-column label="id" prop="id" width="50"/>
                <el-table-column label="r_id" prop="r_id" width="50"/>
                <el-table-column label="menu_level" prop="menu_level"/>
                <el-table-column label="parent_id" prop="parent_id"/>
                <el-table-column label="path" prop="path"/>
                <el-table-column label="name" prop="name"/>
                <el-table-column label="hidden" prop="hidden" width="50"/>
                <el-table-column label="component" prop="component"/>
                <el-table-column label="sort" prop="sort" width="50"/>
                <el-table-column label="title" prop="title"/>
                <el-table-column label="icon" prop="icon"/>
                <el-table-column label="路由操作" width="100">
                    <template #default="{row,$index}">
                        <el-button
                                size="small"
                                type="info"
                                round
                                style="margin-bottom: 5px"
                                @click="updateMenu(row)"
                        >
                            修改
                        </el-button>
                        <el-button
                                size="small"
                                type="danger"
                                round
                                @click="deleteMenu"
                        >
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-pagination
                    v-model:current-page="currentPage"
                    v-model:page-size="pageSize"
                    :page-sizes="[5, 6, 7, 8]"
                    :small="true"
                    layout="prev, pager, next, jumper,->,total, sizes,"
                    :total="total"
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    style="margin: 20px 0"
            />
        </el-card>
    </div>
</template>

<script setup lang="ts">
import {onMounted, reactive, ref} from "vue";
import {GetMenusLimit} from "@/api/menu";
import {MenuType} from "@/type/userApi.ts";

let currentPage = ref<number>(1);
let pageSize = ref<number>(5);
let total = ref<number>(1);
let handleSizeChange = (v: number) => {
    pageSize.value = v;
};
let handleCurrentChange = (v: number) => {
    currentPage.value = v;
    getMenuLimit();
};
let data = ref<MenuType[]>([]);

onMounted(() => {
    getMenuLimit();
});

function getMenuLimit() {
    GetMenusLimit({page: currentPage.value, pageSize: pageSize.value}).then((res) => {
        data.value = res.data.menus;
        total.value = res.data.total;
    });
}


let dialogFormVisible = ref<boolean>(false);

let dialogTitle = ref<string>("");

let newMenu = ref<MenuType>({
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

const addRouter = () => {
    dialogTitle.value = "添加路由";
    dialogFormVisible.value = true;
};

const updateMenu = (menu: MenuType) => {
    newMenu.value = menu;
    dialogTitle.value = "修改路由,谨慎修改!!!";
    dialogFormVisible.value = true;
};

const soreMenu = () => {
    dialogFormVisible.value = false;
};


let deleteMenuDialog = ref<boolean>(false);

const deleteMenu = () => {
    deleteMenuDialog.value = true;
};

const cancelDeleteMenu = () => {
    deleteMenuDialog.value = false;
};

const soreDeleteMenu = () => {
    deleteMenuDialog.value = false;
    // TODO 删除路由

};

</script>

<style scoped lang="scss">
.main {
  width: inherit;
  height: 100vh;
  display: flex;
}

.box-card {
  flex-grow: 1;
}
</style>