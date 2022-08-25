<template>
    <button class="but"><slot></slot></button>
</template>

<script>
export default {
    props: {
        content: String
    },
    data() {
        return {
            test: 1
        }
    },


}
</script>

<style>
.but {
  color: #fff; /* цвет текста */
  text-decoration: none; /* убирать подчёркивание у ссылок */
  user-select: none; /* убирать выделение текста */
  background: #AFCDE7; /* фон кнопки */
  padding: .7em 1.5em; /* отступ от текста */
  outline: none; /* убирать контур в Mozilla */
  border: none;
} 
.but:hover { background: #D8E6F3; } /* при наведении курсора мышки */
.but:active { background: #9cc3e5; } /* при нажатии */
</style>