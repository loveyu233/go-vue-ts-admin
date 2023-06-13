<template>
    <div class="main-top">
        <div class="main-top-left">
            <router-link to="/home/index">
                <p class="main-top-left-title">主页</p>
            </router-link>
        </div>
        <div class="main-top-center">
            <p class="main-top-center-title">数据大屏展示</p>
        </div>
        <div class="main-top-right">
            <p class="main-top-right-title">{{ timeNow }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">

import {onMounted, ref} from "vue";

let timeNow = ref();

onMounted(() => {
    timeNow.value = format();
});

setInterval(() => {
    timeNow.value = format();
}, 1000);

const format = () => {
    //dataString是整数，否则要parseInt转换
    const time = new Date();
    let year = time.getFullYear();
    let month = time.getMonth() + 1;
    let day = time.getDate();
    let hour = time.getHours();
    let minute = time.getMinutes();
    let second = time.getSeconds();
    return year + "-" + (month < 10 ? "0" + month : month) + "-" + (day < 10 ? "0" + day : day) + " " + (hour < 10 ? "0" + hour : hour) + ":" + (minute < 10 ? "0" + minute : minute) + ":" + (second < 10 ? "0" + second : second);
};
</script>

<style scoped lang="scss">
.main-top {
  height: 100px;
  display: flex;

  .main-top-left {
    flex: 1;
    display: flex;
    justify-content: flex-end;
    align-items: center;

    .main-top-left-title {
      color: black;
      font-size: 30px;
      align-content: center;
      line-height: 100%;
    }
  }

  .main-top-center {
    flex: 2;
    display: flex;
    justify-content: center;
    align-items: center;

    .main-top-center-title {
      font-size: 50px;
      align-content: center;
      line-height: 100%;
    }
  }

  .main-top-right {
    flex: 1;
    display: flex;
    align-items: center;

    .main-top-right-title {
      font-size: 30px;
      align-content: center;
      line-height: 100%;
    }
  }
}
</style>