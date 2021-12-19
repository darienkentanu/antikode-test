package brand

import (
	"antikode-test/api/common"
	"antikode-test/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BrandController struct {
	BrandModel models.BrandModel
}

func NewBrandController(brandModel models.BrandModel) *BrandController {
	return &BrandController{BrandModel: brandModel}
}

func (bc *BrandController) PostBrandController(w http.ResponseWriter, r *http.Request) {
	var brandRequest PostBrandRequest
	err := json.NewDecoder(r.Body).Decode(&brandRequest)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Name == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Logo == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Banner == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	brand := models.Brand{
		Name:   brandRequest.Name,
		Logo:   brandRequest.Logo,
		Banner: brandRequest.Banner,
	}
	_, err = bc.BrandModel.Insert(brand)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewInternalServerErrorResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}

func (bc *BrandController) GetAllBrandController(w http.ResponseWriter, r *http.Request) {
	brands, err := bc.BrandModel.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	response := []GetBrandResponse{}
	for _, brand := range brands {
		response = append(response, GetBrandResponse{ID: brand.ID, Name: brand.Name, Logo: brand.Logo, Banner: brand.Banner})
	}
	json.NewEncoder(w).Encode(&response)
}

func (bc *BrandController) EditBrandController(w http.ResponseWriter, r *http.Request) {
	var brandRequest EditBrandRequest
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
	}
	err = json.NewDecoder(r.Body).Decode(&brandRequest)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Name == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Logo == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if brandRequest.Banner == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	brand := models.Brand{
		Name:   brandRequest.Name,
		Logo:   brandRequest.Logo,
		Banner: brandRequest.Banner,
	}
	if _, err := bc.BrandModel.Edit(id, brand); err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}

func (bc *BrandController) DeleteBrandController(w http.ResponseWriter, r *http.Request) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
	}
	if _, err := bc.BrandModel.Delete(id); err != nil {
		json.NewEncoder(w).Encode(common.NewNotFoundResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}
