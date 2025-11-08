package main

import (
	"database/sql"
	"fmt"
	"log"

	"focussprint/internal/matter/handler"
	"focussprint/internal/matter/infra"
	"focussprint/internal/matter/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "matter.db")
	if err != nil {
		log.Fatalf("не удалось открыть БД: %v", err)
	}
	defer db.Close()

	// проверяем подключение
	if err := db.Ping(); err != nil {
		log.Fatalf("не удалось подключиться к БД: %v", err)
	}

	// применяем миграции
	if err := applyMigrations(db); err != nil {
		log.Fatalf("не удалось применить миграции: %v", err)
	}

	// собираем граф зависимостей
	matterInfra := infra.NewMatterInfra(db)
	matterUsecase := usecase.NewMatterUsecase(matterInfra)
	matterHandler := handler.NewMatterHandler(matterUsecase)

	// пример использования
	matter, err := matterHandler.GetMatter(1)
	if err != nil {
		log.Printf("ошибка получения matter: %v", err)
		return
	}

	fmt.Printf("Matter получен: %+v\n", matter)
}
