<template>
    <div class="cpuinfo">
        <el-row><el-col><h2>CPU</h2></el-col></el-row>
        <el-row v-for="core in info.cores" :key="core.core_id">
            <el-col :span="2">
                <el-tag>核心id</el-tag>
                <span>{{ core.core_id }}</span>
            </el-col>
            <el-col :span="8">
                <el-tag>CPU型号</el-tag>
                <span>{{ core.model }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>CPU频率</el-tag>
                <span>{{ core.mhz }}</span>
            </el-col>
            <el-col :span="3">
                <el-tag>AES-NI</el-tag>
                <span>{{ core.flags | aes }}</span>
            </el-col>
            <el-col :span="3">
                <el-tag>VM-x/AMD-V</el-tag>
                <span>{{ core.flags | vm }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>使用率</el-tag>
                <span>{{ cpu_percent[parseInt(core.core_id)] | humanPerc(2) }} %</span>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanByte, humanSec, humanPerc} from "@/common/filters"
    export default {
        name: "cpuinfo",
        props: ['info', 'cpu_percent'],
        computed: {
            core_count: function() {
                return this.info.cores.length
            },
        },
        filters: {
            aes: function(flags) {
                return flags.includes("aes") ? "支持" : "不支持"
            },
            vm: function(flags) {
                return flags.includes("vmx") || flags.includes("svm") ? "支持" : "不支持"
            },
            humanByte: function(size) {
                return humanByte(size)
            },
            humanSec: function(sec) {
                return humanSec(sec)
            },
            humanPerc: function(float, num) {
                return humanPerc(float, num)
            }
        }
    }
</script>
<style>
.el-row {
    margin-top: -1px;
    border-style: solid;
    border-width: 1px;
}
</style>
