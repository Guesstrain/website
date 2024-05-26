package dto

type PersonalInfo struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"size:255"`
	Description string `json:"description" gorm:"type:text"`
	Interest    string `json:"interest" gorm:"type:text"`
	Contact     string `json:"contact" gorm:"size:255"`
}
