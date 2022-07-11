package ressources

import "time"


type Clip struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title"`
	Duration int `json:"duration"`
	Date time.Time `json:"date"`
	Filename string `json:"filename"`
	Resolution string `json:"resolution"`
	Size int `json:"size"`
	ViewCount int `json:"view_count"`
}