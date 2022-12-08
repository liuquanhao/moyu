<template>
    <div class="diskinfo">
        <el-row><el-col><h2>磁盘</h2></el-col></el-row>
        <el-row v-for="partition in partitions" :key="partition.device">
            <el-col :span="6">
                <el-tag>分区</el-tag>
                <span>{{ partition.device }}</span>
            </el-col>
            <el-col :span="8">
                <el-tag>挂载点</el-tag>
                <span>{{ partition.mount_point }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>大小</el-tag>
                <span>{{ partition.size | humanByte }}</span>
            </el-col>
            <el-col :span="6">
                <el-tag>已使用</el-tag>
                <span>{{ partition.used_percent | humanPerc(2) }} %</span>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanByte, humanPerc} from "@/common/filters"
    export default {
        name: "meminfo",
        props: ['partitions'],
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