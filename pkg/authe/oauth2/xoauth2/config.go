package pkgxaouth2

import (
	pkgoauth2 "github.com/teamcubation/teamcandidates/pkg/authe/oauth2"
)

// Config embede la base, así puedes usar la validación genérica
type Config struct {
	pkgoauth2.BaseConfig
}
