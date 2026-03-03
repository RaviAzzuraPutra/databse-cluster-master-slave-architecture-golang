package config

import (
	"databse-cluster-master-slave-architecture-golang/app/config/app_config"
	"databse-cluster-master-slave-architecture-golang/app/config/db_config"
)

func Config() {
	app_config.App_Config()
	db_config.DB_Config()
}
