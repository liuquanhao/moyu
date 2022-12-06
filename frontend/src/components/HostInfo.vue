<template>
    <div class="info">
        <div class="sysinfo">
            <el-row><el-col><h2>系统信息</h2></el-col></el-row>
            <el-row>
                <el-col :span="4">
                    <el-tag>hostname</el-tag>
                    <span>{{ info.hostname }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>uptime</el-tag>
                    <span>{{ info.uptime | humanSec }}</span>
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
            <el-row>
                <el-col :span="8">
                    <el-tag>CPU型号</el-tag>
                    <span>{{ info.model }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>CPU频率</el-tag>
                    <span>{{ info.mhz }}</span>
                </el-col>
                <el-col :span="3">
                    <el-tag>核心数</el-tag>
                    <span>{{ info.core_count }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>AES-NI</el-tag>
                    <span>{{ aes }}</span>
                </el-col>
                <el-col :span="5">
                    <el-tag>VM-x/AMD-V</el-tag>
                    <span>{{ vm }}</span>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="4">
                    <el-tag>内存</el-tag>
                    <span>{{ info.memory | humanByte }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>swap</el-tag>
                    <span>{{ info.swap | humanByte }}</span>
                </el-col>
            </el-row>
        </div>
        <div class="diskinfo">
            <el-row><el-col><h2>磁盘信息</h2></el-col></el-row>
            <el-row v-for="partition in info.partitions" :key="partition.device">
                <el-col :span="6">
                    <el-tag>分区</el-tag>
                    <span>{{ partition.device }}</span>
                </el-col>
                <el-col :span="10">
                    <el-tag>挂载点</el-tag>
                    <span>{{ partition.mount_point }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>大小</el-tag>
                    <span>{{ partition.size | humanByte }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>已使用</el-tag>
                    <span>{{ partition.used_percent | humanPerc(2) }}</span>
                </el-col>
            </el-row>
        </div>  
        <div class="netinfo">
            <el-row><el-col><h2>网络信息</h2></el-col></el-row>
            <el-row v-for="ifce in info.ifces" :key="ifce.name">
                <el-col :span="4">
                    <el-tag>设备名</el-tag>
                    <span>{{ ifce.name }}</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>发送</el-tag>
                    <span>{{ curNetSend(ifce.send_byte) | humanByte }}/s</span>
                </el-col>
                <el-col :span="4">
                    <el-tag>接收</el-tag>
                    <span>{{ curNetRecv(ifce.recv_byte) | humanByte  }}/s</span>
                </el-col>
                <el-col :span="6">
                    <el-tag>总发送</el-tag>
                    <span>{{ ifce.send_byte | humanByte }}</span>
                </el-col>
                <el-col :span="6">
                    <el-tag>总接收</el-tag>
                    <span>{{ ifce.recv_byte | humanByte }}</span>
                </el-col>
            </el-row>
        </div> 
    </div>
</template>

<script>
    export default {
        name: "info",
        props: ['info'],
        computed: {
            vt: function (){
                return this.info.virtual_platform ? this.info.virtual_platform : '未知'
            },
            aes: function() {
                return this.info.flags.includes("aes") ? "支持" : "不支持"
            },
            vm: function() {
                return this.info.flags.includes("vmx") || this.info.flags.includes("svm") ? "支持" : "不支持"
            },
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
        filters: {
            humanByte: function(size) {
                if (size < 1024) {
                    return size + " Byte"
                } else if (size < 1024 * 1024) {
                    return Math.trunc(size / 1024) + " KB"
                } else if (size < 1024 * 1024 * 1024) {
                    return Math.trunc(size / 1024 / 1024) + " MB"
                } else {
                    return (size / 1024 / 1024 / 1024).toFixed(2) + " GB"
                }
            },
            humanSec: function(sec) {
                if (sec < 60) {
                    return Math.trunc(sec) + " 秒";
                } else if (sec < 3600) {
                    var sec = sec % 60
                    var min = Math.trunc(sec / 60);
                    return min + " 分 " + sec + " 秒"
                } else if (sec < 86400) {
                    var min = Math.trunc(sec % 3600 / 60);
                    var hour = Math.trunc(sec / 3600);
                    return hour + " 小时 " + min + " 分"
                } else {
                    var hour = Math.trunc(sec % 86400 / 3600);
                    var day = Math.trunc(sec / 86400);
                    return day + " 天 " + hour + " 小时"
                }
            },
            humanPerc: function(float, dec) {
                return float.toFixed(2) + " %"
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
