<template>
    <div>
        <el-row>
            <el-col>
                <span>Memory({{ info.memory | humanByte }}) | Used({{ usedMem | humanByte }}):</span>
            </el-col>
            <el-col>
                <progress class="nes-progress is-primary" :value="status.memory_percent" max="100"></progress>
            </el-col>
        </el-row>
        <el-row>
            <el-col>
                <span>Swap({{ info.swap | humanByte }}) | Used({{ usedSwap | humanByte }}):</span>
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
        },
        computed: {
            usedMem() {
                return this.info.memory * this.status.memory_percent / 100;
            },
            usedSwap() {
                return this.info.swap * this.status.swap_percent / 100;
            },
        }
    }
</script>