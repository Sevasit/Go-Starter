package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	// Print environment variables for debugging
	fmt.Printf("DB_HOST: %s, DB_PORT: %s, DB_NAME: %s, DB_USER: %s, DB_PASS: %s\n", host, port, dbName, dbUser, password)

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, password, host, port, dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//Migration

	g := gen.NewGenerator(gen.Config{
		OutPath: "C:/GO Lang/CRUD-Api/entity",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	g.UseDB(DB)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()

}
