package outlet

import (
	"antikode-test/api/common"
	"antikode-test/constants"
	"antikode-test/gmaps"
	"antikode-test/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type OutletController struct {
	OutletModel models.OutletModel
	BrandModel  models.BrandModel
}

func NewOutletController(outletModel models.OutletModel, brandModel models.BrandModel) *OutletController {
	return &OutletController{
		OutletModel: outletModel, BrandModel: brandModel,
	}
}

func catch(w http.ResponseWriter) {
	if r := recover(); r != nil {
		fmt.Println("Google Api Key Is Invalid, Please Use A Valid Api Key")
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		w.Write([]byte(`Google Api Key Is Invalid, Please Use A Valid Api Key`))
	} else {
		fmt.Println("Google Api Key Is Valid")
	}
}

func (oc *OutletController) PostOutletController(w http.ResponseWriter, r *http.Request) {
	brandName := mux.Vars(r)["brandname"]
	brandID, err := oc.BrandModel.GetBrandIdByName(brandName)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	var outletRequest PostOutletRequest
	err = json.NewDecoder(r.Body).Decode(&outletRequest)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if outletRequest.Name == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if outletRequest.Picture == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}
	if outletRequest.Address == "" {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
		return
	}

	longitude, latitude, distance := func() (string, string, float64) {
		defer catch(w)
		longitute, latitude := gmaps.Geocoding(outletRequest.Address)
		distance, _ := gmaps.Distancematrix(outletRequest.Address, constants.DESTINATION)
		return longitute, latitude, distance
	}()

	outlet := models.Outlet{
		Name:      outletRequest.Name,
		Picture:   outletRequest.Picture,
		Address:   outletRequest.Address,
		Longitude: longitude,
		Latitude:  latitude,
		BrandID:   brandID,
		Distance:  float32(distance),
	}
	_, err = oc.OutletModel.Insert(outlet)
	if err != nil {
		json.NewEncoder(w).Encode(common.NewInternalServerErrorResponse())
		return
	}
	json.NewEncoder(w).Encode(common.NewSuccessOperationResponse())
}

func (oc *OutletController) GetAllOutletController(w http.ResponseWriter, r *http.Request) {
	outlets, err := oc.OutletModel.GetAll()
	if err != nil {
		json.NewEncoder(w).Encode(common.NewBadRequestResponse())
	}
	response := []GetOutletResponse{}
	for _, outlet := range outlets {
		response = append(response, GetOutletResponse{
			Name: outlet.Name, Picture: outlet.Picture, Address: outlet.Address, Longitute: outlet.Longitude,
			Latitute: outlet.Latitude, Distance: outlet.Distance,
		})
	}
	json.NewEncoder(w).Encode(&response)
}
