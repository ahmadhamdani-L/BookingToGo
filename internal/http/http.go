package http

import (
	"bookingtogo/internal/app/customer"
	"bookingtogo/internal/app/family_list"
	"bookingtogo/internal/app/nationality"
	"bookingtogo/internal/factory"
	"github.com/gorilla/mux"
)


func Init(r *mux.Router, f *factory.Factory ) {
	
	customer.NewHandler(f).Route(r.PathPrefix("/customer").Subrouter())
	family_list.NewHandler(f).Route(r.PathPrefix("/family_list").Subrouter())
	nationality.NewHandler(f).Route(r.PathPrefix("/nationality").Subrouter())
	
}
