package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := "host=" + os.Getenv("POSTGRES_HOST") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " port=" + os.Getenv("POSTGRES_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var version string
	db.Raw("SELECT version()").Scan(&version)

	fmt.Println("Connected!")
	fmt.Println(version)

	
}