package config

type MainConfig struct {
	Database DatabaseConfig
	Midtrans PaymentConfig
}

type DatabaseConfig struct {
	// Isa
}

type PaymentConfig struct {
	// Isa
}

func LoadConfig(path string) (config MainConfig) {
	// Isa
	return config
}
