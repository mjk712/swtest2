package config

type StorageConfig struct {
	TokenSecretKey []byte
	Host           string
	Port           string
	Username       string
	Password       string
	Database       string
	MigrationsPath string
}

func NewStorageConfig() StorageConfig {
	return StorageConfig{
		TokenSecretKey: []byte("GregorySwtest2"),
		Host:           DBHOST,
		Port:           "5432",
		Username:       DBUSERNAME,
		Password:       DBPASSWORD,
		Database:       DBNAME,
		MigrationsPath: MIGRPATH,
	}
}
