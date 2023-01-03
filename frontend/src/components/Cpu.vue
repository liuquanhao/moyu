<template>
    <div>
        <el-row v-for="(percent, index) in cpu_percent" :key="index">
            <el-col>
                <progress class="nes-progress is-success" :value="percent" max="100"></progress>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    export default {
        name: "cpu",
        props: ['cpu_percent'],
    }
</script>
<style>
</style>
