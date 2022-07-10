package router

import (
	"github.com/Gebes/there/v2"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

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
	var clips []int
	// var query database.Clip

	// err := request.Body.BindJson(&query)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind query: %v", err))
	// }
	// err = bindingValidator.Struct(query)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("validation error: %v", err))
	// }

	// // custom ordering query
	// orderParams := request.Params.GetDefault("order", "")
	// if len(orderParams) != 0 {
	// 	order := strings.Split(orderParams, ",")
	// 	if len(order) != 2 {
	// 		return there.Error(there.StatusBadRequest, "Invalid order params. Example: 'date,desc'")
	// 	}
	// 	if !stringInSlice(order[0], []string{"date", "duration", "size"}) {
	// 		return there.Error(there.StatusBadRequest, "Invalid first order param. 'date', 'duration' or 'size'")
	// 	}
	// 	if !stringInSlice(order[1], []string{"asc", "desc"}) {
	// 		return there.Error(there.StatusBadRequest, "Invalid second order param. 'asc' or 'desc'")
	// 	}
	// 	orderParams = strings.Replace(orderParams, ",", " ", -1)
	// }

	// if err := database.GetAllClips(&clips, query, orderParams); err != nil {
	// 	return there.Error(there.StatusNotFound, "No clips found")
	// }

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
	var clip int//database.Clip

	// uuid := request.Params.GetDefault("uuid", "")

	// if err := database.GetOneClip(&clip, uuid); err != nil {
	// 	return there.Error(there.StatusNotFound, "Clip not found")
	// }

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
	// var newClip database.Clip
	// var clip database.Clip

	// err := request.Body.BindJson(&newClip)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind body: %v", err))
	// }

	// if err := database.GetOneClip(&clip, newClip.UUID); err == nil {
	// 	return there.Error(there.StatusBadRequest, "Clip already exists")
	// }

	// if err := database.AddNewClip(&newClip); err != nil {
	// 	return there.Error(there.StatusUnprocessableEntity, "Error while creating the model")
	// }

	return there.Message(there.StatusCreated, "Created")
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
	// var newClip database.Clip
	// uuid := request.Params.GetDefault("uuid", "")

	// err := request.Body.BindJson(&newClip)
	// if err != nil {
	// 	return there.Error(there.StatusBadRequest, fmt.Errorf("unable to bind query: %v", err))
	// }

	// if err := database.PatchClip(&newClip, uuid); err != nil {
	// 	return there.Error(there.StatusUnprocessableEntity, "Error while patching the model")
	// }

	return there.Message(there.StatusOK, "Updated")
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
	// var clip database.Clip
	// uuid := request.Params.GetDefault("uuid", "")

	// if err := database.GetOneClip(&clip, uuid); err != nil {
	// 	return there.Error(there.StatusNotFound, "Clip not found")
	// }

	// if err := database.DeleteClip(&clip, uuid); err != nil {
	// 	return there.Error(there.StatusBadRequest, "Error while deleting the model")
	// }

	return there.Message(there.StatusOK, "Deleted")
}
