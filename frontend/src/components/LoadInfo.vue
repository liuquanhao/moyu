<template>
    <div class="loadinfo">
        <el-row><el-col><h2>负载</h2></el-col></el-row>
        <el-row>
            <el-col :span="4">
                <el-tag>1分钟</el-tag>
                <span>{{ status.load1 }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>5分钟</el-tag>
                <span>{{ status.load5 }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>15分钟</el-tag>
                <span>{{ status.load15 }}</span>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanPerc} from "@/common/filters"
    export default {
        name: "loadinfo",
        props: ['status'],
    }
</script>
<style>
.el-row {
    margin-top: -1px;
    border-style: solid;
    border-width: 1px;
}
</style>
