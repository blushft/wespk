<template>
  <div id="app">
    <sidebar :show="sidebar.opened && !sidebar.hidden"></sidebar>
    <router-view></router-view>
  </div>
</template>

<script>
import { Sidebar } from 'compoents/layout'
export default {
  components: {
    Sidebar
  },
  name: 'app',
  beforeMount () {
    const { body } = document
    const WIDTH = 768
    const RATIO = 3

    const handler = () => {
      if (!document.hidden) {
        let rect = body.getBoundingClientRect()
        let isMobile = rect.width - RATIO < WIDTH
        this.toggleDevice(isMobile ? 'mobile' : 'other')
        this.toggleSidebar({
          opened: !isMobile
        })
      }
    }

    document.addEventListener('visibilitychange', handler)
    window.addEventListener('DOMContentLoaded', handler)
    window.addEventListener('resize', handler)
  }
}
</script>

<style>
#app {
  margin-top: 60px;
}
</style>
