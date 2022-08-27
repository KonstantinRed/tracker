package main

import (
	// пакет для форматированного ввода вывода

	"log" // пакет для логирования
	// пакет для поддержки HTTP протокола
	// пакет для работы с  UTF-8 строками

	"main_mod/internal/serve"

	env "github.com/joho/godotenv"
)

func set_env() {
	err := env.Load()
	if err != nil {
		log.Panicf("Error loading env:%v", err)
	}
}

func main() {
	set_env()

	serve.HttpServe()
}
