package handlers

import (
	_ "github.com/CloudyKit/jet/v6"
	"github.com/simpleittools/celeritas"
	"github.com/simpleittools/celeritastest/data"
	"net/http"
	"time"
)

type Handlers struct {
	App    *celeritas.Celeritas
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.App.LoadTime(time.Now())
	err := h.render(w, r, "home", nil, nil)

	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
