package category

type Category struct {
	ID          int    `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Parent      int    `json:"parent" gorm:"parent"`
	Description string `json:"description" gorm:"description"`
	CreatedDate string `json:"created_date" gorm:"created_date"`
}
