<template>
  <a-layout id="components-layout-demo-fixed-sider">
    <a-layout-sider class="sider-height">
      <div class="logo">
        vue3-admin-nest
      </div>
      <CommonMenu />
    </a-layout-sider>
    <a-layout>
      <CommonHeader />
      <a-layout-content class="layout-content">
        <CommonTabs />
        <div class="content">
          <router-view />
        </div>
      </a-layout-content>
      <a-layout-footer :style="{ textAlign: 'center' }">
        <CommonFooter />
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

import CommonFooter from '@/layout/footer/footer.vue'
import CommonMenu from '@/layout/menu/menu.vue'
import CommonHeader from '@/layout/header/header.vue'
import CommonTabs from '@/layout/tabs/index.vue'

export default defineComponent({
  name: 'LayoutHome',
  components: {
    CommonMenu,
    CommonFooter,
    CommonHeader,
    CommonTabs,
  },
  setup() {
    const selectedKeys: string[] = ['4']
    return {
      selectedKeys,
    }
  },
})
</script>

<style lang="less" scoped>
@import "layout.less";
</style>
