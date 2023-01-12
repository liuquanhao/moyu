<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  watch: {
        $route: {
            immediate: true,
            handler(to, from) {
                document.title = to.meta.title || 'MoYu';
            }
        },
    }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css?family=Press+Start+2P');
#app {
  background: #FFFFFA;
}
</style>
