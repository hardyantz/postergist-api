package post

type Post struct {
	ID          int    `json:"id" gorm:"id,omitempty"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	CreatedDate string `json:"created_date" gorm:"created_date"`
}
