package domain

import "github.com/farwydi/cleanwhale/config"

type Config struct {
	Project config.ProjectConfig

	Transport struct {
		Webhook config.HTTPConfig
	}
}
