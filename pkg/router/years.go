package router

import (
	"github.com/Gebes/there/v2"
)

func GetYears(request there.HttpRequest) there.HttpResponse {
	type Year struct {
		Year  int `json:"year"`
		Count int `json:"count"`
	}
	var years []Year

	// database.Raw("SELECT to_char(vods.date, 'yyyy') AS year, COUNT(to_char(vods.date, 'yyyy')) AS count FROM vods GROUP BY to_char(vods.date, 'yyyy') ORDER BY year DESC").Find(&years)
	// if len(years) < 1 {
	// 	return there.Error(there.StatusBadRequest, "No years found")
	// }

	return there.Json(there.StatusOK, years)
}
