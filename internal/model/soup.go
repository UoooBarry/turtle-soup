package model

import (
	"gorm.io/datatypes"
)

type Soup struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	SoupQuestion string         `json:"soup_question"`
	SoupAnswer   string         `json:"soup_answer"`
	SoupTag      datatypes.JSON `json:"tag" gorm:"type:text"`
	Enabled      bool           `json:"enabled" gorm:"default:true"`
}
