package product

import (
	"antikode-test/api/common"
	"antikode-test/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductController struct {
	ProductModel models.ProductModel
	BrandModel   models.BrandModel
}

func NewProductController(productModels models.ProductModel, brandModel models.BrandModel) *ProductController {
	return &ProductController{ProductModel: productModels, BrandModel: brandModel}
}

func (pc *ProductController) PostProductController(w http.ResponseWriter, r *http.Request) {
	brandName := mux.Vars(r)["brandname"]
	brandID, err := pc.BrandModel.GetBrandIdByName(brandName)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	var productRequest PostProductRequest
	err = json.NewDecoder(r.Body).Decode(&productRequest)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Name == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Picture == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Price == 0 {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	product := models.Product{
		Name:    productRequest.Name,
		Picture: productRequest.Picture,
		Price:   productRequest.Price,
		BrandID: brandID,
	}
	_, err = pc.ProductModel.Insert(product)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewInternalServerErrorResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}

func (pc *ProductController) GetAllProductController(w http.ResponseWriter, r *http.Request) {
	products, err := pc.ProductModel.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	var responseProducts []GetProductResponse
	for _, product := range products {
		responseProducts = append(responseProducts, GetProductResponse{ID: product.ID, Name: product.Name, Picture: product.Picture, Price: product.Price})
	}
	json.NewEncoder(w).Encode(&responseProducts)
}

func (pc *ProductController) EditProductController(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	product, err := pc.ProductModel.GetProductById(id)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	var productRequest EditProductRequest
	err = json.NewDecoder(r.Body).Decode(&productRequest)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Name == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Picture == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if productRequest.Price == 0 {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	product.Name = productRequest.Name
	product.Picture = productRequest.Picture
	product.Price = productRequest.Price
	_, err = pc.ProductModel.Edit(id, product)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewInternalServerErrorResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}

func (pc *ProductController) DeleteProductController(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	_, err = pc.ProductModel.GetProductById(id)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	_, err = pc.ProductModel.Delete(id)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewInternalServerErrorResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}
