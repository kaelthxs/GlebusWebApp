<script>
import AppCard from './AppCard.vue';
export default {
  name: 'AppSection',
  components: {
    AppCard
  },
  data() {
    return {
      tracks: [
        {
          name: 'Track 1',
          songer: 'Artist 1123123123123123123123123123123123123123',
          img: require('../assets/img/image 1.png'),
          price: '90'
        },
        {
          name: 'Track 2',
          songer: 'Artist 2',
          img: require('../assets/img/image 1.png'),
          price: '100'
        },
        {
          name: 'Track 2',
          songer: 'Artist 2',
          img: require('../assets/img/image 1.png'),
          price: '100'
        }
      ]
    }
  },
  props: {
    nameSection: String
  }
};
</script>

<template>
  <section class="section-scroll">
    <p class="section-scroll__title" id="target">{{nameSection}}</p>
    <ul class="section-scroll__list">
      <AppCard
        v-for="(track, index) in tracks"
        :key="index"
        :nameTrack="track.name"
        :songerTrack="track.songer"
        :imgTrack="track.img"
        :priceTrack="track.price"
      />
    </ul>
  </section>
</template>

<style scoped>
  .section-scroll {
    margin: 30px 0 20px 70px;
    flex-direction: column;
    display: flex;
    justify-content: space-between;
  }
  .section-scroll::after {
    content: '';
    position: absolute;
    top: 50%;
    right: 0;
    transform: translateY(-50%);
    width: 5px;
    height: 60%;
  }
  .section-scroll__title {
    font-size: 40px;
    margin-bottom: 40px;
  }
  .section-scroll__list {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
  }
</style>