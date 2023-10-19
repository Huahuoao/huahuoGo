package dto

type RateDto struct {
	Duration string `gorm:"column:duration" json:"duration"`
	Rate     string `gorm:"column:rate" json:"rate"`
}
