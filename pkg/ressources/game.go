package ressources

type Game struct {
	ID     int    `post:"required" patch:"required" json:"id"`
	Name   string `post:"required" patch:"required" json:"name"`
	Boxart string `post:"required" patch:"required" json:"box_art"`
	Clips  []Clip `post:"required" patch:"required" json:"clips"`
}
