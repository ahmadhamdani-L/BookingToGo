package family_list

import (
	// "net/http"

	"github.com/gorilla/mux"
)

func (h *handler) Route(r *mux.Router) {

	r.HandleFunc("/", h.Get).Methods("GET")
	r.HandleFunc("/create", h.Create).Methods("POST")
}
