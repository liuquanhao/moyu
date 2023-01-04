<template>
    <el-container class="content-container">
        <el-header height="80px">
            <div class="logo">
                <span>MoYu</span>
            </div>
            <div class="github-logo">
                <a href="https://github.com/liuquanhao/moyu" target="_blank" rel="noopener">
                    <i class="nes-icon github is-large"></i>
                </a>
            </div>
        </el-header>
        <el-main>
            <div class="nes-container with-title">
                <p class="title">Basic System Information</p>
                <host :info="host_data"/>
                <div class="nes-container with-title is-centered component-container">
                    <p class="title">CPU</p>
                    <cpu :cpu_percent="cpu_percent" />
                </div>
                <div class="nes-container with-title is-centered component-container">
                    <p class="title">Memory</p>
                    <mem :info="memory_info" :status="memory_status" />
                </div>
                <div class="nes-container with-title is-centered component-container">
                    <p class="title">Disk</p>
                    <disk :partitions="partitions" />
                </div>
            </div>
        </el-main>
        <el-footer>
            <p class="powered-by">Powered By: liuxu</p>
        </el-footer>
    </el-container>
</template>
  
<script>
    import MainLayout from '../layouts/Main.vue'
    import Host from '../components/Host.vue'
    import Cpu from '../components/Cpu.vue'
    import Mem from '../components/Mem.vue'
    import Disk from '../components/Disk.vue'
    export default {
        components: {
            MainLayout,
            Host,
            Cpu,
            Mem,
            Disk,
        },
        data: function() {
            return {
                ws: null,
                host_info: {},
                cpu_info: {
                    "count": 0,
                    "cores": [],
                },
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
        computed: {
            host_data: function() {
                let hostname = this.host_info.hostname
                let os = this.host_info.distribution + " (" + this.host_info.kernel + " " + this.host_info.arch + ")"
                let cpu = 'unknown'
                let aes_ni = false
                let vm_x_amd_v = false
                if (this.cpu_info.cores.length > 0) {
                    cpu = this.cpu_info.cores[0].model + " @ " + this.cpu_info.cores[0].mhz
                    aes_ni = this.cpu_info.cores[0].flags.includes("aes")
                    vm_x_amd_v = this.cpu_info.cores[0].flags.includes("vmx")
                }
                let virt = this.host_info.virtual_platform ? this.host_info.virtual_platform : 'unknown'
                let uptime = this.host_info.uptime
                let load_avg = this.load_avg
                let ifces = this.ifces
                return {
                    hostname,
                    os,
                    cpu,
                    aes_ni,
                    vm_x_amd_v,
                    virt,
                    uptime,
                    load_avg,
                    ifces,
                }
            },
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
                let wsProtocol = window.location.protocol == "https:" ? "wss://" : "ws://";
                let wsPort = window.location.port == "" ? "" : ":" + window.location.port;
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

<style scoped>
.logo {
    margin-right: auto;
    margin-left: 10px;
    font-size: 50px;
    height: 10px;
}
.github-logo {
    height: 10px;
    margin-right: 10px;
}
.content-container {
    min-width: 780px;
    max-width: 1260px;
    margin: 10px auto;
    justify-content: center;
}
.el-header {
    left: 0;
    right: 0;
    display: flex;
    flex-direction: row;
    border-bottom: 4px solid #D3D3D3; 
}
.nes-octocat {
    height: 30px;
}
.component-container {
    margin-top: 20px;
}
.el-footer {
    margin-top: 10px;
    margin-left: auto;
}
.powered-by {
    margin-right: 10px;
}
</style>