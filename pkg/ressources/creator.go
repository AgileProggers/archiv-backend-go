package ressources

type Creator struct {
	ID   int    `post:"required" patch:"required" json:"id"`
	Name string `post:"required" patch:"required" json:"name"`
}
