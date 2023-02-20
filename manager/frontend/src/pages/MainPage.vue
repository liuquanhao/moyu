<template>
    <page-layout>
        <div class="pages">
            <el-row type="flex">
                <el-col :span="4"><span>Title</span></el-col>
                <el-col :span="3"><span>Uptime</span></el-col>
                <el-col :span="3"><span>CPU</span></el-col>
                <el-col :span="3"><span>Memory</span></el-col>
                <el-col :span="3"><span>Disk(/)</span></el-col>
                <el-col :span="5"><span>BW(u|d)</span></el-col>
                <el-col :span="2"><button type="button" class="nes-btn" :class="[logged ? 'is-primary' : 'is-disabled']" @click="jumpTo('/add_url')">Add</button></el-col>
            </el-row>
            <el-row type="flex" v-for="page in pages" :key="page.id">
                <el-col :span="4" v-if="logged"><a :href="page.page_url" target="_blank" rel="noopener">{{ page.title }}</a></el-col>
                <el-col :span="4" v-else><span>{{ page.title }}</span></el-col>
                <el-col :span="3"><span v-if="page.page_data">{{ page.page_data.uptime | humanDur }}</span></el-col>
                <el-col :span="3"><span v-if="page.page_data" class="nes-text is-success">{{ page.page_data.cpu_percent | humanPerc }}%</span></el-col>
                <el-col :span="3"><span v-if="page.page_data" class="nes-text is-error">{{ page.page_data.memory_percent | humanPerc }}%</span></el-col>
                <el-col :span="3"><span v-if="page.page_data" class="nes-text is-primary">{{ page.page_data.disk_percent | humanPerc }}%</span></el-col>
                <el-col :span="5">
                    <div v-if="page.page_data" class="nes-badge is-splited net-badge">
                        <span class="is-dark">{{ page.page_data.net_send_rate | humanByte }}</span>
                        <span class="is-warning">{{ page.page_data.net_recv_rate | humanByte }}</span>
                    </div>
                </el-col>
                <el-col :span="2"><button type="button" class="nes-btn" :class="[logged ? 'is-error' : 'is-disabled']" @click="deletePage(page.id, $event)">Delete</button></el-col>
            </el-row>
        </div>
    </page-layout>
</template>

<script>
    import PageLayout from '../layouts/PageLayout.vue'
    export default {
        components: {
            PageLayout,
        },
        data: function() {
            return {
                pages: [],
            }
        },
        created: function() {
            this.initData()
            this.initWs()
        },
        beforeDestroy: function() {
            this.ws.close()
        },
        computed: {
            logged: function() {
                return this.$cookies.get('uid')
            },
        },
        methods: {
            jumpTo: function(url) {
                this.$router.push(url)
            },
            deletePage: function(id, event) {
                this.btnDisabledStatus(event.target, true, 'is-error', 'is-disabled')
                this.axios.delete('./api/page_url/' + id, {withCredentials: true})
                .then(res => {
                    this.btnDisabledStatus(event.target, false, 'is-disabled', 'is-error')
                    this.dropDataPage(id)
                    this.initWs()
                })
                .catch(err => {
                    this.btnDisabledStatus(event.target, false, 'is-disabled', 'is-warning')
                    console.log(err)
                })
            },
            dropDataPage(id) {
                this.pages = this.pages.filter(function(page) {
                    if (page.id == id) {
                        return false
                    }
                    return true
                })
            },
            btnDisabledStatus(element, disabled, removeClass, addClass) {
                element.disabled = disabled
                element.classList.remove(removeClass)
                element.classList.add(addClass)
            },
            initData() {
                this.axios.get('./api/page_url/', {withCredentials: true})
                .then(res => {
                    this.pages = res.data
                })
                .catch(err => {
                    console.log(err)
                })
            },
            initWs() {
                let wsProtocol = window.location.protocol == "https:" ? "wss://" : "ws://"
                let wsPort = window.location.port == "" ? "" : ":" + window.location.port
                this.ws = new WebSocket(wsProtocol + window.location.hostname + wsPort + window.location.pathname + "ws/page_data")
                this.ws.onopen = this.wsOnOpen
                this.ws.onerror = this.wsOnError
                this.ws.onmessage = this.wsOnMessage
                this.ws.onclose = this.wsOnClose
            },
            wsOnOpen() {
                console.log("ws connect success")
            },
            wsOnError() {
                console.log("ws connect fail")
                this.initWs()
            },
            wsOnMessage(e) {
                let pageInfo = JSON.parse(e.data)
                let idx = this.pages.findIndex(page => {
                    return page.id == pageInfo.page_id
                })
                if (idx < 0) {
                    return
                }
                let page = {}
                page.id = this.pages[idx].id
                page.title = this.pages[idx].title
                page.page_url = this.pages[idx].page_url
                pageInfo.page_data.net_send_rate = this.curNetSend(pageInfo.page_data.net_send, pageInfo.page_id, pageInfo.page_data.timestamp)
                pageInfo.page_data.net_recv_rate = this.curNetRecv(pageInfo.page_data.net_recv, pageInfo.page_id, pageInfo.page_data.timestamp)
                page.page_data = pageInfo.page_data
                this.$set(this.pages, idx, page)
            },
            wsOnClose(e) {
                console.log("ws close")
            },
            curNetSend: function(curSend, id, ts) {
                let lastSend = sessionStorage.getItem(id + "-lastSend")
                let lastSendTs = sessionStorage.getItem(id + "-lastSendTs")
                if (!lastSend) {
                    lastSend = curSend
                }
                sessionStorage.setItem(id + "-lastSend", curSend)
                sessionStorage.setItem(id + "-lastSendTs", ts)
                let dur = ts - lastSendTs
                if (dur < 1) {
                    return 0
                }
                return (curSend - lastSend)/(ts - lastSendTs)
            },
            curNetRecv: function(curRecv, id, ts) {
                let lastRecv = sessionStorage.getItem(id + "-lastRecv")
                let lastRecvTs = sessionStorage.getItem(id + "-lastRecvTs")
                if (!lastRecv) {
                    lastRecv = curRecv
                }
                sessionStorage.setItem(id + "-lastRecv", curRecv)
                sessionStorage.setItem(id + "-lastRecvTs", ts)
                let dur = ts - lastRecvTs
                if (dur < 1) {
                    return 0
                }
                return (curRecv - lastRecv)/(ts - lastRecvTs)
            },
        },
        filters: {
            humanByte: function (size, num) {
                if (size < 1024 * 1024) {
                    return Math.trunc(size / 1024).toFixed(num) + "K"
                } else if (size < 1024 * 1024 * 1024) {
                    return (size / 1024 / 1024).toFixed(num) + "M"
                } else {
                    return (size / 1024 / 1024 / 1024).toFixed(num) + "G"
                }
            },
            humanDur: function (sec) {
                if (sec < 60) {
                    return Math.trunc(sec) + " s";
                } else if (sec < 3600) {
                    let min = Math.trunc(sec / 60);
                    return min + " m"
                } else if (sec < 86400) {
                    let hour = Math.trunc(sec / 3600);
                    return hour + " h"
                } else {
                    let day = Math.trunc(sec / 86400);
                    return day + "d"
                }
            },
            humanPerc: function (float) {
                if (!float) {
                    return 0
                }
                return float.toFixed(1)
            },
        }
    }
</script>

<style scoped>
.pages {
    margin: 20px 10px 0px 10px;
}
.el-row {
    border-bottom: 1px;
    border-bottom-style: solid;
    align-items: center;
}
.el-col {
    margin-top: 10px;
    margin-bottom: 10px;
    text-align: center;
}
.el-col button {
    display: block;
    margin: auto;
}
</style>