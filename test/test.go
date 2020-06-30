package test

import "example.org/luksam/kiwi-server/config"

// GetTestConfiguration returns the configuration that is used for testing.
func GetTestConfiguration() config.Config {
	return config.Config{
		Database: config.Database{
			Host:     "localhost",
			Port:     5432,
			User:     "kiwi_server_test_user",
			Password: "dbpassword",
			DbName:   "kiwi_server_test",
		},
	}
}
