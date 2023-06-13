<template>
    <div>
        <el-carousel :interval="4000" type="card" height="200px">
            <el-carousel-item v-for="item in images" :key="item.id">
                <img :src="item.url" style="width: 100%;height: 100%"/>
            </el-carousel-item>
        </el-carousel>
        <el-card
                style="margin-top: 40px"
                shadow="hover"
        >
            <img :src="userStore.info.avatar"/>
            <span
                    style="font-size: 40px;margin-left: 40px"
            >
                {{ GetTime() }}好呀!
                <span
                        style="color: #00f2fe"
                >
                    {{ userStore.info.username }}
                </span>
            </span>
        </el-card>

    </div>
</template>

<script setup lang="ts">
import {rootStore} from "@/store";
import {GetTime} from "@/utils/time.ts";
import {onMounted, ref} from "vue";
import {GetCarouselAll} from "@/api/carousel";
import {CarouselType} from "@/type/userApi.ts";

let userStore = rootStore();
let images = ref<CarouselType[]>();

onMounted(() => {
    GetCarouselAll().then(res => {
        images.value = res.data;
        console.log(images.value);
    });
});
</script>

<style scoped>

</style>