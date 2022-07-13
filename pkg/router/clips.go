package router

import (
	"context"
	"fmt"
	"strconv"

	"github.com/AgileProggers/archiv-backend-go/pkg/database"
	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
	"github.com/AgileProggers/archiv-backend-go/pkg/ressources"
	"github.com/Gebes/there/v2"
)

// GetClips godoc
// @Summary Get all clips
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {array} database.Clip
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /clips/ [get]
// @Param uuid query string false "The uuid of a clip"
// @Param title query string false "The title of a clip"
// @Param duration query int false "The duration of a clip"
// @Param date query string false "The date of a clip"
// @Param filename query string false "The filename of a clip"
// @Param resolution query string false "The resolution of a clip"
// @Param size query int false "The size of a clip"
// @Param viewcount query int false "The viewcount of a clip"
// @Param creator query int false "The creator id of a clip"
// @Param game query int false "The game id of a clip"
// @Param vod query string false "The vod id of a clip"
// @Param order query string false "Set order direction divided by comma. Possible ordering values: 'date', 'duration', 'size'. Possible directions: 'asc', 'desc'. Example: 'date,desc'"
func GetClips(request there.HttpRequest) there.HttpResponse {
	var clips []*ent.Clip
	params := map[string][]string(*request.Params)

	clips, err := database.ClipsByQuery(params)
	if err != nil {
		return there.Error(there.StatusBadRequest, fmt.Errorf("unable to filter clips: %v", err))
	}

	return there.Json(there.StatusOK, clips)
}

// GetClipByUUID godoc
// @Summary Get clips by uuid
// @Tags Clips
// @Produce json
// @Success 200 {object} database.Clip
// @Failure 404 {string} string
// @Router /clips/{uuid} [get]
// @Param uuid path string true "Unique Identifier"
func GetClipByUUID(request there.HttpRequest) there.HttpResponse {
	uuid := request.RouteParams.GetDefault("uuid", "")  

	id, err := strconv.Atoi(uuid)
	if err != nil {
		return there.Error(there.StatusBadRequest, err)
	}

	clip, err := database.ClipById(id)
	if err != nil {
		status := there.StatusInternalServerError
		if ent.IsNotFound(err) {
			status = there.StatusNotFound

		}

		return there.Error(status, err)
	}

	return there.Json(there.StatusOK, clip)
}

// CreateClip godoc
// @Summary Create clip
// @Tags Clips
// @Accept json
// @Produce json
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /clips/ [post]
// @Param Body body database.Clip true "Clip obj"
func CreateClip(request there.HttpRequest) there.HttpResponse {
	var clip ressources.Clip

	err := request.Body.BindJson(&clip)
	if err != nil {
		return there.Error(there.StatusBadRequest, err )
	}

	err = postValidator.Struct(clip)
	if err != nil {
		return there.Error(there.StatusBadRequest, err )
	}

	newClip, err := database.CreateClip(clip)
	if err != nil {
		status := there.StatusInternalServerError

		if ent.IsConstraintError(err) {
			status = there.StatusConflict
		}

		if ent.IsValidationError(err) {
			status = there.StatusBadRequest
		}

		return there.Error(status, err )
	}

	return there.Json(there.StatusCreated, newClip)
}

// PatchClip godoc
// @Summary Patch clip
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 422 {string} string
// @Router /clips/{uuid} [patch]
// @Param uuid path string true "Unique Identifier"
// @Param Body body database.Clip true "Clip obj"
func PatchClip(request there.HttpRequest) there.HttpResponse {
	var clip ressources.Clip
	uuid := request.RouteParams.GetDefault("uuid", "")  

	id, convErr := strconv.Atoi(uuid)
	if convErr != nil {
		return there.Error(there.StatusBadRequest, convErr)
	}

	clip.ID = id

	err := request.Body.BindJson(&clip)
	if err != nil {
		return there.Error(there.StatusBadRequest, err )
	}


	err = patchValidator.Struct(clip)
	if err != nil {
		return there.Error(there.StatusBadRequest, err )
	}

	newClip, err := database.PatchClip(id).
		SetDate(clip.Date).
		SetDuration(clip.Duration).
		SetFilename(clip.Filename).
		SetResolution(clip.Resolution).
		SetSize(clip.Size).
		SetTitle(clip.Title).
		SetViewCount(clip.ViewCount).
		Save(context.Background())

	if err != nil {
		return there.Error(there.StatusBadRequest, "Unable to create Clip" )
	}

	return there.Json(there.StatusCreated, newClip)
}

// DeleteClip godoc
// @Summary Delete clip
// @Tags Clips
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Router /clips/{uuid} [delete]
// @Param uuid path string true "Unique Identifier"
func DeleteClip(request there.HttpRequest) there.HttpResponse {
	uuid := request.RouteParams.GetDefault("uuid", "")

	id, convErr := strconv.Atoi(uuid)
	if convErr != nil {
		return there.Error(there.StatusBadRequest, convErr)
	}

	err := database.DeleteClip(id)
	if err != nil {
		return there.Error(there.StatusBadRequest, err )
	}

	return there.Message(there.StatusOK, "Clip Deleted succesfully")
}
