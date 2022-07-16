package ressources

import "time"


type Clip struct {
	ID int `json:"id,omitempty"`
	Title string `post:"required" patch:"required" json:"title"`
	Duration int `post:"required" patch:"required" json:"duration"`
	Date time.Time `post:"required" patch:"required" json:"date"`
	Filename string `post:"required" patch:"required" json:"filename"`
	Resolution string `post:"required" patch:"required" json:"resolution"`
	Size int `post:"required" patch:"required" json:"size"`
	ViewCount int `post:"required" patch:"required" json:"view_count"`
	VodID int `post:"required" patch:"required" json:"vod_id"`
}