package main

import (
	"github.com/simpleittools/celeritas"
	"github.com/simpleittools/celeritastest/data"
	"github.com/simpleittools/celeritastest/handlers"
	"github.com/simpleittools/celeritastest/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
