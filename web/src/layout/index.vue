<template>
    <div class="main">
        <div class="left" :class="{fold:store.fold}">
            <Index></Index>
            <el-menu
                    style="border:0!important;"
                    :collapse="store.fold"
                    :router="true"
                    background-color="linear-gradient(120deg, #84fab0 0%, #8fd3f4 100%)"
                    text-color="black"
                    :default-active="router.currentRoute.value.path"
            >
                <Menu></Menu>
            </el-menu>
        </div>
        <div class="right" :class="{fold:!store.fold}">
            <div class="right_top">
                <Tabbar style="width: 100%;height: 100%"></Tabbar>
            </div>
            <div class="right_bottom">
                <Main></Main>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Index from "@/layout/titleIcon/index.vue";
import Menu from "@/layout/menu/index.vue";
import Main from "@/layout/main/index.vue";
import Tabbar from "@/layout/tabbar/index.vue";
import {useRouter} from "vue-router";
import {layoutStore} from "@/store/layout.ts";
import {onMounted, ref} from "vue";
import {rootStore} from "@/store";

let router = useRouter();
let store = layoutStore();
let userStore = rootStore();

onMounted(() => {
    userStore.getUserInfo();
    userStore.getMenusIist();
});


</script>

<style scoped lang="scss">
.main {
  width: 100%;
  height: 100vh;
  //background-image: linear-gradient(-225deg, #E3FDF5 0%, #FFE6FA 100%); //background-color: #1296db;
  display: flex;

  .left {
    //flex: 1;
    //background-image: linear-gradient(-225deg, #20E2D7 0%, #F9FEA5 100%);

    //background-image: v-bind(color);
    width: 260px;
    //background-color: #7e8c8d;
    //min-width: $layoutLeftMinWidthSize;
    overflow: auto;
    transition: all 0.5s linear;

    &.fold {
      width: 70px;
    }
  }

  .right {
    //background-color: aquamarine;
    display: flex;
    flex-direction: column;
    flex-grow: 1;

    // 必须添加,要不然右侧窗口不会缩小
    &.fold {
      width: 400px;
    }

    .right_top {
      flex: 1;
      //background-color: red;
      min-height: $layoutRightTopMinHeightSize;
    }

    .right_bottom {
      flex: 9;
      //color: black;
      //background-color: black;
      min-width: $layoutRightBottomMinWidthSize;
      overflow: auto;
      padding: 30px;
    }
  }
}
</style>