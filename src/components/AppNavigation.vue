<script>
  export default {
    name: 'AppNavigation',
  };
</script>

<template>
  <div class="wrapper-nav">
    <nav class="nav">
      <ul class="nav__list">
        <li class="nav__item nav__item--header">
          <a href="#!" class="nav-header">OOO<br> GLEBUS</a>
          
          <a href="#!" class="burger-menu__wrapper">
            <div class="burger-menu__stick"></div>
            <div class="burger-menu__stick"></div>
            <div class="burger-menu__stick"></div>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--singer">
              <use href="../assets/img/sprite.svg#man"></use>
            </svg>
            <span class="nav__link-text">РЕП ПЕВЦЫ</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--tracks">
              <use href="../assets/img/sprite.svg#fire"></use>
            </svg>
            <span class="nav__link-text">ТРЕКИ</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--price">
              <use href="../assets/img/sprite.svg#like"></use>
            </svg>
            <span class="nav__link-text">ОЦЕНКИ</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--rating">
              <use href="../assets/img/sprite.svg#crown"></use>
            </svg>
            <span class="nav__link-text">РЕЙТИНГ</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--albums">
              <use href="../assets/img/sprite.svg#hand"></use>
            </svg>
            <span class="nav__link-text">АЛЬБОМЫ</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--about">
              <use href="../assets/img/sprite.svg#quastion"></use>
            </svg>
            <span class="nav__link-text">ООО НАС</span>
          </a>
        </li>
        <li class="nav__item">
          <a href="#!" class="nav__link">
            <svg class="nav__svg nav__svg--last">
              <use href="../assets/img/sprite.svg#heart"></use>
            </svg>
            <span class="nav__link-text">ИДИ НАХУЙ</span>
          </a>
        </li>
        <!-- <li class="nav__item">
          <img src="../assets/img/nav-png.png" style="cursor: pointer;">
        </li> -->
      </ul>
    </nav>
  </div>
</template>

<style scoped>
  .wrapper-nav {
    margin-top: 50px;
    min-width: 350px;
  }
  .nav {
    flex-direction: column;
    display: flex;
    align-items: center;
    background-color: #1A1919;
    width: 80px;
    height: calc(100% - 40px);
    margin-bottom: 40px;
    border-radius: 15px;
    transition: .3s;
    overflow: hidden;
  }
  .burger-menu {
    width: 100%;
  }
  .burger-menu__wrapper {
    flex-direction: column;
    display: flex;
    justify-content: center;
    align-items: center;
    /* padding: 0 1px; */
  }
  .burger-menu__stick {
    box-sizing: border-box;
    width: 100%;
    height: 5px;
    background: #fff;
    z-index: 100;
  }
  .burger-menu__stick:nth-child(2) {
    margin: 15px 0;
  }
  .nav:hover {
    width: 360px;
    transition-delay: .1s;
  }
  .nav:hover .nav__item--header {
    width: 100%;
    transition-delay: .1s;
  }
  .nav:hover .nav__svg {
    margin-right: 15px;
    transition-delay: .1s;
    width: 50px;
    height: 50px;
  }
  .nav:hover .nav__link-text {
    /* margin-left: 25px; */
    width: 100%;
    opacity: 1;
    transition-delay: .1s;
  }
  .nav__list {
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: space-around;
  }
  .nav-header {
    display: inline-block;
    font-family: 'Porter Sans Block'; 
    font-size: 30px;
    text-align: center;
    /*  */
    overflow: hidden;
  }
  .nav__item {
    display: flex;
    align-items: center;
  }
  .nav__item--header {
    width: 0;
    padding: 5px 0;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: #1d1d1d;
    border-radius: 10px;
    transition: .3s;
  }
  .nav__item:last-child {
    justify-content: center;
  }
  .nav__svg {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 50px;
    height: 50px;
    margin: 0;
    fill: #fff;
    transition: .3s;
    transition-delay: .3s;
  }
  .nav__svg use {
    display: block;
    width: 100%;
    height: 100%;
  }
  .nav__link-text {
    transition: .3s;
    white-space: nowrap;
    width: 0;
    opacity: 0;
  }
  .nav__link {
    display: flex;
    align-items: center;
    font-size: 32px;
    color: #fff;
  }
  /* .nav__link:hover .nav__link-text{
    transform: scale(0.9);
  }
  .nav__link:hover .nav__svg--singer{
    transform: scale(0.95);
    fill: #400E60;
  }
  .nav__link:hover .nav__svg--tracks{
    transform: scale(0.95);
    fill: #FF5C00;
  }
  .nav__link:hover .nav__svg--price{
    transform: scale(0.95);
    fill: #2D861E;
  }
  .nav__link:hover .nav__svg--rating{
    transform: scale(0.95);
    fill: #FFD600;
  }
  .nav__link:hover .nav__svg--albums{
    transform: scale(0.95);
    fill: #FFBC7E;
  }
  .nav__link:hover .nav__svg--about{
    transform: scale(0.95);
    fill: #3352F4;
  }
  .nav__link:hover .nav__svg--last{
    transform: scale(0.95);
    fill: #FF0000;
  } */
</style>