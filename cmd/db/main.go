package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/config"
	"github.com/SE-Curriculum-Planner/Long-Plan-Backend/infrastructure"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var fs embed.FS

func init() {
	config.InitConfig()
	infrastructure.InitDB()
}

func main() {
	method := os.Getenv("method")

	db, _ := infrastructure.DB.DB()
	databaseDrv, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	sourceDrv, err := iofs.New(fs, "migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance(
		"iofs", sourceDrv,
		"lwa", databaseDrv)
	if err != nil {
		panic(err)
	}

	switch method {
	case "migrate":
		if err := m.Up(); err != nil {
			panic(err)
		}
		println("migrated!")

	case "up":
		if err := m.Steps(1); err != nil {
			panic(err)
		}
		println("UP(1 version)!")

	case "down":
		if err := m.Steps(-1); err != nil {
			panic(err)
		}
		println("DOWN(1 version)!")

	case "drop":
		if err := m.Down(); err != nil {
			panic(err)
		}
		println("dropped!")

	case "reset":
		// Reset by dropping and migrating again
		if err := m.Down(); err != nil {
			panic(err)
		}
		if err := m.Up(); err != nil {
			panic(err)
		}
		println("reset!")

	default:
		panic(fmt.Errorf("unsupported migration method: %s", method))
	}
}
