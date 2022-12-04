<template>
    <div>
        <el-descriptions title="CPU信息">
            <el-descriptions-item label="型号">{{ model }}</el-descriptions-item>
            <el-descriptions-item label="线程数">{{ threads }}</el-descriptions-item>
            <el-descriptions-item label="频率">{{ freq }}</el-descriptions-item>
            <el-descriptions-item label="缓存">{{ cache }}</el-descriptions-item>
            <el-descriptions-item label="指令集">{{ flags }}</el-descriptions-item>
        </el-descriptions>
    </div>
</template>

<script>
    export default {
        name: "cpuInfo",
        data() {
            return {
                model: "",
                threads: 0,
                freq: 0,
                cache: 0,
                flags: "",
            }
        },
        mounted() {
            this.initData()
        },
        methods: {
            initData() {
                this.axios.get("http://149.62.46.168/hardware_info").then(
                    res => {
                        console.log(res)
                        this.model = res.data.cpu.model
                        this.threads = res.data.cpu.threads
                        this.freq = res.data.cpu.speed
                        this.cache = res.data.cpu.cache
                        this.flags = ""
                    }
                ).catch(res => {
                    console.log(res)
                })
            }
        }
    }
</script>

<style lang="css">

</style>
