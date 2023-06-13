<template>
    <div class="bottom-left">
        <div class="bottom-left-top" ref="bottomLeftTop">
        </div>
        <div class="bottom-left-center" ref="bottomLeftCenter">
            adadasdas
        </div>
        <div class="bottom-left-bottom" ref="bottomLeftBottom">
            dadasda
        </div>
    </div>
</template>

<script setup lang="ts">
import * as echarts from "echarts";
import "echarts-liquidfill";
import {onMounted, ref} from "vue";
import {HistogramOption, PieChartOption, WaterBalloonDiagramOption} from "@/style/echarts.ts";

let bottomLeftTop = ref();
let bottomLeftCenter = ref();
let bottomLeftBottom = ref();

onMounted(() => {
    let waterBalloonDiagram = echarts.init(bottomLeftTop.value);
    waterBalloonDiagram.setOption(WaterBalloonDiagramOption);

    let histogram = echarts.init(bottomLeftCenter.value);
    histogram.setOption(HistogramOption);

    let pieChart = echarts.init(bottomLeftBottom.value);
    pieChart.setOption(PieChartOption);
});

</script>

<style scoped lang="scss">

.bottom-left {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;

  .bottom-left-top {
    flex: 1.2;
  }

  .bottom-left-center {
    flex: 2;
    padding: 20px 0;
  }

  .bottom-left-bottom {
    flex: 1.2;
  }
}
</style>