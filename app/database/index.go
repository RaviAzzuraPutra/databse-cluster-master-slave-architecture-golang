package database

import (
	"databse-cluster-master-slave-architecture-golang/app/config/db_config"
	"databse-cluster-master-slave-architecture-golang/app/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB_Cluster struct {
	Master        *gorm.DB
	SlaveCases    *gorm.DB
	SlaveSuspects *gorm.DB
}

var dbCluster *DB_Cluster

func Connect() *DB_Cluster {

	var ErrorConnect error

	dsnMaster := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().MASTER_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	dsnSlave1 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().SLAVE1_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	dsnSlave2 := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s client_encoding=UTF8",
		db_config.DB_Config().SLAVE2_HOST, db_config.DB_Config().DB_USER, db_config.DB_Config().DB_PASSWORD, db_config.DB_Config().DB_NAME,
		db_config.DB_Config().DB_PORT, db_config.DB_Config().DB_SSLMODE, db_config.DB_Config().DB_TIMEZONE,
	)

	master, ErrorConnect := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	slave1, ErrorConnect := gorm.Open(postgres.Open(dsnSlave1), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	slave2, ErrorConnect := gorm.Open(postgres.Open(dsnSlave2), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if ErrorConnect != nil {
		panic("An error occurred while trying to connect to the database!! " + ErrorConnect.Error())
	}

	errMigrate := master.AutoMigrate(&models.Cases{}, &models.Suspects{})

	if errMigrate != nil {
		panic("Failed to Migrate Database!! " + errMigrate.Error())
	}

	fmt.Println("=========================================")
	fmt.Println("🚀 Database Cluster Status:")
	fmt.Println("✅ Master Connection: OK!")
	fmt.Println("✅ Slave 1 Connection: OK!")
	fmt.Println("✅ Slave 2 Connection: OK!")
	fmt.Println("✅ Auto Migration: Success!")
	fmt.Println("=========================================")

	dbCluster = &DB_Cluster{
		Master:        master,
		SlaveCases:    slave1,
		SlaveSuspects: slave2,
	}

	return dbCluster
}

func GetInstanceDbCluster() *DB_Cluster {
	if dbCluster == nil {
		panic("Database cluster is not initialized. Please call Connect() first.")
	}
	return dbCluster
}
