<template>
  <!-- <router-view /> -->
  <router-view v-if="isRouterAlive"></router-view>
</template>


<script>
export default {
  name: 'App',
   components: {},
  provide(){
    return {
      reload: this.reload
    }
  },
  data() {
    return {
      isRouterAlive: true,
    };
  },
  cread() {},
  methods: {
    reload(){
      this.isRouterAlive = false;
      this.$nextTick(function(){
        this.isRouterAlive = true;
      })
    }
  },
  mounted() {
  },
}
</script>

<style>
body{
  margin: 0 0;
  font-family: Arial;
}
</style>
