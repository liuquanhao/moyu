<template>
    <main-layout>
        <host-info :info="hostInfo"></host-info>
    </main-layout>
  </template>
  
  <script>
    import MainLayout from '../layouts/Main.vue'
    import HostInfo from '../components/HostInfo.vue'
    export default {
        components: {
            MainLayout,
            HostInfo,
        },
        data: function() {
            return {
                hostname: "",
                uptime: 0,
                arch: "",
                distribution: "",
                kernel: "",
                core_count: 0,
                model: "",
                mhz: 0.0,
                flags: "",
                memory: 0,
                swap: 0,
                virtual_platform: "",
                partitions: null,
                ifces: null,
                timestamp: 0,
            }
        },
        computed: {
            hostInfo: function() {
                return {
                    hostname: this.hostname,
                    uptime: this.uptime,
                    arch: this.arch,
                    distribution: this.distribution,
                    kernel: this.kernel,
                    core_count: this.core_count,
                    model: this.model,
                    mhz: this.mhz,
                    flags: this.flags,
                    memory: this.memory,
                    swap: this.swap,
                    virtual_platform: this.virtual_platform,
                    partitions: this.partitions,
                    ifces: this.ifces,
                }
            }
        },
        mounted() {
            this.initData()
        },
        methods: {
            initData() {
                this.axios.get("http://134.195.88.86:8081/sys_info").then(
                    res => {
                        this.hostname = res.data.hostname
                        this.uptime = res.data.uptime
                        this.arch = res.data.arch
                        this.distribution = res.data.distribution
                        this.kernel = res.data.kernel
                        this.memory = res.data.memory
                        this.virtual_platform = res.data.virtual_platform
                        this.core_count = res.data.cpu_info.count
                        this.model = res.data.cpu_info.cores[0].model
                        this.mhz = res.data.cpu_info.cores[0].mhz
                        this.flags = res.data.cpu_info.cores[0].flags
                        this.partitions = res.data.disk_info.partitions
                        this.ifces = res.data.network_info.ifces
                        this.timestamp = res.data.timestamp
                    }
                ).catch(res => {
                    console.log(res)
                })
            }
        }
    }
  </script>