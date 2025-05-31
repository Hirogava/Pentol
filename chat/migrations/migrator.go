package db

import (
	"fmt"
	"github.com/Hirogava/pentol/internal/repository"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(manager *repository.Manager, dbName string) {
	driver, err := postgres.WithInstance(manager.Conn, &postgres.Config{})
	if err != nil {
		panic(fmt.Sprintf("Не удалось создать драйвер миграции: %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://files/migrations/files/%s", dbName),
		"postgres",
		driver,
	)
	if err != nil {
		panic(fmt.Sprintf("Не удалось создать мигратора: %v", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(fmt.Sprintf("Не удалось применить миграции: %v", err))
	}
}