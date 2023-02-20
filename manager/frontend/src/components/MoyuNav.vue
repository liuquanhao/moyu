<template>
    <div class="moyu-nav">
        <div class="nav-left">
            <div class="logo">
                <a href="/#/"><span>MoYu</span></a>
            </div>
        </div>
        <div class="nav-right">
            <div class="btn" v-if="logged">
                <button type="button" class="nes-btn" @click="logout">Logout</button>
            </div>
            <div class="btn" v-else>
                <button type="button" class="nes-btn is-primary" @click="login">Login</button>
            </div>
            <div class="github-logo">
                <a href="https://github.com/liuquanhao/moyu" target="_blank" rel="noopener">
                    <i class="nes-icon github is-medium"></i>
                </a>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "moyu-nav",
        data: function() {
            return {
                isClick: true
            }
        },
        methods: {
            login: function() {
                this.$router.push('/login')
            },
            logout: function() {
                this.isClick = true
                this.axios.post('./api/logout', {}, {withCredentials: true}).then(res => {
                    this.$cookies.remove('uid')
                    this.isClick = false
                    if (this.$route.name == 'MainPage') {
                        this.$router.go(0)
                    } else {
                        this.$router.push('/')
                    }
                }).catch(err => {
                    this.$cookies.remove('uid')
                    this.isClick = false
                    if (this.$route.name == 'MainPage') {
                        this.$router.go(0)
                    } else {
                        this.$router.push('/')
                    }
                });
            }
        },
        computed: {
            logged: function() {
                return this.$cookies.get('uid')
            }
        }
    }
</script>

<style scoped>
.logo a {
    color: black;
    text-decoration: none;
}
.moyu-nav {
    display: flex;
    margin-top: 10px;
    margin-bottom: 10px;
    border-bottom: 4px solid #D3D3D3;
}
.nav-left {
    display: flex;
    width: 50%;
}
.nav-right {
    display: flex;
    justify-content: flex-end;
    width: 50%;
}
.logo {
    margin-left: 10px;
    font-size: 30px;
}
.btn {
    margin-right: 15px;
    margin-bottom: 5px;
}
.github-logo {
    margin-right: 10px;
}
</style>