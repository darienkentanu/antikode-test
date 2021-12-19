package router

import (
	"antikode-test/api/controller/brand"
	"antikode-test/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterBrandPath(r *mux.Router, db *gorm.DB) {
	bm := models.NewBrandModel(db)
	bc := brand.NewBrandController(bm)
	// ------------------------------------------------------------------
	// CRUD Brand
	// ------------------------------------------------------------------
	r.HandleFunc("/brands/add", bc.PostBrandController).Methods(http.MethodPost)
	r.HandleFunc("/brands/getall", bc.GetAllBrandController).Methods(http.MethodGet)
	r.HandleFunc("/brands/edit/{id}", bc.EditBrandController).Methods(http.MethodPut)
	r.HandleFunc("/brands/delete/{id}", bc.DeleteBrandController).Methods(http.MethodDelete)
}
