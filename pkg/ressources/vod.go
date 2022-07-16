package ressources

import (
	"time"

	"github.com/AgileProggers/archiv-backend-go/pkg/ent"
)

type Vod struct {
	Id         string  	  `post:"required" patch:"required" json:"id,omitempty"`
	Title      string     `post:"required" patch:"required" json:"title"`
	Duration   int        `post:"required" patch:"required" json:"duration"`
	Date       time.Time  `post:"required" patch:"required" json:"date"`
	Filename   string     `post:"required" patch:"required" json:"filename"`
	Resolution string     `post:"required" patch:"required" json:"resolution"`
	Fps        float32    `post:"required" patch:"required" json:"fps"`
	Size       int        `post:"required" patch:"required" json:"size"`
	Publish    bool       `post:"required" patch:"required" json:"publish"`
	Clips      []ent.Clip `post:"required" patch:"required" json:"clips"`
}