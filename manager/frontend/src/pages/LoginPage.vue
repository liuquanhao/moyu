<template>
    <page-layout>
        <div class="nes-container with-title is-centered">
            <p class="title">Login</p>
            <el-row><el-col><div class="nes-field is-inline"><label for="user-ipt">User:</label><input type="text" id="user-ipt" class="nes-input" :class="{'is-error': userFail}" v-model="user" /></div></el-col></el-row>
            <el-row><el-col><div class="nes-field is-inline"><label for="token-ipt">Token:</label><input type="password" id="token-ipt" class="nes-input" :class="{'is-error': tokenFail}" v-model="token" /></div></el-col></el-row>
            <el-row>
                <el-col :span="6" :offset="8">
                    <button type="button" class="nes-btn login-btn" :class="[isClick ? 'is-disabled': 'is-primary']" @click="login">Login</button>
                </el-col>
                <el-col :span="8" class="forget-token-col">
                    <a href="https://github.com/liuquanhao/moyu" target="_blank" rel="noopener">
                        <span class="nes-text is-disabled forget-token-txt">Forget Token?</span>
                    </a>
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
                user: "",
                token: "",
                userFail: false,
                tokenFail: false,
                isClick: false,
            }
        },
        methods: {
            login: function() {
                this.userFail = false
                this.tokenFail = false
                if (!this.user) {
                    this.userFail = true
                    return
                }
                if (!this.token) {
                    this.tokenFail = true
                    return
                }
                this.isClick = true
                this.axios.post('./api/login', {
                    user: this.user,
                    token: this.token
                }, {withCredentials: true}).then(res => {
                    this.isClick = false
                    this.$router.push('/')
                }).catch(err => {
                    this.isClick = false
                    this.userFail = true
                    this.tokenFail = true
                });
            }
        }
    }
</script>

<style scoped>
.nes-container {
    margin: 200px auto;
    width: 60%;
}
.el-row {
    margin-bottom: 10px;
}
.login-btn {
    margin-top: 10px;
}
.forget-token-col {
    margin-top: 30px;
}
.forget-token-txt {
    font-size: 5px;
    text-decoration: underline;
}
.nes-input {
    border-image-repeat: stretch;
}
</style>