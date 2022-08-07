package middleware

import (
	"github.com/simpleittools/celeritas"
	"github.com/simpleittools/celeritastest/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
