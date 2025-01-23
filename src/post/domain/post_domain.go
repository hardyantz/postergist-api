package post

type Post struct {
	ID          int    `json:"id" gorm:"id,omitempty"`
	Title       string `json:"title" gorm:"title"`
	Description string `json:"description" gorm:"description"`
	CategoryID  int    `json:"category_id" gorm:"category_id"`
	CreatedDate string `json:"created_date" gorm:"created_date"`
}
