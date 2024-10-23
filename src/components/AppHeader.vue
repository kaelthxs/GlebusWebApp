<template>
  <header class="header">
    <label for="search-block" class="header__search-label">
      <input type="text" class="header__search-input" placeholder="Поиск..." id="search-block">
			<button type="submit">adawdaw</button>
      <svg class="header__search-svg">
        <use href="svgs/sprite.svg#man"></use>
      </svg>
    </label>
    <a href="#!" class="header__search-link">SIGN IN</a>
  </header>
</template>

<script>
export default {
  name: 'AppHeader' // Название вашего компонента
  // data() {
  //   return {
  //     // Здесь можно определить реактивные данные
  //     message: 'Привет, мир!',
  //   };
  // },
  // methods: {
  //   // Здесь можно определить методы вашего компонента
  //   greet() {
  //     alert(this.message);
  //   },
  // },
  // mounted() {
  //   // Код, который выполнится при монтировании компонента
  //   console.log('Компонент смонтирован');
  // },
};
</script>
<style>
.header {
	position:sticky;
	top: 0;
	margin-top: 50px;
	margin-right: 60px;
	z-index: 999;
	display: flex;
	align-items: center;
	justify-content: flex-end;
}
.header__search-label {
	display: flex;
	align-items: center;
	box-shadow: -22px 9px 0 1px rgba(0, 0, 0, 0.25);
	background: #1a1919;
	border-radius: 10px;
	width: 935px;
	padding: 10px 15px;
	position: relative;
}
.header__search-input {
	font-size: 22px;
	color: #444;
	background-color: #242424;
	border: none;
	outline: none;
	width: 100%;
	padding: 12px 20px;
	border-radius: 5px;
}
.header__search-svg {
	width: 40px;
	height: 40px;
	fill: #fff;
	position: absolute;
	top: 50%;
	right: 12px;
	transform: translate(-50%, -50%);
}
.header__search-link {
	font-size: 30px;
	text-align: center;
	color: #fff;
	font-family: "Audiowide", sans-serif;
	text-shadow: #bd00ff 1px 0 15px;
	margin-left: 80px;
}
</style>