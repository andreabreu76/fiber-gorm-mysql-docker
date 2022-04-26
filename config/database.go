package config

import (
	"ackfinance/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDB := os.Getenv("MYSQL_DATABASE")

	DatabaseUri := mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Fprintf(os.Stdout, "Connecting to database: %s\n", DatabaseUri)

	var err error

	Database, err = gorm.Open(mysql.Open(DatabaseUri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic(err)
	}

	err = Database.AutoMigrate(
		&models.AccountStatus{},
		&models.Account{},
		&models.MovimentType{},
		&models.Moviments{},
	)
	if err != nil {
		log.Fatalf("Error on AutoMigrate: %s", err)
	}

	AccountStatuses := []models.AccountStatus{
		{Name: "Ativo"},
		{Name: "Inativo"},
	}
	Database.Create(&AccountStatuses)
	MovimentsTypes := []models.MovimentType{
		{Name: "Depósito"},
		{Name: "Saque"},
		{Name: "Transferência"},
	}
	Database.Create(&MovimentsTypes)

	return nil
}
