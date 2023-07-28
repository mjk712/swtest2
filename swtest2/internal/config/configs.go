package config

type StorageConfig struct {
	TokenSecretKey []byte
	DbString       string
	MigrationsPath string
}

func NewStorageConfig() StorageConfig {
	return StorageConfig{
		TokenSecretKey: []byte("GregorySwtest2"),
		DbString:       DBSTRING,
		MigrationsPath: MIGRPATH,
	}
}
