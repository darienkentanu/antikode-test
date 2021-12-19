package router

import (
	"antikode-test/api/controller/outlet"
	"antikode-test/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterOutletPath(r *mux.Router, db *gorm.DB) {
	om := models.NewOutletModel(db)
	bm := models.NewBrandModel(db)
	oc := outlet.NewOutletController(om, bm)
	// ------------------------------------------------------------------
	// CRUD Outlet
	// ------------------------------------------------------------------
	r.HandleFunc("/outlets/add/{brandname}", oc.PostOutletController).Methods(http.MethodPost)
	r.HandleFunc("/outlets/getall", oc.GetAllOutletController).Methods(http.MethodGet)
	r.HandleFunc("/outlets/edit/{id}", oc.EditOutletController).Methods(http.MethodPut)
	r.HandleFunc("/outlets/delete/{id}", oc.DeleteOutletController).Methods(http.MethodDelete)
}
