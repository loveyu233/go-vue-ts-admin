<template>
    <template v-for="m in store.menus" :key="m.path">
        <Tree :menu="m" v-if="m.hidden===0">
        </Tree>
    </template>
</template>

<script setup lang="ts">
import {rootStore} from "@/store/index.ts";
import {ref, watch} from "vue";
import {layoutStore} from "@/store/layout.ts";
import Tree from "./tree/index.vue";

let store = rootStore();

let storel = layoutStore();

let color = ref("black");

watch(() => storel.dark, (value, oldValue, onCleanup) => {
    if (value) {
        color.value = "white";
    } else {
        color.value = "black";
    }
});


</script>

<style scoped lang="scss">
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