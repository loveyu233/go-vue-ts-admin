<template>
    <el-menu-item
            v-if="!menu.children"
            :index="menu.path"
    >
        <el-icon>
            <component :is="menu.icon"></component>
        </el-icon>
        <span>{{ menu.title }}</span>
    </el-menu-item>
    <el-sub-menu
            v-else
            :index="menu.path"
    >
        <template #title>
            <el-icon>
                <component :is="menu.icon"></component>
            </el-icon>
            <span>{{ menu.title }}</span>
        </template>

        <template
                v-for="m in menu.children"
                :key="m.path"
        >
            <Index :menu="m"/>
        </template>
    </el-sub-menu>
</template>

<script lang="ts" setup>
import {MenuType} from "@/type/userApi.ts";
import {ref, watch} from "vue";
import {layoutStore} from "@/store/layout.ts";

let storel = layoutStore();

let color = ref("black");

watch(() => storel.dark, (value, oldValue, onCleanup) => {
    if (value) {
        color.value = "white";
    } else {
        color.value = "black";
    }
});


defineProps<{
    menu: MenuType
}>();
</script>

<style lang="scss" scoped>
.el-menu-item, :deep(.el-sub-menu__title span), :deep(.el-icon) {
  color: v-bind(color);
}

.el-menu-item:hover {
  background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
}

:deep(.el-sub-menu__title:hover) {
  background: linear-gradient(to right, #4facfe 0%, #00f2fe 100%);
}
</style>

