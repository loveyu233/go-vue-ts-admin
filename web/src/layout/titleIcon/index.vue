<template>
    <div>
        <div class="top">
            <div class="left">
                <img :src="setting.icon">
            </div>
            <div class="right">
                <p>
                    {{ setting.title }}
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import setting from "@/settings.ts";

</script>

<style scoped lang="scss">
.top {
  display: flex;
  align-items: center;

  .left {
    img {
      border-radius: 50%;
      width: 50px;
      height: 50px;
    }

    padding: 10px;
  }

  .right {
    color: white;
    font-size: 30px;
  }
}
</style>