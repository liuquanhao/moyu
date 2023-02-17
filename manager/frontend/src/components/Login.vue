<template>
    <div class="nes-container with-title is-centered">
        <p class="title">Login</p>
        <div class="nes-field is-inline">
            <label for="inline_field">Username:</label>
            <input type="text" id="inline_field" class="nes-input is-success">
        </div>
        <div class="nes-field is-inline">
            <label for="inline_field">Password:</label>
            <input type="text" id="inline_field" class="nes-input is-success">
        </div>
    </div>
</template>
  
<script>
    export default {
        name: 'HelloWorld',
        data () {
            return {
            msg: 'Welcome to Your Vue.js App'
            }
        }
    }
</script>

<style scoped>
</style>