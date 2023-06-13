<template>
    <div class="main">
        <div class="screen" ref="curren">
            <div class="top">
                <Index/>
            </div>
            <div class="bottom">
                <div class="left">
                    <LeftIndex/>
                </div>
                <div class="center">
                    <CenterIndex/>
                </div>
                <div class="right">
                    <RightIndex/>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">

import {onMounted, ref} from "vue";
import Index from "./top/index.vue";

import LeftIndex from "./bottom/letf/index.vue";
import CenterIndex from "./bottom/center/index.vue";
import RightIndex from "./bottom/right/index.vue";

let curren = ref();

function getScale(w = 1920, h = 1080) {
    const ww = window.innerWidth / w;
    const wh = window.innerHeight / h;
    return ww < wh ? ww : wh;
}

onMounted(() => {
    document.documentElement.requestFullscreen();
    curren.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`;
});
//监听视口变化
window.onresize = () => {
    curren.value.style.transform = `scale(${getScale()}) translate(-50%,-50%)`;
};
</script>

<style scoped lang="scss">
.main {
  width: 100vw;
  height: 100vh;
  //background: url("./images/bg.png") no-repeat;
  //background-size: cover;
  //background-image: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%);

  .screen {
    position: fixed;
    width: 1920px;
    height: 1080px;
    left: 50%;
    top: 50%;
    transform-origin: top left;
    //background-color: #75e5e5;
    display: flex;
    flex-direction: column;

    .top {
      width: 100%;
      flex: 1;
      //background-color: #7e8c8d;
    }

    .bottom {
      flex: 9;
      //background-color: #656b06;
      display: flex;

      .left {
        flex: 1;
        //background-color: #75e5e5;
      }

      .center {
        flex: 2;
        //background-color: white;
      }

      .right {
        flex: 1;
        //background-color: #20E2D7;
      }
    }
  }
}
</style>