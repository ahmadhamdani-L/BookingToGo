package main

import (
	db "bookingtogo/internal/database"
	// "bookingtogo/internal/database/migration"
	"bookingtogo/internal/factory"
	https "bookingtogo/internal/http"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)



func main() {
	db.Init()
	r := mux.NewRouter()
	f := factory.NewFactory()
	https.Init(r,f)
	// migration.Init()
	fmt.Println("Aplikasi Test BookingToGo start Port 3000...")
	http.ListenAndServe(":3000", r)
}
