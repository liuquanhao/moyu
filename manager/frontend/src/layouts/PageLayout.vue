<template>
    <el-container class="container">
        <el-header>
            <moyu-nav></moyu-nav>
        </el-header>
        <el-main>
            <slot></slot>
        </el-main>
        <el-footer>
            <moyu-footer></moyu-footer>
        </el-footer>
    </el-container>
</template>
  
<script>
    import MoyuNav from '../components/MoyuNav.vue'
    import MoyuFooter from '../components/MoyuFooter.vue'
    export default {
        components: {
            MoyuNav,
            MoyuFooter,
        },
        data: function() {
            return {}
        }
    }
</script>
  
<style scoped>
.container {
    min-width: 780px;
    max-width: 1260px;
    margin: 0px auto;
    justify-content: center;
}
.el-main {
    min-height: 800px;
}
</style>