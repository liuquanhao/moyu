<template>
    <div>
        <el-row v-for="partition in partitions" :key="partition.device">
            <el-col>
                <span>{{ partition.device }}({{ partition.size | humanByte }}) | Used({{ usedDisk(partition.size, partition.used_percent) | humanByte }}) | Mounted({{ partition.mount_point }}):</span>
            </el-col>
            <el-col>
                <progress class="nes-progress is-primary" :value="partition.used_percent" max="100"></progress>
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
            },
        },
        methods: {
            usedDisk(diskSize, percent) {
                return diskSize * percent / 100;
            },
        }
    }
</script>