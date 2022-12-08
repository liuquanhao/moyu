<template>
    <div class="meminfo">
        <el-row><el-col><h2>内存</h2></el-col></el-row>
        <el-row>
            <el-col :span="4">
                <el-tag>内存</el-tag>
                <span>{{ info.memory | humanByte }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>已使用</el-tag>
                <span>{{ status.memory_percent | humanPerc(2) }} %</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <el-tag>swap</el-tag>
                <span>{{ info.swap | humanByte }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>已使用</el-tag>
                <span>{{ status.swap_percent | humanPerc(2) }} %</span>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanByte, humanPerc} from "@/common/filters"
    export default {
        name: "meminfo",
        props: ['info', 'status'],
        filters: {
            humanByte: function(size) {
                return humanByte(size)
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