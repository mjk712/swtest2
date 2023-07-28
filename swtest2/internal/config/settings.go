package config

const (
	//DBSTRING   = "user=postgres password=postgres host=host.docker.internal port=5432 database=travelagency sslmode=disable"
	DBHOST     = "host.docker.internal"
	DBUSERNAME = "postgres"
	DBPASSWORD = "postgres"
	DBNAME     = "travelagency"
	MIGRPATH   = "./internal/database/migrations"
	//MIGRPATH = "/Users/grigorijmatukov/projects/swtest2/internal/database/migrations"
)
