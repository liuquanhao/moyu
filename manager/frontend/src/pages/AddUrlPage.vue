<template>
    <page-layout>
        <div class="nes-container with-title">
            <p class="title">Add Page Url</p>
            <el-row><el-col><div class="nes-field is-inline"><label for="title-ipt">Title:</label><input type="text" class="nes-input" :class="{'is-error': titleFail}" v-model="title" placeholder="MoyuPage" /></div></el-col></el-row>
            <el-row><el-col><div class="nes-field is-inline"><label>Url:</label><input type="text" id="page-url-ipt" class="nes-input" :class="{'is-error': pageUrlFail, 'is-success': pageUrlSuccess}" v-model="pageUrl" placeholder="https://moyu.linux.plus/" /></div></el-col></el-row>
            <el-row>
                <el-col :span="6" :offset="6">
                    <button type="button" class="nes-btn test-btn" :class="[isClickTest ? 'is-disabled': 'is-warning']" @click="urlTest">Test</button>
                </el-col>
                <el-col :span="6" :offset="4">
                    <button type="button" class="nes-btn add-btn" :class="[isClickAdd ? 'is-disabled': 'is-primary']" @click="urlAdd">Add</button>
                </el-col>
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
                title: "",
                pageUrl: "",
                titleFail: false,
                pageUrlFail: false,
                pageUrlSuccess: false,
                isClickTest: false,
                isClickAdd: false,
            }
        },
        methods: {
            urlTest: function() {
                this.titleFail = false
                this.pageUrlFail = false
                if (!this.title) {
                    this.titleFail = true
                    return
                }
                if (!this.pageUrl) {
                    this.pageUrlFail = true
                    return
                }
                this.isClickTest = true
                this.axios.get(this.pageUrl, {withCredentials: false}).then(res => {
                    this.isClickTest = false
                    this.pageUrlSuccess = true
                }).catch(err => {
                    this.isClickTest = false
                    this.pageUrlFail = true
                });
            },
            urlAdd: function() {
                this.titleFail = false
                this.pageUrlFail = false
                if (!this.title) {
                    this.titleFail = true
                    return
                }
                if (!this.pageUrl) {
                    this.pageUrlFail = true
                    return
                }
                this.isClickAdd = true
                this.axios.post('./api/page_url', {
                    'title': this.title,
                    'page_url': this.pageUrl,
                }, {withCredentials: true}).then(res => {
                    this.$router.push('/')
                }).catch(err => {
                    this.isClickAdd = false
                    this.titleFail = true
                    this.pageUrlFail = true
                });
            }
        }
    }
</script>

<style scoped>
.nes-container {
    margin: 200px auto;
    width: 80%;
}
.nes-input {
    border-image-repeat: stretch;
}
.el-row {
    margin-bottom: 10px;
}
</style>