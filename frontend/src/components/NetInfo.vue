<template>
    <div class="netinfo">
        <el-row><el-col><h2>网络</h2></el-col></el-row>
        <el-row v-for="ifce in ifces" :key="ifce.name">
            <el-col :span="4">
                <el-tag>设备名</el-tag>
                <span>{{ ifce.name }}</span>
            </el-col>
            <el-col :span="6">
                <el-tag>总发送</el-tag>
                <span>{{ ifce.send_byte | humanByte }}</span>
            </el-col>
            <el-col :span="6">
                <el-tag>总接收</el-tag>
                <span>{{ ifce.recv_byte | humanByte }}</span>
            </el-col>
            <el-col :span="4">
                <el-tag>发送</el-tag>
                <span>{{ curNetSend(ifce.send_byte) | humanByte }}/s</span>
            </el-col>
            <el-col :span="4">
                <el-tag>接收</el-tag>
                <span>{{ curNetRecv(ifce.recv_byte) | humanByte  }}/s</span>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanByte, humanPerc} from "@/common/filters"
    export default {
        name: "netinfo",
        props: ['ifces'],
        filters: {
            humanByte: function(size) {
                return humanByte(size)
            },
            humanPerc: function(float, num) {
                return humanPerc(float, num)
            }
        },
        methods: {
            curNetSend: function(curTotalSend) {
                var lastTotalSend = sessionStorage.getItem("lastTotalSend")
                if (!lastTotalSend) {
                    lastTotalSend = curTotalSend
                }
                sessionStorage.setItem("lastTotalSend", curTotalSend)
                return curTotalSend - lastTotalSend
            },
            curNetRecv: function(curTotalRecv) {
                var lastTotalRecv = sessionStorage.getItem("lastTotalRecv")
                if (!lastTotalRecv) {
                    lastTotalRecv = curTotalRecv
                }
                sessionStorage.setItem("lastTotalRecv", curTotalRecv)
                return curTotalRecv - lastTotalRecv
            }
        },
    }
</script>
<style>
.el-row {
    margin-top: -1px;
    border-style: solid;
    border-width: 1px;
}
</style>