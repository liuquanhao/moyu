<template>
    <main-layout>
        <host-info :info="host_info" :uptime="uptime" />
        <cpu-info :info="cpu_info" :cpu_percent="cpu_percent" />
        <load-info :status="load_avg" />
        <mem-info :info="memory_info" :status="memory_status" />
        <net-info :ifces="ifces" />
        <disk-info :partitions="partitions" />
    </main-layout>
  </template>
  
  <script>
    import MainLayout from '../layouts/Main.vue'
    import HostInfo from '../components/HostInfo.vue'
    import CpuInfo from '../components/CpuInfo.vue'
    import LoadInfo from '../components/LoadInfo.vue'
    import MemInfo from '../components/MemInfo.vue'
    import NetInfo from '../components/NetInfo.vue'
    import DiskInfo from '../components/DiskInfo.vue'
    export default {
        components: {
            MainLayout,
            HostInfo,
            CpuInfo,
            LoadInfo,
            MemInfo,
            NetInfo,
            DiskInfo,
        },
        data: function() {
            return {
                ws: null,
                host_info: {},
                cpu_info: {},
                cpu_percent: {},
                load_avg: {},
                memory_info: {},
                memory_status: {},
                ifces: {},
                partitions: {},
                uptime: {},
                timestamp: 0,
            }
        },
        created() {
            this.initData()
            this.initWs()
        },
        destroyed() {
            this.ws.close()
        },
        methods: {
            initData() {
                this.axios.get("/sys_info").then(
                    res => {
                        this.host_info = res.data.host_info
                        this.cpu_info = res.data.cpu_info
                        this.load_avg = res.data.system_status.load_avg
                        this.memory_info = res.data.memory_info
                        this.ifces = res.data.network_info.ifces
                        this.partitions = res.data.disk_info.partitions
                        this.cpu_percent = res.data.system_status.cpu_percent
                        this.memory_status = res.data.system_status.memory_status
                        this.timestamp = res.data.timestamp
                    }
                ).catch(res => {
                    console.log(res)
                })
            },
            initWs() {
                var wsProtocol = window.location.protocol == "https:" ? "wss://" : "ws://";
                var wsPort = window.location.port == 80 ? "" : ":" + window.location.port;
                this.ws = new WebSocket(wsProtocol + window.location.hostname + wsPort + "/ws/sys_status")
                this.ws.onopen = this.wsOnOpen
                this.ws.onerror = this.wsOnError
                this.ws.onmessage = this.wsOnMessage
                this.ws.onclose = this.wsOnClose
            },
            wsOnOpen() {
                console.log("ws连接成功")
            },
            wsOnError() {
                console.log("ws连接错误")
                this.initWs()
            },
            wsOnMessage(e) {
                var systemStatus = JSON.parse(e.data)
                this.cpu_percent = systemStatus.cpu_percent
                this.load_avg = systemStatus.load_avg
                this.memory_status = systemStatus.memory_status
                this.ifces = systemStatus.network_info.ifces
                this.partitions = systemStatus.disk_info.partitions
                this.uptime = systemStatus.uptime
            },
            wsOnClose(e) {
                console.log("ws断开连接")
            }
        }
    }
  </script>