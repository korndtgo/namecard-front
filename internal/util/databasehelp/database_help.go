package databasehelp

import "card-service/internal/config"

type DatabaseHelp struct {
	config *config.Configuration
}

func NewDatabaseHelp(c *config.Configuration) *DatabaseHelp {
	return &DatabaseHelp{
		config: c,
	}
}
