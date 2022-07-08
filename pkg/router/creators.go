package router

import (
	"fmt"
	"github.com/AgileProggers/archiv-backend-go/pkg/database"
	"github.com/Gebes/there/v2"
	"strconv"
)

// GetCreators godoc
// @Summary Get all creators
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {array} database.Creator
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/ [get]
// @Param uuid query int false "The uuid of a creator"
// @Param name query string false "The name of a creator"
func GetCreators(request there.HttpRequest) there.HttpResponse {
	var creators []database.Creator
	var query database.Creator

	err := request.Body.BindJson(&query)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind query: %v", err))
	}
	err = bindingValidator.Struct(query)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("validation error: %v", err))
	}

	if err := database.GetAllCreators(&creators, query); err != nil {
		return there.Error(there.StatusNotFound, "No creators found")
	}

	return there.Json(there.StatusOK, creators)
}

// GetCreatorByUUID godoc
// @Summary Get creator by uuid
// @Tags Creators
// @Produce json
// @Success 200 {object} database.Creator
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/{uuid} [get]
// @Param uuid path int true "Unique Identifyer"
func GetCreatorByUUID(request there.HttpRequest) there.HttpResponse {
	var creator database.Creator

	uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	}

	if err := database.GetOneCreator(&creator, uuid); err != nil {
		return there.Error(there.StatusNotFound, "Creator not found")
	}

	return there.Json(there.StatusOK, creator)
}

// CreateCreator godoc
// @Summary Create creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /creators/ [post]
// @Param Body body database.Creator true "Creator obj"
func CreateCreator(request there.HttpRequest) there.HttpResponse {
	var newCreator database.Creator
	var creator database.Creator

	err := request.Body.BindJson(&newCreator)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	}

	if err := database.GetOneCreator(&creator, newCreator.UUID); err == nil {
		return there.Error(there.StatusBadRequest, "Creator already exists")
	}

	if err := database.AddNewCreator(&newCreator); err != nil {
		return there.Error(there.StatusUnprocessableEntity, "Error while creating the model")
	}

	return there.Error(there.StatusCreated, "Created")
}

// PatchCreator godoc
// @Summary Patch creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /creators/{uuid} [patch]
// @Param uuid path int true "Unique Identifier"
// @Param Body body database.Creator true "Creator obj"
func PatchCreator(request there.HttpRequest) there.HttpResponse {
	var newCreator database.Creator

	uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	}

	err = request.Body.BindJson(&newCreator)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	}

	if err := database.PatchCreator(&newCreator, uuid); err != nil {
		return there.Error(there.StatusUnprocessableEntity, "Error while patching the model")
	}

	return there.Error(there.StatusOK, "Updated")
}

// DeleteCreator godoc
// @Summary Delete creator
// @Tags Creators
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /creators/{uuid} [delete]
// @Param uuid path string true "Unique Identifier"
func DeleteCreator(request there.HttpRequest) there.HttpResponse {
	var creator database.Creator

	uuid, err := strconv.Atoi(request.Params.GetDefault("uuid", ""))
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("uuid is invalid: %v", err))
	}

	if err := database.GetOneCreator(&creator, uuid); err != nil {
		return there.Error(there.StatusNotFound, "Creator not found")
	}

	if err := database.DeleteCreator(&creator, uuid); err != nil {
		return there.Error(there.StatusBadRequest, "Error while deleting the model")
	}

	return there.Error(there.StatusOK, "Deleted")
}
