<template>
    <div class="bottom-right-box">
        <div class="bottom-right-box-1" ref="bottomRightBox1"></div>
        <div class="bottom-right-box-2" ref="bottomRightBox2"></div>
        <div class="bottom-right-box-3" ref="bottomRightBox3"></div>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import * as echarts from "echarts";
import {Box1Option, Box2Option, Box3Option} from "@/style/echarts.ts";

let bottomRightBox1 = ref();
let bottomRightBox2 = ref();
let bottomRightBox3 = ref();

onMounted(() => {
    let box1 = echarts.init(bottomRightBox1.value);
    box1.setOption(Box1Option);
    let box2 = echarts.init(bottomRightBox2.value);
    box2.setOption(Box2Option);
    let box3 = echarts.init(bottomRightBox3.value);
    box3.setOption(Box3Option);
});
</script>

<style scoped lang="scss">
.bottom-right-box {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .bottom-right-box-1 {
    flex: 1;
  }

  .bottom-right-box-2 {
    flex: 1;
    margin: 20px 0;
  }

  .bottom-right-box-3 {
    flex: 1;
  }
}
</style>