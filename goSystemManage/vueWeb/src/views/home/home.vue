<template>
  <div>
    <head-nav v-if="displayApp"></head-nav>
    <div class="left-fixed-right-auto" v-if="displayApp">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import HeadNav from '@/views/head-nav/head-nav'
export default {
  name: 'home',
  data () {
    return {
      displayApp: true
    }
  },
  components: {HeadNav}
}
</script>
<style>

</style>
