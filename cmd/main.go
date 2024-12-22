package main

import (
	"example/internal/api"
	"example/internal/business_logic"
	"example/internal/db"
	"example/internal/worker"
)

func main() {
	db := db.New()

	// business_logic.New() принимает интерфейс,
	// но передаем конкретную реализацию, которую можно заменить
	logic := business_logic.New(db)

	// просто пример что логика может быть вызвана разными триггерами так скажем
	go worker.New(logic).Work()

	api.Run(logic)
}
