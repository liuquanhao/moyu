<template>
    <div>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>Hostname:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.hostname }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>OS:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.os }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>CPU:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.cpu }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>AES-NI:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.aes_ni }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>VM-x/AMD-V:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.vm_x_amd_v }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>Virt:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.virt }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>Uptime:</span></el-col>
            <el-col :xs="17" :sm="17" :md="19" :lg="20" :xl="20"><span>{{ info.uptime | humanSec }}</span></el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>Load:</span></el-col>
            <el-col :xs="6" :sm="6" :md="5" :lg="4" :xl="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">1min</span>
                    <span class="is-warning">{{ info.load_avg.load1 }}</span>
                </div>
            </el-col>
            <el-col :xs="6" :sm="6" :md="5" :lg="4" :xl="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">5min</span>
                    <span class="is-warning">{{ info.load_avg.load5 }}</span>
                </div>
            </el-col>
            <el-col :xs="5" :sm="5" :md="5" :lg="4" :xl="4">
                <div class="nes-badge is-splited">
                    <span class="is-dark">15min</span>
                    <span class="is-warning">{{ info.load_avg.load15 }}</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>BW(u|d):</span></el-col>
            <el-col :xs="9" :sm="10" :md="9" :lg="8" :xl="8" v-for="ifce in info.ifces" :key="ifce.name">
                <div class="nes-badge is-splited net-badge">
                    <span class="is-success">{{ curNetSend(ifce.send_byte) | humanByte }}/s</span>
                    <span class="is-warning">{{ curNetRecv(ifce.recv_byte) | humanByte  }}/s</span>
                </div>
            </el-col>
        </el-row>
        <el-row>
            <el-col :xs="7" :sm="7" :md="5" :lg="4" :xl="4"><span>Traffic:</span></el-col>
            <el-col :xs="9" :sm="10" :md="9" :lg="8" :xl="8" v-for="ifce in info.ifces" :key="ifce.name">
                <div class="nes-badge is-splited net-badge">
                    <span class="is-success">{{ ifce.send_byte | humanByte }}</span>
                    <span class="is-warning">{{ ifce.recv_byte | humanByte  }}</span>
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
<style scoped>
.el-row {
    margin-top: 5px;
}
.net-badge {
    width: 330px;
}
</style>
