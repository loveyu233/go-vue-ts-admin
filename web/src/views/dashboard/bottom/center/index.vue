<template>
    <div class="bottom-center">
        <div class="bottom-center-top" ref="bottomCenterTop"></div>
        <div class="bottom-center-bottom" ref="bottomCenterBottom"></div>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import * as echarts from "echarts";
import chinaMapJson from "./json/chinaMap.json";
import {ChinaMapOption, LineChartOption} from "@/style/echarts.ts";

let bottomCenterTop = ref();
let bottomCenterBottom = ref();

echarts.registerMap("china", chinaMapJson);

onMounted(() => {
    let mapEcharts = echarts.init(bottomCenterTop.value);
    mapEcharts.setOption(ChinaMapOption);

    let LineChart = echarts.init(bottomCenterBottom.value);
    LineChart.setOption(LineChartOption);
});

</script>

<style scoped lang="scss">
.bottom-center {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .bottom-center-top {
    flex: 3;
  }

  .bottom-center-bottom {
    flex: 1;
  }
}
</style>