<template>
    <div>
        <el-row>
            <el-col :span="4">
                <span>Hostname:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.hostname }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>OS:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.os }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>CPU:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.cpu }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>AES-NI:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.aes_ni }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>VM-x/AMD-V:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.vm_x_amd_v }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>Virt:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.virt }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>Uptime:</span>
            </el-col>
            <el-col :span="20">
                <span>{{ info.uptime | humanSec }}</span>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>Load:</span>
            </el-col>
            <el-col :span="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">1min</span>
                    <span class="is-warning">{{ info.load_avg.load1 }}</span>
                </div>
            </el-col>
            <el-col :span="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">5min</span>
                    <span class="is-warning">{{ info.load_avg.load5 }}</span>
                </div>
            </el-col>
            <el-col :span="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">15min</span>
                    <span class="is-warning">{{ info.load_avg.load15 }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="4">
                <span>Net(u|d):</span>
            </el-col>
            <el-col :span="4" v-for="ifce in info.ifces" :key="ifce.name">
                <div class="nes-badge is-splited net-badge">
                    <span class="is-success">{{ curNetSend(ifce.send_byte) | humanByte }}/s</span>
                    <span class="is-warning">{{ curNetRecv(ifce.recv_byte) | humanByte  }}/s</span>
                </div>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanSec, humanByte} from "@/common/filters"
    export default {
        name: "host",
        props: ['info'],
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
            humanSec: function(sec) {
                return humanSec(sec)
            },
            humanByte: function(size) {
                return humanByte(size)
            }
        },
    }
</script>
<style>
.el-row {
    margin-top: 5px;
}
.net-badge {
    width: 330px;
}
</style>
