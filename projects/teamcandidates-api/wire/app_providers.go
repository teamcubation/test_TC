package wire

import (
	config "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/config"
)

func ProvideConfigLoader() (config.Loader, error) {
	return config.NewConfigLoader()
}
