package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joshuabl97/urlShort/data"
)

// redirects the user to the endpoint specified in the url/path
func (h *HandlerHelper) Redirect(w http.ResponseWriter, r *http.Request) {
	endpoint := chi.URLParam(r, "endpoint")
	h.l.Info().Msg("Redirect endpoint: " + endpoint + " being processed")

	ok, url := data.CheckEndpoint(h.l, h.db, endpoint)

	if !ok {
		h.l.Error().Msg("Endpoint not found")
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
