package api

import (
	"example.org/luksam/kiwi-server/api/external"
)

// Setup and configure the API endpoints
func Setup() error {
	return external.Setup()
}
