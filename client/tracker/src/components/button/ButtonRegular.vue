<template>
    <button class="knopka"><slot></slot></button>
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
.knopka {
  color: #fff; /* цвет текста */
  text-decoration: none; /* убирать подчёркивание у ссылок */
  user-select: none; /* убирать выделение текста */
  background: rgb(212,75,56); /* фон кнопки */
  padding: .7em 1.5em; /* отступ от текста */
  outline: none; /* убирать контур в Mozilla */
  border: none;
} 
.knopka:hover { background: rgb(232,95,76); } /* при наведении курсора мышки */
.knopka:active { background: rgb(152,15,0); } /* при нажатии */
</style>