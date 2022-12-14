<template>
    <div class="info">
        <div class="sysinfo">
            <el-row><el-col><h2>系统</h2></el-col></el-row>
            <el-row>
                <el-col :span="4">
                    <el-tag>hostname</el-tag>
                    <span>{{ info.hostname }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>uptime</el-tag>
                    <span>{{ ut | humanSec }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>系统版本</el-tag>
                    <span>{{ info.distribution }}</span>
                </el-col>
                <el-col :span="5">
                    <el-tag>内核</el-tag>
                    <span>{{ info.kernel }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>硬件架构</el-tag>
                    <span>{{ info.arch }}</span>
                </el-col>
                <el-col :span="3">
                    <el-tag>虚拟化</el-tag>
                    <span>{{ vt }}</span>
                </el-col>
            </el-row>
        </div>
    </div>
</template>

<script>
    import {humanSec} from "@/common/filters"
    export default {
        name: "hostinfo",
        props: ['info', 'uptime'],
        computed: {
            vt: function() {
                return this.info.virtual_platform ? this.info.virtual_platform : '未知'
            },
            ut: function() {
                return this.uptime ? this.uptime : this.info.uptime
            }
        },
        filters: {
            humanSec: function(sec) {
                return humanSec(sec)
            },
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
