<template>
    <div>
        <el-row>
            <el-col>
                <span>Memory | {{ info.memory | humanByte }}:</span>
            </el-col>
            <el-col>
                <progress class="nes-progress is-primary" :value="status.memory_percent" max="100"></progress>
            </el-col>
        </el-row>
        <el-row>
            <el-col>
                <span>Swap | {{ info.swap | humanByte }}:</span>
            </el-col>
            <el-col>
                <progress class="nes-progress is-primary" :value="status.swap_percent" max="100"></progress>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import {humanByte, humanPerc} from "@/common/filters"
    export default {
        name: "mee",
        props: ['info', 'status'],
        filters: {
            humanByte: function(size) {
                return humanByte(size)
            },
            humanPerc: function(float, num) {
                return humanPerc(float, num)
            }
        }
    }
</script>