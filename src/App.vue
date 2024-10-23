<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
  <HomePage msg="Welcome to our site!"/>
</template>

<script>
import HomePage from './components/HomePage.vue'

export default {
  name: 'App',
  components: {
    HomePage
  }
}
</script>

<style>
#app {
  display: flex;
  max-height: 100vh;
  overflow: hidden;
  margin: 0;
  padding: 0;
  width: 100%;
}
</style>
