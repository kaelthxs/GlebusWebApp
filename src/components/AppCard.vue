<script>
export default {
  name: 'AppCard',
  data() {
    return {

    };
  },
  props: {
    nameTrack: String,
    songerTrack: String,
    imgTrack: String,
    priceTrack: String,
    text: {
      type: String,
      required: true
    },
  }
  // beforeDestroy() {
  //   clearInterval(this.intervalId);
  // }
};
</script>

<template>
  <li class="section-scroll__item">
    <img :src="imgTrack" class="section-scroll__img">
    <div class="section-scroll__info">
      <div class="section-scroll__about-autor">
        <p class="about-autor__name-track">{{nameTrack}}</p>
        <p class="about-autor__name-autor">{{songerTrack}}</p>
      </div>
      <p class="section-scroll__info-price">
        <svg class="section-scroll__info-price-svg">
          <use href="../assets/img/sprite.svg#crown"></use>
        </svg>
        <span class="section-scroll__info-price-text">{{priceTrack}}</span>
      </p>
    </div>
  </li>
</template>

<style scoped>
  .section-scroll__item {
    padding: 30px;
    box-sizing: border-box;
    flex-direction: column;
    display: flex;
    align-items: center;
    justify-content: space-around;
    border-radius: 20px;
    width: 370px;
    height: 460px;
    box-shadow: -18px 21px 11px -8px rgba(0, 0, 0, 0.25);
    background: linear-gradient(180deg, #212121 0%, #ffc940 100%);
  }
  .section-scroll__img {
    width: 300px;
    height: 300px;
    border-radius: 30px;
    pointer-events: none;
    user-drag: none;
    user-select: none;
  }
  .section-scroll__info {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
  .section-scroll__about-autor {
  }
  .about-autor__name-track {
    font-size: 25px;
    color: #fff;
  }
  .about-autor__name-autor {
    white-space: nowrap;
    max-width: 220px;
    overflow: hidden;

    font-size: 20px;
    color: #fff;
    font-family: 'Arima'; 
  }
  .section-scroll__info-price {
    position: relative;
  }
  .section-scroll__info-price-svg {
    position: absolute;
    top: 0;
    right: 0;
    transform: translate(90%, -50%) rotate(25deg);
    width: 20px;
    height: 20px;
    fill: #fff;
  }
  .section-scroll__info-price-text {
    font-size: 48px;
  }
</style>
