<template>
    <router-view v-slot="{Component}">
        <transition
                enter-active-class="animate__animated animate__backInRight"
                :duration="{enter:1000,leave:0}"
        >
            <component
                    :is="Component"
                    v-if="flag"
            >
            </component>
        </transition>
    </router-view>
</template>

<script setup lang="ts">

import {layoutStore} from "@/store/layout.ts";
import {nextTick, ref, watch} from "vue";

let store = layoutStore();

// 通过flag控制v-if,v-if是可以销毁组件的,从而达到强制刷新的效果
let flag = ref(true);

watch(() => store.refresh, () => {
    flag.value = false;
    // 页面重新渲染完后再加载出来
    nextTick(() => {
        flag.value = true;
    });
});

</script>

<style scoped>

</style>