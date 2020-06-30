package config

// GetConfiguration returns the hardcoded configuration.
func GetConfiguration() Config {
	return Config{
		ChirpStackApplicationServer: ChirpStackApplicationServer{
			Address:  "localhost:8080",
			Username: "admin",
			Password: "admin",
		},
		Database: Database{
			Host:     "localhost",
			Port:     5432,
			User:     "chirpstack_as_events",
			Password: "dbpassword",
			DbName:   "chirpstack_as_events",
		},
	}
}

// Config contains the whole configuration.
type Config struct {
	ChirpStackApplicationServer ChirpStackApplicationServer
	Database                    Database
}

// ChirpStackApplicationServer contains the configuration required to access the ChirpStack Application Server API.
type ChirpStackApplicationServer struct {
	Address  string
	Username string
	Password string
}

// Database contains the configuration required for the database access.
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}
