package router

import (
	"fmt"
	"github.com/AgileProggers/archiv-backend-go/pkg/database"
	"github.com/Gebes/there/v2"
	"strings"
)

// GetVods godoc
// @Summary Get all vods
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {array} database.Vod
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /vods/ [get]
// @Param uuid query string false "The uuid of a vod"
// @Param title query string false "The title of a vod"
// @Param duration query int false "The duration of a vod"
// @Param date query string false "The date of a vod"
// @Param filename query string false "The filename of a vod"
// @Param resolution query string false "The resolution of a vod"
// @Param fps query int false "The fps of a vod"
// @Param size query int false "The size of a vod"
// @Param order query string false "Set order direction divided by comma. Possible ordering values: 'date', 'duration', 'size'. Possible directions: 'asc', 'desc'. Example: 'date,desc'"
func GetVods(request there.HttpRequest) there.HttpResponse {
	var vods []database.Vod

	var query database.Vod
	err := request.Body.BindJson(&query)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind query: %v", err))
	}
	err = bindingValidator.Struct(query)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("validation error: %v", err))
	}

	// custom ordering query
	orderParams := request.Params.GetDefault("order", "")
	if len(orderParams) != 0 {
		order := strings.Split(orderParams, ",")
		if len(order) != 2 {
			return there.Error(there.StatusBadRequest, "Invalid order params. Example: 'date,desc'")
		}
		if !stringInSlice(order[0], []string{"date", "duration", "size"}) {
			return there.Error(there.StatusBadRequest, "Invalid first order param. 'date', 'duration' or 'size'")
		}
		if !stringInSlice(order[1], []string{"asc", "desc"}) {
			return there.Error(there.StatusBadRequest, "Invalid second order param. 'asc' or 'desc'")
		}
		orderParams = strings.Replace(orderParams, ",", " ", -1)
	}

	if err := database.GetAllVods(&vods, query, orderParams); err != nil {
		return there.Error(there.StatusNotFound, "No vods found")
	}

	return there.Json(there.StatusOK, vods)
}

// GetVodByUUID godoc
// @Summary Get vod by uuid
// @Tags Vods
// @Produce json
// @Success 200 {object} database.Vod
// @Failure 404 {string} string
// @Router /vods/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetVodByUUID(request there.HttpRequest) there.HttpResponse {
	var vod database.Vod

	uuid := request.Params.GetDefault("uuid", "")
	if err := database.GetOneVod(&vod, uuid); err != nil {
		return there.Error(there.StatusNotFound, "Vod not found")
	}

	return there.Json(there.StatusOK, vod)
}

// CreateVod godoc
// @Summary Create vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /vods/ [post]
// @Param Body body database.Vod true "Vod obj"
func CreateVod(request there.HttpRequest) there.HttpResponse {
	var newVod database.Vod
	var vod database.Vod

	err := request.Body.BindJson(&vod)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	}

	if err := database.GetOneVod(&vod, newVod.UUID); err == nil {
		return there.Error(there.StatusBadRequest, "Vod already exists")
	}

	if err := database.AddNewVod(&newVod); err != nil {
		return there.Error(there.StatusUnprocessableEntity, "Error while creating the model")
	}

	return there.Error(there.StatusCreated, "Created")
}

// PatchVod godoc
// @Summary Patch vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /vods/{uuid} [patch]
// @Param uuid path string true "Unique Identifier"
// @Param Body body database.Vod true "Vod obj"
func PatchVod(request there.HttpRequest) there.HttpResponse {
	var newVod database.Vod
	uuid := request.Params.GetDefault("uuid", "")

	err := request.Body.BindJson(&newVod)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	}

	if err := database.PatchVod(&newVod, uuid); err != nil {
		return there.Error(there.StatusUnprocessableEntity, "Error while patching the model")
	}

	return there.Error(there.StatusOK, "Updated")
}

// DeleteVod godoc
// @Summary Delete vod
// @Tags Vods
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /vods/{uuid} [delete]
// @Param uuid path string true "Unique Identifier"
func DeleteVod(request there.HttpRequest) there.HttpResponse {
	var vod database.Vod

	uuid := request.Params.GetDefault("uuid", "")

	if err := database.GetOneVod(&vod, uuid); err != nil {
		return there.Error(there.StatusNotFound, "Vod not found")
	}

	if err := database.DeleteVod(&vod, uuid); err != nil {
		return there.Error(there.StatusBadRequest, "Error while deleting the model")
	}

	return there.Error(there.StatusOK, "Deleted")
}
