<script>
import '../assets/styles/HomePage.css';
import AppHeader from './AppHeader.vue';
import AppNagitaion from './AppNavigation.vue';
import AppSection from './AppSection.vue';

export default {
  name: 'HomePage',
  components: {
    AppHeader,
    AppNagitaion,
    AppSection
  },
  data(){
    return {
      sections: [
        {
          name: 'NameSection1',
        },
        {
          name: 'NameSection2',
        },
        {
          name: 'NameSection3',
        }
      ]
    }
  }
  // props: {
  //   msg: String
  // }
}
</script>

<template>
  <AppNagitaion/>
  <div class="wrapper-main">
    <main class="main">
      <AppHeader/>
      <div class="wrapper-scrollsection">
        <AppSection v-for="(section, index) in sections"
        :key="index"
        :nameSection="section.name"
        />
      </div>
    </main>
  </div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
  .wrapper-scrollsection {
    overflow-y: scroll;
    box-sizing: border-box;
    height: 100vh;
    margin-top: -150px;
    padding-bottom: 150px;
    padding-top: 150px;
    margin-right: 1px;
  }
  /* .wrapper-scrolsection::-webkit-scrollbar {
    width: 0;
  }
  .wrapper-scrolsection::-webkit-scrollbar-track {
    background: none;
  }
  .wrapper-scrolsection::-webkit-scrollbar-thumb {
    background-color: none;
    border-radius: 20px;
  } */
</style>
