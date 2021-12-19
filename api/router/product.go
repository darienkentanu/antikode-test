package router

import (
	"antikode-test/api/controller/product"
	"antikode-test/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterProductPath(r *mux.Router, db *gorm.DB) {
	pm := models.NewProductModel(db)
	bm := models.NewBrandModel(db)
	pc := product.NewProductController(pm, bm)
	// ------------------------------------------------------------------
	// CRUD Product
	// ------------------------------------------------------------------
	r.HandleFunc("/products/add/{brandname}", pc.PostProductController).Methods(http.MethodPost)
	r.HandleFunc("/products/getall", pc.GetAllProductController).Methods(http.MethodGet)
	r.HandleFunc("/products/edit/{id}", pc.EditProductController).Methods(http.MethodPut)
	r.HandleFunc("/products/delete/{id}", pc.DeleteProductController).Methods(http.MethodDelete)
}
