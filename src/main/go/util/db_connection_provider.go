package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func GetDbHandle() (*gorm.DB, func()) {
	fmt.Println("Getting db handle")
	var err error
	config := getDbCredentials()
	log.Println("config = ", config)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", config.HostName, config.UserName, config.Password, config.DefaultDB)
	log.Println("dsn = ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		log.Fatal(err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// Ping
	err = sqlDB.Ping()
	if err != nil {
		log.Println("Db connection failed")
		log.Fatal(err)
	} else {
		log.Println("found db connection")
		fmt.Println(sqlDB.Stats())
	}

	// Close

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(1)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//sqlDB.SetConnMaxLifetime(time.Hour)

	//sqlDB.Close()

	return db, func() { sqlDB.Close() }

}
